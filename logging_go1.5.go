// +build go1.5

package logging

import (
	"io"
	"log"
)

// Output -> log.Output
func (this *Logger) Output(calldepth int, s string) error {
	if this == nil {
		return log.Output(calldepth, s)
	}
	this.Calls++
	return this.Logger.Output(calldepth, s)
}

// SetOutput -> log.SetOutput
func (this *Logger) SetOutput(w io.Writer) {
	if this == nil {
		log.SetOutput(w)
	} else {
		this.Logger.SetOutput(w)
	}
}
