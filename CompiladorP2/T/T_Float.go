package t

import (
	compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"
	"strconv"
)

type T_Float struct {
	Num     string
	linea   int
	columna int
}

// Constructor for T_Float
func NewT_Float(num string, linea int, columna int) *T_Float {
	o := new(T_Float)
	o.Num = num
	o.linea = linea
	o.columna = columna
	return o
}

func (t T_Float) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	_, num := strconv.ParseFloat(t.Num, 64)
	if num != nil {
		ctx.PushError("Valor Para Float Incorrecto", t.linea, t.columna)
	}
	return &compiladorp2.Atributos{
		EV:        make([]string, 0, 0),
		EF:        make([]string, 0, 0),
		Resultado: compiladorp2.NewFloatLiteral(t.Num),
	}
}
