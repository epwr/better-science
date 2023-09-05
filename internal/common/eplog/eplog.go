// log provides a structured logging solution that has the same basic interface as slog, but
// is intended to be simpler (convention > configuration) and to output jsonl lines to power
// GCP's logging system.
//
// NOTE: currently, log can only handle the basic primitives as args to the log functions.
// Check addArgs for a complete list of what is usable.
//
// Copywrite 2023 Eric Power - All Rights Reserved
package log

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type log_level string
const (
	DEBUG log_level = "DEBUG"
	INFO = "INFO"
	WARN = "WARN"
	ERROR = "ERROR"
	CRITICAL = "CRITICAL"
)

// code_location specifies a location in the code.
//
// NOTE: fields are exported so they can be encoded into JSON. 
type code_location struct {
	Function string `json:"function"`
	File string `json:"file"`
	Line int `json:"line"`
}

// log_record stores the core information needed to write to the logs
//
// NOTE: fields are exported so they can be encoded into JSON.
type log_record struct {
	Level string `json:"level"`
	Message string `json:"message"`
	Caller code_location `json:"location"`
	Data map[string]any `json:"data"`
}

// addArgs parses the args list into a map, then sets the record's Data field
// to the map. It follows the behaviour of slog, in that:
//
// 1. If an argument is a string and this is not the last argument, the following argument is
// treated as the value
// 
// 2. Otherwise, the argument is treated as a value with key "!BADKEY".
func (r *log_record) addArgs(args... any) {

	var data = make(map[string]any)

	for len(args) > 1 {

		switch key := args[0].(type) {
		case string:
			value := args[1]
			data[key] = value
			args = args[2:]
		default:
			keyStr, value := createBadKeyPair(args[0])
			data[keyStr] = value
			args = args[1:]
		}
	}
	if len(args) == 1 {
		key, value := createBadKeyPair(args[0])
		data[key] = value
	}

	r.Data = data
}


// createBadKeyPair returns a key, value pair where the key
// is set to "!BADKEY" to match slog's behaviour.
func createBadKeyPair(argument any) (string, any) {

	return "!BADKEY", argument	
}


// Info writes a message to the logs with a severity Level of INFO.
//
// The first argument is a message string, and the remaining arguments
// should be an even numbered amount of values that represent key:value
// pairs (eg. args[0] is a key, args[1] is a value).
func Info(msg string, args... any) {

	code_location := get_location_of_caller()

	record := log_record {
		Level: INFO,
		Message: msg,
		Caller: code_location,
	}
	
	record.addArgs(args...)
	write_log_record(record)
}

// Warn writes a message to the logs with a severity Level of WARN.
//
// The first argument is a message string, and the remaining arguments
// should be an even numbered amount of values that represent key:value
// pairs (eg. args[0] is a key, args[1] is a value).
func Warn(msg string, args... any) {
	
	code_location := get_location_of_caller()

	record := log_record {
		Level: WARN,
		Message: msg,
		Caller: code_location,
	}

	record.addArgs(args...)
	write_log_record(record)
}

// Error writes a message to the logs with a severity Level of ERROR.
//
// The first argument is a message string, and the remaining arguments
// should be an even numbered amount of values that represent key:value
// pairs (eg. args[0] is a key, args[1] is a value).
func Error(msg string, args... any) {
	
	code_location := get_location_of_caller()

	record := log_record {
		Level: ERROR,
		Message: msg,
		Caller: code_location,
	}

	record.addArgs(args...)
	write_log_record(record)
}


// Critical writes a message to the logs with a severity Level of CRITICAL.
//
// The first argument is a message string, and the remaining arguments
// should be an even numbered amount of values that represent key:value
// pairs (eg. args[0] is a key, args[1] is a value).
func Critical(msg string, args... any) {

	code_location := get_location_of_caller()

	record := log_record {
		Level: CRITICAL,
		Message: msg,
		Caller: code_location,
	}
	
	record.addArgs(args...)
	write_log_record(record)
}


// write_log_record writes out the contents of a log_record to stdout.
func write_log_record(record log_record) {
	
	jsonData, err := json.Marshal(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}



// get_location_of_caller returns the code_location (function name, file, and line number) of where this function's
// caller was called from (2 levels back down the call stack). 
func get_location_of_caller() code_location {

	prog_counters := make([]uintptr, 1)
	entry_count := runtime.Callers(3, prog_counters)
	frames := runtime.CallersFrames(prog_counters[:entry_count])

	frame, _ := frames.Next()

	code_location := code_location {
		Function: frame.Function,
		File: frame.File,
		Line: frame.Line,
	}

	return code_location
}
