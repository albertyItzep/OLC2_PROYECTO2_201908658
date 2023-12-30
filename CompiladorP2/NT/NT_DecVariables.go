package nt

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type NT_Variables struct {
	Id        string
	Tipo      compiladorp2.CAbstrExpr
	expresion compiladorp2.CAbstrExpr
	linea     int
	columna   int
}

// Constructor for NT_Variables
func NewNT_Variables(expresion compiladorp2.CAbstrExpr, linea int, columna int) *NT_Variables {
	o := new(NT_Variables)
	o.expresion = expresion
	o.linea = linea
	o.columna = columna
	return o
}

//nt NT_Variables
