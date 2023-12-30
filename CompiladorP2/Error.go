package compiladorp2

type ErrorCompiler struct {
	Description string
	linea       int
	columna     int
}

// Constructor for ErrorCompiler
func NewErrorCompiler(description string, linea int, columna int) *ErrorCompiler {
	o := new(ErrorCompiler)
	o.Description = description
	o.linea = linea
	o.columna = columna
	return o
}
