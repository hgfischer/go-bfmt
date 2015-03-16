package bfmt_test

import (
	"testing"

	"github.com/hgfischer/go-bfmt"
	"github.com/stretchr/testify/assert"
)

func TestBoolFlagSetTrue(t *testing.T) {
	b := bfmt.Bool(false)
	err := b.Set("true")
	assert.Nil(t, err)
	assert.True(t, bool(b))
}

func TestBoolFlagSetFalse(t *testing.T) {
	b := bfmt.Bool(true)
	err := b.Set("false")
	assert.Nil(t, err)
	assert.False(t, bool(b))
}

func TestBoolFlagSetErr(t *testing.T) {
	b := bfmt.Bool(true)
	err := b.Set("error")
	assert.NotNil(t, err)
}

func TestBoolFlagGet(t *testing.T) {
	a := bfmt.Bool(true)
	assert.True(t, a.Get().(bool))
	b := bfmt.Bool(false)
	assert.False(t, b.Get().(bool))
}

func TestBoolFlagString(t *testing.T) {
	a := bfmt.Bool(true)
	assert.Equal(t, "true", a.String())
	b := bfmt.Bool(false)
	assert.Equal(t, "false", b.String())
}

func TestBoolFlagIsBoolFlag(t *testing.T) {
	a := bfmt.Bool(true)
	assert.True(t, a.IsBoolFlag())
	b := bfmt.Bool(false)
	assert.True(t, b.IsBoolFlag())
}
