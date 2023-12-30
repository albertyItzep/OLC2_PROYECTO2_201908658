// Generated from Control.g4 by ANTLR 4.13.0
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ControlParser}.
 */
public interface ControlListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the {@code ProgInicial}
	 * labeled alternative in {@link ControlParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProgInicial(ControlParser.ProgInicialContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ProgInicial}
	 * labeled alternative in {@link ControlParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProgInicial(ControlParser.ProgInicialContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BlockContent}
	 * labeled alternative in {@link ControlParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlockContent(ControlParser.BlockContentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BlockContent}
	 * labeled alternative in {@link ControlParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlockContent(ControlParser.BlockContentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracion(ControlParser.SentenciaDeclaracionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracion(ControlParser.SentenciaDeclaracionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracionVector}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracionVector(ControlParser.SentenciaDeclaracionVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracionVector}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracionVector(ControlParser.SentenciaDeclaracionVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracionMatrix}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracionMatrix(ControlParser.SentenciaDeclaracionMatrixContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracionMatrix}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracionMatrix(ControlParser.SentenciaDeclaracionMatrixContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracionFuncion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracionFuncion(ControlParser.SentenciaDeclaracionFuncionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracionFuncion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracionFuncion(ControlParser.SentenciaDeclaracionFuncionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaAsignacion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaAsignacion(ControlParser.SentenciaAsignacionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaAsignacion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaAsignacion(ControlParser.SentenciaAsignacionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaIf}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaIf(ControlParser.SentenciaIfContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaIf}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaIf(ControlParser.SentenciaIfContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaSwitch}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaSwitch(ControlParser.SentenciaSwitchContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaSwitch}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaSwitch(ControlParser.SentenciaSwitchContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaWhile}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaWhile(ControlParser.SentenciaWhileContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaWhile}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaWhile(ControlParser.SentenciaWhileContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaFor}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaFor(ControlParser.SentenciaForContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaFor}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaFor(ControlParser.SentenciaForContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaGuard}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaGuard(ControlParser.SentenciaGuardContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaGuard}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaGuard(ControlParser.SentenciaGuardContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaBreak}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaBreak(ControlParser.SentenciaBreakContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaBreak}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaBreak(ControlParser.SentenciaBreakContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaContinue}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaContinue(ControlParser.SentenciaContinueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaContinue}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaContinue(ControlParser.SentenciaContinueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaReturn}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaReturn(ControlParser.SentenciaReturnContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaReturn}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaReturn(ControlParser.SentenciaReturnContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaFuncVectoriales}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaFuncVectoriales(ControlParser.SentenciaFuncVectorialesContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaFuncVectoriales}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaFuncVectoriales(ControlParser.SentenciaFuncVectorialesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaLLamadaFuncion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaLLamadaFuncion(ControlParser.SentenciaLLamadaFuncionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaLLamadaFuncion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaLLamadaFuncion(ControlParser.SentenciaLLamadaFuncionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaPrint}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaPrint(ControlParser.SentenciaPrintContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaPrint}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaPrint(ControlParser.SentenciaPrintContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaStruct}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaStruct(ControlParser.SentenciaStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaStruct}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaStruct(ControlParser.SentenciaStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracionStruct}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracionStruct(ControlParser.SentenciaDeclaracionStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracionStruct}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracionStruct(ControlParser.SentenciaDeclaracionStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaLLamadaMetodos}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaLLamadaMetodos(ControlParser.SentenciaLLamadaMetodosContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaLLamadaMetodos}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaLLamadaMetodos(ControlParser.SentenciaLLamadaMetodosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionIf}
	 * labeled alternative in {@link ControlParser#insIf}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionIf(ControlParser.InstruccionIfContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionIf}
	 * labeled alternative in {@link ControlParser#insIf}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionIf(ControlParser.InstruccionIfContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionSwitch}
	 * labeled alternative in {@link ControlParser#insSwitch}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionSwitch(ControlParser.InstruccionSwitchContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionSwitch}
	 * labeled alternative in {@link ControlParser#insSwitch}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionSwitch(ControlParser.InstruccionSwitchContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionCase}
	 * labeled alternative in {@link ControlParser#instCase}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionCase(ControlParser.InstruccionCaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionCase}
	 * labeled alternative in {@link ControlParser#instCase}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionCase(ControlParser.InstruccionCaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionDefault}
	 * labeled alternative in {@link ControlParser#instDefault}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionDefault(ControlParser.InstruccionDefaultContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionDefault}
	 * labeled alternative in {@link ControlParser#instDefault}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionDefault(ControlParser.InstruccionDefaultContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionWhile}
	 * labeled alternative in {@link ControlParser#instWhile}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionWhile(ControlParser.InstruccionWhileContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionWhile}
	 * labeled alternative in {@link ControlParser#instWhile}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionWhile(ControlParser.InstruccionWhileContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionFor}
	 * labeled alternative in {@link ControlParser#instFor}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionFor(ControlParser.InstruccionForContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionFor}
	 * labeled alternative in {@link ControlParser#instFor}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionFor(ControlParser.InstruccionForContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionGuard}
	 * labeled alternative in {@link ControlParser#instGuard}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionGuard(ControlParser.InstruccionGuardContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionGuard}
	 * labeled alternative in {@link ControlParser#instGuard}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionGuard(ControlParser.InstruccionGuardContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionBreak}
	 * labeled alternative in {@link ControlParser#instBreak}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionBreak(ControlParser.InstruccionBreakContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionBreak}
	 * labeled alternative in {@link ControlParser#instBreak}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionBreak(ControlParser.InstruccionBreakContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionContinue}
	 * labeled alternative in {@link ControlParser#instContinue}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionContinue(ControlParser.InstruccionContinueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionContinue}
	 * labeled alternative in {@link ControlParser#instContinue}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionContinue(ControlParser.InstruccionContinueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionReturnSimple}
	 * labeled alternative in {@link ControlParser#instReturn}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionReturnSimple(ControlParser.InstruccionReturnSimpleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionReturnSimple}
	 * labeled alternative in {@link ControlParser#instReturn}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionReturnSimple(ControlParser.InstruccionReturnSimpleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InstruccionReturnExpresion}
	 * labeled alternative in {@link ControlParser#instReturn}.
	 * @param ctx the parse tree
	 */
	void enterInstruccionReturnExpresion(ControlParser.InstruccionReturnExpresionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InstruccionReturnExpresion}
	 * labeled alternative in {@link ControlParser#instReturn}.
	 * @param ctx the parse tree
	 */
	void exitInstruccionReturnExpresion(ControlParser.InstruccionReturnExpresionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DecVector_ExpresionLista}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void enterDecVector_ExpresionLista(ControlParser.DecVector_ExpresionListaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DecVector_ExpresionLista}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void exitDecVector_ExpresionLista(ControlParser.DecVector_ExpresionListaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DecVector_ObjetLista}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void enterDecVector_ObjetLista(ControlParser.DecVector_ObjetListaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DecVector_ObjetLista}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void exitDecVector_ObjetLista(ControlParser.DecVector_ObjetListaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DecVector_Id}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void enterDecVector_Id(ControlParser.DecVector_IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DecVector_Id}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void exitDecVector_Id(ControlParser.DecVector_IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DecVector_ValCor}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void enterDecVector_ValCor(ControlParser.DecVector_ValCorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DecVector_ValCor}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void exitDecVector_ValCor(ControlParser.DecVector_ValCorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DecVectorConst_ExpresionLista}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void enterDecVectorConst_ExpresionLista(ControlParser.DecVectorConst_ExpresionListaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DecVectorConst_ExpresionLista}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void exitDecVectorConst_ExpresionLista(ControlParser.DecVectorConst_ExpresionListaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DecVectorConst_ObjetLista}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void enterDecVectorConst_ObjetLista(ControlParser.DecVectorConst_ObjetListaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DecVectorConst_ObjetLista}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void exitDecVectorConst_ObjetLista(ControlParser.DecVectorConst_ObjetListaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DecVectorConst_Id}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void enterDecVectorConst_Id(ControlParser.DecVectorConst_IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DecVectorConst_Id}
	 * labeled alternative in {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void exitDecVectorConst_Id(ControlParser.DecVectorConst_IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectFunc_Append}
	 * labeled alternative in {@link ControlParser#funcsVectoriales}.
	 * @param ctx the parse tree
	 */
	void enterVectFunc_Append(ControlParser.VectFunc_AppendContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectFunc_Append}
	 * labeled alternative in {@link ControlParser#funcsVectoriales}.
	 * @param ctx the parse tree
	 */
	void exitVectFunc_Append(ControlParser.VectFunc_AppendContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectFunc_Remove}
	 * labeled alternative in {@link ControlParser#funcsVectoriales}.
	 * @param ctx the parse tree
	 */
	void enterVectFunc_Remove(ControlParser.VectFunc_RemoveContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectFunc_Remove}
	 * labeled alternative in {@link ControlParser#funcsVectoriales}.
	 * @param ctx the parse tree
	 */
	void exitVectFunc_Remove(ControlParser.VectFunc_RemoveContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectFunc_RemoveLast}
	 * labeled alternative in {@link ControlParser#funcsVectoriales}.
	 * @param ctx the parse tree
	 */
	void enterVectFunc_RemoveLast(ControlParser.VectFunc_RemoveLastContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectFunc_RemoveLast}
	 * labeled alternative in {@link ControlParser#funcsVectoriales}.
	 * @param ctx the parse tree
	 */
	void exitVectFunc_RemoveLast(ControlParser.VectFunc_RemoveLastContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#decMatriz}.
	 * @param ctx the parse tree
	 */
	void enterDecMatriz(ControlParser.DecMatrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#decMatriz}.
	 * @param ctx the parse tree
	 */
	void exitDecMatriz(ControlParser.DecMatrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#tipoMatriz}.
	 * @param ctx the parse tree
	 */
	void enterTipoMatriz(ControlParser.TipoMatrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#tipoMatriz}.
	 * @param ctx the parse tree
	 */
	void exitTipoMatriz(ControlParser.TipoMatrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#defMatrix}.
	 * @param ctx the parse tree
	 */
	void enterDefMatrix(ControlParser.DefMatrixContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#defMatrix}.
	 * @param ctx the parse tree
	 */
	void exitDefMatrix(ControlParser.DefMatrixContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaValoresMat}.
	 * @param ctx the parse tree
	 */
	void enterListaValoresMat(ControlParser.ListaValoresMatContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaValoresMat}.
	 * @param ctx the parse tree
	 */
	void exitListaValoresMat(ControlParser.ListaValoresMatContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaValoresMat2}.
	 * @param ctx the parse tree
	 */
	void enterListaValoresMat2(ControlParser.ListaValoresMat2Context ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaValoresMat2}.
	 * @param ctx the parse tree
	 */
	void exitListaValoresMat2(ControlParser.ListaValoresMat2Context ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#simpleMax}.
	 * @param ctx the parse tree
	 */
	void enterSimpleMax(ControlParser.SimpleMaxContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#simpleMax}.
	 * @param ctx the parse tree
	 */
	void exitSimpleMax(ControlParser.SimpleMaxContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instStruct}.
	 * @param ctx the parse tree
	 */
	void enterInstStruct(ControlParser.InstStructContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instStruct}.
	 * @param ctx the parse tree
	 */
	void exitInstStruct(ControlParser.InstStructContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaAtributos}.
	 * @param ctx the parse tree
	 */
	void enterListaAtributos(ControlParser.ListaAtributosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaAtributos}.
	 * @param ctx the parse tree
	 */
	void exitListaAtributos(ControlParser.ListaAtributosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#decStruct}.
	 * @param ctx the parse tree
	 */
	void enterDecStruct(ControlParser.DecStructContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#decStruct}.
	 * @param ctx the parse tree
	 */
	void exitDecStruct(ControlParser.DecStructContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#lDupla}.
	 * @param ctx the parse tree
	 */
	void enterLDupla(ControlParser.LDuplaContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#lDupla}.
	 * @param ctx the parse tree
	 */
	void exitLDupla(ControlParser.LDuplaContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#llamAtributos}.
	 * @param ctx the parse tree
	 */
	void enterLlamAtributos(ControlParser.LlamAtributosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#llamAtributos}.
	 * @param ctx the parse tree
	 */
	void exitLlamAtributos(ControlParser.LlamAtributosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#llamadaMetodos}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaMetodos(ControlParser.LlamadaMetodosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#llamadaMetodos}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaMetodos(ControlParser.LlamadaMetodosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_Funcion}
	 * labeled alternative in {@link ControlParser#decFunciones}.
	 * @param ctx the parse tree
	 */
	void enterDec_Funcion(ControlParser.Dec_FuncionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_Funcion}
	 * labeled alternative in {@link ControlParser#decFunciones}.
	 * @param ctx the parse tree
	 */
	void exitDec_Funcion(ControlParser.Dec_FuncionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaParametros}.
	 * @param ctx the parse tree
	 */
	void enterListaParametros(ControlParser.ListaParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaParametros}.
	 * @param ctx the parse tree
	 */
	void exitListaParametros(ControlParser.ListaParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#llamadaFuncion}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaFuncion(ControlParser.LlamadaFuncionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#llamadaFuncion}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaFuncion(ControlParser.LlamadaFuncionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaLlamadaParametros}.
	 * @param ctx the parse tree
	 */
	void enterListaLlamadaParametros(ControlParser.ListaLlamadaParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaLlamadaParametros}.
	 * @param ctx the parse tree
	 */
	void exitListaLlamadaParametros(ControlParser.ListaLlamadaParametrosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IntruccionPrint}
	 * labeled alternative in {@link ControlParser#instPrint}.
	 * @param ctx the parse tree
	 */
	void enterIntruccionPrint(ControlParser.IntruccionPrintContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntruccionPrint}
	 * labeled alternative in {@link ControlParser#instPrint}.
	 * @param ctx the parse tree
	 */
	void exitIntruccionPrint(ControlParser.IntruccionPrintContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Inst_Casteos}
	 * labeled alternative in {@link ControlParser#instCasteos}.
	 * @param ctx the parse tree
	 */
	void enterInst_Casteos(ControlParser.Inst_CasteosContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Inst_Casteos}
	 * labeled alternative in {@link ControlParser#instCasteos}.
	 * @param ctx the parse tree
	 */
	void exitInst_Casteos(ControlParser.Inst_CasteosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ListaExpresiones}
	 * labeled alternative in {@link ControlParser#expresionList}.
	 * @param ctx the parse tree
	 */
	void enterListaExpresiones(ControlParser.ListaExpresionesContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ListaExpresiones}
	 * labeled alternative in {@link ControlParser#expresionList}.
	 * @param ctx the parse tree
	 */
	void exitListaExpresiones(ControlParser.ListaExpresionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#objectsList}.
	 * @param ctx the parse tree
	 */
	void enterObjectsList(ControlParser.ObjectsListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#objectsList}.
	 * @param ctx the parse tree
	 */
	void exitObjectsList(ControlParser.ObjectsListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_VTipoValor}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_VTipoValor(ControlParser.Dec_VTipoValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_VTipoValor}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_VTipoValor(ControlParser.Dec_VTipoValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_VTipo}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_VTipo(ControlParser.Dec_VTipoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_VTipo}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_VTipo(ControlParser.Dec_VTipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_VVal}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_VVal(ControlParser.Dec_VValContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_VVal}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_VVal(ControlParser.Dec_VValContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_CVal}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_CVal(ControlParser.Dec_CValContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_CVal}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_CVal(ControlParser.Dec_CValContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_CTipoValor}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_CTipoValor(ControlParser.Dec_CTipoValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_CTipoValor}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_CTipoValor(ControlParser.Dec_CTipoValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_Aumento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_Aumento(ControlParser.Asignacion_AumentoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_Aumento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_Aumento(ControlParser.Asignacion_AumentoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_Decremento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_Decremento(ControlParser.Asignacion_DecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_Decremento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_Decremento(ControlParser.Asignacion_DecrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_ValorGen}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_ValorGen(ControlParser.Asignacion_ValorGenContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_ValorGen}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_ValorGen(ControlParser.Asignacion_ValorGenContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_VectorAumento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_VectorAumento(ControlParser.Asignacion_VectorAumentoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_VectorAumento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_VectorAumento(ControlParser.Asignacion_VectorAumentoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_VectorDecremento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_VectorDecremento(ControlParser.Asignacion_VectorDecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_VectorDecremento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_VectorDecremento(ControlParser.Asignacion_VectorDecrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_MatrixAumento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_MatrixAumento(ControlParser.Asignacion_MatrixAumentoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_MatrixAumento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_MatrixAumento(ControlParser.Asignacion_MatrixAumentoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_MatrixDecremento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_MatrixDecremento(ControlParser.Asignacion_MatrixDecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_MatrixDecremento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_MatrixDecremento(ControlParser.Asignacion_MatrixDecrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_VectorGeneral}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_VectorGeneral(ControlParser.Asignacion_VectorGeneralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_VectorGeneral}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_VectorGeneral(ControlParser.Asignacion_VectorGeneralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_MatrixGeneral}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_MatrixGeneral(ControlParser.Asignacion_MatrixGeneralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_MatrixGeneral}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_MatrixGeneral(ControlParser.Asignacion_MatrixGeneralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_LlAtribGeneral}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_LlAtribGeneral(ControlParser.Asignacion_LlAtribGeneralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_LlAtribGeneral}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_LlAtribGeneral(ControlParser.Asignacion_LlAtribGeneralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_LlAtribAumento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_LlAtribAumento(ControlParser.Asignacion_LlAtribAumentoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_LlAtribAumento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_LlAtribAumento(ControlParser.Asignacion_LlAtribAumentoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_LlAtribDecremento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_LlAtribDecremento(ControlParser.Asignacion_LlAtribDecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_LlAtribDecremento}
	 * labeled alternative in {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_LlAtribDecremento(ControlParser.Asignacion_LlAtribDecrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Decimal}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Decimal(ControlParser.Expr_DecimalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Decimal}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Decimal(ControlParser.Expr_DecimalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_InstCasteo}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_InstCasteo(ControlParser.Expr_InstCasteoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_InstCasteo}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_InstCasteo(ControlParser.Expr_InstCasteoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaIgualDif}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaIgualDif(ControlParser.Expr_ValidaIgualDifContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaIgualDif}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaIgualDif(ControlParser.Expr_ValidaIgualDifContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_OpMulDiv}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_OpMulDiv(ControlParser.Expr_OpMulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_OpMulDiv}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_OpMulDiv(ControlParser.Expr_OpMulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaAnd}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaAnd(ControlParser.Expr_ValidaAndContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaAnd}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaAnd(ControlParser.Expr_ValidaAndContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_LlamAtributos}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_LlamAtributos(ControlParser.Expr_LlamAtributosContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_LlamAtributos}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_LlamAtributos(ControlParser.Expr_LlamAtributosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Nil}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Nil(ControlParser.Expr_NilContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Nil}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Nil(ControlParser.Expr_NilContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_PosMatrix}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_PosMatrix(ControlParser.Expr_PosMatrixContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_PosMatrix}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_PosMatrix(ControlParser.Expr_PosMatrixContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Boolean}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Boolean(ControlParser.Expr_BooleanContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Boolean}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Boolean(ControlParser.Expr_BooleanContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_InstRango}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_InstRango(ControlParser.Expr_InstRangoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_InstRango}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_InstRango(ControlParser.Expr_InstRangoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValNumNeg}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValNumNeg(ControlParser.Expr_ValNumNegContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValNumNeg}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValNumNeg(ControlParser.Expr_ValNumNegContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaOr}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaOr(ControlParser.Expr_ValidaOrContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaOr}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaOr(ControlParser.Expr_ValidaOrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValNegacion}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValNegacion(ControlParser.Expr_ValNegacionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValNegacion}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValNegacion(ControlParser.Expr_ValNegacionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ID}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ID(ControlParser.Expr_IDContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ID}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ID(ControlParser.Expr_IDContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaMayQue}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaMayQue(ControlParser.Expr_ValidaMayQueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaMayQue}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaMayQue(ControlParser.Expr_ValidaMayQueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Conteo}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Conteo(ControlParser.Expr_ConteoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Conteo}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Conteo(ControlParser.Expr_ConteoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_IsEmpty}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_IsEmpty(ControlParser.Expr_IsEmptyContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_IsEmpty}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_IsEmpty(ControlParser.Expr_IsEmptyContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_OpSumRes}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_OpSumRes(ControlParser.Expr_OpSumResContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_OpSumRes}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_OpSumRes(ControlParser.Expr_OpSumResContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Entero}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Entero(ControlParser.Expr_EnteroContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Entero}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Entero(ControlParser.Expr_EnteroContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ParExpre}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ParExpre(ControlParser.Expr_ParExpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ParExpre}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ParExpre(ControlParser.Expr_ParExpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_StringChar}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_StringChar(ControlParser.Expr_StringCharContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_StringChar}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_StringChar(ControlParser.Expr_StringCharContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaMenQue}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaMenQue(ControlParser.Expr_ValidaMenQueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaMenQue}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaMenQue(ControlParser.Expr_ValidaMenQueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_PosVector}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_PosVector(ControlParser.Expr_PosVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_PosVector}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_PosVector(ControlParser.Expr_PosVectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#tipovariable}.
	 * @param ctx the parse tree
	 */
	void enterTipovariable(ControlParser.TipovariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#tipovariable}.
	 * @param ctx the parse tree
	 */
	void exitTipovariable(ControlParser.TipovariableContext ctx);
}