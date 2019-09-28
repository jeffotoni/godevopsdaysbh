package main

import (
	"log"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost == r.Method {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"error method!"}`))
		return
	}
	time.Sleep(time.Millisecond * 300)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"success"}`))
	return
}

// Middleware Logger
func Logger(name string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			log.Printf(
				"%s %s %s %s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		})
	}
}

type Adapter func(http.Handler) http.Handler

// Middleware
func Middleware(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func main() {

	mux := http.NewServeMux()

	println("Server Run... 8181")
	//mux.Handle("/hello", http.HandlerFunc(Hello))
	mux.Handle("/hello",
		Middleware(
			http.HandlerFunc(Hello),
			Logger("/hello"),
		))

	mux.Handle("/hello2", http.HandlerFunc(Hello))
	mux.Handle("/hello3", http.HandlerFunc(Hello))
	mux.Handle("/hello4", http.HandlerFunc(Hello))
	mux.Handle("/hello5", http.HandlerFunc(Hello))
	mux.Handle("/hello6", http.HandlerFunc(Hello))

	sm := &http.Server{
		Addr:    ":8181",
		Handler: mux,
	}

	sm.ListenAndServe()
}
