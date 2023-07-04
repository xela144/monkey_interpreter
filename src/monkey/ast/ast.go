package ast

import "monkey/token"


type Node interface {
    TokenLiteral() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

// Root node of every AST that the parser produces:
type Program struct {
    Statements []Statement
}


func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}


type Identifier struct{
    Token token.Token  // token.IDENT
    Value string
}

type LetStatement struct {
    // e.g.: "let x = 5;
    Token token.Token  // let
    Name *Identifier   // x
    Value Expression   // x = 5;
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }



