package evaluator

import (
	"monkey/ast"
	"monkey/object"
	"monkey/token"
	"strconv"
)

func quote(node ast.Node, env *object.Environment) object.Object {
	node = evalUnquoteCalls(node, env)
	return &object.Quote{Node: node}
}

func evalUnquoteCalls(quoted ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if isUnquoteCall(node) {
			call, _ := node.(*ast.CallExpression)
			unquoted := Eval(call.Arguments[0], env)
			return convertObjectToASTNode(unquoted)
		}
		return node
	})
}

func isUnquoteCall(node ast.Node) bool {
	call, ok := node.(*ast.CallExpression)
	if ok {
		return call.Function.String() == "unquote" && len(call.Arguments) == 1
	}
	return false
}

func convertObjectToASTNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		return &ast.IntegerLiteral{Token: token.Token{Type: token.INT, Literal: strconv.Itoa(int(obj.Value))}, Value: obj.Value}
	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{Type: token.TRUE, Literal: "true"}
		} else {
			t = token.Token{Type: token.FALSE, Literal: "false"}
		}
		return &ast.Boolean{Token: t, Value: obj.Value}
	case *object.Quote:
		return obj.Node
	default:
		return nil
	}
}
