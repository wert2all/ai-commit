package project

import (
	"bytes"

	"github.com/wert2all/ai-commit/changes"
)

type (
	ContextBuilder interface {
		AddChanges() ContextBuilder
		Build() (string, error)
	}
	contextBuilderImpl struct {
		errors  []error
		changes changes.Changes
	}
)

// AddChanges implements ContextBuilder.
func (c contextBuilderImpl) AddChanges() ContextBuilder {
	changes, err := changes.NewChanges()
	if err != nil {
		c.errors = append(c.errors, err)
	}
	c.changes = changes
	return c
}

// Build implements ContextBuilder.
func (c contextBuilderImpl) Build() (string, error) {
	if len(c.errors) != 0 {
		return "", c.errors[0]
	}
	var context bytes.Buffer

	if c.changes != nil {
		context.WriteString("=== Changes ===\n")
		context.Write(c.changes.Value())
	}

	return context.String(), nil
}

func NewBuilder() ContextBuilder {
	return contextBuilderImpl{errors: make([]error, 0)}
}
