package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// var (
// 	RequestTotal = func() *prometheus.CounterVec {
// 		return prometheus.NewCounterVec(
// 			prometheus.CounterOpts{
// 				Name: "promhttp_metric_handler_requests_total",
// 				Help: "Total number of scrapes by HTTP status code.",
// 			},
// 			[]string{"code"},
// 		)
// 	}()
// )

// func init() {
// 	prometheus.MustRegister(RequestTotal)
// }

func printSortedHeader(w http.ResponseWriter, r *http.Request) {

	headers := make(map[string]interface{})

	for k, v := range r.Header {
		headers[strings.ToLower(k)] = string(v[0])
	}

	keys := make([]string, 0, len(headers))
	for k := range headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Fprintln(w, k, ": ", headers[k])
	}

	// commented code for parsing map
	// temp := `<table style="width:100%"> hahahaha
	// 	{{range $key, $Value := .headers -}}
	// 	<tr> {{$key}}: {{$Value}} neww </tr>
	// 	{{end}}
	// 	</table>`

	// templ, err := template.New("headers").Parse(temp)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = templ.Execute(w, headers)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func main() {
	finish := make(chan bool)

	go func() {
		headerserver := http.NewServeMux()
		headerserver.HandleFunc("/", printSortedHeader)
		http.ListenAndServe(":8080", headerserver)
	}()

	go func() {
		//RequestTotal.With(prometheus.Labels{"code": "200"}).Inc()
		metricserver := http.NewServeMux()
		metricserver.Handle("/", promhttp.Handler())
		http.ListenAndServe(":9110", metricserver)
	}()

	<-finish
}
