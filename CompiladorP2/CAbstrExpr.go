package compiladorp2

type CAbstrExpr interface {
	Compilar(ctx *Contexto) *Atributos
}
