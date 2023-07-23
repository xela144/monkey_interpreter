package evaluator

import (
    "monkey/ast"
    "monkey/object"
)


var (
    TRUE = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
)


func Eval(node ast.Node) object.Object {
    switch node := node.(type) {

    // Statements
    case *ast.Program:
        return evalStatements(node.Statements)

    case *ast.ExpressionStatement:
        return Eval(node.Expression)

    // Experssions
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}

    case *ast.Boolean:
        return nativeBoolToBooleanObject(node.Value)
    }

    return nil
}


func evalStatements(stmts []ast.Statement) object.Object {
    var result object.Object
    for _, statement := range stmts {
        result = Eval(statement)
    }


    return result
}


func nativeBoolToBooleanObject(value bool) *object.Boolean {
    if value {
        return TRUE
    }

    return FALSE
}
