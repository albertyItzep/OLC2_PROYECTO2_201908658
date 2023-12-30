package t

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type T_TipoVariable struct {
	Tipo string
}

// Constructor for T_TipoVariable
func NewT_TipoVariable(tipo string) *T_TipoVariable {
	o := new(T_TipoVariable)
	o.Tipo = tipo
	return o
}

func (tTipo *T_TipoVariable) GetTipoD() compiladorp2.TipoD {
	switch tTipo.Tipo {
	case "String":
		return compiladorp2.String
	case "Float":
		return compiladorp2.Float
	case "Int":
		return compiladorp2.Integer
	case "Bool":
		return compiladorp2.Bool
	case "Character":
		return compiladorp2.Character
	default:
		return compiladorp2.Nil
	}
}
