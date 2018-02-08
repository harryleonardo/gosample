package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/tokopedia/gosample/hello"
	grace "gopkg.in/paytm/grace.v1"
	logging "gopkg.in/tokopedia/logging.v1"
)

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	hwm := hello.NewHelloWorldModule()

	http.HandleFunc("/", hwm.RenderHtml)
	http.HandleFunc("/jsonData", hwm.JsonData)
	// http.HandlerFunc("/",hwm.)

	// go logging.StatsLog()

	tracer.Init(&tracer.Config{Port: 8700, Enabled: true})

	log.Fatal(grace.Serve(":9000", nil))
}
