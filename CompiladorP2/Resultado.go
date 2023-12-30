package compiladorp2

type TipoD uint

const (
	Nil       TipoD = 0
	Character TipoD = 1
	Bool      TipoD = 2
	Integer   TipoD = 3
	Float     TipoD = 4
	String    TipoD = 5
	Break     TipoD = 6
	Continue  TipoD = 7
	Rango     TipoD = 8
	Vector    TipoD = 9
	Compuesto TipoD = 10
)

func (t TipoD) String() string {
	switch t {
	case Nil:
		return "NIL"
	case Character:
		return "Character"
	case Integer:
		return "Integer"
	case Float:
		return "Float"
	case String:
		return "String"
	case Bool:
		return "Bool"
	case Break:
		return "Break"
	case Continue:
		return "Continue"
	case Rango:
		return "Rango"
	case Compuesto:
		return "Compuesto"
	default:
		return "Desconocido"
	}
}

type Resultado struct {
	Valor      string
	TipoDato   TipoD
	TipoVector TipoD
	Rango      []int
	Vector     []Resultado
	Return     bool
}

// Constructor for Resultado vacio no nil
func NewResultado() *Resultado {
	o := new(Resultado)
	return o
}

func NewNilLiteral() *Resultado {
	return &Resultado{
		Valor:      "",
		TipoDato:   Nil,
		TipoVector: Nil,
		Rango:      make([]int, 0),
		Vector:     make([]Resultado, 0),
		Return:     false,
	}
}

func NewIntLiteral(valor string) *Resultado {
	return &Resultado{
		Valor:      valor,
		TipoDato:   Integer,
		TipoVector: Nil,
		Rango:      make([]int, 0),
		Vector:     make([]Resultado, 0),
		Return:     false,
	}
}

func NewFloatLiteral(valor string) *Resultado {
	return &Resultado{
		Valor:      valor,
		TipoDato:   Float,
		TipoVector: Nil,
		Rango:      make([]int, 0),
		Vector:     make([]Resultado, 0),
		Return:     false,
	}
}
func NewStringLiteral(valor string) *Resultado {
	return &Resultado{
		Valor:      valor,
		TipoDato:   String,
		TipoVector: Nil,
		Rango:      make([]int, 0),
		Vector:     make([]Resultado, 0),
		Return:     false,
	}
}
func NewCharLiteral(valor string) *Resultado {
	return &Resultado{
		Valor:      valor,
		TipoDato:   Character,
		TipoVector: Nil,
		Rango:      make([]int, 0),
		Vector:     make([]Resultado, 0),
		Return:     false,
	}
}
func NewBoolLiteral(valor string) *Resultado {
	return &Resultado{
		Valor:      valor,
		TipoDato:   Bool,
		TipoVector: Nil,
		Rango:      make([]int, 0),
		Vector:     make([]Resultado, 0),
		Return:     false,
	}
}
