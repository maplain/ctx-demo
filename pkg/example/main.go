package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/maplain/ctx-demo/pkg/lang"
)

func main() {
	flag.Parse()
	http.HandleFunc("/number", handleNumber)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleNumber(w http.ResponseWriter, req *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	done := make(chan struct{})
	defer close(done)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires.
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	go func() {
		select {
		case <-time.After(timeout):
			cancel()
		case <-done:
			return
		}
	}()

	// Check the search query.
	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	if language, err := lang.FromRequest(req); err == nil {
		ctx = lang.NewContext(ctx, language)
	}

	// Run the Google search and print the results.
	start := time.Now()
	n, err := strconv.Atoi(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if n <= 0 || n > 10 {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	results, err := lang.Search(ctx, n)
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := resultsTemplate.Execute(w, struct {
		Results          lang.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
}

var resultsTemplate = template.Must(template.New("results").Parse(`{{range .Results}}{{.Language}}-{{.Number}}
{{end}}{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}
`))
