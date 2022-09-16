package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	nodeName := os.Getenv("CURRENT_NODE_NAME")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "This request Scheduled on node %s \n", nodeName)
	})
	log.Fatalln(http.ListenAndServe(":80", nil))
}
