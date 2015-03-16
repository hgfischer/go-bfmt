package bfmt

import (
	"fmt"
	"strconv"
)

// The following are for providing flag.Getter compatibility

func (b *Bool) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = Bool(v)
	return err
}

func (b *Bool) Get() interface{} {
	return bool(*b)
}

func (b *Bool) String() string {
	return fmt.Sprintf("%v", *b)
}

func (b *Bool) IsBoolFlag() bool {
	return true
}
