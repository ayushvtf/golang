package main

import (
	"fmt"
	"net/http"
	"sort"

	// "github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func printSortedHeader(w http.ResponseWriter, r *http.Request) {

	// print header
	/*
		for k, vals := range r.Header {
			fmt.Print(k, " : ")
			for _, v := range vals {
				fmt.Println(v)
			}
		}
	*/

	var keyList []string
	for key := range r.Header {
		keyList = append(keyList, key)
	}
	sort.Strings(keyList)

	// m := make(map[string]string)
	// var tinderMatch = make(map[string]string)
	// for _, key := range keyList {
	// 	m[key] = r.Header[key]
	// }

	//fmt.Print(string(m))
	for _, key := range keyList {
		fmt.Fprintln(w, key, ":", r.Header[key], "</br>")
	}

	// t, err := template.ParseFiles("files/test.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // items := struct {
	// // 	Country string
	// // 	City    string
	// // }{
	// // 	Country: "Australia",
	// // 	City:    "Paris",
	// // }
	// t.Execute(w, keyList)

}

func prometheusCount(w http.ResponseWriter, r *http.Request) {
	http.Handle("/", promhttp.Handler())
}

func main() {
	finish := make(chan bool)

	server8080 := http.NewServeMux()
	server8080.HandleFunc("/", printSortedHeader)

	// server9110 := http.NewServeMux()
	// server9110.HandleFunc("/prometheus", prometheusCount)

	go func() {
		http.ListenAndServe(":8080", server8080)
	}()

	// go func() {
	// 	http.Handle("/metrics", promhttp.Handler())
	// 	http.ListenAndServe(":9110", server9110)
	// }()

	<-finish
}
