package sn

import (
	"fmt"
	"strings"

	"github.com/kenkyu392/go-sn/rule"
)

var _ fmt.Stringer = new(Generator)

const (
	// DefaultDelimiter is the delimiter character
	// used in the Generator created by NewGenerator.
	DefaultDelimiter = "-"
)

// NewGenerator returns a new generator based on rules.
func NewGenerator(rules ...rule.Rule) *Generator {
	return &Generator{
		Delimiter: DefaultDelimiter,
		Rules:     rules,
	}
}

// Generator has a String function that generates data based on rules.
type Generator struct {
	Delimiter string
	Rules     []rule.Rule
}

// String implements fmt.Stringer interface.
func (g *Generator) String() string {
	var elems = make([]string, len(g.Rules))
	for i, rule := range g.Rules {
		elems[i] = string(rule())
	}
	return strings.Join(elems, g.Delimiter)
}
