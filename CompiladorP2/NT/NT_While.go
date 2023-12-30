package nt

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type NT_InstWhile struct {
	expr    compiladorp2.CAbstrExpr
	block   compiladorp2.CAbstrExpr
	linea   int
	columna int
}

// Constructor for NT_InstWhile
func NewNT_InstWhile(expr compiladorp2.CAbstrExpr, block compiladorp2.CAbstrExpr, linea int, columna int) *NT_InstWhile {
	o := new(NT_InstWhile)
	o.expr = expr
	o.block = block
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_InstWhile) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {

	LInicio := ctx.NewEtq()
	ctx.Gen(LInicio + ":")
	expr := nt.expr.Compilar(ctx)
	ctx.ImprimirEtiquetas(expr.EV)
	nt.block.Compilar(ctx)
	ctx.Gen("goto " + LInicio + ";")
	ctx.ImprimirEtiquetas(expr.EF)
	return compiladorp2.NewAtributosNil()
}
