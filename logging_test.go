package logging

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLoggingWithNilReferenceProducesTraditionalBehavior(t *testing.T) {
	defer restore()
	out := prepare()

	thing := new(ThingUnderTest)
	thing.Action()

	assertCallDepthAccuracy(t, out)
	assertCorrectLogMessage(t, out)
}

func TestLoggingWithLoggerCapturesOutput(t *testing.T) {
	defer restore()
	out := prepare()

	thing := new(ThingUnderTest)
	thing.log = Capture()
	thing.Action()

	assertCallDepthAccuracy(t, thing.log.Log)
	assertCorrectLogMessage(t, thing.log.Log)
	assertNothingLogged(t, out)
}

func TestLoggingWithLoggerCapturesOutputInAllProvidedWriters(t *testing.T) {
	defer restore()

	out0 := prepare()
	out1 := new(bytes.Buffer)
	out2 := new(bytes.Buffer)

	thing := new(ThingUnderTest)
	thing.log = Capture(out1, out2)
	thing.Action()

	assertCorrectLogMessage(t, thing.log.Log)
	assertCorrectLogMessage(t, out1)
	assertCorrectLogMessage(t, out2)
	assertNothingLogged(t, out0)
}

func TestLogCallsAreCounted(t *testing.T) {
	defer restore()

	callCount := 10

	thing := new(ThingUnderTest)
	thing.log = Capture()

	for x := 0; x < callCount; x++ {
		thing.Action()
	}

	assertLogCallsCounted(t, thing.log.Calls, callCount)
}

func TestLoggingWithDiscard(t *testing.T) {
	defer restore()
	out := prepare()

	thing := new(ThingUnderTest)
	thing.log = Discard()
	thing.Action()

	assertNothingLogged(t, thing.log.Log)
	assertNothingLogged(t, out)
}

func assertLogCallsCounted(t *testing.T, actual int, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("\n"+
			"Expected: %d\n"+
			"Actual:   %d",
			expected, actual)
	}
}
func assertCorrectLogMessage(t *testing.T, out *bytes.Buffer) {
	t.Helper()
	expected := "Hello, World!\n"
	actual := out.String()
	if !strings.HasSuffix(actual, expected) {
		t.Errorf("\n"+
			"Expected Suffix: %s\n"+
			"Actual Log:      %s",
			expected, actual)
	}
}
func assertCallDepthAccuracy(t *testing.T, out *bytes.Buffer) {
	t.Helper()
	expected := "thing_under_test.go:8"
	actual := out.String()
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("\n"+
			"Expected Prefix: %s\n"+
			"Actual Log:      %s",
			expected, actual)
	}
}
func assertNothingLogged(t *testing.T, out *bytes.Buffer) {
	t.Helper()
	length := out.Len()
	if length > 0 {
		t.Errorf("Expected log to be empty, but it had a length of %d.", length)
	}
}

/////////////////////////////////////////////////////////////////////////////////

func restore() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags)
}

func prepare() *bytes.Buffer {
	out := new(bytes.Buffer)
	log.SetOutput(out)
	log.SetFlags(log.Lshortfile)
	return out
}
