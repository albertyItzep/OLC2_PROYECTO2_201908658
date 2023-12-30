package nt

import compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"

type NT_BloqueContenido struct {
	nSentencias []compiladorp2.CAbstrExpr
}

// Constructor for NT_BloqueContenido
func NewNT_BloqueContenido() *NT_BloqueContenido {
	return &NT_BloqueContenido{
		nSentencias: make([]compiladorp2.CAbstrExpr, 0),
	}
}

func (nt *NT_BloqueContenido) AddSentencia(sentencia compiladorp2.CAbstrExpr) {
	nt.nSentencias = append(nt.nSentencias, sentencia)
}

func (nt *NT_BloqueContenido) Compilar(ctx *compiladorp2.Contexto) *compiladorp2.Atributos {
	for _, sentencia := range nt.nSentencias {
		tmp := sentencia.Compilar(ctx)
		if tmp.Resultado.TipoDato == compiladorp2.Break {
			break
		} else if tmp.Resultado.TipoDato == compiladorp2.Continue {
			continue
		} else if tmp.Resultado.Return {
			return tmp
		}
	}
	return compiladorp2.NewAtributosNil()
}
