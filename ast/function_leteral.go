package ast

import (
	"bytes"
	"strings"

	"github.com/watariRyo/go-interpreter/token"
)

type FunctionLiteral struct {
	Token     token.Token
	Paramters []*Identifier
	Body      *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Paramters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral() + " ")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}
