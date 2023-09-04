// A package that provides basic structured logging.
package log

// TODOs:
// - Add structured logging (add in data dictionaries, and display all non-message keys & values.
// - Add timestamps
// - Decide on default format (LEVEL - timestamp - message - record)? Just record? 
// - Can CallerLocation and LogRecord not be public?
// - Get logging level from the config system

import (
	"fmt"
	"runtime"
	"strconv"
)

type log_level string
const (
	DEBUG log_level = "DEBUG"
	INFO = "INFO"
	WARN = "WARN"
	ERROR = "ERROR"
	CRITICAL = "CRITICAL"
)

// A location in the code.
// Should not be used externally to this package.
type CallerLocation struct {
	Function string
	File string
	Line int
}

// A record that stores the core information for logging purposes.
// Should not be used externally to this package.
type LogRecord struct {
	Level string
	Message string
	Caller CallerLocation
}


// Writes an information  message to the logs.
func Info(msg string) {

	caller_location := get_caller_location()

	record := LogRecord {
		Level: INFO,
		Message: msg,
		Caller: caller_location,
	}

	write_log_record(record)
}

// Writes a warning message to the logs.
func Warn(msg string) {
	
	caller_location := get_caller_location()

	record := LogRecord {
		Level: INFO,
		Message: msg,
		Caller: caller_location,
	}

	write_log_record(record)
}

// Writes an error message to the logs.
func Error(msg string) {
	
	caller_location := get_caller_location()

	record := LogRecord {
		Level: INFO,
		Message: msg,
		Caller: caller_location,
	}

	write_log_record(record)
}


// Writes a critical error message to the logs.
func Critical(msg string) {

	caller_location := get_caller_location()

	record := LogRecord {
		Level: INFO,
		Message: msg,
		Caller: caller_location,
	}
	write_log_record(record)
}


// A single function to write out the contents of a LogRecord.
func write_log_record(record LogRecord) {

	location_string := record.Caller.File + " - " + strconv.Itoa(record.Caller.Line)
	fmt.Println(record.Level + " - " +  location_string + ": " + record.Message)
}



// Get the location (function name, file, and line number) of where this function's
// caller was called from (2 levels back down the call stack). 
func get_caller_location() CallerLocation {

	prog_counters := make([]uintptr, 1)
	entry_count := runtime.Callers(3, prog_counters)
	frames := runtime.CallersFrames(prog_counters[:entry_count])

	frame, _ := frames.Next()

	caller_location := CallerLocation {
		Function: frame.Function,
		File: frame.File,
		Line: frame.Line,
	}

	return caller_location
}
