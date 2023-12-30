package compiladorp2

type Atributos struct {
	EV        []string
	EF        []string
	TipoRes   TipoD
	Resultado *Resultado
}

func NewAtributosNil() *Atributos {
	return &Atributos{
		EV:        make([]string, 0),
		EF:        make([]string, 0),
		Resultado: NewNilLiteral(),
	}
}
