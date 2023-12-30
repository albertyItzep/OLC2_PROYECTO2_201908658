package nt

import (
	compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"
)

type NT_MultDiv struct {
	ExprIzq compiladorp2.CAbstrExpr
	ExprDer compiladorp2.CAbstrExpr
	oper    string
	linea   int
	columna int
}

// Constructor for NT_Mult
func NewNT_MultDiv(xprIzq compiladorp2.CAbstrExpr, xprDer compiladorp2.CAbstrExpr, oper string, linea int, columna int) *NT_MultDiv {
	o := new(NT_MultDiv)
	o.ExprIzq = xprIzq
	o.ExprDer = xprDer
	o.oper = oper
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_MultDiv) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	expIzq := nt.ExprIzq.Compilar(ctx)
	expDer := nt.ExprDer.Compilar(ctx)

	if expIzq.Resultado.TipoDato != expDer.Resultado.TipoDato {
		if expIzq.Resultado.TipoDato > expDer.Resultado.TipoDato {
			expDer.Resultado = ctx.Conversor.Ampliar(expDer.Resultado, expIzq.Resultado.TipoDato, nt.linea, nt.columna)
		} else {
			expIzq.Resultado = ctx.Conversor.Ampliar(expIzq.Resultado, expDer.Resultado.TipoDato, nt.linea, nt.columna)
		}
	}

	if expDer.Resultado.TipoDato == compiladorp2.Nil || expIzq.Resultado.TipoDato == compiladorp2.Nil {
		ctx.PushError("Suma: Valor NIL no compatible", nt.linea, nt.columna)
	}

	switch expIzq.Resultado.TipoDato {
	case compiladorp2.Nil:
		ctx.PushError("Multiplicacion: No se pueden multiplicar expresiones de tipo Nil", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Bool:
		ctx.PushError("Multiplicacion: No se pueden multiplicar expresiones de tipo Bool", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Integer:
		ctx.Gen("Operacion multiplicacion Integers")
		tmp := ctx.NewTemp()
		ctx.Gen(tmp + " = " + expIzq.Resultado.Valor + " " + nt.oper + " " + expDer.Resultado.Valor)
		return &compiladorp2.Atributos{
			EV:        make([]string, 0, 0),
			EF:        make([]string, 0, 0),
			Resultado: compiladorp2.NewIntLiteral(tmp),
		}
	case compiladorp2.Float:
		ctx.Gen("Operacion multiplicacion Floats")
		tmp := ctx.NewTemp()
		ctx.Gen(tmp + " = " + expIzq.Resultado.Valor + " " + nt.oper + " " + expDer.Resultado.Valor)
		return &compiladorp2.Atributos{
			EV:        make([]string, 0, 0),
			EF:        make([]string, 0, 0),
			Resultado: compiladorp2.NewFloatLiteral(tmp),
		}
	case compiladorp2.String:
		ctx.PushError("Multiplicacion: No se pueden multiplicar expresiones de tipo String", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Character:
		ctx.PushError("Multiplicacion: No se pueden multiplicar expresiones de tipo Character", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	default:
		ctx.PushError("Multiplicacion: Error de tipos desconocidos", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	}
}
