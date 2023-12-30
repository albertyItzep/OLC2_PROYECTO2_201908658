// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import "github.com/antlr4-go/antlr/v4"

// ControlListener is a complete listener for a parse tree produced by ControlParser.
type ControlListener interface {
	antlr.ParseTreeListener

	// EnterProgInicial is called when entering the ProgInicial production.
	EnterProgInicial(c *ProgInicialContext)

	// EnterBlockContent is called when entering the BlockContent production.
	EnterBlockContent(c *BlockContentContext)

	// EnterSentenciaDeclaracion is called when entering the SentenciaDeclaracion production.
	EnterSentenciaDeclaracion(c *SentenciaDeclaracionContext)

	// EnterSentenciaDeclaracionVector is called when entering the SentenciaDeclaracionVector production.
	EnterSentenciaDeclaracionVector(c *SentenciaDeclaracionVectorContext)

	// EnterSentenciaDeclaracionMatrix is called when entering the SentenciaDeclaracionMatrix production.
	EnterSentenciaDeclaracionMatrix(c *SentenciaDeclaracionMatrixContext)

	// EnterSentenciaDeclaracionFuncion is called when entering the SentenciaDeclaracionFuncion production.
	EnterSentenciaDeclaracionFuncion(c *SentenciaDeclaracionFuncionContext)

	// EnterSentenciaAsignacion is called when entering the SentenciaAsignacion production.
	EnterSentenciaAsignacion(c *SentenciaAsignacionContext)

	// EnterSentenciaIf is called when entering the SentenciaIf production.
	EnterSentenciaIf(c *SentenciaIfContext)

	// EnterSentenciaSwitch is called when entering the SentenciaSwitch production.
	EnterSentenciaSwitch(c *SentenciaSwitchContext)

	// EnterSentenciaWhile is called when entering the SentenciaWhile production.
	EnterSentenciaWhile(c *SentenciaWhileContext)

	// EnterSentenciaFor is called when entering the SentenciaFor production.
	EnterSentenciaFor(c *SentenciaForContext)

	// EnterSentenciaGuard is called when entering the SentenciaGuard production.
	EnterSentenciaGuard(c *SentenciaGuardContext)

	// EnterSentenciaBreak is called when entering the SentenciaBreak production.
	EnterSentenciaBreak(c *SentenciaBreakContext)

	// EnterSentenciaContinue is called when entering the SentenciaContinue production.
	EnterSentenciaContinue(c *SentenciaContinueContext)

	// EnterSentenciaReturn is called when entering the SentenciaReturn production.
	EnterSentenciaReturn(c *SentenciaReturnContext)

	// EnterSentenciaFuncVectoriales is called when entering the SentenciaFuncVectoriales production.
	EnterSentenciaFuncVectoriales(c *SentenciaFuncVectorialesContext)

	// EnterSentenciaLLamadaFuncion is called when entering the SentenciaLLamadaFuncion production.
	EnterSentenciaLLamadaFuncion(c *SentenciaLLamadaFuncionContext)

	// EnterSentenciaPrint is called when entering the SentenciaPrint production.
	EnterSentenciaPrint(c *SentenciaPrintContext)

	// EnterSentenciaStruct is called when entering the SentenciaStruct production.
	EnterSentenciaStruct(c *SentenciaStructContext)

	// EnterSentenciaDeclaracionStruct is called when entering the SentenciaDeclaracionStruct production.
	EnterSentenciaDeclaracionStruct(c *SentenciaDeclaracionStructContext)

	// EnterSentenciaLLamadaMetodos is called when entering the SentenciaLLamadaMetodos production.
	EnterSentenciaLLamadaMetodos(c *SentenciaLLamadaMetodosContext)

	// EnterInstruccionIf is called when entering the InstruccionIf production.
	EnterInstruccionIf(c *InstruccionIfContext)

	// EnterInstruccionSwitch is called when entering the InstruccionSwitch production.
	EnterInstruccionSwitch(c *InstruccionSwitchContext)

	// EnterInstruccionCase is called when entering the InstruccionCase production.
	EnterInstruccionCase(c *InstruccionCaseContext)

	// EnterInstruccionDefault is called when entering the InstruccionDefault production.
	EnterInstruccionDefault(c *InstruccionDefaultContext)

	// EnterInstruccionWhile is called when entering the InstruccionWhile production.
	EnterInstruccionWhile(c *InstruccionWhileContext)

	// EnterInstruccionFor is called when entering the InstruccionFor production.
	EnterInstruccionFor(c *InstruccionForContext)

	// EnterInstruccionGuard is called when entering the InstruccionGuard production.
	EnterInstruccionGuard(c *InstruccionGuardContext)

	// EnterInstruccionBreak is called when entering the InstruccionBreak production.
	EnterInstruccionBreak(c *InstruccionBreakContext)

	// EnterInstruccionContinue is called when entering the InstruccionContinue production.
	EnterInstruccionContinue(c *InstruccionContinueContext)

	// EnterInstruccionReturnSimple is called when entering the InstruccionReturnSimple production.
	EnterInstruccionReturnSimple(c *InstruccionReturnSimpleContext)

	// EnterInstruccionReturnExpresion is called when entering the InstruccionReturnExpresion production.
	EnterInstruccionReturnExpresion(c *InstruccionReturnExpresionContext)

	// EnterDecVector_ExpresionLista is called when entering the DecVector_ExpresionLista production.
	EnterDecVector_ExpresionLista(c *DecVector_ExpresionListaContext)

	// EnterDecVector_ObjetLista is called when entering the DecVector_ObjetLista production.
	EnterDecVector_ObjetLista(c *DecVector_ObjetListaContext)

	// EnterDecVector_Id is called when entering the DecVector_Id production.
	EnterDecVector_Id(c *DecVector_IdContext)

	// EnterDecVector_ValCor is called when entering the DecVector_ValCor production.
	EnterDecVector_ValCor(c *DecVector_ValCorContext)

	// EnterDecVectorConst_ExpresionLista is called when entering the DecVectorConst_ExpresionLista production.
	EnterDecVectorConst_ExpresionLista(c *DecVectorConst_ExpresionListaContext)

	// EnterDecVectorConst_ObjetLista is called when entering the DecVectorConst_ObjetLista production.
	EnterDecVectorConst_ObjetLista(c *DecVectorConst_ObjetListaContext)

	// EnterDecVectorConst_Id is called when entering the DecVectorConst_Id production.
	EnterDecVectorConst_Id(c *DecVectorConst_IdContext)

	// EnterVectFunc_Append is called when entering the VectFunc_Append production.
	EnterVectFunc_Append(c *VectFunc_AppendContext)

	// EnterVectFunc_Remove is called when entering the VectFunc_Remove production.
	EnterVectFunc_Remove(c *VectFunc_RemoveContext)

	// EnterVectFunc_RemoveLast is called when entering the VectFunc_RemoveLast production.
	EnterVectFunc_RemoveLast(c *VectFunc_RemoveLastContext)

	// EnterDecMatriz is called when entering the decMatriz production.
	EnterDecMatriz(c *DecMatrizContext)

	// EnterTipoMatriz is called when entering the tipoMatriz production.
	EnterTipoMatriz(c *TipoMatrizContext)

	// EnterDefMatrix is called when entering the defMatrix production.
	EnterDefMatrix(c *DefMatrixContext)

	// EnterListaValoresMat is called when entering the listaValoresMat production.
	EnterListaValoresMat(c *ListaValoresMatContext)

	// EnterListaValoresMat2 is called when entering the listaValoresMat2 production.
	EnterListaValoresMat2(c *ListaValoresMat2Context)

	// EnterSimpleMax is called when entering the simpleMax production.
	EnterSimpleMax(c *SimpleMaxContext)

	// EnterInstStruct is called when entering the instStruct production.
	EnterInstStruct(c *InstStructContext)

	// EnterListaAtributos is called when entering the listaAtributos production.
	EnterListaAtributos(c *ListaAtributosContext)

	// EnterDecStruct is called when entering the decStruct production.
	EnterDecStruct(c *DecStructContext)

	// EnterLDupla is called when entering the lDupla production.
	EnterLDupla(c *LDuplaContext)

	// EnterLlamAtributos is called when entering the llamAtributos production.
	EnterLlamAtributos(c *LlamAtributosContext)

	// EnterLlamadaMetodos is called when entering the llamadaMetodos production.
	EnterLlamadaMetodos(c *LlamadaMetodosContext)

	// EnterDec_Funcion is called when entering the Dec_Funcion production.
	EnterDec_Funcion(c *Dec_FuncionContext)

	// EnterListaParametros is called when entering the listaParametros production.
	EnterListaParametros(c *ListaParametrosContext)

	// EnterLlamadaFuncion is called when entering the llamadaFuncion production.
	EnterLlamadaFuncion(c *LlamadaFuncionContext)

	// EnterListaLlamadaParametros is called when entering the listaLlamadaParametros production.
	EnterListaLlamadaParametros(c *ListaLlamadaParametrosContext)

	// EnterIntruccionPrint is called when entering the IntruccionPrint production.
	EnterIntruccionPrint(c *IntruccionPrintContext)

	// EnterInst_Casteos is called when entering the Inst_Casteos production.
	EnterInst_Casteos(c *Inst_CasteosContext)

	// EnterListaExpresiones is called when entering the ListaExpresiones production.
	EnterListaExpresiones(c *ListaExpresionesContext)

	// EnterObjectsList is called when entering the objectsList production.
	EnterObjectsList(c *ObjectsListContext)

	// EnterDec_VTipoValor is called when entering the Dec_VTipoValor production.
	EnterDec_VTipoValor(c *Dec_VTipoValorContext)

	// EnterDec_VTipo is called when entering the Dec_VTipo production.
	EnterDec_VTipo(c *Dec_VTipoContext)

	// EnterDec_VVal is called when entering the Dec_VVal production.
	EnterDec_VVal(c *Dec_VValContext)

	// EnterDec_CVal is called when entering the Dec_CVal production.
	EnterDec_CVal(c *Dec_CValContext)

	// EnterDec_CTipoValor is called when entering the Dec_CTipoValor production.
	EnterDec_CTipoValor(c *Dec_CTipoValorContext)

	// EnterAsignacion_Aumento is called when entering the Asignacion_Aumento production.
	EnterAsignacion_Aumento(c *Asignacion_AumentoContext)

	// EnterAsignacion_Decremento is called when entering the Asignacion_Decremento production.
	EnterAsignacion_Decremento(c *Asignacion_DecrementoContext)

	// EnterAsignacion_ValorGen is called when entering the Asignacion_ValorGen production.
	EnterAsignacion_ValorGen(c *Asignacion_ValorGenContext)

	// EnterAsignacion_VectorAumento is called when entering the Asignacion_VectorAumento production.
	EnterAsignacion_VectorAumento(c *Asignacion_VectorAumentoContext)

	// EnterAsignacion_VectorDecremento is called when entering the Asignacion_VectorDecremento production.
	EnterAsignacion_VectorDecremento(c *Asignacion_VectorDecrementoContext)

	// EnterAsignacion_MatrixAumento is called when entering the Asignacion_MatrixAumento production.
	EnterAsignacion_MatrixAumento(c *Asignacion_MatrixAumentoContext)

	// EnterAsignacion_MatrixDecremento is called when entering the Asignacion_MatrixDecremento production.
	EnterAsignacion_MatrixDecremento(c *Asignacion_MatrixDecrementoContext)

	// EnterAsignacion_VectorGeneral is called when entering the Asignacion_VectorGeneral production.
	EnterAsignacion_VectorGeneral(c *Asignacion_VectorGeneralContext)

	// EnterAsignacion_MatrixGeneral is called when entering the Asignacion_MatrixGeneral production.
	EnterAsignacion_MatrixGeneral(c *Asignacion_MatrixGeneralContext)

	// EnterAsignacion_LlAtribGeneral is called when entering the Asignacion_LlAtribGeneral production.
	EnterAsignacion_LlAtribGeneral(c *Asignacion_LlAtribGeneralContext)

	// EnterAsignacion_LlAtribAumento is called when entering the Asignacion_LlAtribAumento production.
	EnterAsignacion_LlAtribAumento(c *Asignacion_LlAtribAumentoContext)

	// EnterAsignacion_LlAtribDecremento is called when entering the Asignacion_LlAtribDecremento production.
	EnterAsignacion_LlAtribDecremento(c *Asignacion_LlAtribDecrementoContext)

	// EnterExpr_Decimal is called when entering the Expr_Decimal production.
	EnterExpr_Decimal(c *Expr_DecimalContext)

	// EnterExpr_InstCasteo is called when entering the Expr_InstCasteo production.
	EnterExpr_InstCasteo(c *Expr_InstCasteoContext)

	// EnterExpr_ValidaIgualDif is called when entering the Expr_ValidaIgualDif production.
	EnterExpr_ValidaIgualDif(c *Expr_ValidaIgualDifContext)

	// EnterExpr_OpMulDiv is called when entering the Expr_OpMulDiv production.
	EnterExpr_OpMulDiv(c *Expr_OpMulDivContext)

	// EnterExpr_ValidaAnd is called when entering the Expr_ValidaAnd production.
	EnterExpr_ValidaAnd(c *Expr_ValidaAndContext)

	// EnterExpr_LlamAtributos is called when entering the Expr_LlamAtributos production.
	EnterExpr_LlamAtributos(c *Expr_LlamAtributosContext)

	// EnterExpr_Nil is called when entering the Expr_Nil production.
	EnterExpr_Nil(c *Expr_NilContext)

	// EnterExpr_PosMatrix is called when entering the Expr_PosMatrix production.
	EnterExpr_PosMatrix(c *Expr_PosMatrixContext)

	// EnterExpr_Boolean is called when entering the Expr_Boolean production.
	EnterExpr_Boolean(c *Expr_BooleanContext)

	// EnterExpr_InstRango is called when entering the Expr_InstRango production.
	EnterExpr_InstRango(c *Expr_InstRangoContext)

	// EnterExpr_ValNumNeg is called when entering the Expr_ValNumNeg production.
	EnterExpr_ValNumNeg(c *Expr_ValNumNegContext)

	// EnterExpr_ValidaOr is called when entering the Expr_ValidaOr production.
	EnterExpr_ValidaOr(c *Expr_ValidaOrContext)

	// EnterExpr_ValNegacion is called when entering the Expr_ValNegacion production.
	EnterExpr_ValNegacion(c *Expr_ValNegacionContext)

	// EnterExpr_ID is called when entering the Expr_ID production.
	EnterExpr_ID(c *Expr_IDContext)

	// EnterExpr_ValidaMayQue is called when entering the Expr_ValidaMayQue production.
	EnterExpr_ValidaMayQue(c *Expr_ValidaMayQueContext)

	// EnterExpr_Conteo is called when entering the Expr_Conteo production.
	EnterExpr_Conteo(c *Expr_ConteoContext)

	// EnterExpr_IsEmpty is called when entering the Expr_IsEmpty production.
	EnterExpr_IsEmpty(c *Expr_IsEmptyContext)

	// EnterExpr_OpSumRes is called when entering the Expr_OpSumRes production.
	EnterExpr_OpSumRes(c *Expr_OpSumResContext)

	// EnterExpr_Entero is called when entering the Expr_Entero production.
	EnterExpr_Entero(c *Expr_EnteroContext)

	// EnterExpr_ParExpre is called when entering the Expr_ParExpre production.
	EnterExpr_ParExpre(c *Expr_ParExpreContext)

	// EnterExpr_StringChar is called when entering the Expr_StringChar production.
	EnterExpr_StringChar(c *Expr_StringCharContext)

	// EnterExpr_ValidaMenQue is called when entering the Expr_ValidaMenQue production.
	EnterExpr_ValidaMenQue(c *Expr_ValidaMenQueContext)

	// EnterExpr_PosVector is called when entering the Expr_PosVector production.
	EnterExpr_PosVector(c *Expr_PosVectorContext)

	// EnterTipovariable is called when entering the tipovariable production.
	EnterTipovariable(c *TipovariableContext)

	// ExitProgInicial is called when exiting the ProgInicial production.
	ExitProgInicial(c *ProgInicialContext)

	// ExitBlockContent is called when exiting the BlockContent production.
	ExitBlockContent(c *BlockContentContext)

	// ExitSentenciaDeclaracion is called when exiting the SentenciaDeclaracion production.
	ExitSentenciaDeclaracion(c *SentenciaDeclaracionContext)

	// ExitSentenciaDeclaracionVector is called when exiting the SentenciaDeclaracionVector production.
	ExitSentenciaDeclaracionVector(c *SentenciaDeclaracionVectorContext)

	// ExitSentenciaDeclaracionMatrix is called when exiting the SentenciaDeclaracionMatrix production.
	ExitSentenciaDeclaracionMatrix(c *SentenciaDeclaracionMatrixContext)

	// ExitSentenciaDeclaracionFuncion is called when exiting the SentenciaDeclaracionFuncion production.
	ExitSentenciaDeclaracionFuncion(c *SentenciaDeclaracionFuncionContext)

	// ExitSentenciaAsignacion is called when exiting the SentenciaAsignacion production.
	ExitSentenciaAsignacion(c *SentenciaAsignacionContext)

	// ExitSentenciaIf is called when exiting the SentenciaIf production.
	ExitSentenciaIf(c *SentenciaIfContext)

	// ExitSentenciaSwitch is called when exiting the SentenciaSwitch production.
	ExitSentenciaSwitch(c *SentenciaSwitchContext)

	// ExitSentenciaWhile is called when exiting the SentenciaWhile production.
	ExitSentenciaWhile(c *SentenciaWhileContext)

	// ExitSentenciaFor is called when exiting the SentenciaFor production.
	ExitSentenciaFor(c *SentenciaForContext)

	// ExitSentenciaGuard is called when exiting the SentenciaGuard production.
	ExitSentenciaGuard(c *SentenciaGuardContext)

	// ExitSentenciaBreak is called when exiting the SentenciaBreak production.
	ExitSentenciaBreak(c *SentenciaBreakContext)

	// ExitSentenciaContinue is called when exiting the SentenciaContinue production.
	ExitSentenciaContinue(c *SentenciaContinueContext)

	// ExitSentenciaReturn is called when exiting the SentenciaReturn production.
	ExitSentenciaReturn(c *SentenciaReturnContext)

	// ExitSentenciaFuncVectoriales is called when exiting the SentenciaFuncVectoriales production.
	ExitSentenciaFuncVectoriales(c *SentenciaFuncVectorialesContext)

	// ExitSentenciaLLamadaFuncion is called when exiting the SentenciaLLamadaFuncion production.
	ExitSentenciaLLamadaFuncion(c *SentenciaLLamadaFuncionContext)

	// ExitSentenciaPrint is called when exiting the SentenciaPrint production.
	ExitSentenciaPrint(c *SentenciaPrintContext)

	// ExitSentenciaStruct is called when exiting the SentenciaStruct production.
	ExitSentenciaStruct(c *SentenciaStructContext)

	// ExitSentenciaDeclaracionStruct is called when exiting the SentenciaDeclaracionStruct production.
	ExitSentenciaDeclaracionStruct(c *SentenciaDeclaracionStructContext)

	// ExitSentenciaLLamadaMetodos is called when exiting the SentenciaLLamadaMetodos production.
	ExitSentenciaLLamadaMetodos(c *SentenciaLLamadaMetodosContext)

	// ExitInstruccionIf is called when exiting the InstruccionIf production.
	ExitInstruccionIf(c *InstruccionIfContext)

	// ExitInstruccionSwitch is called when exiting the InstruccionSwitch production.
	ExitInstruccionSwitch(c *InstruccionSwitchContext)

	// ExitInstruccionCase is called when exiting the InstruccionCase production.
	ExitInstruccionCase(c *InstruccionCaseContext)

	// ExitInstruccionDefault is called when exiting the InstruccionDefault production.
	ExitInstruccionDefault(c *InstruccionDefaultContext)

	// ExitInstruccionWhile is called when exiting the InstruccionWhile production.
	ExitInstruccionWhile(c *InstruccionWhileContext)

	// ExitInstruccionFor is called when exiting the InstruccionFor production.
	ExitInstruccionFor(c *InstruccionForContext)

	// ExitInstruccionGuard is called when exiting the InstruccionGuard production.
	ExitInstruccionGuard(c *InstruccionGuardContext)

	// ExitInstruccionBreak is called when exiting the InstruccionBreak production.
	ExitInstruccionBreak(c *InstruccionBreakContext)

	// ExitInstruccionContinue is called when exiting the InstruccionContinue production.
	ExitInstruccionContinue(c *InstruccionContinueContext)

	// ExitInstruccionReturnSimple is called when exiting the InstruccionReturnSimple production.
	ExitInstruccionReturnSimple(c *InstruccionReturnSimpleContext)

	// ExitInstruccionReturnExpresion is called when exiting the InstruccionReturnExpresion production.
	ExitInstruccionReturnExpresion(c *InstruccionReturnExpresionContext)

	// ExitDecVector_ExpresionLista is called when exiting the DecVector_ExpresionLista production.
	ExitDecVector_ExpresionLista(c *DecVector_ExpresionListaContext)

	// ExitDecVector_ObjetLista is called when exiting the DecVector_ObjetLista production.
	ExitDecVector_ObjetLista(c *DecVector_ObjetListaContext)

	// ExitDecVector_Id is called when exiting the DecVector_Id production.
	ExitDecVector_Id(c *DecVector_IdContext)

	// ExitDecVector_ValCor is called when exiting the DecVector_ValCor production.
	ExitDecVector_ValCor(c *DecVector_ValCorContext)

	// ExitDecVectorConst_ExpresionLista is called when exiting the DecVectorConst_ExpresionLista production.
	ExitDecVectorConst_ExpresionLista(c *DecVectorConst_ExpresionListaContext)

	// ExitDecVectorConst_ObjetLista is called when exiting the DecVectorConst_ObjetLista production.
	ExitDecVectorConst_ObjetLista(c *DecVectorConst_ObjetListaContext)

	// ExitDecVectorConst_Id is called when exiting the DecVectorConst_Id production.
	ExitDecVectorConst_Id(c *DecVectorConst_IdContext)

	// ExitVectFunc_Append is called when exiting the VectFunc_Append production.
	ExitVectFunc_Append(c *VectFunc_AppendContext)

	// ExitVectFunc_Remove is called when exiting the VectFunc_Remove production.
	ExitVectFunc_Remove(c *VectFunc_RemoveContext)

	// ExitVectFunc_RemoveLast is called when exiting the VectFunc_RemoveLast production.
	ExitVectFunc_RemoveLast(c *VectFunc_RemoveLastContext)

	// ExitDecMatriz is called when exiting the decMatriz production.
	ExitDecMatriz(c *DecMatrizContext)

	// ExitTipoMatriz is called when exiting the tipoMatriz production.
	ExitTipoMatriz(c *TipoMatrizContext)

	// ExitDefMatrix is called when exiting the defMatrix production.
	ExitDefMatrix(c *DefMatrixContext)

	// ExitListaValoresMat is called when exiting the listaValoresMat production.
	ExitListaValoresMat(c *ListaValoresMatContext)

	// ExitListaValoresMat2 is called when exiting the listaValoresMat2 production.
	ExitListaValoresMat2(c *ListaValoresMat2Context)

	// ExitSimpleMax is called when exiting the simpleMax production.
	ExitSimpleMax(c *SimpleMaxContext)

	// ExitInstStruct is called when exiting the instStruct production.
	ExitInstStruct(c *InstStructContext)

	// ExitListaAtributos is called when exiting the listaAtributos production.
	ExitListaAtributos(c *ListaAtributosContext)

	// ExitDecStruct is called when exiting the decStruct production.
	ExitDecStruct(c *DecStructContext)

	// ExitLDupla is called when exiting the lDupla production.
	ExitLDupla(c *LDuplaContext)

	// ExitLlamAtributos is called when exiting the llamAtributos production.
	ExitLlamAtributos(c *LlamAtributosContext)

	// ExitLlamadaMetodos is called when exiting the llamadaMetodos production.
	ExitLlamadaMetodos(c *LlamadaMetodosContext)

	// ExitDec_Funcion is called when exiting the Dec_Funcion production.
	ExitDec_Funcion(c *Dec_FuncionContext)

	// ExitListaParametros is called when exiting the listaParametros production.
	ExitListaParametros(c *ListaParametrosContext)

	// ExitLlamadaFuncion is called when exiting the llamadaFuncion production.
	ExitLlamadaFuncion(c *LlamadaFuncionContext)

	// ExitListaLlamadaParametros is called when exiting the listaLlamadaParametros production.
	ExitListaLlamadaParametros(c *ListaLlamadaParametrosContext)

	// ExitIntruccionPrint is called when exiting the IntruccionPrint production.
	ExitIntruccionPrint(c *IntruccionPrintContext)

	// ExitInst_Casteos is called when exiting the Inst_Casteos production.
	ExitInst_Casteos(c *Inst_CasteosContext)

	// ExitListaExpresiones is called when exiting the ListaExpresiones production.
	ExitListaExpresiones(c *ListaExpresionesContext)

	// ExitObjectsList is called when exiting the objectsList production.
	ExitObjectsList(c *ObjectsListContext)

	// ExitDec_VTipoValor is called when exiting the Dec_VTipoValor production.
	ExitDec_VTipoValor(c *Dec_VTipoValorContext)

	// ExitDec_VTipo is called when exiting the Dec_VTipo production.
	ExitDec_VTipo(c *Dec_VTipoContext)

	// ExitDec_VVal is called when exiting the Dec_VVal production.
	ExitDec_VVal(c *Dec_VValContext)

	// ExitDec_CVal is called when exiting the Dec_CVal production.
	ExitDec_CVal(c *Dec_CValContext)

	// ExitDec_CTipoValor is called when exiting the Dec_CTipoValor production.
	ExitDec_CTipoValor(c *Dec_CTipoValorContext)

	// ExitAsignacion_Aumento is called when exiting the Asignacion_Aumento production.
	ExitAsignacion_Aumento(c *Asignacion_AumentoContext)

	// ExitAsignacion_Decremento is called when exiting the Asignacion_Decremento production.
	ExitAsignacion_Decremento(c *Asignacion_DecrementoContext)

	// ExitAsignacion_ValorGen is called when exiting the Asignacion_ValorGen production.
	ExitAsignacion_ValorGen(c *Asignacion_ValorGenContext)

	// ExitAsignacion_VectorAumento is called when exiting the Asignacion_VectorAumento production.
	ExitAsignacion_VectorAumento(c *Asignacion_VectorAumentoContext)

	// ExitAsignacion_VectorDecremento is called when exiting the Asignacion_VectorDecremento production.
	ExitAsignacion_VectorDecremento(c *Asignacion_VectorDecrementoContext)

	// ExitAsignacion_MatrixAumento is called when exiting the Asignacion_MatrixAumento production.
	ExitAsignacion_MatrixAumento(c *Asignacion_MatrixAumentoContext)

	// ExitAsignacion_MatrixDecremento is called when exiting the Asignacion_MatrixDecremento production.
	ExitAsignacion_MatrixDecremento(c *Asignacion_MatrixDecrementoContext)

	// ExitAsignacion_VectorGeneral is called when exiting the Asignacion_VectorGeneral production.
	ExitAsignacion_VectorGeneral(c *Asignacion_VectorGeneralContext)

	// ExitAsignacion_MatrixGeneral is called when exiting the Asignacion_MatrixGeneral production.
	ExitAsignacion_MatrixGeneral(c *Asignacion_MatrixGeneralContext)

	// ExitAsignacion_LlAtribGeneral is called when exiting the Asignacion_LlAtribGeneral production.
	ExitAsignacion_LlAtribGeneral(c *Asignacion_LlAtribGeneralContext)

	// ExitAsignacion_LlAtribAumento is called when exiting the Asignacion_LlAtribAumento production.
	ExitAsignacion_LlAtribAumento(c *Asignacion_LlAtribAumentoContext)

	// ExitAsignacion_LlAtribDecremento is called when exiting the Asignacion_LlAtribDecremento production.
	ExitAsignacion_LlAtribDecremento(c *Asignacion_LlAtribDecrementoContext)

	// ExitExpr_Decimal is called when exiting the Expr_Decimal production.
	ExitExpr_Decimal(c *Expr_DecimalContext)

	// ExitExpr_InstCasteo is called when exiting the Expr_InstCasteo production.
	ExitExpr_InstCasteo(c *Expr_InstCasteoContext)

	// ExitExpr_ValidaIgualDif is called when exiting the Expr_ValidaIgualDif production.
	ExitExpr_ValidaIgualDif(c *Expr_ValidaIgualDifContext)

	// ExitExpr_OpMulDiv is called when exiting the Expr_OpMulDiv production.
	ExitExpr_OpMulDiv(c *Expr_OpMulDivContext)

	// ExitExpr_ValidaAnd is called when exiting the Expr_ValidaAnd production.
	ExitExpr_ValidaAnd(c *Expr_ValidaAndContext)

	// ExitExpr_LlamAtributos is called when exiting the Expr_LlamAtributos production.
	ExitExpr_LlamAtributos(c *Expr_LlamAtributosContext)

	// ExitExpr_Nil is called when exiting the Expr_Nil production.
	ExitExpr_Nil(c *Expr_NilContext)

	// ExitExpr_PosMatrix is called when exiting the Expr_PosMatrix production.
	ExitExpr_PosMatrix(c *Expr_PosMatrixContext)

	// ExitExpr_Boolean is called when exiting the Expr_Boolean production.
	ExitExpr_Boolean(c *Expr_BooleanContext)

	// ExitExpr_InstRango is called when exiting the Expr_InstRango production.
	ExitExpr_InstRango(c *Expr_InstRangoContext)

	// ExitExpr_ValNumNeg is called when exiting the Expr_ValNumNeg production.
	ExitExpr_ValNumNeg(c *Expr_ValNumNegContext)

	// ExitExpr_ValidaOr is called when exiting the Expr_ValidaOr production.
	ExitExpr_ValidaOr(c *Expr_ValidaOrContext)

	// ExitExpr_ValNegacion is called when exiting the Expr_ValNegacion production.
	ExitExpr_ValNegacion(c *Expr_ValNegacionContext)

	// ExitExpr_ID is called when exiting the Expr_ID production.
	ExitExpr_ID(c *Expr_IDContext)

	// ExitExpr_ValidaMayQue is called when exiting the Expr_ValidaMayQue production.
	ExitExpr_ValidaMayQue(c *Expr_ValidaMayQueContext)

	// ExitExpr_Conteo is called when exiting the Expr_Conteo production.
	ExitExpr_Conteo(c *Expr_ConteoContext)

	// ExitExpr_IsEmpty is called when exiting the Expr_IsEmpty production.
	ExitExpr_IsEmpty(c *Expr_IsEmptyContext)

	// ExitExpr_OpSumRes is called when exiting the Expr_OpSumRes production.
	ExitExpr_OpSumRes(c *Expr_OpSumResContext)

	// ExitExpr_Entero is called when exiting the Expr_Entero production.
	ExitExpr_Entero(c *Expr_EnteroContext)

	// ExitExpr_ParExpre is called when exiting the Expr_ParExpre production.
	ExitExpr_ParExpre(c *Expr_ParExpreContext)

	// ExitExpr_StringChar is called when exiting the Expr_StringChar production.
	ExitExpr_StringChar(c *Expr_StringCharContext)

	// ExitExpr_ValidaMenQue is called when exiting the Expr_ValidaMenQue production.
	ExitExpr_ValidaMenQue(c *Expr_ValidaMenQueContext)

	// ExitExpr_PosVector is called when exiting the Expr_PosVector production.
	ExitExpr_PosVector(c *Expr_PosVectorContext)

	// ExitTipovariable is called when exiting the tipovariable production.
	ExitTipovariable(c *TipovariableContext)
}
