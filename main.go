// Copyright © 2020 morgulbrut
// This work is free. You can redistribute it and/or modify it under the
// terms of the Do What The Fuck You Want To Public License, Version 2,
// as published by Sam Hocevar. See the LICENSE file for more details.

// main.go server: https://gist.github.com/Boerworz/b683e46ae0761056a636

package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/fuzzyserver/handles"
)

func main() {
	// We need to cast handleRoot to a http.HandlerFunc since wrapHandlerWithLogging
	// takes a http.Handler, not an arbitrary function.
	handler := wrapHandlerWithLogging(http.HandlerFunc(handleRoot))
	http.Handle("/", handler)
	http.Handle("/r100", wrapHandlerWithLogging(http.HandlerFunc(handles.R100)))
	http.Handle("/r200", wrapHandlerWithLogging(http.HandlerFunc(handles.R200)))
	http.Handle("/r300", wrapHandlerWithLogging(http.HandlerFunc(handles.R300)))
	http.Handle("/r400", wrapHandlerWithLogging(http.HandlerFunc(handles.R400)))
	http.Handle("/r500", wrapHandlerWithLogging(http.HandlerFunc(handles.R500)))

	rand.Seed(time.Now().UnixNano())

	log.Fatal(http.ListenAndServe(":1337", nil))
}

// ---- Handlers

func handleRoot(w http.ResponseWriter, req *http.Request) {
	handle := rand.Intn(5)
	if handle == 0 {
		handles.R100(w, req)
	}
	if handle == 1 {
		handles.R200(w, req)
	}
	if handle == 2 {
		handles.R300(w, req)
	}
	if handle == 3 {
		handles.R400(w, req)
	}
	if handle == 4 {
		handles.R500(w, req)
	}

}

// ---- Logging

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf(color256.HiYellow("--> %s %s %s %s", req.Method, req.URL.Path, req.RemoteAddr, req.UserAgent()))

		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		statusCode := lrw.statusCode
		log.Printf(color256.HiCyan("<-- %d %s", statusCode, http.StatusText(statusCode)))
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
