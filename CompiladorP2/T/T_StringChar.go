package t

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type T_String struct {
	val     string
	linea   int
	columna int
}

// Constructor for T_StringChar
func NewT_String(val string, linea int, columna int) *T_String {
	o := new(T_String)
	o.val = val
	o.linea = linea
	o.columna = columna
	return o
}

func (t T_String) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	return &compiladorp2.Atributos{
		EV:        make([]string, 0, 0),
		EF:        make([]string, 0, 0),
		Resultado: compiladorp2.NewStringLiteral(t.val),
	}
}

type T_Char struct {
	val     string
	linea   int
	columna int
}

// Constructor for T_StringChar
func NewT_Char(val string, linea int, columna int) *T_Char {
	o := new(T_Char)
	o.val = val
	o.linea = linea
	o.columna = columna
	return o
}

func (t T_Char) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	return &compiladorp2.Atributos{
		EV:        make([]string, 0, 0),
		EF:        make([]string, 0, 0),
		Resultado: compiladorp2.NewCharLiteral(t.val),
	}
}
