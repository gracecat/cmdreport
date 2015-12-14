package main

import (
	"encoding/json"
	"fmt"
	"github.com/gracecat/cmdreport/report"
	"github.com/kr/pretty"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("cmdreport testserver")

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.Method == "GET" {
			http.Error(rw, `{"error":"GET Method not available"}`, http.StatusBadRequest)
			return
		}

		// get the posted data and parse it...
		bodyBytes, _ := ioutil.ReadAll(req.Body)
		t := report.Report{}
		err := json.Unmarshal(bodyBytes, &t)
		if err != nil {
			http.Error(rw, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}
		fmt.Printf("%# v", pretty.Formatter(t))
		fmt.Fprintf(rw, `{"status":"ok"}`)
	})

	panic(http.ListenAndServe(":9000", nil))
}
