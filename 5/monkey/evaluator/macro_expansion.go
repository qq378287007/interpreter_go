package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	definitions := []int{}
	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)
			definitions = append(definitions, i)
		}
	}
	for i := len(definitions) - 1; i >= 0; i-- {
		definitionIndex := definitions[i]
		program.Statements = append(program.Statements[:definitionIndex],
			program.Statements[definitionIndex+1:]...)
	}
}

func isMacroDefinition(node ast.Statement) bool {
	letStatement, ok := node.(*ast.LetStatement)
	if ok {
		_, ok = letStatement.Value.(*ast.MacroLiteral)
	}
	return ok
}

func addMacro(stmt ast.Statement, env *object.Environment) {
	letStatement, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStatement.Value.(*ast.MacroLiteral)
	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env:        env,
		Body:       macroLiteral.Body,
	}
	env.Set(letStatement.Name.String(), macro)
}

func ExpandMacros(program ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		if callExpression, ok := node.(*ast.CallExpression); ok {
			if macro, ok := isMacroCall(callExpression, env); ok {
				args := quoteArgs(callExpression)
				evalEnv := extendMacroEnv(macro, args)
				evaluated := Eval(macro.Body, evalEnv)
				if quote, ok := evaluated.(*object.Quote); ok {
					return quote.Node
				}
				panic("we only support returning AST-nodes from macros")
			}
		}
		return node
	})
}

func isMacroCall(exp *ast.CallExpression, env *object.Environment) (*object.Macro, bool) {
	if identifier, ok := exp.Function.(*ast.Identifier); ok {
		if obj, ok := env.Get(identifier.String()); ok {
			if macro, ok := obj.(*object.Macro); ok {
				return macro, true
			}
		}
	}
	return nil, false
}

func quoteArgs(exp *ast.CallExpression) []*object.Quote {
	args := []*object.Quote{}
	for _, a := range exp.Arguments {
		args = append(args, &object.Quote{Node: a})
	}
	return args
}

func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	extended := object.NewEnclosedEnvironment(macro.Env)
	for paramIdx, param := range macro.Parameters {
		extended.Set(param.String(), args[paramIdx])
	}
	return extended
}
