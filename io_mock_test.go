package bfmt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type IOWriterMock struct {
	Called   bool
	Expected string
	T        *testing.T
}

func (m *IOWriterMock) Write(p []byte) (n int, err error) {
	m.Called = true
	assert.Equal(m.T, string(p), m.Expected)
	return len(p), nil
}
