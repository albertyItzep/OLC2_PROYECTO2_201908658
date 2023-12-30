// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import "github.com/antlr4-go/antlr/v4"

// BaseControlListener is a complete listener for a parse tree produced by ControlParser.
type BaseControlListener struct{}

var _ ControlListener = &BaseControlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseControlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseControlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseControlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseControlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgInicial is called when production ProgInicial is entered.
func (s *BaseControlListener) EnterProgInicial(ctx *ProgInicialContext) {}

// ExitProgInicial is called when production ProgInicial is exited.
func (s *BaseControlListener) ExitProgInicial(ctx *ProgInicialContext) {}

// EnterBlockContent is called when production BlockContent is entered.
func (s *BaseControlListener) EnterBlockContent(ctx *BlockContentContext) {}

// ExitBlockContent is called when production BlockContent is exited.
func (s *BaseControlListener) ExitBlockContent(ctx *BlockContentContext) {}

// EnterSentenciaDeclaracion is called when production SentenciaDeclaracion is entered.
func (s *BaseControlListener) EnterSentenciaDeclaracion(ctx *SentenciaDeclaracionContext) {}

// ExitSentenciaDeclaracion is called when production SentenciaDeclaracion is exited.
func (s *BaseControlListener) ExitSentenciaDeclaracion(ctx *SentenciaDeclaracionContext) {}

// EnterSentenciaDeclaracionVector is called when production SentenciaDeclaracionVector is entered.
func (s *BaseControlListener) EnterSentenciaDeclaracionVector(ctx *SentenciaDeclaracionVectorContext) {
}

// ExitSentenciaDeclaracionVector is called when production SentenciaDeclaracionVector is exited.
func (s *BaseControlListener) ExitSentenciaDeclaracionVector(ctx *SentenciaDeclaracionVectorContext) {
}

// EnterSentenciaDeclaracionMatrix is called when production SentenciaDeclaracionMatrix is entered.
func (s *BaseControlListener) EnterSentenciaDeclaracionMatrix(ctx *SentenciaDeclaracionMatrixContext) {
}

// ExitSentenciaDeclaracionMatrix is called when production SentenciaDeclaracionMatrix is exited.
func (s *BaseControlListener) ExitSentenciaDeclaracionMatrix(ctx *SentenciaDeclaracionMatrixContext) {
}

// EnterSentenciaDeclaracionFuncion is called when production SentenciaDeclaracionFuncion is entered.
func (s *BaseControlListener) EnterSentenciaDeclaracionFuncion(ctx *SentenciaDeclaracionFuncionContext) {
}

// ExitSentenciaDeclaracionFuncion is called when production SentenciaDeclaracionFuncion is exited.
func (s *BaseControlListener) ExitSentenciaDeclaracionFuncion(ctx *SentenciaDeclaracionFuncionContext) {
}

// EnterSentenciaAsignacion is called when production SentenciaAsignacion is entered.
func (s *BaseControlListener) EnterSentenciaAsignacion(ctx *SentenciaAsignacionContext) {}

// ExitSentenciaAsignacion is called when production SentenciaAsignacion is exited.
func (s *BaseControlListener) ExitSentenciaAsignacion(ctx *SentenciaAsignacionContext) {}

// EnterSentenciaIf is called when production SentenciaIf is entered.
func (s *BaseControlListener) EnterSentenciaIf(ctx *SentenciaIfContext) {}

// ExitSentenciaIf is called when production SentenciaIf is exited.
func (s *BaseControlListener) ExitSentenciaIf(ctx *SentenciaIfContext) {}

// EnterSentenciaSwitch is called when production SentenciaSwitch is entered.
func (s *BaseControlListener) EnterSentenciaSwitch(ctx *SentenciaSwitchContext) {}

// ExitSentenciaSwitch is called when production SentenciaSwitch is exited.
func (s *BaseControlListener) ExitSentenciaSwitch(ctx *SentenciaSwitchContext) {}

// EnterSentenciaWhile is called when production SentenciaWhile is entered.
func (s *BaseControlListener) EnterSentenciaWhile(ctx *SentenciaWhileContext) {}

// ExitSentenciaWhile is called when production SentenciaWhile is exited.
func (s *BaseControlListener) ExitSentenciaWhile(ctx *SentenciaWhileContext) {}

// EnterSentenciaFor is called when production SentenciaFor is entered.
func (s *BaseControlListener) EnterSentenciaFor(ctx *SentenciaForContext) {}

// ExitSentenciaFor is called when production SentenciaFor is exited.
func (s *BaseControlListener) ExitSentenciaFor(ctx *SentenciaForContext) {}

// EnterSentenciaGuard is called when production SentenciaGuard is entered.
func (s *BaseControlListener) EnterSentenciaGuard(ctx *SentenciaGuardContext) {}

// ExitSentenciaGuard is called when production SentenciaGuard is exited.
func (s *BaseControlListener) ExitSentenciaGuard(ctx *SentenciaGuardContext) {}

// EnterSentenciaBreak is called when production SentenciaBreak is entered.
func (s *BaseControlListener) EnterSentenciaBreak(ctx *SentenciaBreakContext) {}

// ExitSentenciaBreak is called when production SentenciaBreak is exited.
func (s *BaseControlListener) ExitSentenciaBreak(ctx *SentenciaBreakContext) {}

// EnterSentenciaContinue is called when production SentenciaContinue is entered.
func (s *BaseControlListener) EnterSentenciaContinue(ctx *SentenciaContinueContext) {}

// ExitSentenciaContinue is called when production SentenciaContinue is exited.
func (s *BaseControlListener) ExitSentenciaContinue(ctx *SentenciaContinueContext) {}

// EnterSentenciaReturn is called when production SentenciaReturn is entered.
func (s *BaseControlListener) EnterSentenciaReturn(ctx *SentenciaReturnContext) {}

// ExitSentenciaReturn is called when production SentenciaReturn is exited.
func (s *BaseControlListener) ExitSentenciaReturn(ctx *SentenciaReturnContext) {}

// EnterSentenciaFuncVectoriales is called when production SentenciaFuncVectoriales is entered.
func (s *BaseControlListener) EnterSentenciaFuncVectoriales(ctx *SentenciaFuncVectorialesContext) {}

// ExitSentenciaFuncVectoriales is called when production SentenciaFuncVectoriales is exited.
func (s *BaseControlListener) ExitSentenciaFuncVectoriales(ctx *SentenciaFuncVectorialesContext) {}

// EnterSentenciaLLamadaFuncion is called when production SentenciaLLamadaFuncion is entered.
func (s *BaseControlListener) EnterSentenciaLLamadaFuncion(ctx *SentenciaLLamadaFuncionContext) {}

// ExitSentenciaLLamadaFuncion is called when production SentenciaLLamadaFuncion is exited.
func (s *BaseControlListener) ExitSentenciaLLamadaFuncion(ctx *SentenciaLLamadaFuncionContext) {}

// EnterSentenciaPrint is called when production SentenciaPrint is entered.
func (s *BaseControlListener) EnterSentenciaPrint(ctx *SentenciaPrintContext) {}

// ExitSentenciaPrint is called when production SentenciaPrint is exited.
func (s *BaseControlListener) ExitSentenciaPrint(ctx *SentenciaPrintContext) {}

// EnterSentenciaStruct is called when production SentenciaStruct is entered.
func (s *BaseControlListener) EnterSentenciaStruct(ctx *SentenciaStructContext) {}

// ExitSentenciaStruct is called when production SentenciaStruct is exited.
func (s *BaseControlListener) ExitSentenciaStruct(ctx *SentenciaStructContext) {}

// EnterSentenciaDeclaracionStruct is called when production SentenciaDeclaracionStruct is entered.
func (s *BaseControlListener) EnterSentenciaDeclaracionStruct(ctx *SentenciaDeclaracionStructContext) {
}

// ExitSentenciaDeclaracionStruct is called when production SentenciaDeclaracionStruct is exited.
func (s *BaseControlListener) ExitSentenciaDeclaracionStruct(ctx *SentenciaDeclaracionStructContext) {
}

// EnterSentenciaLLamadaMetodos is called when production SentenciaLLamadaMetodos is entered.
func (s *BaseControlListener) EnterSentenciaLLamadaMetodos(ctx *SentenciaLLamadaMetodosContext) {}

// ExitSentenciaLLamadaMetodos is called when production SentenciaLLamadaMetodos is exited.
func (s *BaseControlListener) ExitSentenciaLLamadaMetodos(ctx *SentenciaLLamadaMetodosContext) {}

// EnterInstruccionIf is called when production InstruccionIf is entered.
func (s *BaseControlListener) EnterInstruccionIf(ctx *InstruccionIfContext) {}

// ExitInstruccionIf is called when production InstruccionIf is exited.
func (s *BaseControlListener) ExitInstruccionIf(ctx *InstruccionIfContext) {}

// EnterInstruccionSwitch is called when production InstruccionSwitch is entered.
func (s *BaseControlListener) EnterInstruccionSwitch(ctx *InstruccionSwitchContext) {}

// ExitInstruccionSwitch is called when production InstruccionSwitch is exited.
func (s *BaseControlListener) ExitInstruccionSwitch(ctx *InstruccionSwitchContext) {}

// EnterInstruccionCase is called when production InstruccionCase is entered.
func (s *BaseControlListener) EnterInstruccionCase(ctx *InstruccionCaseContext) {}

// ExitInstruccionCase is called when production InstruccionCase is exited.
func (s *BaseControlListener) ExitInstruccionCase(ctx *InstruccionCaseContext) {}

// EnterInstruccionDefault is called when production InstruccionDefault is entered.
func (s *BaseControlListener) EnterInstruccionDefault(ctx *InstruccionDefaultContext) {}

// ExitInstruccionDefault is called when production InstruccionDefault is exited.
func (s *BaseControlListener) ExitInstruccionDefault(ctx *InstruccionDefaultContext) {}

// EnterInstruccionWhile is called when production InstruccionWhile is entered.
func (s *BaseControlListener) EnterInstruccionWhile(ctx *InstruccionWhileContext) {}

// ExitInstruccionWhile is called when production InstruccionWhile is exited.
func (s *BaseControlListener) ExitInstruccionWhile(ctx *InstruccionWhileContext) {}

// EnterInstruccionFor is called when production InstruccionFor is entered.
func (s *BaseControlListener) EnterInstruccionFor(ctx *InstruccionForContext) {}

// ExitInstruccionFor is called when production InstruccionFor is exited.
func (s *BaseControlListener) ExitInstruccionFor(ctx *InstruccionForContext) {}

// EnterInstruccionGuard is called when production InstruccionGuard is entered.
func (s *BaseControlListener) EnterInstruccionGuard(ctx *InstruccionGuardContext) {}

// ExitInstruccionGuard is called when production InstruccionGuard is exited.
func (s *BaseControlListener) ExitInstruccionGuard(ctx *InstruccionGuardContext) {}

// EnterInstruccionBreak is called when production InstruccionBreak is entered.
func (s *BaseControlListener) EnterInstruccionBreak(ctx *InstruccionBreakContext) {}

// ExitInstruccionBreak is called when production InstruccionBreak is exited.
func (s *BaseControlListener) ExitInstruccionBreak(ctx *InstruccionBreakContext) {}

// EnterInstruccionContinue is called when production InstruccionContinue is entered.
func (s *BaseControlListener) EnterInstruccionContinue(ctx *InstruccionContinueContext) {}

// ExitInstruccionContinue is called when production InstruccionContinue is exited.
func (s *BaseControlListener) ExitInstruccionContinue(ctx *InstruccionContinueContext) {}

// EnterInstruccionReturnSimple is called when production InstruccionReturnSimple is entered.
func (s *BaseControlListener) EnterInstruccionReturnSimple(ctx *InstruccionReturnSimpleContext) {}

// ExitInstruccionReturnSimple is called when production InstruccionReturnSimple is exited.
func (s *BaseControlListener) ExitInstruccionReturnSimple(ctx *InstruccionReturnSimpleContext) {}

// EnterInstruccionReturnExpresion is called when production InstruccionReturnExpresion is entered.
func (s *BaseControlListener) EnterInstruccionReturnExpresion(ctx *InstruccionReturnExpresionContext) {
}

// ExitInstruccionReturnExpresion is called when production InstruccionReturnExpresion is exited.
func (s *BaseControlListener) ExitInstruccionReturnExpresion(ctx *InstruccionReturnExpresionContext) {
}

// EnterDecVector_ExpresionLista is called when production DecVector_ExpresionLista is entered.
func (s *BaseControlListener) EnterDecVector_ExpresionLista(ctx *DecVector_ExpresionListaContext) {}

// ExitDecVector_ExpresionLista is called when production DecVector_ExpresionLista is exited.
func (s *BaseControlListener) ExitDecVector_ExpresionLista(ctx *DecVector_ExpresionListaContext) {}

// EnterDecVector_ObjetLista is called when production DecVector_ObjetLista is entered.
func (s *BaseControlListener) EnterDecVector_ObjetLista(ctx *DecVector_ObjetListaContext) {}

// ExitDecVector_ObjetLista is called when production DecVector_ObjetLista is exited.
func (s *BaseControlListener) ExitDecVector_ObjetLista(ctx *DecVector_ObjetListaContext) {}

// EnterDecVector_Id is called when production DecVector_Id is entered.
func (s *BaseControlListener) EnterDecVector_Id(ctx *DecVector_IdContext) {}

// ExitDecVector_Id is called when production DecVector_Id is exited.
func (s *BaseControlListener) ExitDecVector_Id(ctx *DecVector_IdContext) {}

// EnterDecVector_ValCor is called when production DecVector_ValCor is entered.
func (s *BaseControlListener) EnterDecVector_ValCor(ctx *DecVector_ValCorContext) {}

// ExitDecVector_ValCor is called when production DecVector_ValCor is exited.
func (s *BaseControlListener) ExitDecVector_ValCor(ctx *DecVector_ValCorContext) {}

// EnterDecVectorConst_ExpresionLista is called when production DecVectorConst_ExpresionLista is entered.
func (s *BaseControlListener) EnterDecVectorConst_ExpresionLista(ctx *DecVectorConst_ExpresionListaContext) {
}

// ExitDecVectorConst_ExpresionLista is called when production DecVectorConst_ExpresionLista is exited.
func (s *BaseControlListener) ExitDecVectorConst_ExpresionLista(ctx *DecVectorConst_ExpresionListaContext) {
}

// EnterDecVectorConst_ObjetLista is called when production DecVectorConst_ObjetLista is entered.
func (s *BaseControlListener) EnterDecVectorConst_ObjetLista(ctx *DecVectorConst_ObjetListaContext) {}

// ExitDecVectorConst_ObjetLista is called when production DecVectorConst_ObjetLista is exited.
func (s *BaseControlListener) ExitDecVectorConst_ObjetLista(ctx *DecVectorConst_ObjetListaContext) {}

// EnterDecVectorConst_Id is called when production DecVectorConst_Id is entered.
func (s *BaseControlListener) EnterDecVectorConst_Id(ctx *DecVectorConst_IdContext) {}

// ExitDecVectorConst_Id is called when production DecVectorConst_Id is exited.
func (s *BaseControlListener) ExitDecVectorConst_Id(ctx *DecVectorConst_IdContext) {}

// EnterVectFunc_Append is called when production VectFunc_Append is entered.
func (s *BaseControlListener) EnterVectFunc_Append(ctx *VectFunc_AppendContext) {}

// ExitVectFunc_Append is called when production VectFunc_Append is exited.
func (s *BaseControlListener) ExitVectFunc_Append(ctx *VectFunc_AppendContext) {}

// EnterVectFunc_Remove is called when production VectFunc_Remove is entered.
func (s *BaseControlListener) EnterVectFunc_Remove(ctx *VectFunc_RemoveContext) {}

// ExitVectFunc_Remove is called when production VectFunc_Remove is exited.
func (s *BaseControlListener) ExitVectFunc_Remove(ctx *VectFunc_RemoveContext) {}

// EnterVectFunc_RemoveLast is called when production VectFunc_RemoveLast is entered.
func (s *BaseControlListener) EnterVectFunc_RemoveLast(ctx *VectFunc_RemoveLastContext) {}

// ExitVectFunc_RemoveLast is called when production VectFunc_RemoveLast is exited.
func (s *BaseControlListener) ExitVectFunc_RemoveLast(ctx *VectFunc_RemoveLastContext) {}

// EnterDecMatriz is called when production decMatriz is entered.
func (s *BaseControlListener) EnterDecMatriz(ctx *DecMatrizContext) {}

// ExitDecMatriz is called when production decMatriz is exited.
func (s *BaseControlListener) ExitDecMatriz(ctx *DecMatrizContext) {}

// EnterTipoMatriz is called when production tipoMatriz is entered.
func (s *BaseControlListener) EnterTipoMatriz(ctx *TipoMatrizContext) {}

// ExitTipoMatriz is called when production tipoMatriz is exited.
func (s *BaseControlListener) ExitTipoMatriz(ctx *TipoMatrizContext) {}

// EnterDefMatrix is called when production defMatrix is entered.
func (s *BaseControlListener) EnterDefMatrix(ctx *DefMatrixContext) {}

// ExitDefMatrix is called when production defMatrix is exited.
func (s *BaseControlListener) ExitDefMatrix(ctx *DefMatrixContext) {}

// EnterListaValoresMat is called when production listaValoresMat is entered.
func (s *BaseControlListener) EnterListaValoresMat(ctx *ListaValoresMatContext) {}

// ExitListaValoresMat is called when production listaValoresMat is exited.
func (s *BaseControlListener) ExitListaValoresMat(ctx *ListaValoresMatContext) {}

// EnterListaValoresMat2 is called when production listaValoresMat2 is entered.
func (s *BaseControlListener) EnterListaValoresMat2(ctx *ListaValoresMat2Context) {}

// ExitListaValoresMat2 is called when production listaValoresMat2 is exited.
func (s *BaseControlListener) ExitListaValoresMat2(ctx *ListaValoresMat2Context) {}

// EnterSimpleMax is called when production simpleMax is entered.
func (s *BaseControlListener) EnterSimpleMax(ctx *SimpleMaxContext) {}

// ExitSimpleMax is called when production simpleMax is exited.
func (s *BaseControlListener) ExitSimpleMax(ctx *SimpleMaxContext) {}

// EnterInstStruct is called when production instStruct is entered.
func (s *BaseControlListener) EnterInstStruct(ctx *InstStructContext) {}

// ExitInstStruct is called when production instStruct is exited.
func (s *BaseControlListener) ExitInstStruct(ctx *InstStructContext) {}

// EnterListaAtributos is called when production listaAtributos is entered.
func (s *BaseControlListener) EnterListaAtributos(ctx *ListaAtributosContext) {}

// ExitListaAtributos is called when production listaAtributos is exited.
func (s *BaseControlListener) ExitListaAtributos(ctx *ListaAtributosContext) {}

// EnterDecStruct is called when production decStruct is entered.
func (s *BaseControlListener) EnterDecStruct(ctx *DecStructContext) {}

// ExitDecStruct is called when production decStruct is exited.
func (s *BaseControlListener) ExitDecStruct(ctx *DecStructContext) {}

// EnterLDupla is called when production lDupla is entered.
func (s *BaseControlListener) EnterLDupla(ctx *LDuplaContext) {}

// ExitLDupla is called when production lDupla is exited.
func (s *BaseControlListener) ExitLDupla(ctx *LDuplaContext) {}

// EnterLlamAtributos is called when production llamAtributos is entered.
func (s *BaseControlListener) EnterLlamAtributos(ctx *LlamAtributosContext) {}

// ExitLlamAtributos is called when production llamAtributos is exited.
func (s *BaseControlListener) ExitLlamAtributos(ctx *LlamAtributosContext) {}

// EnterLlamadaMetodos is called when production llamadaMetodos is entered.
func (s *BaseControlListener) EnterLlamadaMetodos(ctx *LlamadaMetodosContext) {}

// ExitLlamadaMetodos is called when production llamadaMetodos is exited.
func (s *BaseControlListener) ExitLlamadaMetodos(ctx *LlamadaMetodosContext) {}

// EnterDec_Funcion is called when production Dec_Funcion is entered.
func (s *BaseControlListener) EnterDec_Funcion(ctx *Dec_FuncionContext) {}

// ExitDec_Funcion is called when production Dec_Funcion is exited.
func (s *BaseControlListener) ExitDec_Funcion(ctx *Dec_FuncionContext) {}

// EnterListaParametros is called when production listaParametros is entered.
func (s *BaseControlListener) EnterListaParametros(ctx *ListaParametrosContext) {}

// ExitListaParametros is called when production listaParametros is exited.
func (s *BaseControlListener) ExitListaParametros(ctx *ListaParametrosContext) {}

// EnterLlamadaFuncion is called when production llamadaFuncion is entered.
func (s *BaseControlListener) EnterLlamadaFuncion(ctx *LlamadaFuncionContext) {}

// ExitLlamadaFuncion is called when production llamadaFuncion is exited.
func (s *BaseControlListener) ExitLlamadaFuncion(ctx *LlamadaFuncionContext) {}

// EnterListaLlamadaParametros is called when production listaLlamadaParametros is entered.
func (s *BaseControlListener) EnterListaLlamadaParametros(ctx *ListaLlamadaParametrosContext) {}

// ExitListaLlamadaParametros is called when production listaLlamadaParametros is exited.
func (s *BaseControlListener) ExitListaLlamadaParametros(ctx *ListaLlamadaParametrosContext) {}

// EnterIntruccionPrint is called when production IntruccionPrint is entered.
func (s *BaseControlListener) EnterIntruccionPrint(ctx *IntruccionPrintContext) {}

// ExitIntruccionPrint is called when production IntruccionPrint is exited.
func (s *BaseControlListener) ExitIntruccionPrint(ctx *IntruccionPrintContext) {}

// EnterInst_Casteos is called when production Inst_Casteos is entered.
func (s *BaseControlListener) EnterInst_Casteos(ctx *Inst_CasteosContext) {}

// ExitInst_Casteos is called when production Inst_Casteos is exited.
func (s *BaseControlListener) ExitInst_Casteos(ctx *Inst_CasteosContext) {}

// EnterListaExpresiones is called when production ListaExpresiones is entered.
func (s *BaseControlListener) EnterListaExpresiones(ctx *ListaExpresionesContext) {}

// ExitListaExpresiones is called when production ListaExpresiones is exited.
func (s *BaseControlListener) ExitListaExpresiones(ctx *ListaExpresionesContext) {}

// EnterObjectsList is called when production objectsList is entered.
func (s *BaseControlListener) EnterObjectsList(ctx *ObjectsListContext) {}

// ExitObjectsList is called when production objectsList is exited.
func (s *BaseControlListener) ExitObjectsList(ctx *ObjectsListContext) {}

// EnterDec_VTipoValor is called when production Dec_VTipoValor is entered.
func (s *BaseControlListener) EnterDec_VTipoValor(ctx *Dec_VTipoValorContext) {}

// ExitDec_VTipoValor is called when production Dec_VTipoValor is exited.
func (s *BaseControlListener) ExitDec_VTipoValor(ctx *Dec_VTipoValorContext) {}

// EnterDec_VTipo is called when production Dec_VTipo is entered.
func (s *BaseControlListener) EnterDec_VTipo(ctx *Dec_VTipoContext) {}

// ExitDec_VTipo is called when production Dec_VTipo is exited.
func (s *BaseControlListener) ExitDec_VTipo(ctx *Dec_VTipoContext) {}

// EnterDec_VVal is called when production Dec_VVal is entered.
func (s *BaseControlListener) EnterDec_VVal(ctx *Dec_VValContext) {}

// ExitDec_VVal is called when production Dec_VVal is exited.
func (s *BaseControlListener) ExitDec_VVal(ctx *Dec_VValContext) {}

// EnterDec_CVal is called when production Dec_CVal is entered.
func (s *BaseControlListener) EnterDec_CVal(ctx *Dec_CValContext) {}

// ExitDec_CVal is called when production Dec_CVal is exited.
func (s *BaseControlListener) ExitDec_CVal(ctx *Dec_CValContext) {}

// EnterDec_CTipoValor is called when production Dec_CTipoValor is entered.
func (s *BaseControlListener) EnterDec_CTipoValor(ctx *Dec_CTipoValorContext) {}

// ExitDec_CTipoValor is called when production Dec_CTipoValor is exited.
func (s *BaseControlListener) ExitDec_CTipoValor(ctx *Dec_CTipoValorContext) {}

// EnterAsignacion_Aumento is called when production Asignacion_Aumento is entered.
func (s *BaseControlListener) EnterAsignacion_Aumento(ctx *Asignacion_AumentoContext) {}

// ExitAsignacion_Aumento is called when production Asignacion_Aumento is exited.
func (s *BaseControlListener) ExitAsignacion_Aumento(ctx *Asignacion_AumentoContext) {}

// EnterAsignacion_Decremento is called when production Asignacion_Decremento is entered.
func (s *BaseControlListener) EnterAsignacion_Decremento(ctx *Asignacion_DecrementoContext) {}

// ExitAsignacion_Decremento is called when production Asignacion_Decremento is exited.
func (s *BaseControlListener) ExitAsignacion_Decremento(ctx *Asignacion_DecrementoContext) {}

// EnterAsignacion_ValorGen is called when production Asignacion_ValorGen is entered.
func (s *BaseControlListener) EnterAsignacion_ValorGen(ctx *Asignacion_ValorGenContext) {}

// ExitAsignacion_ValorGen is called when production Asignacion_ValorGen is exited.
func (s *BaseControlListener) ExitAsignacion_ValorGen(ctx *Asignacion_ValorGenContext) {}

// EnterAsignacion_VectorAumento is called when production Asignacion_VectorAumento is entered.
func (s *BaseControlListener) EnterAsignacion_VectorAumento(ctx *Asignacion_VectorAumentoContext) {}

// ExitAsignacion_VectorAumento is called when production Asignacion_VectorAumento is exited.
func (s *BaseControlListener) ExitAsignacion_VectorAumento(ctx *Asignacion_VectorAumentoContext) {}

// EnterAsignacion_VectorDecremento is called when production Asignacion_VectorDecremento is entered.
func (s *BaseControlListener) EnterAsignacion_VectorDecremento(ctx *Asignacion_VectorDecrementoContext) {
}

// ExitAsignacion_VectorDecremento is called when production Asignacion_VectorDecremento is exited.
func (s *BaseControlListener) ExitAsignacion_VectorDecremento(ctx *Asignacion_VectorDecrementoContext) {
}

// EnterAsignacion_MatrixAumento is called when production Asignacion_MatrixAumento is entered.
func (s *BaseControlListener) EnterAsignacion_MatrixAumento(ctx *Asignacion_MatrixAumentoContext) {}

// ExitAsignacion_MatrixAumento is called when production Asignacion_MatrixAumento is exited.
func (s *BaseControlListener) ExitAsignacion_MatrixAumento(ctx *Asignacion_MatrixAumentoContext) {}

// EnterAsignacion_MatrixDecremento is called when production Asignacion_MatrixDecremento is entered.
func (s *BaseControlListener) EnterAsignacion_MatrixDecremento(ctx *Asignacion_MatrixDecrementoContext) {
}

// ExitAsignacion_MatrixDecremento is called when production Asignacion_MatrixDecremento is exited.
func (s *BaseControlListener) ExitAsignacion_MatrixDecremento(ctx *Asignacion_MatrixDecrementoContext) {
}

// EnterAsignacion_VectorGeneral is called when production Asignacion_VectorGeneral is entered.
func (s *BaseControlListener) EnterAsignacion_VectorGeneral(ctx *Asignacion_VectorGeneralContext) {}

// ExitAsignacion_VectorGeneral is called when production Asignacion_VectorGeneral is exited.
func (s *BaseControlListener) ExitAsignacion_VectorGeneral(ctx *Asignacion_VectorGeneralContext) {}

// EnterAsignacion_MatrixGeneral is called when production Asignacion_MatrixGeneral is entered.
func (s *BaseControlListener) EnterAsignacion_MatrixGeneral(ctx *Asignacion_MatrixGeneralContext) {}

// ExitAsignacion_MatrixGeneral is called when production Asignacion_MatrixGeneral is exited.
func (s *BaseControlListener) ExitAsignacion_MatrixGeneral(ctx *Asignacion_MatrixGeneralContext) {}

// EnterAsignacion_LlAtribGeneral is called when production Asignacion_LlAtribGeneral is entered.
func (s *BaseControlListener) EnterAsignacion_LlAtribGeneral(ctx *Asignacion_LlAtribGeneralContext) {}

// ExitAsignacion_LlAtribGeneral is called when production Asignacion_LlAtribGeneral is exited.
func (s *BaseControlListener) ExitAsignacion_LlAtribGeneral(ctx *Asignacion_LlAtribGeneralContext) {}

// EnterAsignacion_LlAtribAumento is called when production Asignacion_LlAtribAumento is entered.
func (s *BaseControlListener) EnterAsignacion_LlAtribAumento(ctx *Asignacion_LlAtribAumentoContext) {}

// ExitAsignacion_LlAtribAumento is called when production Asignacion_LlAtribAumento is exited.
func (s *BaseControlListener) ExitAsignacion_LlAtribAumento(ctx *Asignacion_LlAtribAumentoContext) {}

// EnterAsignacion_LlAtribDecremento is called when production Asignacion_LlAtribDecremento is entered.
func (s *BaseControlListener) EnterAsignacion_LlAtribDecremento(ctx *Asignacion_LlAtribDecrementoContext) {
}

// ExitAsignacion_LlAtribDecremento is called when production Asignacion_LlAtribDecremento is exited.
func (s *BaseControlListener) ExitAsignacion_LlAtribDecremento(ctx *Asignacion_LlAtribDecrementoContext) {
}

// EnterExpr_Decimal is called when production Expr_Decimal is entered.
func (s *BaseControlListener) EnterExpr_Decimal(ctx *Expr_DecimalContext) {}

// ExitExpr_Decimal is called when production Expr_Decimal is exited.
func (s *BaseControlListener) ExitExpr_Decimal(ctx *Expr_DecimalContext) {}

// EnterExpr_InstCasteo is called when production Expr_InstCasteo is entered.
func (s *BaseControlListener) EnterExpr_InstCasteo(ctx *Expr_InstCasteoContext) {}

// ExitExpr_InstCasteo is called when production Expr_InstCasteo is exited.
func (s *BaseControlListener) ExitExpr_InstCasteo(ctx *Expr_InstCasteoContext) {}

// EnterExpr_ValidaIgualDif is called when production Expr_ValidaIgualDif is entered.
func (s *BaseControlListener) EnterExpr_ValidaIgualDif(ctx *Expr_ValidaIgualDifContext) {}

// ExitExpr_ValidaIgualDif is called when production Expr_ValidaIgualDif is exited.
func (s *BaseControlListener) ExitExpr_ValidaIgualDif(ctx *Expr_ValidaIgualDifContext) {}

// EnterExpr_OpMulDiv is called when production Expr_OpMulDiv is entered.
func (s *BaseControlListener) EnterExpr_OpMulDiv(ctx *Expr_OpMulDivContext) {}

// ExitExpr_OpMulDiv is called when production Expr_OpMulDiv is exited.
func (s *BaseControlListener) ExitExpr_OpMulDiv(ctx *Expr_OpMulDivContext) {}

// EnterExpr_ValidaAnd is called when production Expr_ValidaAnd is entered.
func (s *BaseControlListener) EnterExpr_ValidaAnd(ctx *Expr_ValidaAndContext) {}

// ExitExpr_ValidaAnd is called when production Expr_ValidaAnd is exited.
func (s *BaseControlListener) ExitExpr_ValidaAnd(ctx *Expr_ValidaAndContext) {}

// EnterExpr_LlamAtributos is called when production Expr_LlamAtributos is entered.
func (s *BaseControlListener) EnterExpr_LlamAtributos(ctx *Expr_LlamAtributosContext) {}

// ExitExpr_LlamAtributos is called when production Expr_LlamAtributos is exited.
func (s *BaseControlListener) ExitExpr_LlamAtributos(ctx *Expr_LlamAtributosContext) {}

// EnterExpr_Nil is called when production Expr_Nil is entered.
func (s *BaseControlListener) EnterExpr_Nil(ctx *Expr_NilContext) {}

// ExitExpr_Nil is called when production Expr_Nil is exited.
func (s *BaseControlListener) ExitExpr_Nil(ctx *Expr_NilContext) {}

// EnterExpr_PosMatrix is called when production Expr_PosMatrix is entered.
func (s *BaseControlListener) EnterExpr_PosMatrix(ctx *Expr_PosMatrixContext) {}

// ExitExpr_PosMatrix is called when production Expr_PosMatrix is exited.
func (s *BaseControlListener) ExitExpr_PosMatrix(ctx *Expr_PosMatrixContext) {}

// EnterExpr_Boolean is called when production Expr_Boolean is entered.
func (s *BaseControlListener) EnterExpr_Boolean(ctx *Expr_BooleanContext) {}

// ExitExpr_Boolean is called when production Expr_Boolean is exited.
func (s *BaseControlListener) ExitExpr_Boolean(ctx *Expr_BooleanContext) {}

// EnterExpr_InstRango is called when production Expr_InstRango is entered.
func (s *BaseControlListener) EnterExpr_InstRango(ctx *Expr_InstRangoContext) {}

// ExitExpr_InstRango is called when production Expr_InstRango is exited.
func (s *BaseControlListener) ExitExpr_InstRango(ctx *Expr_InstRangoContext) {}

// EnterExpr_ValNumNeg is called when production Expr_ValNumNeg is entered.
func (s *BaseControlListener) EnterExpr_ValNumNeg(ctx *Expr_ValNumNegContext) {}

// ExitExpr_ValNumNeg is called when production Expr_ValNumNeg is exited.
func (s *BaseControlListener) ExitExpr_ValNumNeg(ctx *Expr_ValNumNegContext) {}

// EnterExpr_ValidaOr is called when production Expr_ValidaOr is entered.
func (s *BaseControlListener) EnterExpr_ValidaOr(ctx *Expr_ValidaOrContext) {}

// ExitExpr_ValidaOr is called when production Expr_ValidaOr is exited.
func (s *BaseControlListener) ExitExpr_ValidaOr(ctx *Expr_ValidaOrContext) {}

// EnterExpr_ValNegacion is called when production Expr_ValNegacion is entered.
func (s *BaseControlListener) EnterExpr_ValNegacion(ctx *Expr_ValNegacionContext) {}

// ExitExpr_ValNegacion is called when production Expr_ValNegacion is exited.
func (s *BaseControlListener) ExitExpr_ValNegacion(ctx *Expr_ValNegacionContext) {}

// EnterExpr_ID is called when production Expr_ID is entered.
func (s *BaseControlListener) EnterExpr_ID(ctx *Expr_IDContext) {}

// ExitExpr_ID is called when production Expr_ID is exited.
func (s *BaseControlListener) ExitExpr_ID(ctx *Expr_IDContext) {}

// EnterExpr_ValidaMayQue is called when production Expr_ValidaMayQue is entered.
func (s *BaseControlListener) EnterExpr_ValidaMayQue(ctx *Expr_ValidaMayQueContext) {}

// ExitExpr_ValidaMayQue is called when production Expr_ValidaMayQue is exited.
func (s *BaseControlListener) ExitExpr_ValidaMayQue(ctx *Expr_ValidaMayQueContext) {}

// EnterExpr_Conteo is called when production Expr_Conteo is entered.
func (s *BaseControlListener) EnterExpr_Conteo(ctx *Expr_ConteoContext) {}

// ExitExpr_Conteo is called when production Expr_Conteo is exited.
func (s *BaseControlListener) ExitExpr_Conteo(ctx *Expr_ConteoContext) {}

// EnterExpr_IsEmpty is called when production Expr_IsEmpty is entered.
func (s *BaseControlListener) EnterExpr_IsEmpty(ctx *Expr_IsEmptyContext) {}

// ExitExpr_IsEmpty is called when production Expr_IsEmpty is exited.
func (s *BaseControlListener) ExitExpr_IsEmpty(ctx *Expr_IsEmptyContext) {}

// EnterExpr_OpSumRes is called when production Expr_OpSumRes is entered.
func (s *BaseControlListener) EnterExpr_OpSumRes(ctx *Expr_OpSumResContext) {}

// ExitExpr_OpSumRes is called when production Expr_OpSumRes is exited.
func (s *BaseControlListener) ExitExpr_OpSumRes(ctx *Expr_OpSumResContext) {}

// EnterExpr_Entero is called when production Expr_Entero is entered.
func (s *BaseControlListener) EnterExpr_Entero(ctx *Expr_EnteroContext) {}

// ExitExpr_Entero is called when production Expr_Entero is exited.
func (s *BaseControlListener) ExitExpr_Entero(ctx *Expr_EnteroContext) {}

// EnterExpr_ParExpre is called when production Expr_ParExpre is entered.
func (s *BaseControlListener) EnterExpr_ParExpre(ctx *Expr_ParExpreContext) {}

// ExitExpr_ParExpre is called when production Expr_ParExpre is exited.
func (s *BaseControlListener) ExitExpr_ParExpre(ctx *Expr_ParExpreContext) {}

// EnterExpr_StringChar is called when production Expr_StringChar is entered.
func (s *BaseControlListener) EnterExpr_StringChar(ctx *Expr_StringCharContext) {}

// ExitExpr_StringChar is called when production Expr_StringChar is exited.
func (s *BaseControlListener) ExitExpr_StringChar(ctx *Expr_StringCharContext) {}

// EnterExpr_ValidaMenQue is called when production Expr_ValidaMenQue is entered.
func (s *BaseControlListener) EnterExpr_ValidaMenQue(ctx *Expr_ValidaMenQueContext) {}

// ExitExpr_ValidaMenQue is called when production Expr_ValidaMenQue is exited.
func (s *BaseControlListener) ExitExpr_ValidaMenQue(ctx *Expr_ValidaMenQueContext) {}

// EnterExpr_PosVector is called when production Expr_PosVector is entered.
func (s *BaseControlListener) EnterExpr_PosVector(ctx *Expr_PosVectorContext) {}

// ExitExpr_PosVector is called when production Expr_PosVector is exited.
func (s *BaseControlListener) ExitExpr_PosVector(ctx *Expr_PosVectorContext) {}

// EnterTipovariable is called when production tipovariable is entered.
func (s *BaseControlListener) EnterTipovariable(ctx *TipovariableContext) {}

// ExitTipovariable is called when production tipovariable is exited.
func (s *BaseControlListener) ExitTipovariable(ctx *TipovariableContext) {}
