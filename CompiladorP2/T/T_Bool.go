package t

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type T_Bool struct {
	val     string
	linea   int
	columna int
}

// Constructor for T_Bool
func NewT_Bool(val string, linea int, columna int) *T_Bool {
	o := new(T_Bool)
	o.val = val
	o.linea = linea
	o.columna = columna
	return o
}

func (t T_Bool) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	EV := ctx.NewEtq()
	EF := ctx.NewEtq()

	if t.val == "false" {
		ctx.Gen("goto " + EF)
	} else {
		ctx.Gen("goto " + EV)
	}
	EVs := make([]string, 1, 1)
	EVs = append(EVs, EV)

	EFs := make([]string, 1, 1)
	EFs = append(EFs, EF)

	return &compiladorp2.Atributos{
		EV:        EVs,
		EF:        EFs,
		Resultado: compiladorp2.NewNilLiteral(),
	}
}
