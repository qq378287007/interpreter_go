package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Node interface { // 节点
	String() string
}

type Statement interface { // 语句
	Node
}

type Expression interface { // 表达式
	Node
}

type Program struct { // 程序
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct { // let语句
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

type ReturnStatement struct { // return语句
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

type ExpressionStatement struct { // expression语句
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	return es.Expression.String() + ";" + "\n"
}

type BlockStatement struct { // block语句
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

type Identifier struct { // 标识符
	Token token.Token // 词法单元
}

func (i *Identifier) String() string { return i.Token.Literal }

type Boolean struct { // 布尔字面量
	Token token.Token
	Value bool
}

func (b *Boolean) String() string { return b.Token.Literal }

type IntegerLiteral struct { // 整数字面量
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) String() string { return il.Token.Literal }

type PrefixExpression struct { // 前缀表达式
	Token token.Token
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

type InfixExpression struct { // 中缀表达式
	Token token.Token
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

type IfExpression struct { // if表达式
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

type FunctionLiteral struct { // 函数字面量
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

type CallExpression struct { // 调用表达式
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

type StringLiteral struct { // 字符串字面量
	Token token.Token
}

func (sl *StringLiteral) String() string { return sl.Token.Literal }

type ArrayLiteral struct { // 数组字面量
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

type IndexExpression struct { // 索引表达式
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

type HashLiteral struct { // 哈希字面量
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
