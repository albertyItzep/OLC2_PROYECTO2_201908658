package nt

import (
	compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"
)

// Estructura que crea el C3d para condiciones sin agrupamiento == != > >= < <=
type NT_Condicional struct {
	ExprIzq compiladorp2.CAbstrExpr
	ExprDer compiladorp2.CAbstrExpr
	oper    string
	linea   int
	columna int
}

// Constructor for NT_Condicional
func NewNT_Condicional(xprIzq compiladorp2.CAbstrExpr, xprDer compiladorp2.CAbstrExpr, oper string, linea int, columna int) *NT_Condicional {
	o := new(NT_Condicional)
	o.ExprIzq = xprIzq
	o.ExprDer = xprDer
	o.oper = oper
	o.linea = linea
	o.columna = columna
	return o
}

func (nt *NT_Condicional) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	exprIzq := nt.ExprIzq.Compilar(ctx)
	exprDer := nt.ExprDer.Compilar(ctx)
	//creamos las direcciones a las que viajara con verdadero o falso
	if exprIzq.Resultado.TipoDato == exprDer.Resultado.TipoDato {
		EV := ctx.NewEtq() //L1
		EF := ctx.NewEtq() //L2
		EVs := make([]string, 0, 0)
		EFs := make([]string, 0, 0)
		EVs = append(EVs, EV)
		EFs = append(EFs, EF)

		ctx.Gen("// C3d Condicional simple: " + nt.oper)
		ctx.Gen("if (" + exprIzq.Resultado.Valor + " " + nt.oper + " " +
			exprDer.Resultado.Valor + ") goto " + EV + ";")
		ctx.Gen("goto " + EF + ";")

		return &compiladorp2.Atributos{
			EV:        EVs,
			EF:        EFs,
			TipoRes:   exprDer.Resultado.TipoDato,
			Resultado: compiladorp2.NewNilLiteral(),
		}
	}
	ctx.PushError("Condicion: Valores de comparacion incompatibles", nt.linea, nt.columna)
	return compiladorp2.NewAtributosNil()
}

//Estructura que valida || && en otras palabras AND y OR

type NT_CondAnd struct {
	ExprIzq compiladorp2.CAbstrExpr
	ExprDer compiladorp2.CAbstrExpr
	textIz  string
	textDer string
	linea   int
	columna int
}

// Constructor for NT_CondAnd
func NewNT_CondAnd(xprIzq compiladorp2.CAbstrExpr, xprDer compiladorp2.CAbstrExpr, linea int, columna int) *NT_CondAnd {
	o := new(NT_CondAnd)
	o.ExprIzq = xprIzq
	o.ExprDer = xprDer
	o.linea = linea
	o.columna = columna
	return o
}

func (nt *NT_CondAnd) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	//and y or son diferentes por la salida del valor
	exprIzq := nt.ExprIzq.Compilar(ctx)
	ctx.ImprimirEtiquetas(exprIzq.EV)
	exprDer := nt.ExprDer.Compilar(ctx)

	if exprIzq.TipoRes != exprDer.TipoRes {
		ctx.PushError("AND: Tipos incompatibles", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	}

	return &compiladorp2.Atributos{
		EV:        exprDer.EV,
		EF:        ctx.Unir(exprIzq.EF, exprDer.EF),
		TipoRes:   exprDer.Resultado.TipoDato,
		Resultado: compiladorp2.NewNilLiteral(),
	}
}

type NT_OR struct {
	ExprIzq compiladorp2.CAbstrExpr
	ExprDer compiladorp2.CAbstrExpr
	textIz  string
	textDer string
	linea   int
	columna int
}

// Constructor for NT_OR
func NewNT_OR(textIz string, textDer string, linea int, columna int) *NT_OR {
	o := new(NT_OR)
	o.textIz = textIz
	o.textDer = textDer
	o.linea = linea
	o.columna = columna
	return o
}

func (nt *NT_OR) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	//and y or son diferentes por la salida del valor
	exprIzq := nt.ExprIzq.Compilar(ctx)
	ctx.ImprimirEtiquetas(exprIzq.EV)
	exprDer := nt.ExprDer.Compilar(ctx)

	if exprIzq.TipoRes != exprDer.TipoRes {
		ctx.PushError("AND: Tipos incompatibles", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	}

	return &compiladorp2.Atributos{
		EV:        ctx.Unir(exprIzq.EV, exprDer.EV),
		EF:        exprDer.EF,
		TipoRes:   exprDer.Resultado.TipoDato,
		Resultado: compiladorp2.NewNilLiteral(),
	}
}

type NT_NegExp struct {
	expOpe  compiladorp2.CAbstrExpr
	linea   int
	columna int
}

// Constructor for NT_NegExp
func NewT_NegExp(expOpe compiladorp2.CAbstrExpr, linea int, columna int) *NT_NegExp {
	o := new(NT_NegExp)
	o.expOpe = expOpe
	o.linea = linea
	o.columna = columna
	return o
}

func (nt *NT_NegExp) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	tmpExp := nt.expOpe.Compilar(ctx)
	return &compiladorp2.Atributos{
		EV:        tmpExp.EF,
		EF:        tmpExp.EV,
		TipoRes:   tmpExp.Resultado.TipoDato,
		Resultado: compiladorp2.NewNilLiteral(),
	}
}
