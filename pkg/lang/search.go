package lang

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"context"
)

var timeoutLower, timeoutUpper, workerNum int

func init() {
	flag.IntVar(&timeoutLower, "lo", 100, "lower limit of timeout value(millisecond)")
	flag.IntVar(&timeoutUpper, "hi", 500, "upper limit of timeout value(millisecond)")
	flag.IntVar(&workerNum, "wn", 5, "number of workers")
}

// Results is an ordered list of search results.
type Results []Result

// A Result contains the title and URL of a search result.
type Result struct {
	Language, Number string
}

type Request struct {
	Ctx      context.Context
	Lang     Lang
	Database []string
	Query    int
	ErrChan  chan<- error
}

func Search(ctx context.Context, query int) (Results, error) {
	var results Results
	lang, ok := FromContext(ctx)
	if ok {
		results = append(results, Result{Language: string(lang), Number: numbers[lang][query-1]})
		return results, nil
	}

	// errChan is used by downstream worker to notify us cancel the whole job
	errChan := make(chan error, workerNum)
	defer close(errChan)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Println("Generate queue")
	queue := genRequests(ctx, query, errChan)

	// fan out
	var resultChans []<-chan Result
	for i := 0; i < workerNum; i++ {
		fmt.Println("Generate worker", i)
		resultChans = append(resultChans, processRequest(ctx, queue))
	}

	// fan in
	fmt.Println("merge worker")
	resultSink := merge(ctx, resultChans...)

	var wg sync.WaitGroup
	// collect
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range resultSink {
			results = append(results, result)
		}
	}()

	var err error
	// make sure err is correctly written if we receive any
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		case err = <-errChan:
			cancel()
			return
		}
	}()

	wg.Wait()
	return results, err
}

func genRequests(ctx context.Context, q int, errChan chan<- error) <-chan Request {
	out := make(chan Request)
	go func() {
		defer close(out)
		for l, nums := range numbers {
			select {
			case out <- Request{Ctx: ctx, Lang: l, Database: nums, Query: q, ErrChan: errChan}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func processRequest(ctx context.Context, queue <-chan Request) <-chan Result {
	out := make(chan Result)
	go func() {
		defer close(out)
		for req := range queue {
			select {
			case <-ctx.Done():
				return
			default:
				result, err := workerDo(req)
				if err != nil {
					req.ErrChan <- err
					return
				}
				out <- result
			}
		}
	}()
	return out
}

func merge(ctx context.Context, inputs ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	out := make(chan Result)

	output := func(c <-chan Result) {
		for n := range c {
			select {
			case out <- n:
			case <-ctx.Done():
			}
		}
		wg.Done()
	}
	wg.Add(len(inputs))
	for _, c := range inputs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func workerDo(req Request) (Result, error) {
	numberOfChunks := rand.Intn(timeoutUpper-timeoutLower) + timeoutLower
	for i := 0; i < numberOfChunks; i++ {
		select {
		case <-req.Ctx.Done(): // check whether this work is cancelled
			return Result{}, errors.New(fmt.Sprintf("timeout while working, cost time %dms", (i+1)+timeoutLower))
		default:
			time.Sleep(time.Millisecond * 1) // mimic working
		}
	}
	return Result{Language: string(req.Lang), Number: req.Database[req.Query-1]}, nil
}
