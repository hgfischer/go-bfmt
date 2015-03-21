package bfmt_test

import (
	"os"
	"testing"

	"github.com/hgfischer/go-bfmt"
	"github.com/stretchr/testify/assert"
)

var (
	trueBool  = bfmt.Bool(true)
	falseBool = bfmt.Bool(false)
	expected  = "test"
)

const (
	ShouldHavePrinted    = "Should have printed the expected string with a true Bool"
	ShouldNotHavePrinted = "Nothing should be printed with a false Bool"
)

func TestTrueBoolFprintf(t *testing.T) {
	mock := &IOWriterMock{Expected: expected, T: t}
	trueBool.Fprintf(mock, "%s", expected)
	assert.True(t, mock.Called, ShouldHavePrinted)
}

func TestFalseBoolFprintf(t *testing.T) {
	mock := &IOWriterMock{T: t}
	falseBool.Fprintf(mock, "%s", expected)
	assert.False(t, mock.Called, ShouldNotHavePrinted)
}

func TestTrueBoolPrintf(t *testing.T) {
	mock := &IOWriterMock{Expected: expected, T: t}
	bfmt.BoolWriter = mock
	trueBool.Printf("%s", expected)
	assert.True(t, mock.Called, ShouldHavePrinted)
	bfmt.BoolWriter = os.Stdout
}

func TestFalseBoolPrintf(t *testing.T) {
	mock := &IOWriterMock{T: t}
	bfmt.BoolWriter = mock
	falseBool.Printf("%s", expected)
	assert.False(t, mock.Called, ShouldNotHavePrinted)
	bfmt.BoolWriter = os.Stdout
}

func TestTrueBoolSprintf(t *testing.T) {
	res := trueBool.Sprintf("%s", expected)
	assert.Equal(t, expected, res)
}

func TestFalseBoolSprintf(t *testing.T) {
	res := falseBool.Sprintf("%s", "empty")
	assert.Equal(t, "", res)
}

func TestTrueBoolErrorf(t *testing.T) {
	err := trueBool.Errorf("%s", expected)
	assert.Equal(t, expected, err.Error())
}

func TestFalseBoolErrorf(t *testing.T) {
	err := falseBool.Errorf("%s", "should be nil")
	assert.Nil(t, err)
}

func TestTrueBoolFprint(t *testing.T) {
	mock := &IOWriterMock{Expected: expected, T: t}
	trueBool.Fprint(mock, expected)
	assert.True(t, mock.Called, ShouldHavePrinted)
}

func TestFalseBoolFprint(t *testing.T) {
	mock := &IOWriterMock{T: t}
	falseBool.Fprint(mock, expected)
	assert.False(t, mock.Called, ShouldNotHavePrinted)
}

func TestTrueBoolPrint(t *testing.T) {
	mock := &IOWriterMock{Expected: expected, T: t}
	bfmt.BoolWriter = mock
	trueBool.Print(expected)
	assert.True(t, mock.Called, ShouldHavePrinted)
	bfmt.BoolWriter = os.Stdout
}

func TestFalseBoolPrint(t *testing.T) {
	mock := &IOWriterMock{T: t}
	bfmt.BoolWriter = mock
	falseBool.Print(expected)
	assert.False(t, mock.Called, ShouldNotHavePrinted)
	bfmt.BoolWriter = os.Stdout
}

func TestTrueBoolSprint(t *testing.T) {
	res := trueBool.Sprint(expected)
	assert.Equal(t, expected, res)
}

func TestFalseBoolSprint(t *testing.T) {
	res := falseBool.Sprint("empty")
	assert.Equal(t, "", res)
}

func TestTrueBoolFprintln(t *testing.T) {
	mock := &IOWriterMock{Expected: expected + "\n", T: t}
	trueBool.Fprintln(mock, expected)
	assert.True(t, mock.Called, ShouldHavePrinted)
}

func TestFalseBoolFprintln(t *testing.T) {
	mock := &IOWriterMock{T: t}
	falseBool.Fprintln(mock, expected)
	assert.False(t, mock.Called, ShouldNotHavePrinted)
}

func TestTrueBoolPrintln(t *testing.T) {
	mock := &IOWriterMock{Expected: expected + "\n", T: t}
	bfmt.BoolWriter = mock
	trueBool.Println(expected)
	assert.True(t, mock.Called, ShouldHavePrinted)
	bfmt.BoolWriter = os.Stdout
}

func TestFalseBoolPrintln(t *testing.T) {
	mock := &IOWriterMock{T: t}
	bfmt.BoolWriter = mock
	falseBool.Println(expected)
	assert.False(t, mock.Called, ShouldNotHavePrinted)
	bfmt.BoolWriter = os.Stdout
}

func TestTrueBoolSprintln(t *testing.T) {
	res := trueBool.Sprintln(expected)
	assert.Equal(t, expected+"\n", res)
}

func TestFalseBoolSprintln(t *testing.T) {
	res := falseBool.Sprintln("empty")
	assert.Equal(t, "", res)
}
