package main

import "net/http"

// func Racer(a, b string) string {
// 	if measureResponseTime(a) < measureResponseTime(b) {
// 		return a
// 	}
// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	startTime := time.Now()
// 	http.Get(url)
// 	return time.Since(startTime)
// }

func Racer(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
