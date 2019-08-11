package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

// ServerLogEntry ... Server log entry
type ServerLogEntry struct {
	Date          string
	DateEpoch     int64
	Method        string
	URI           string
	SrcIP         net.IP
	SrcPort       string
	Host          string
	Status        int
	Size          int64
	UserAgent     string
	RequestHeader http.Header
}

// LoggingResponseWriter ...
type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// EventLogger ...
type EventLogger struct {
	Description  string
	FunctionName string
	EventType    string //(INFO,ERROR)
	ErrorMessage error
}

// WriteWebLog ... Writes web logs to accessLogs.json file
func WriteWebLog(lg ServerLogEntry, file string) {

	// open log file
	logfile, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Failed to open log file ", logfile, ":", err)
	}

	// 3. Marshal  Log to JSON
	logJSON, err := json.Marshal(lg)
	if err != nil {
		fmt.Println(err)
	}

	// Write log to file
	_, err = logfile.WriteString(string(logJSON))
	if err != nil {
		fmt.Println(err)
	}

	// add new line
	_, err = logfile.WriteString("\n")
	if err != nil {
		fmt.Println(err)
	}
	logfile.Close()
}

// WriteEventLog ... Writes event logs to eventLogs.json file
func WriteEventLog(ev EventLogger, file string) {

	// open log file
	logfile, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Failed to open log file ", logfile, ":", err)
	}

	// 3. Marshal  Log to JSON
	logJSON, err := json.Marshal(ev)
	if err != nil {
		fmt.Println(err)
	}

	// Write log to file
	_, err = logfile.WriteString(string(logJSON))
	if err != nil {
		fmt.Println(err)
	}

	// add new line
	_, err = logfile.WriteString("\n")
	if err != nil {
		fmt.Println(err)
	}
	logfile.Close()
}

// HumanReadableToEpoch ...
func HumanReadableToEpoch(date time.Time) int64 {
	return date.UnixNano() / 1e6 // seconds
}

// NewWebLogEntry ...
func NewWebLogEntry(req *http.Request, w *LoggingResponseWriter, date time.Time) ServerLogEntry {

	var temp ServerLogEntry

	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		// handle err
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		// handle err
	}

	temp.Date = date.UTC().String()
	temp.DateEpoch = HumanReadableToEpoch(date.UTC())
	temp.Method = req.Method
	temp.URI = req.URL.Path
	temp.SrcIP = userIP
	temp.SrcPort = port
	temp.Status = w.StatusCode
	temp.Host = req.Host
	temp.UserAgent = req.UserAgent()
	temp.Size = req.ContentLength
	temp.RequestHeader = req.Header

	return temp
}

// NewEventLogEntry ...
func NewEventLogEntry(desc string, funcName string, eventType string, errorMessage error) EventLogger {

	var temp EventLogger

	temp.Description = desc
	temp.FunctionName = funcName
	temp.EventType = eventType
	temp.ErrorMessage = errorMessage

	return temp
}
