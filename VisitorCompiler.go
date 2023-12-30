package main

import (
	"OLC2_PROYECTO2_201908658/Background"
	compiladorp2 "OLC2_PROYECTO2_201908658/CompiladorP2"
	nt "OLC2_PROYECTO2_201908658/CompiladorP2/NT"
	t "OLC2_PROYECTO2_201908658/CompiladorP2/T"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type CompilerVisit struct {
	antlr.ParseTreeVisitor
	Raiz compiladorp2.CAbstrExpr
}

//fmt.Printf("%+v\n", nodoTmp) mostrar un objeto

func NewCompilerVisit() Background.ControlVisitor {
	return &CompilerVisit{
		ParseTreeVisitor: Background.BaseControlVisitor{},
		Raiz:             nil,
	}
}

func (p CompilerVisit) Visit(tree antlr.ParseTree) interface{} {
	switch nodo := tree.(type) {
	case *antlr.ErrorNodeImpl:
		return compiladorp2.NewErrorCompiler(nodo.GetText(), 0, 0)
	default:
		nodoActual, resStatus := tree.Accept(p).(compiladorp2.CAbstrExpr)
		if resStatus {
			return nodoActual
		}
		return compiladorp2.NewErrorCompiler("Nodo desconocido", 0, 0)
	}
}

/*
func (p CompilerVisit) VisitChildren(node antlr.RuleNode) interface{} {
	panic("not implemented") // TODO: Implement
}

func (p CompilerVisit) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("not implemented") // TODO: Implement
}*/

func (p CompilerVisit) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return compiladorp2.NewErrorCompiler("Error desconocido", 0, 0)
}

// Visit a parse tree produced by ControlParser#ProgInicial.
func (p CompilerVisit) VisitProgInicial(ctx *Background.ProgInicialContext) interface{} {
	nodo := ctx.Block().Accept(p).(compiladorp2.CAbstrExpr)
	p.Raiz = nodo
	nodoProgInicial := nt.NewNT_ProgInicial(nodo)
	return nodoProgInicial
}

// Visit a parse tree produced by ControlParser#BlockContent.
func (p CompilerVisit) VisitBlockContent(ctx *Background.BlockContentContext) interface{} {
	blockContent := nt.NewNT_BloqueContenido()
	sentencias := ctx.AllSentencias()
	if len(sentencias) > 0 {
		for _, v := range sentencias {
			nodoTmp := v.Accept(p).(compiladorp2.CAbstrExpr)
			blockContent.AddSentencia(nodoTmp)
		}
	} else {
	}
	return blockContent
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracion.
func (p CompilerVisit) VisitSentenciaDeclaracion(ctx *Background.SentenciaDeclaracionContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracionVector.
func (p CompilerVisit) VisitSentenciaDeclaracionVector(ctx *Background.SentenciaDeclaracionVectorContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracionMatrix.
func (p CompilerVisit) VisitSentenciaDeclaracionMatrix(ctx *Background.SentenciaDeclaracionMatrixContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracionFuncion.
func (p CompilerVisit) VisitSentenciaDeclaracionFuncion(ctx *Background.SentenciaDeclaracionFuncionContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaAsignacion.
func (p CompilerVisit) VisitSentenciaAsignacion(ctx *Background.SentenciaAsignacionContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaIf.
func (p CompilerVisit) VisitSentenciaIf(ctx *Background.SentenciaIfContext) interface{} {
	return ctx.InsIf().Accept(p).(compiladorp2.CAbstrExpr)
}

// Visit a parse tree produced by ControlParser#SentenciaSwitch.
func (p CompilerVisit) VisitSentenciaSwitch(ctx *Background.SentenciaSwitchContext) interface{} {
	return ctx.InsSwitch().Accept(p).(compiladorp2.CAbstrExpr)
}

// Visit a parse tree produced by ControlParser#SentenciaWhile.
func (p CompilerVisit) VisitSentenciaWhile(ctx *Background.SentenciaWhileContext) interface{} {
	return ctx.InstWhile().Accept(p).(compiladorp2.CAbstrExpr)
}

// Visit a parse tree produced by ControlParser#SentenciaFor.
func (p CompilerVisit) VisitSentenciaFor(ctx *Background.SentenciaForContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaGuard.
func (p CompilerVisit) VisitSentenciaGuard(ctx *Background.SentenciaGuardContext) interface{} {
	return ctx.InstGuard().Accept(p).(compiladorp2.CAbstrExpr)
}

// Visit a parse tree produced by ControlParser#SentenciaBreak.
func (p CompilerVisit) VisitSentenciaBreak(ctx *Background.SentenciaBreakContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaContinue.
func (p CompilerVisit) VisitSentenciaContinue(ctx *Background.SentenciaContinueContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaReturn.
func (p CompilerVisit) VisitSentenciaReturn(ctx *Background.SentenciaReturnContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaFuncVectoriales.
func (p CompilerVisit) VisitSentenciaFuncVectoriales(ctx *Background.SentenciaFuncVectorialesContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaLLamadaFuncion.
func (p CompilerVisit) VisitSentenciaLLamadaFuncion(ctx *Background.SentenciaLLamadaFuncionContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaPrint.
func (p CompilerVisit) VisitSentenciaPrint(ctx *Background.SentenciaPrintContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaStruct.
func (p CompilerVisit) VisitSentenciaStruct(ctx *Background.SentenciaStructContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracionStruct.
func (p CompilerVisit) VisitSentenciaDeclaracionStruct(ctx *Background.SentenciaDeclaracionStructContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaLLamadaMetodos.
func (p CompilerVisit) VisitSentenciaLLamadaMetodos(ctx *Background.SentenciaLLamadaMetodosContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#InstruccionIf.
func (p CompilerVisit) VisitInstruccionIf(ctx *Background.InstruccionIfContext) interface{} {
	expr := ctx.Expr().Accept(p).(compiladorp2.CAbstrExpr)
	bloques := ctx.AllBlock()
	var bloque1 compiladorp2.CAbstrExpr
	var bloque2 compiladorp2.CAbstrExpr

	fmt.Println(len(bloques))
	if bloques != nil {
		if len(bloques) == 2 {
			bloque1 = ctx.Block(0).Accept(p).(compiladorp2.CAbstrExpr)
			bloque2 = ctx.Block(1).Accept(p).(compiladorp2.CAbstrExpr)
		} else {
			bloque1 = ctx.Block(0).Accept(p).(compiladorp2.CAbstrExpr)
			fmt.Printf("%+v\n", bloque1)
		}
	}

	if ctx.InsIf() != nil {
		bloque2 = ctx.InsIf().Accept(p).(compiladorp2.CAbstrExpr)
	}
	tmp := nt.NewNT_If(expr, bloque1, bloque2, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return tmp
}

// Visit a parse tree produced by ControlParser#InstruccionSwitch.
func (p CompilerVisit) VisitInstruccionSwitch(ctx *Background.InstruccionSwitchContext) interface{} {
	exp := ctx.Expr().Accept(p).(compiladorp2.CAbstrExpr)
	cases := ctx.AllInstCase()

	sentenciaSwitch := nt.NewNT_InstSwitch(exp, ctx.Expr().GetStart().GetLine(), ctx.Expr().GetStart().GetColumn())

	if ctx.InstDefault() != nil {
		sentenciaSwitch.AddDefault(ctx.InstDefault().Accept(p).(compiladorp2.CAbstrExpr))
	}

	for _, sentencia := range cases {
		nodo := sentencia.Accept(p).(*nt.NT_InstCase)
		sentenciaSwitch.AddCase(*nodo)
	}
	return sentenciaSwitch
}

// Visit a parse tree produced by ControlParser#InstruccionCase.
func (p CompilerVisit) VisitInstruccionCase(ctx *Background.InstruccionCaseContext) interface{} {
	return nt.NewNT_InstCase(
		ctx.Expr().Accept(p).(compiladorp2.CAbstrExpr),
		ctx.Block().Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#InstruccionDefault.
func (p CompilerVisit) VisitInstruccionDefault(ctx *Background.InstruccionDefaultContext) interface{} {
	return nt.NewNT_InstDefault(
		ctx.Block().Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#InstruccionWhile.
func (p CompilerVisit) VisitInstruccionWhile(ctx *Background.InstruccionWhileContext) interface{} {
	expr := ctx.Expr().Accept(p).(compiladorp2.CAbstrExpr)
	block := ctx.Block().Accept(p).(compiladorp2.CAbstrExpr)
	return nt.NewNT_InstWhile(expr, block, ctx.Expr().GetStart().GetLine(), ctx.Expr().GetStart().GetColumn())
}

// Visit a parse tree produced by ControlParser#InstruccionFor.
func (p CompilerVisit) VisitInstruccionFor(ctx *Background.InstruccionForContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#InstruccionGuard.
func (p CompilerVisit) VisitInstruccionGuard(ctx *Background.InstruccionGuardContext) interface{} {
	return nt.NewNT_Guard(
		ctx.Expr().Accept(p).(compiladorp2.CAbstrExpr),
		ctx.Block().Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#InstruccionBreak.
func (p CompilerVisit) VisitInstruccionBreak(ctx *Background.InstruccionBreakContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#InstruccionContinue.
func (p CompilerVisit) VisitInstruccionContinue(ctx *Background.InstruccionContinueContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#InstruccionReturnSimple.
func (p CompilerVisit) VisitInstruccionReturnSimple(ctx *Background.InstruccionReturnSimpleContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#InstruccionReturnExpresion.
func (p CompilerVisit) VisitInstruccionReturnExpresion(ctx *Background.InstruccionReturnExpresionContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#DecVector_ExpresionLista.
func (p CompilerVisit) VisitDecVector_ExpresionLista(ctx *Background.DecVector_ExpresionListaContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#DecVector_ObjetLista.
func (p CompilerVisit) VisitDecVector_ObjetLista(ctx *Background.DecVector_ObjetListaContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#DecVector_Id.
func (p CompilerVisit) VisitDecVector_Id(ctx *Background.DecVector_IdContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#DecVector_ValCor.
func (p CompilerVisit) VisitDecVector_ValCor(ctx *Background.DecVector_ValCorContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#DecVectorConst_ExpresionLista.
func (p CompilerVisit) VisitDecVectorConst_ExpresionLista(ctx *Background.DecVectorConst_ExpresionListaContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#DecVectorConst_ObjetLista.
func (p CompilerVisit) VisitDecVectorConst_ObjetLista(ctx *Background.DecVectorConst_ObjetListaContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#DecVectorConst_Id.
func (p CompilerVisit) VisitDecVectorConst_Id(ctx *Background.DecVectorConst_IdContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#VectFunc_Append.
func (p CompilerVisit) VisitVectFunc_Append(ctx *Background.VectFunc_AppendContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#VectFunc_Remove.
func (p CompilerVisit) VisitVectFunc_Remove(ctx *Background.VectFunc_RemoveContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#VectFunc_RemoveLast.
func (p CompilerVisit) VisitVectFunc_RemoveLast(ctx *Background.VectFunc_RemoveLastContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#decMatriz.
func (p CompilerVisit) VisitDecMatriz(ctx *Background.DecMatrizContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#tipoMatriz.
func (p CompilerVisit) VisitTipoMatriz(ctx *Background.TipoMatrizContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#defMatrix.
func (p CompilerVisit) VisitDefMatrix(ctx *Background.DefMatrixContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaValoresMat.
func (p CompilerVisit) VisitListaValoresMat(ctx *Background.ListaValoresMatContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaValoresMat2.
func (p CompilerVisit) VisitListaValoresMat2(ctx *Background.ListaValoresMat2Context) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#simpleMax.
func (p CompilerVisit) VisitSimpleMax(ctx *Background.SimpleMaxContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instStruct.
func (p CompilerVisit) VisitInstStruct(ctx *Background.InstStructContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaAtributos.
func (p CompilerVisit) VisitListaAtributos(ctx *Background.ListaAtributosContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#decStruct.
func (p CompilerVisit) VisitDecStruct(ctx *Background.DecStructContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#lDupla.
func (p CompilerVisit) VisitLDupla(ctx *Background.LDuplaContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#llamAtributos.
func (p CompilerVisit) VisitLlamAtributos(ctx *Background.LlamAtributosContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#llamadaMetodos.
func (p CompilerVisit) VisitLlamadaMetodos(ctx *Background.LlamadaMetodosContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Dec_Funcion.
func (p CompilerVisit) VisitDec_Funcion(ctx *Background.Dec_FuncionContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaParametros.
func (p CompilerVisit) VisitListaParametros(ctx *Background.ListaParametrosContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#llamadaFuncion.
func (p CompilerVisit) VisitLlamadaFuncion(ctx *Background.LlamadaFuncionContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaLlamadaParametros.
func (p CompilerVisit) VisitListaLlamadaParametros(ctx *Background.ListaLlamadaParametrosContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#IntruccionPrint.
func (p CompilerVisit) VisitIntruccionPrint(ctx *Background.IntruccionPrintContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Inst_Casteos.
func (p CompilerVisit) VisitInst_Casteos(ctx *Background.Inst_CasteosContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#ListaExpresiones.
func (p CompilerVisit) VisitListaExpresiones(ctx *Background.ListaExpresionesContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#objectsList.
func (p CompilerVisit) VisitObjectsList(ctx *Background.ObjectsListContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Dec_VTipoValor.
func (p CompilerVisit) VisitDec_VTipoValor(ctx *Background.Dec_VTipoValorContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Dec_VTipo.
func (p CompilerVisit) VisitDec_VTipo(ctx *Background.Dec_VTipoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Dec_VVal.
func (p CompilerVisit) VisitDec_VVal(ctx *Background.Dec_VValContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Dec_CVal.
func (p CompilerVisit) VisitDec_CVal(ctx *Background.Dec_CValContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Dec_CTipoValor.
func (p CompilerVisit) VisitDec_CTipoValor(ctx *Background.Dec_CTipoValorContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_Aumento.
func (p CompilerVisit) VisitAsignacion_Aumento(ctx *Background.Asignacion_AumentoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_Decremento.
func (p CompilerVisit) VisitAsignacion_Decremento(ctx *Background.Asignacion_DecrementoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_ValorGen.
func (p CompilerVisit) VisitAsignacion_ValorGen(ctx *Background.Asignacion_ValorGenContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_VectorAumento.
func (p CompilerVisit) VisitAsignacion_VectorAumento(ctx *Background.Asignacion_VectorAumentoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_VectorDecremento.
func (p CompilerVisit) VisitAsignacion_VectorDecremento(ctx *Background.Asignacion_VectorDecrementoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_MatrixAumento.
func (p CompilerVisit) VisitAsignacion_MatrixAumento(ctx *Background.Asignacion_MatrixAumentoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_MatrixDecremento.
func (p CompilerVisit) VisitAsignacion_MatrixDecremento(ctx *Background.Asignacion_MatrixDecrementoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_VectorGeneral.
func (p CompilerVisit) VisitAsignacion_VectorGeneral(ctx *Background.Asignacion_VectorGeneralContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_MatrixGeneral.
func (p CompilerVisit) VisitAsignacion_MatrixGeneral(ctx *Background.Asignacion_MatrixGeneralContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_LlAtribGeneral.
func (p CompilerVisit) VisitAsignacion_LlAtribGeneral(ctx *Background.Asignacion_LlAtribGeneralContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_LlAtribAumento.
func (p CompilerVisit) VisitAsignacion_LlAtribAumento(ctx *Background.Asignacion_LlAtribAumentoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_LlAtribDecremento.
func (p CompilerVisit) VisitAsignacion_LlAtribDecremento(ctx *Background.Asignacion_LlAtribDecrementoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_Decimal.
func (p CompilerVisit) VisitExpr_Decimal(ctx *Background.Expr_DecimalContext) interface{} {
	return t.NewT_Float(
		ctx.DECIMAL().GetText(),
		ctx.DECIMAL().GetSymbol().GetLine(),
		ctx.DECIMAL().GetSymbol().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_InstCasteo.
func (p CompilerVisit) VisitExpr_InstCasteo(ctx *Background.Expr_InstCasteoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_ValidaIgualDif.
func (p CompilerVisit) VisitExpr_ValidaIgualDif(ctx *Background.Expr_ValidaIgualDifContext) interface{} {
	return nt.NewNT_Condicional(
		ctx.Expr(0).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.Expr(1).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetOp().GetText(),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_OpMulDiv.
func (p CompilerVisit) VisitExpr_OpMulDiv(ctx *Background.Expr_OpMulDivContext) interface{} {
	return nt.NewNT_MultDiv(
		ctx.Expr(0).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.Expr(1).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetOp().GetText(),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_ValidaAnd.
func (p CompilerVisit) VisitExpr_ValidaAnd(ctx *Background.Expr_ValidaAndContext) interface{} {
	return nt.NewNT_CondAnd(
		ctx.Expr(0).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.Expr(1).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_LlamAtributos.
func (p CompilerVisit) VisitExpr_LlamAtributos(ctx *Background.Expr_LlamAtributosContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_Nil.
func (p CompilerVisit) VisitExpr_Nil(ctx *Background.Expr_NilContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_PosMatrix.
func (p CompilerVisit) VisitExpr_PosMatrix(ctx *Background.Expr_PosMatrixContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_Boolean.
func (p CompilerVisit) VisitExpr_Boolean(ctx *Background.Expr_BooleanContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_InstRango.
func (p CompilerVisit) VisitExpr_InstRango(ctx *Background.Expr_InstRangoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_ValNumNeg.
func (p CompilerVisit) VisitExpr_ValNumNeg(ctx *Background.Expr_ValNumNegContext) interface{} {
	return nt.NewNT_NegUnaria(
		ctx.Expr().Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_ValidaOr.
func (p CompilerVisit) VisitExpr_ValidaOr(ctx *Background.Expr_ValidaOrContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_ValNegacion.
func (p CompilerVisit) VisitExpr_ValNegacion(ctx *Background.Expr_ValNegacionContext) interface{} {
	return nt.NewT_NegExp(
		ctx.Expr().Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_ID.
func (p CompilerVisit) VisitExpr_ID(ctx *Background.Expr_IDContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_ValidaMayQue.
func (p CompilerVisit) VisitExpr_ValidaMayQue(ctx *Background.Expr_ValidaMayQueContext) interface{} {
	return nt.NewNT_Condicional(
		ctx.Expr(0).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.Expr(1).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetOp().GetText(),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_Conteo.
func (p CompilerVisit) VisitExpr_Conteo(ctx *Background.Expr_ConteoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_IsEmpty.
func (p CompilerVisit) VisitExpr_IsEmpty(ctx *Background.Expr_IsEmptyContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_OpSumRes.
func (p CompilerVisit) VisitExpr_OpSumRes(ctx *Background.Expr_OpSumResContext) interface{} {
	return nt.NewNT_Sum(
		ctx.Expr(0).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.Expr(1).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetOp().GetText(),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_Entero.
func (p CompilerVisit) VisitExpr_Entero(ctx *Background.Expr_EnteroContext) interface{} {
	return t.NewT_Entero(
		ctx.INT().GetText(),
		ctx.INT().GetSymbol().GetLine(),
		ctx.INT().GetSymbol().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_ParExpre.
func (p CompilerVisit) VisitExpr_ParExpre(ctx *Background.Expr_ParExpreContext) interface{} {
	return ctx.Expr().Accept(p).(compiladorp2.CAbstrExpr)
}

// Visit a parse tree produced by ControlParser#Expr_StringChar.
func (p CompilerVisit) VisitExpr_StringChar(ctx *Background.Expr_StringCharContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_ValidaMenQue.
func (p CompilerVisit) VisitExpr_ValidaMenQue(ctx *Background.Expr_ValidaMenQueContext) interface{} {
	return nt.NewNT_Condicional(
		ctx.Expr(0).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.Expr(1).Accept(p).(compiladorp2.CAbstrExpr),
		ctx.GetOp().GetText(),
		ctx.GetStart().GetLine(),
		ctx.GetStart().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_PosVector.
func (p CompilerVisit) VisitExpr_PosVector(ctx *Background.Expr_PosVectorContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#tipovariable.
func (p CompilerVisit) VisitTipovariable(ctx *Background.TipovariableContext) interface{} {
	panic("not implemented") // TODO: Implement
}
