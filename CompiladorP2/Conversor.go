package compiladorp2

type Conversor struct {
	ctx *Contexto
}

func NewConversor(ctx *Contexto) *Conversor {
	return &Conversor{
		ctx: ctx,
	}
}

func (c *Conversor) Ampliar(res *Resultado, tipo TipoD, linea int, columna int) *Resultado {
	switch tipo {
	case Nil:
		c.ctx.PushError("No se puede realizar conversiones a Nil", linea, columna)
	case Bool:

	case Integer:
		switch res.TipoDato {
		case Nil:
			c.ctx.PushError("No se puede realizar conversiones a Nil", linea, columna)
			return NewNilLiteral()
		case Bool:
			c.ctx.PushError("No se puede realizar conversion de Bool a Int", linea, columna)
			return NewNilLiteral()
		case Integer:
			return res
		case String:
			res.Valor = "(int)" + res.Valor
			res.TipoDato = Integer
			return res
		case Character:
			c.ctx.PushError("No se puede realizar la conversion de Character a Integer", linea, columna)
			return NewNilLiteral()
		case Float:
			res.Valor = "(int)" + res.Valor
			res.TipoDato = Integer
			return res
		}
	case String:
		switch res.TipoDato {
		case Nil:
			c.ctx.PushError("No se puede realizar la conversion de Nil a String", linea, columna)
			return NewNilLiteral()
		case Bool:
			//gen code
		case Integer:
			//gen code
		case String:
			return res
		case Character:
		case Float:
			//gen code
		}
	case Character:
	case Float:
		switch res.TipoDato {
		case Nil:
			c.ctx.PushError("No se puede realizar la conversion de Nil a Float", linea, columna)
			return NewNilLiteral()
		case Bool:
			c.ctx.PushError("No se puede realizar la conversion de bool a Float", linea, columna)
			return NewNilLiteral()
		case Integer:
			res.Valor = "(float)" + res.Valor
			res.TipoDato = Float
			return res
		case String:
			res.Valor = "(float)" + res.Valor
			res.TipoDato = Float
			return res
		case Character:
			c.ctx.PushError("No se puede realizar la conversion de character a Float", linea, columna)
			return NewNilLiteral()
		case Float:
			return res
		}
	}
	c.ctx.PushError("Conversion ilegal"+
		tipo.String()+" y"+res.TipoDato.String(), linea, columna)
	return NewNilLiteral()
}
