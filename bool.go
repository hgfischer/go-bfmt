package bfmt

import (
	"fmt"
	"io"
	"os"
)

type Bool bool

var BoolWriter io.Writer = os.Stdout

func (b Bool) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprintf(w, format, a...)
	}
	return 0, nil
}

func (b Bool) Printf(format string, a ...interface{}) {
	if b {
		fmt.Fprintf(BoolWriter, format, a...)
	}
}

func (b Bool) Sprintf(format string, a ...interface{}) string {
	if b {
		return fmt.Sprintf(format, a...)
	}
	return ""
}

func (b Bool) Errorf(format string, a ...interface{}) error {
	if b {
		return fmt.Errorf(format, a...)
	}
	return nil
}

func (b Bool) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprint(w, a...)
	}
	return 0, nil
}

func (b Bool) Print(a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprint(BoolWriter, a...)
	}
	return 0, nil
}

func (b Bool) Sprint(a ...interface{}) string {
	if b {
		return fmt.Sprint(a...)
	}
	return ""
}

func (b Bool) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprintln(w, a...)
	}
	return 0, nil
}

func (b Bool) Println(a ...interface{}) (n int, err error) {
	if b {
		return fmt.Fprintln(BoolWriter, a...)
	}
	return 0, nil
}

func (b Bool) Sprintln(a ...interface{}) string {
	if b {
		return fmt.Sprintln(a...)
	}
	return ""
}
