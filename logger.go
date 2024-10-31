package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)


func LogHTTPError(err error, r *http.Request) {
	headers := ""
	for key, values := range r.Header {
		headers += key + ": " + strings.Join(values, ", ") + "; "
	}
	
	body, _ := io.ReadAll(r.Body)

	log.Error().Str("method", r.Method).Str("path", r.URL.Path).Str("query", r.URL.RawQuery).Str("headers", headers).Str("body", string(body)).Msg(err.Error())
}

func LogHTTPInfo(msg string, r *http.Request) {
	headers := ""
	for key, values := range r.Header {
		headers += key + ": " + strings.Join(values, ", ") + "; "
	}
	
	body, _ := io.ReadAll(r.Body)

	log.Info().Str("method", r.Method).Str("path", r.URL.Path).Str("query", r.URL.RawQuery).Str("headers", headers).Str("body", string(body)).Msg(msg)
}

func LogFatal(msg string) {
	log.Fatal().Msg(msg)
}

func LogInfo(msg string) {
	log.Info().Msg(msg)
}
