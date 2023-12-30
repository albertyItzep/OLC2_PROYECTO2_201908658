package nt

import (
	compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"
)

type NT_Sum struct {
	ExprIzq compiladorp2.CAbstrExpr
	ExprDer compiladorp2.CAbstrExpr
	oper    string
	linea   int
	columna int
}

// Constructor for NT_SumRes
func NewNT_Sum(exprIzq compiladorp2.CAbstrExpr, exprDer compiladorp2.CAbstrExpr, oper string, linea int, columna int) *NT_Sum {
	o := new(NT_Sum)
	o.ExprIzq = exprIzq
	o.ExprDer = exprDer
	o.oper = oper
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_Sum) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
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
		ctx.PushError("Suma: No se pueden sumar expresiones de tipo Nil", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Bool:
		ctx.PushError("Suma: No se pueden sumar expresiones de tipo Bool", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Integer:
		ctx.Gen("Operacion Suma Integers")
		tmp := ctx.NewTemp()
		ctx.Gen(tmp + " = " + expIzq.Resultado.Valor + " " + nt.oper + " " + expDer.Resultado.Valor)
		return &compiladorp2.Atributos{
			EV:        make([]string, 0, 0),
			EF:        make([]string, 0, 0),
			Resultado: compiladorp2.NewIntLiteral(tmp),
		}
	case compiladorp2.String:
		ctx.Gen("Operacion Suma Strings")
		tmp := ctx.NewTemp()
		ctx.Gen(tmp + " = " + expIzq.Resultado.Valor + " " + nt.oper + " " + expDer.Resultado.Valor)
		return &compiladorp2.Atributos{
			EV:        make([]string, 0, 0),
			EF:        make([]string, 0, 0),
			Resultado: compiladorp2.NewIntLiteral(tmp),
		}
	case compiladorp2.Character:
		ctx.PushError("Suma: No se pueden sumar expresiones de tipo Character", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Float:
		ctx.Gen("Operacion Suma Float")
		tmp := ctx.NewTemp()
		ctx.Gen(tmp + " = " + expIzq.Resultado.Valor + " " + nt.oper + " " + expDer.Resultado.Valor)
		return &compiladorp2.Atributos{
			EV:        make([]string, 0, 0),
			EF:        make([]string, 0, 0),
			Resultado: compiladorp2.NewFloatLiteral(tmp),
		}
	default:
		ctx.PushError("Suma: Error Tipos desconocidos", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	}

}

type NT_Res struct {
	ExprIzq compiladorp2.CAbstrExpr
	ExprDer compiladorp2.CAbstrExpr
	oper    string
	linea   int
	columna int
}

// Constructor for NT_Res
func NewNT_Res(xprIzq compiladorp2.CAbstrExpr, xprDer compiladorp2.CAbstrExpr, oper string, linea int, columna int) *NT_Res {
	o := new(NT_Res)
	o.ExprIzq = xprIzq
	o.ExprDer = xprDer
	o.oper = oper
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_Res) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
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
		ctx.PushError("Resta: Valor NIL no compatible", nt.linea, nt.columna)
	}

	switch expIzq.Resultado.TipoDato {
	case compiladorp2.Nil:
		ctx.PushError("Resta: No se pueden restar expresiones de tipo Nil", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Bool:
		ctx.PushError("Resta: No se pueden restar expresiones de tipo Bool", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Integer:
		ctx.Gen("Operacion Resta Integers")
		tmp := ctx.NewTemp()
		ctx.Gen(tmp + " = " + expIzq.Resultado.Valor + " " + nt.oper + " " + expDer.Resultado.Valor)
		return &compiladorp2.Atributos{
			EV:        make([]string, 0, 0),
			EF:        make([]string, 0, 0),
			Resultado: compiladorp2.NewIntLiteral(tmp),
		}
	case compiladorp2.String:
		ctx.PushError("Resta: No se pueden restar expresiones de tipo string", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Character:
		ctx.PushError("Resta: No se pueden restar expresiones de tipo Character", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Float:
		ctx.Gen("Operacion Resta Float")
		tmp := ctx.NewTemp()
		ctx.Gen(tmp + " = " + expIzq.Resultado.Valor + " " + nt.oper + " " + expDer.Resultado.Valor)
		return &compiladorp2.Atributos{
			EV:        make([]string, 0, 0),
			EF:        make([]string, 0, 0),
			Resultado: compiladorp2.NewFloatLiteral(tmp),
		}
	default:
		ctx.PushError("Resta: Error Tipos desconocidos", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	}
}

type NT_Modulo struct {
	ExprIzq compiladorp2.CAbstrExpr
	ExprDer compiladorp2.CAbstrExpr
	oper    string
	linea   int
	columna int
}

// Constructor for NT_Modulo
func NewNT_Modulo(xprIzq compiladorp2.CAbstrExpr, xprDer compiladorp2.CAbstrExpr, oper string, linea int, columna int) *NT_Modulo {
	o := new(NT_Modulo)
	o.ExprIzq = xprIzq
	o.ExprDer = xprDer
	o.oper = oper
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_Modulo) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
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
		ctx.PushError("Modulo: Valor NIL no compatible", nt.linea, nt.columna)
	}

	switch expIzq.Resultado.TipoDato {
	case compiladorp2.Nil:
		ctx.PushError("Modulo: No se pueden Modulo expresiones de tipo Nil", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Bool:
		ctx.PushError("Modulo: No se pueden Modulo expresiones de tipo Bool", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Integer:
		ctx.Gen("Operacion Resta Integers")
		tmp := ctx.NewTemp()
		ctx.Gen(tmp + " = " + expIzq.Resultado.Valor + " " + nt.oper + " " + expDer.Resultado.Valor)
		return &compiladorp2.Atributos{
			EV:        make([]string, 0, 0),
			EF:        make([]string, 0, 0),
			Resultado: compiladorp2.NewIntLiteral(tmp),
		}
	case compiladorp2.String:
		ctx.PushError("Modulo: No se pueden Modulo expresiones de tipo string", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Character:
		ctx.PushError("Modulo: No se pueden Modulo expresiones de tipo Character", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	case compiladorp2.Float:
		ctx.PushError("Modulo: No se pueden Modulo expresiones de tipo Float", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	default:
		ctx.PushError("Modulo: Error Tipos desconocidos", nt.linea, nt.columna)
		return compiladorp2.NewAtributosNil()
	}
}

type NT_NegUnaria struct {
	expr    compiladorp2.CAbstrExpr
	linea   int
	columna int
}

// Constructor for NT_NegUnaria
func NewNT_NegUnaria(expr compiladorp2.CAbstrExpr, linea int, columna int) *NT_NegUnaria {
	o := new(NT_NegUnaria)
	o.expr = expr
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_NegUnaria) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	expr := nt.expr.Compilar(ctx)
	expr.Resultado.Valor = "-1(" + expr.Resultado.Valor + ")"
	return &compiladorp2.Atributos{
		EV:        expr.EV,
		EF:        expr.EF,
		TipoRes:   expr.TipoRes,
		Resultado: expr.Resultado,
	}
}
