package template

import "strings"

type IMessageRetriever interface {
	Message() string
}

type ITemplate interface {
	first() string
	third() string
	ExecuteAlgorithm(IMessageRetriever) string
}

type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return "hello"
}
func (t *TemplateImpl) third() string {
	return "template"
}
func (t *TemplateImpl) ExecuteAlgorithm(m IMessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type AnonymousTemplate struct{}
type MessageRetrieverFunc func() string

func (a *AnonymousTemplate) first() string {
	return "hello"
}
func (a *AnonymousTemplate) third() string {
	return "template"
}
func (a *AnonymousTemplate) ExecuteAlgorithm(f MessageRetrieverFunc) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

type TemplateAdapter struct {
	myFunc func() string
}

func (a *TemplateAdapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}
func MessageRetrieverAdater(f MessageRetrieverFunc) IMessageRetriever {
	return &TemplateAdapter{myFunc: f}
}
