package main

import (
	"fmt"
	"net/http"
	"sort"
	// "github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promauto"
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

	for _, key := range keyList {
		fmt.Fprintln(w, key, ":", r.Header[key], "</br>")
	}

}

func main() {
	//recordMetrics()
	http.HandleFunc("/", printSortedHeader)
	//http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":2112", nil)
}
