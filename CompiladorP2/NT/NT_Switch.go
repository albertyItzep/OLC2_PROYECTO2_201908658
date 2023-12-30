package nt

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type NT_InstSwitch struct {
	exp         compiladorp2.CAbstrExpr
	InstCase    []NT_InstCase
	InstDefault compiladorp2.CAbstrExpr
	linea       int
	columna     int
}

// Constructor for NT_InstSwitch
func NewNT_InstSwitch(exp compiladorp2.CAbstrExpr, linea int, columna int) *NT_InstSwitch {
	o := new(NT_InstSwitch)
	o.exp = exp
	o.InstCase = make([]NT_InstCase, 0)
	o.InstDefault = nil
	o.linea = linea
	o.columna = columna
	return o
}

func (ntSwitch *NT_InstSwitch) AddCase(cases NT_InstCase) {
	ntSwitch.InstCase = append(ntSwitch.InstCase, cases)
}

func (ntSwitch *NT_InstSwitch) AddDefault(exp compiladorp2.CAbstrExpr) {
	ntSwitch.InstDefault = exp
}

func (nt NT_InstSwitch) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	expr := nt.exp.Compilar(ctx)
	lSalida := ctx.NewEtq()
	for num, val := range nt.InstCase {
		tmpExV := val.exp.Compilar(ctx)
		tEV := ctx.NewEtq()
		tEF := ctx.NewEtq()
		ctx.Gen("if(" + expr.Resultado.Valor + "==" + tmpExV.Resultado.Valor + ") goto " + tEV + ";")
		ctx.Gen("goto " + tEF + ";")
		ctx.Gen(tEV + ":")
		val.Compilar(ctx)
		ctx.Gen("goto " + lSalida + ";")
		ctx.Gen(tEF + ":")
		if len(nt.InstCase)-1 == num {
			nt.InstDefault.Compilar(ctx)
		}
	}
	ctx.Gen(lSalida + ":")
	return compiladorp2.NewAtributosNil()
}

type NT_InstCase struct {
	exp     compiladorp2.CAbstrExpr
	Block   compiladorp2.CAbstrExpr
	linea   int
	columna int
}

// Constructor for NT_InstCase
func NewNT_InstCase(exp compiladorp2.CAbstrExpr, block compiladorp2.CAbstrExpr, linea int, columna int) *NT_InstCase {
	o := new(NT_InstCase)
	o.exp = exp
	o.Block = block
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_InstCase) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	nt.Block.Compilar(ctx)
	return compiladorp2.NewAtributosNil()
}

type NT_InstDefault struct {
	block   compiladorp2.CAbstrExpr
	linea   int
	columna int
}

// Constructor for NT_InstDefault
func NewNT_InstDefault(block compiladorp2.CAbstrExpr, linea int, columna int) *NT_InstDefault {
	o := new(NT_InstDefault)
	o.block = block
	o.linea = linea
	o.columna = columna
	return o
}

func (nt NT_InstDefault) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	nt.block.Compilar(ctx)
	return compiladorp2.NewAtributosNil()
}
