package report

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"runtime"
	"time"
)

type Report struct {
	Date        string   `json:"date"`         // report create date
	Author      string   `json:"author"`       // the report author
	OS          string   `json:"os"`           // type of the operating system
	Command     string   `json:"command"`      // the command that was executed
	CommandArgs []string `json:"command_args"` // the command arguments
	Stdout      []string `json:"stdout"`       // the message commig from stdout
	Stderr      []string `json:"stderr"`       // the message commig from stderr
}

// NewReport initialize and return a Report object
func NewReport(author string) *Report {
	r := Report{}
	r.Date = time.Now().Format("20060102-150405") // set to current time
	r.Author = author
	r.OS = runtime.GOOS
	return &r
}

// Post encode the data as json and send a POST request
func Post(url, author, cmd string, cmdArgs, stdout, stderr []string) (resp *http.Response, err error) {
	r := NewReport(author)
	r.Command = cmd
	r.CommandArgs = cmdArgs
	r.Stdout = stdout
	r.Stderr = stderr

	tmpDataJSON, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	req, errReq := http.NewRequest("POST", url, bytes.NewBuffer(tmpDataJSON))
	if errReq == nil {
		req.Header.Set("X-Custom-Header", "cmdreport")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			return nil, errors.New("Aborted! Server Offline")
		}
	}
	return resp, nil
}
