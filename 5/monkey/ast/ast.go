package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// 基础节点接口
type Node interface {
	String() string
}

// 语句
type Statement interface {
	Node
}

// 表达式
type Expression interface {
	Node
}

// 程序
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// let语句
type LetStatement struct {
	Name  *Identifier // 标识符
	Value Expression  // 右侧表达式
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString("let ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	out.WriteString(ls.Value.String())
	out.WriteString(";")
	out.WriteString("\n")

	return out.String()
}

// return语句
type ReturnStatement struct {
	ReturnValue Expression //return右边表达式
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString("return ")
	out.WriteString(rs.ReturnValue.String())
	out.WriteString(";")
	out.WriteString("\n")

	return out.String()
}

// expression语句
type ExpressionStatement struct {
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	return es.Expression.String() + ";" + "\n"
}

// block语句
type BlockStatement struct {
	Statements []Statement
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	out.WriteString("\n")
	for _, s := range bs.Statements {
		out.WriteString("\t" + s.String())
	}
	out.WriteString("}")

	return out.String()
}

// 标识符
type Identifier struct {
	Token token.Token // 词法单元
}

func (i *Identifier) String() string { return i.Token.Literal }

// 布尔字面量
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) String() string { return b.Token.Literal }

// 整数字面量
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) String() string { return il.Token.Literal }

// 前缀表达式
type PrefixExpression struct {
	Token token.Token // The prefix token, e.g. !
	Right Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Token.Literal)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// 中缀表达式
type InfixExpression struct {
	Token token.Token // The operator token, e.g. +
	Left  Expression
	Right Expression
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Token.Literal + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// if表达式
type IfExpression struct {
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// 函数字面量
type FunctionLiteral struct {
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// 调用表达式
type CallExpression struct {
	Function  Expression // 标识符或函数字面量
	Arguments []Expression
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// 字符串字面量
type StringLiteral struct {
	Token token.Token
}

func (sl *StringLiteral) String() string { return sl.Token.Literal }

// 数组字面量
type ArrayLiteral struct {
	Elements []Expression
}

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// 索引表达式
type IndexExpression struct {
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

// 哈希字面量
type HashLiteral struct {
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+" : "+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// 宏字面量
type MacroLiteral struct {
	Parameters []*Identifier
	Body       *BlockStatement
}

func (ml *MacroLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range ml.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("macro")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(ml.Body.String())

	return out.String()
}
