package nt

import (
	compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"
)

type NT_If struct {
	expr        compiladorp2.CAbstrExpr
	bloque      compiladorp2.CAbstrExpr
	bloqueFalso compiladorp2.CAbstrExpr
	linea       int
	columna     int
}

// Constructor for NT_If
func NewNT_If(expr compiladorp2.CAbstrExpr, bloque compiladorp2.CAbstrExpr, bloqueFalso compiladorp2.CAbstrExpr, linea int, columna int) *NT_If {
	o := new(NT_If)
	o.expr = expr
	o.bloque = bloque
	o.bloqueFalso = bloqueFalso
	o.linea = linea
	o.columna = columna
	return o
}

func (nt *NT_If) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	ctx.Gen("//C3d If")
	expr := nt.expr.Compilar(ctx)
	ctx.ImprimirEtiquetas(expr.EV)
	nt.bloque.Compilar(ctx)
	tmpSalida := ctx.NewEtq()
	ctx.Gen("goto " + tmpSalida + ";")
	if nt.bloqueFalso != nil {
		nt.bloqueFalso.Compilar(ctx)
	}
	ctx.ImprimirEtiquetas(expr.EF)
	ctx.Gen("//Valores falsos")
	ctx.Gen(tmpSalida + ":")

	return &compiladorp2.Atributos{
		EV:        make([]string, 0),
		EF:        make([]string, 0),
		TipoRes:   compiladorp2.Nil,
		Resultado: compiladorp2.NewNilLiteral(),
	}
}

/*
	tmpSalida := ctx.NewEtq()

	tmpSize := len(expr.EV)
	if tmpSize <= 0 {
		ctx.PushError("IF: Error No existe salto verdadero", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	}

	ctx.Gen(expr.EV[tmpSize-1] + ":")
	ctx.Gen("//C3d Bloque if verdadero")
	nt.bloque.Compilar(ctx)
	ctx.Gen("goto " + tmpSalida + ": s")
	ctx.ImprimirEtiquetas(expr.EF)
	ctx.Gen("//C3d Bloque if falso")
	if nt.bloqueFalso != nil {
		expr.EF = append(expr.EF, tmpSalida)
		nt.bloqueFalso.Compilar(ctx)
	} else {
		ctx.Gen(tmpSalida + ":")
	}
*/
