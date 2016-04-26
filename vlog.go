// Package vlog add leveled log on std log(golang.org/pkg/log/)
// It implements most std log functions(except logger), variables
// and add provides V-style logging controlled by the -v flag or SetLogLevel()
// If flag.Parse be called before any logging, -v flag(default 0) use automaticlly.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Basic examples:
//  vlog.SetLogLevel(3)
//  vlog.GetLogLevel()
//
//	vlog.Println("Prepare to repel boarders")
//
//	vlog.Fatalf("Initialization failed: %s", err)
//
// See the documentation for the V function for an explanation of these examples:
//
//	if vlog.V(2) {
//		vlog.Print("Starting transaction...")
//	}
//
//	vlog.V(2).Println("Processed", nItems, "elements")
package vlog

import (
	"flag"
	"io"
	"log"
)

// These flags define which text to prefix to each log entry generated by the Logger.
const (
	// Bits or'ed together to control what's printed.
	// There is no control over the order they appear (the order listed
	// here) or the format they present (as described in the comments).
	// The prefix is followed by a colon only when Llongfile or Lshortfile
	// is specified.
	// For example, flags Ldate | Ltime (or LstdFlags) produce,
	//	2009/01/23 01:23:23 message
	// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = log.Ldate             // the date in the local time zone: 2009/01/23
	Ltime         = log.Ltime             // the time in the local time zone: 01:23:23
	Lmicroseconds = log.Lmicroseconds     // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile     = log.Llongfile         // full file name and line number: /a/b/c/d.go:23
	Lshortfile    = log.Lshortfile        // final file name element and line number: d.go:23. overrides Llongfile
	LUTC          = log.LUTC              // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = log.Ldate | log.Ltime // initial values for the standard logger
)

var level uint

// Verbose is a boolean type that implements Print like function.
// See the documentation of V for more information.
type Verbose bool

func init() {
	flag.UintVar(&level, "v", 0, "log level for V logs")
}

// Set log level, just use at parameter initialize zone.
func SetLogLevel(v uint) {
	level = v
}

// Get log level, just use at parameter initialize zone.
func GetLogLevel() uint {
	return level
}

func (vb Verbose) Printf(format string, v ...interface{}) {
	if vb {
		log.Printf(format, v)
	}
}

func (vb Verbose) Println(v ...interface{}) {
	if vb {
		log.Println(v)
	}
}

func (vb Verbose) Print(v ...interface{}) {
	if vb {
		log.Print(v)
	}
}

// Whether an individual call to V generates a log record depends on the setting of level.
func V(v uint) Verbose {
	if v <= level {
		return true
	}
	return false
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
func Flags() int {
	return log.Flags()
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	log.Flags()
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return log.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

// These functions write to the standard logger.

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	log.Print(v)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	log.Println(v)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	log.Fatal(v)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	log.Fatalln()
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	log.Panic(v)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	log.Panicf(format, v)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	log.Panicln(v)
}

// Output writes the output for a logging event.  The string s contains
// the text to print after the prefix specified by the flags of the
// Logger.  A newline is appended if the last character of s is not
// already a newline.  Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if Llongfile or Lshortfile is set; a value of 1 will print the details
// for the caller of Output.
func Output(calldepth int, s string) error {
	return log.Output(calldepth, s)
}
