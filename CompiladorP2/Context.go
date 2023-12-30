package compiladorp2

import (
	"fmt"
	"strconv"
)

type Contexto struct {
	tmp       int
	lb        int
	lbFin     []int
	display   []string
	Errors    []ErrorCompiler
	CodGen    []string
	Conversor *Conversor
}

func NewContexto() *Contexto {
	c := &Contexto{
		tmp:       1,
		lb:        1,
		lbFin:     make([]int, 0),
		display:   make([]string, 0),
		Errors:    make([]ErrorCompiler, 0),
		CodGen:    make([]string, 0),
		Conversor: nil,
	}
	c.Conversor = NewConversor(c)
	return c
}

func (ctx *Contexto) PushError(desc string, linea int, columna int) {
	ctx.Errors = append(ctx.Errors, *NewErrorCompiler(desc, linea, columna))
}

// agregar etiqueta al display
func (ctx *Contexto) PushDisplay() string {
	nuevaEtiqueta := ctx.NewEtq()
	ctx.display = append(ctx.display, nuevaEtiqueta)
	return nuevaEtiqueta
}

// ver display
func (ctx *Contexto) PeekDislay() string {
	if len(ctx.display) >= 1 {
		return ctx.display[len(ctx.display)-1]
	} else {
		panic("error display vacío")
	}
}

// eliminar de display
func (ctx *Contexto) PopDisplay() string {
	etiqueta := ctx.PeekDislay()
	if len(ctx.display) > 0 {
		ctx.display = ctx.display[:len(ctx.display)-1]
	}
	return etiqueta
}

// generar salida
func (ctx *Contexto) Gen(out string) {
	ctx.CodGen = append(ctx.CodGen, out)
	fmt.Println(out)
}

// crear nuevo temporal
func (ctx *Contexto) NewTemp() string {
	ctx.tmp = ctx.tmp + 1
	return "t" + strconv.Itoa(ctx.tmp)
}

// generar nueva etiqueta
func (ctx *Contexto) NewEtq() string {
	ctx.lb = ctx.lb + 1
	return "L" + strconv.Itoa(ctx.lb)
}

// imprimir etiqueta
func (ctx *Contexto) ImprimirEtiquetas(etiquetas []string) {
	for _, etiqueta := range etiquetas {
		ctx.Gen(etiqueta + ":")
	}
}

// unir etiquetas
func (ctx *Contexto) Unir(etq1 []string, etq2 []string) []string {
	for _, etq := range etq2 {
		etq1 = append(etq1, etq)
	}
	return etq1
}
