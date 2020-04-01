package logging

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestLoggingWithNilReferenceProducesTraditionalBehavior(t *testing.T) {
	defer restore()
	out := prepare()

	thing := new(ThingUnderTest)
	thing.Action()

	assertions.New(t).So(out.String(), should.Equal, "Hello, World!\n")
}

func TestLoggingWithLoggerCapturesOutput(t *testing.T) {
	defer restore()
	out := prepare()

	thing := new(ThingUnderTest)
	thing.log = Capture()
	thing.Action()

	assertions.New(t).So(thing.log.Log.String(), should.Equal, "Hello, World!\n")
	assertNothingLoggedToStandardLogOutput(t, out)
}

func TestLoggingWithLoggerCapturesOutputInAllProvidedWriters(t *testing.T) {
	defer restore()

	out0 := prepare()
	out1 := new(bytes.Buffer)
	out2 := new(bytes.Buffer)

	thing := new(ThingUnderTest)
	thing.log = Capture(out1, out2)
	thing.Action()

	log0 := thing.log.Log.String()
	log1 := out1.String()
	log2 := out2.String()

	assertions.New(t).So(log0, should.Equal, "Hello, World!\n")
	assertions.New(t).So(log1, should.Equal, "Hello, World!\n")
	assertions.New(t).So(log2, should.Equal, "Hello, World!\n")
	assertNothingLoggedToStandardLogOutput(t, out0)
}

func TestLogCallsAreCounted(t *testing.T) {
	defer restore()

	thing := new(ThingUnderTest)
	thing.log = Capture()
	for x := 0; x < 10; x++ {
		thing.Action()
	}
	assertions.New(t).So(thing.log.Calls, should.Equal, 10)
}

func TestLoggingWithDiscard(t *testing.T) {
	defer restore()
	out := prepare()

	thing := new(ThingUnderTest)
	thing.log = Discard()
	thing.Action()

	assertions.New(t).So(thing.log.Log.Len(), should.Equal, 0)
	assertNothingLoggedToStandardLogOutput(t, out)
}

func assertNothingLoggedToStandardLogOutput(t *testing.T, out *bytes.Buffer) bool {
	return assertions.New(t).So(out.Len(), should.Equal, 0)
}

/////////////////////////////////////////////////////////////////////////////////

func restore() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags)
}

func prepare() *bytes.Buffer {
	out := new(bytes.Buffer)
	log.SetOutput(out)
	log.SetFlags(0)
	return out
}

/////////////////////////////////////////////////////////////////////////////////

type ThingUnderTest struct {
	log *Logger
}

func (this *ThingUnderTest) Action() {
	this.log.Printf("Hello, World!")
}
