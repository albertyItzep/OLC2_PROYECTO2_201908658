package nt

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type NT_Guard struct {
	expr    compiladorp2.CAbstrExpr
	bloque  compiladorp2.CAbstrExpr
	linea   int
	columna int
}

// Constructor for NT_Guard
func NewNT_Guard(expr compiladorp2.CAbstrExpr, bloque compiladorp2.CAbstrExpr, linea int, columna int) *NT_Guard {
	o := new(NT_Guard)
	o.expr = expr
	o.bloque = bloque
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_Guard) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	ctx.Gen("//C3d If")
	expr := nt.expr.Compilar(ctx)

	ctx.ImprimirEtiquetas(expr.EF)
	nt.bloque.Compilar(ctx)
	tmpSalida := ctx.NewEtq()
	ctx.Gen("goto " + tmpSalida + ";")
	ctx.ImprimirEtiquetas(expr.EV)
	ctx.Gen("//Valores falsos")
	ctx.Gen(tmpSalida + ":")
	return &compiladorp2.Atributos{
		EV:        make([]string, 0),
		EF:        make([]string, 0),
		TipoRes:   compiladorp2.Nil,
		Resultado: compiladorp2.NewNilLiteral(),
	}
}
