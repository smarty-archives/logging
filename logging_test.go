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
	out := new(bytes.Buffer)
	log.SetOutput(out)
	log.SetFlags(0)
	defer log.SetOutput(os.Stdout)
	defer log.SetFlags(log.LstdFlags)

	thing := new(ThingUnderTest)
	thing.Action()

	assertions.New(t).So(out.String(), should.Equal, "Hello, World!\n")
}

func TestLoggingWithLoggerCapturesOutput(t *testing.T) {
	out := new(bytes.Buffer)
	log.SetOutput(out)
	log.SetFlags(0)
	defer log.SetOutput(os.Stdout)
	defer log.SetFlags(log.LstdFlags)

	thing := new(ThingUnderTest)
	thing.log = Capture()
	thing.Action()

	assertions.New(t).So(thing.log.Log.String(), should.Equal, "Hello, World!\n")
	assertions.New(t).So(out.Len(), should.Equal, 0)
}

/////////////////////////////////////////////////////////////////////////////////

type ThingUnderTest struct {
	log *Logger
}

func (this *ThingUnderTest) Action() {
	this.log.Println("Hello, World!")
}
