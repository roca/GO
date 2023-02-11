package main

import (
	"fmt"
	"strings"
)

// Strategy Design Pattern

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(buffer *strings.Builder)
	End(buffer *strings.Builder)
	AddListItem(buffer *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(buffer *strings.Builder) {
}

func (m *MarkdownListStrategy) End(buffer *strings.Builder) {
}

func (m *MarkdownListStrategy) AddListItem(buffer *strings.Builder, item string) {
	buffer.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct{}

func (h *HtmlListStrategy) Start(buffer *strings.Builder) {
	buffer.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(buffer *strings.Builder) {
	buffer.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(buffer *strings.Builder, item string) {
	buffer.WriteString("  <li>" + item + "</li>\n")
}

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{listStrategy: listStrategy}
}

func (t *TextProcessor) SetOutputFormat(format OutputFormat) {
	switch format {
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}
	case Html:
		t.listStrategy = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	t.listStrategy.Start(&t.builder)
	for _, item := range items {
		t.listStrategy.AddListItem(&t.builder, item)
	}
	t.listStrategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp.String())
	tp.Reset()

	tp.SetOutputFormat(Html)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp.String())
	tp.Reset()
}
