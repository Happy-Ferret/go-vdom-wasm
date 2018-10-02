package vn

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHWithStringChildren(t *testing.T) {
	cases := []struct {
		TagName        string
		childrenString string
		want           *Vnode
	}{
		{"div", "Hello world", &Vnode{TagName: "div", Attrs: nil, Children: &TextNode{Value: "Hello world", Element: nil}, Element: nil}},
		{"span", "Hello", &Vnode{TagName: "span", Attrs: nil, Children: &TextNode{Value: "Hello", Element: nil}, Element: nil}},
	}
	for _, c := range cases {
		got := H(c.TagName, nil, c.childrenString)

		assert.Equal(t, c.want, got)
	}
}

func TestHWithChildren(t *testing.T) {
	childrenString := "Hello Children"

	expectedChild := &Vnode{TagName: "div", Attrs: nil, Children: &TextNode{childrenString, nil}, Element: nil}
	expectedVnode := &Vnode{TagName: "div", Attrs: nil, Children: Children{expectedChild}, Element: nil}

	currentVnode := H("div", nil, Children{H("div", nil, childrenString)})

	assert.Equal(t, expectedVnode, currentVnode)
}

func TestHWithAttributes(t *testing.T) {
	ev := &Ev{"f": func(args []js.Value) {}}
	cases := []struct {
		Attrs *Attrs
		want  *Vnode
	}{
		{&Attrs{}, &Vnode{TagName: "div", Attrs: &Attrs{}, Children: &TextNode{Value: "Hello world", Element: nil}, Element: nil}},
		{&Attrs{Props: &Props{"class": "hello"}}, &Vnode{TagName: "div", Attrs: &Attrs{Props: &Props{"class": "hello"}}, Children: &TextNode{Value: "Hello world", Element: nil}, Element: nil}},
		{&Attrs{Events: ev}, &Vnode{TagName: "div", Attrs: &Attrs{Events: ev}, Children: &TextNode{Value: "Hello world", Element: nil}, Element: nil}},
		{&Attrs{Events: ev, Props: &Props{"class": "hello"}}, &Vnode{TagName: "div", Attrs: &Attrs{Events: ev, Props: &Props{"class": "hello"}}, Children: &TextNode{Value: "Hello world", Element: nil}, Element: nil}},
	}
	for _, c := range cases {
		got := H("div", c.Attrs, "Hello world")

		assert.Equal(t, c.want, got)
	}
}