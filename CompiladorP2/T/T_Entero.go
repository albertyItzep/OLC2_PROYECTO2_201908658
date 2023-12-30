package t

import (
	compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"
)

type T_Entero struct {
	num     string
	linea   int
	columna int
}

// Constructor for T_Entero
func NewT_Entero(num string, linea int, columna int) *T_Entero {
	o := new(T_Entero)
	o.num = num
	o.linea = linea
	o.columna = columna
	return o
}

func (t T_Entero) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	return &compiladorp2.Atributos{
		EV:        make([]string, 0, 0),
		EF:        make([]string, 0, 0),
		Resultado: compiladorp2.NewIntLiteral(t.num),
	}
}
