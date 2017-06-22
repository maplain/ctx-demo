package lang

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"
)

func FromRequest(req *http.Request) (Lang, error) {
	// Check the search query.
	lang := req.FormValue("lang")
	if lang == "" {
		return "", errors.New("no language setting")
	}
	_, ok := numbers[Lang(lang)]
	if !ok {
		return "", errors.New("unsupported language")
	}
	return Lang(lang), nil
}

type key int

const langKey key = 0

func NewContext(ctx context.Context, lang Lang) context.Context {
	return context.WithValue(ctx, langKey, lang)
}

func FromContext(ctx context.Context) (Lang, bool) {
	lang, ok := ctx.Value(langKey).(Lang)
	return lang, ok
}
