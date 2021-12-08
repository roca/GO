package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []*HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", (indent+1)*indentSize))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, e := range e.elements {
		sb.WriteString(e.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     *HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root:     &HtmlElement{name: rootName},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := &HtmlElement{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := &HtmlElement{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	// <ul><li>hello</li><li>world</li></ul>

	sb.WriteString("<ul>")
	for _, word := range words {
		sb.WriteString("<li>")
		sb.WriteString(word)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	b := NewHtmlBuilder("ul").
		AddChildFluent("li", "hello").
		AddChildFluent("li", "world")

	fmt.Println(b.String())
}
