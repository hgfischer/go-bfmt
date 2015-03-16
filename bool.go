package bfmt

import (
	"fmt"
	"io"
	"os"
)

// bfmt.Bool is just a bool type, that conditionally mimics the fmt package. If its value is `true`
// bfmt.Bool methods do what they are meant to be done, otherwise, they do nothing.
type Bool bool

// This is the default io.Writes interface used by bfmt.Bool to output things
var BoolWriter io.Writer = os.Stdout

// Conditionally and formatted output to a io.Writer. See fmt.Fprintf. Return (0, nil) if false.
func (b Bool) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprintf(w, format, a...)
	}
	return 0, nil
}

// Conditionally and formatted output. See fmt.Printf.
func (b Bool) Printf(format string, a ...interface{}) {
	if b {
		fmt.Fprintf(BoolWriter, format, a...)
	}
}

// Conditionally and formatted output to a string. See fmt.Sprintf. Return "" if false.
func (b Bool) Sprintf(format string, a ...interface{}) string {
	if b {
		return fmt.Sprintf(format, a...)
	}
	return ""
}

// Conditionally and formatted output to a error. See fmt.Errorf. Return nil if false.
func (b Bool) Errorf(format string, a ...interface{}) error {
	if b {
		return fmt.Errorf(format, a...)
	}
	return nil
}

// Conditionally print output to an io.Writer. See fmt.Fprint. Return (0, nil) if false.
func (b Bool) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprint(w, a...)
	}
	return 0, nil
}

// Conditionally print output. See fmt.Print. Return (0, nil) if false.
func (b Bool) Print(a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprint(BoolWriter, a...)
	}
	return 0, nil
}

// Conditionally print output to a string. See fmt.Sprint. Return "" if false.
func (b Bool) Sprint(a ...interface{}) string {
	if b {
		return fmt.Sprint(a...)
	}
	return ""
}

// Conditionally print output with carriage return to an io.Writer. See fmt.Fprintln. Return (0, nil) if false.
func (b Bool) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprintln(w, a...)
	}
	return 0, nil
}

// Conditionally print output with carriage return. See fmt.Println. Return (0, nil) if false.
func (b Bool) Println(a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprintln(BoolWriter, a...)
	}
	return 0, nil
}

// Conditionally print output with carriage return to a string. See fmt.Sprintln. Return "" if false.
func (b Bool) Sprintln(a ...interface{}) string {
	if b {
		return fmt.Sprintln(a...)
	}
	return ""
}
