package bfmt

import (
	"fmt"
	"strconv"
)

// The following are for providing flag.Getter compatibility

// Set parse a boolean string value and change the method receiver value
func (b *Bool) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = Bool(v)
	return err
}

// Get returns a bool that have the same value as the method receiver
func (b *Bool) Get() interface{} {
	return bool(*b)
}

// String returns a string representation of the method receiver
func (b *Bool) String() string {
	return fmt.Sprintf("%v", *b)
}

// IsBoolFlag returns true
func (b *Bool) IsBoolFlag() bool {
	return true
}
