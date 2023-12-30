package nt

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type NT_ProgInicial struct {
	lbloques compiladorp2.CAbstrExpr
}

// Constructor for NT_ProgInicial
func NewNT_ProgInicial(lbloques compiladorp2.CAbstrExpr) *NT_ProgInicial {
	o := new(NT_ProgInicial)
	o.lbloques = lbloques
	return o
}

func (nt NT_ProgInicial) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	atributos := nt.lbloques.Compilar(ctx)
	return atributos
}
