package vn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameVNode(t *testing.T) {
	a := H("div", nil, "Hello world")
	b := H("div", nil, "Hello world")
	c := H("span", nil, "Hello world")

	d := H("span", &Attrs{Props: &Props{"hello": "world"}}, "Hello world")
	e := H("span", &Attrs{Props: &Props{"hello": "world"}}, "Hello world")
	f := H("div", &Attrs{Props: &Props{"hello": "world"}}, "Hello worldx")

	g := H("span", &Attrs{}, "Hello world")
	h := H("span", &Attrs{}, "Hello world")

	assert.Equal(t, true, a.isSame(b))
	assert.Equal(t, false, a.isSame(c))
	assert.Equal(t, true, g.isSame(h))
	assert.Equal(t, true, d.isSame(e))
	assert.Equal(t, false, d.isSame(f))
}