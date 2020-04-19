package logging

type ThingUnderTest struct {
	log *Logger
}

func (this *ThingUnderTest) Action() {
	this.log.Printf("Hello, World!") // NOTE: This statement must reside on line 8
	// of this file. There are tests that assert that the correct call depth is used
	// in order to get the log.Lshortfile output right.
}
