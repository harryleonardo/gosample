package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/tokopedia/gosample/hello"
	grace "gopkg.in/paytm/grace.v1"
	logging "gopkg.in/tokopedia/logging.v1"
)

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	// gops helps us get stack trace if something wrong/slow in production
	// if err := agent.Listen(nil); err != nil {
	// 	log.Fatal(err)
	// }

	hwm := hello.NewHelloWorldModule()

	http.HandleFunc("/", hwm.RenderHtml)
	http.HandleFunc("/jsonData", hwm.JsonData)
	// http.HandlerFunc("/",hwm.)

	// go logging.StatsLog()

	log.Fatal(grace.Serve(":9000", nil))

	// config := nsq.NewConfig()
	// w, _ := nsq.NewProducer("devel-go.tkpd:4150", config)

	// err := w.Publish("Harry", []byte("Hello"))
	// if err != nil {
	// 	log.Panic("Could not connect")
	// }

	// w.Stop()
}
