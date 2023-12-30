grammar Control;

// tokens
ID      : [a-zA-Z_][a-zA-Z0-9_]*;
INT     : [0-9]+ ;
DECIMAL : [0-9]+'.'[0-9]+;
STRING  : '"' ('\\"'|~["])* '"';
WS      :[ \t\r\n] -> skip;
LINEALCOMENT: '//' ~[\r\n]* -> channel(HIDDEN)  ;
BlockComment: '/*' .*? '*/' -> skip;
//rules

prog
    :block EOF              #ProgInicial
;

block: (sentencias)*        #BlockContent       
   ;

sentencias
    : declaracion           #SentenciaDeclaracion
    | decVector             #SentenciaDeclaracionVector
    | decMatriz             #SentenciaDeclaracionMatrix
    | decFunciones          #SentenciaDeclaracionFuncion
    | asignacion            #SentenciaAsignacion
    | insIf                 #SentenciaIf
    | insSwitch             #SentenciaSwitch
    | instWhile             #SentenciaWhile
    | instFor               #SentenciaFor
    | instGuard             #SentenciaGuard
    | instBreak             #SentenciaBreak
    | instContinue          #SentenciaContinue
    | instReturn            #SentenciaReturn
    | funcsVectoriales      #SentenciaFuncVectoriales
    | llamadaFuncion        #SentenciaLLamadaFuncion
    | instPrint             #SentenciaPrint
    | instStruct            #SentenciaStruct
    | decStruct             #SentenciaDeclaracionStruct
    | llamadaMetodos        #SentenciaLLamadaMetodos
    ;

//if
insIf
    : 'if'expr '{' block '}' ( 'else' ( insIf | '{' block '}' ) )?      #InstruccionIf
    ;
//switch
insSwitch
    : 'switch' expr '{' instCase* instDefault? '}'         #InstruccionSwitch
    ;

instCase
    : 'case' expr ':' block                                 #InstruccionCase
    ;
instDefault
    : 'default' ':' block                                   #InstruccionDefault
    ;

//while
instWhile
    : 'while' expr '{' block '}'                            #InstruccionWhile
    ;
//for
instFor
    : 'for' ID 'in' expr '{' block '}'                      #InstruccionFor
    ;
//guard
instGuard
    :'guard' expr 'else' '{' block '}'                      #InstruccionGuard
    ;
//break
instBreak
    : 'break'                                               #InstruccionBreak                  
    ;
//continue
instContinue
    : 'continue'                                            #InstruccionContinue
    ;
//return
instReturn
    : 'return'                                              #InstruccionReturnSimple
    | 'return' expr                                         #InstruccionReturnExpresion
    ;
//vectores
decVector
    : 'var' ID (':' '[' (tipovariable|ID) ']')? '=' '[' expresionList* ']' #DecVector_ExpresionLista
    | 'var' ID (':' '[' (tipovariable|ID) ']')? '=' '[' objectsList ']'    #DecVector_ObjetLista
    | 'var' ID (':' '[' (tipovariable|ID) ']') '=' ID                     #DecVector_Id
    | 'var' ID (':' '[' (tipovariable|ID) ']') '='  '['']'                 #DecVector_ValCor
    | 'let' ID (':' '[' (tipovariable|ID) ']')? '=' '[' expresionList* ']'    #DecVectorConst_ExpresionLista
    | 'let' ID (':' '[' (tipovariable|ID) ']')? '='  '[' objectsList ']'      #DecVectorConst_ObjetLista
    | 'let' ID (':' '[' (tipovariable|ID) ']') '='  ID                       #DecVectorConst_Id
    ;

funcsVectoriales
    : ID '.' 'append' '('(expr | lDupla)')'             #VectFunc_Append
    | ID '.' 'remove' '(''at' ':' expr')'               #VectFunc_Remove
    | ID '.' 'removeLast' '(' ')'                       #VectFunc_RemoveLast
    //| ID '.' ID '('expr')' //cualquier funcion con punto
    ;

// matrices
decMatriz 
    : 'var' ID ':' tipoMatriz '=' defMatrix
    ;

tipoMatriz
    : '[' tipoMatriz ']'
    | '[' tipovariable ']'
    ;

defMatrix
    : listaValoresMat
    | simpleMax
    ;

listaValoresMat
    : '[' listaValoresMat2 ']'
    ;

listaValoresMat2
    : listaValoresMat2 ',' listaValoresMat
    | listaValoresMat
    | expresionList
    | objectsList
    ;

simpleMax
    : tipoMatriz '(' 'repeating' ':' simpleMax ',' 'count' ':' INT ')'
    | tipoMatriz '(' 'repeating' ':' expr ',' 'count' ':' INT ')'
    ;

//structs
instStruct
    : 'struct' ID '{' listaAtributos* '}'
    ;

listaAtributos
    : ('let'|'var') ID (':' (tipovariable|ID))? ('=' expr)? ';'?
    | 'mutating'? decFunciones
    ;

//creacion de structs
decStruct
    : ('var'|'let') ID (':'ID)? '=' lDupla
    | ('var'|'let') ID (':'ID)? '=' ID?
    | ('var'|'let') ID (':'ID)? '=' llamadaFuncion
    ;
lDupla
    : ID '('(','? ID ':' (expr|llamadaFuncion))* ')'
    ;

//llamada atributos dentro de la funcion con self y fuera
llamAtributos
    : ID ('[' INT ']')? ('.' ID ('[' INT ']')?)+
    | 'self' '.' ID ('[' INT ']')? ('.' ID ('[' INT ']')?)*
    ;
llamadaMetodos
    : ID ('[' INT ']')? ('.' ID('[' INT ']')?)+ '(' listaLlamadaParametros? ')'
    | 'self' '.' ID ('[' INT ']')?('.' ID ('[' INT ']')?)* '(' listaLlamadaParametros? ')'
    ;


//funciones
decFunciones
    : 'func' ID '(' listaParametros* ')' ('->' (tipovariable|ID))? '{' block '}'  #Dec_Funcion
    ;

listaParametros
    : ','? (ID| '_')? ID ':' 'inout'? (tipovariable | '[' tipovariable ']' | ID)
    ;


//Llamada de funciones
llamadaFuncion
    : ID '(' listaLlamadaParametros ')'
    ;

listaLlamadaParametros
    : listaLlamadaParametros ',' (ID ':')? '&'? expr
    | (ID ':')? '&'? expr
    ;

//funciones embebidas
instPrint
    : 'print' '(' (expr|llamadaFuncion) (',' (expr|llamadaFuncion))* ')' #IntruccionPrint
    ;

instCasteos
    : tipovariable '(' expr ')' #Inst_Casteos
    ;

expresionList
    : (expr ','?)               #ListaExpresiones
    ;

objectsList
    : objectsList ',' lDupla
    | lDupla
    ;

declaracion
    : 'var' ID ':' tipovariable '=' expr    #Dec_VTipoValor
    | 'var' ID ':' tipovariable '?'         #Dec_VTipo
    | 'var' ID '=' expr                     #Dec_VVal
    | 'let' ID '=' expr                     #Dec_CVal
    | 'let' ID ':' tipovariable '=' expr    #Dec_CTipoValor
    ;
asignacion
    : ID '+=' expr                                              #Asignacion_Aumento
    | ID '-=' expr                                              #Asignacion_Decremento
    | ID '=' expr                                               #Asignacion_ValorGen
    | ID '[' expr ']' '+=' expr                             #Asignacion_VectorAumento
    | ID '[' expr ']' '-=' expr                             #Asignacion_VectorDecremento
    | ID '[' expr ']' ('[' expr ']')+ '-=' expr         #Asignacion_MatrixAumento
    | ID '[' expr ']' ('[' expr ']')+ '+=' expr         #Asignacion_MatrixDecremento
    | ID '[' expr ']' '=' expr                              #Asignacion_VectorGeneral
    | ID '[' expr ']' ('[' expr ']')+ '=' expr          #Asignacion_MatrixGeneral
    | llamAtributos '=' expr                                    #Asignacion_LlAtribGeneral
    | llamAtributos '+=' expr                                   #Asignacion_LlAtribAumento
    | llamAtributos '-=' expr                                   #Asignacion_LlAtribDecremento
    ;
expr
    : '-' expr                                  #Expr_ValNumNeg
    | '!' rigth = expr                          #Expr_ValNegacion
    | left=expr op=('*'|'/') rigth=expr         #Expr_OpMulDiv
    | left=expr op=('+'| '-'|'%') rigth=expr    #Expr_OpSumRes
    | left=expr op=('>='|'>') right=expr        #Expr_ValidaMayQue
    | left=expr op=('<='|'<') right=expr        #Expr_ValidaMenQue
    | left=expr op=('=='|'!=') right=expr       #Expr_ValidaIgualDif
    | left=expr '&&' right=expr                 #Expr_ValidaAnd
    | left=expr '||' right=expr                 #Expr_ValidaOr
    | '(' expr ')'                              #Expr_ParExpre
    | INT                                       #Expr_Entero
    | DECIMAL                                   #Expr_Decimal
    | STRING                                    #Expr_StringChar
    | ('true' | 'false')                        #Expr_Boolean
    | 'nil'                                     #Expr_Nil
    | ID                                        #Expr_ID
    | ID '.' 'count'                            #Expr_Conteo
    | ID '.' 'isEmpty' '(' ')'                  #Expr_IsEmpty
    | ID '[' (expr) ']'                       #Expr_PosVector        
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+   #Expr_PosMatrix
    | expr '...' expr                             #Expr_InstRango
    | instCasteos                               #Expr_InstCasteo
    | llamAtributos                             #Expr_LlamAtributos
    ;

tipovariable
    : 'String'
    | 'Float'
    | 'Int'
    | 'Bool'
    | 'Character'
    ;
