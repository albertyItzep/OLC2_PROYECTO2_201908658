// Generated from Control.g4 by ANTLR 4.13.0
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class ControlLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.0", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		T__45=46, T__46=47, T__47=48, T__48=49, T__49=50, T__50=51, T__51=52, 
		T__52=53, T__53=54, T__54=55, T__55=56, T__56=57, T__57=58, T__58=59, 
		T__59=60, T__60=61, T__61=62, T__62=63, T__63=64, T__64=65, T__65=66, 
		T__66=67, ID=68, INT=69, DECIMAL=70, STRING=71, WS=72, LINEALCOMENT=73, 
		BlockComment=74;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
			"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32", 
			"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40", 
			"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48", 
			"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56", 
			"T__57", "T__58", "T__59", "T__60", "T__61", "T__62", "T__63", "T__64", 
			"T__65", "T__66", "ID", "INT", "DECIMAL", "STRING", "WS", "LINEALCOMENT", 
			"BlockComment"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'if'", "'{'", "'}'", "'else'", "'switch'", "'case'", "':'", "'default'", 
			"'while'", "'for'", "'in'", "'guard'", "'break'", "'continue'", "'return'", 
			"'var'", "'['", "']'", "'='", "'let'", "'.'", "'append'", "'('", "')'", 
			"'remove'", "'at'", "'removeLast'", "','", "'repeating'", "'count'", 
			"'struct'", "';'", "'mutating'", "'self'", "'func'", "'->'", "'_'", "'inout'", 
			"'&'", "'print'", "'?'", "'+='", "'-='", "'-'", "'!'", "'*'", "'/'", 
			"'+'", "'%'", "'>='", "'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", 
			"'true'", "'false'", "'nil'", "'isEmpty'", "'...'", "'String'", "'Float'", 
			"'Int'", "'Bool'", "'Character'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "ID", "INT", "DECIMAL", 
			"STRING", "WS", "LINEALCOMENT", "BlockComment"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ControlLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Control.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000J\u01f5\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002"+
		"\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002"+
		"\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002"+
		"\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002"+
		"\u001b\u0007\u001b\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002"+
		"\u001e\u0007\u001e\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007"+
		"!\u0002\"\u0007\"\u0002#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007"+
		"&\u0002\'\u0007\'\u0002(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007"+
		"+\u0002,\u0007,\u0002-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u0007"+
		"0\u00021\u00071\u00022\u00072\u00023\u00073\u00024\u00074\u00025\u0007"+
		"5\u00026\u00076\u00027\u00077\u00028\u00078\u00029\u00079\u0002:\u0007"+
		":\u0002;\u0007;\u0002<\u0007<\u0002=\u0007=\u0002>\u0007>\u0002?\u0007"+
		"?\u0002@\u0007@\u0002A\u0007A\u0002B\u0007B\u0002C\u0007C\u0002D\u0007"+
		"D\u0002E\u0007E\u0002F\u0007F\u0002G\u0007G\u0002H\u0007H\u0002I\u0007"+
		"I\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010"+
		"\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0016"+
		"\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001f\u0001\u001f\u0001 \u0001 \u0001 \u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0001 \u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001#\u0001#\u0001#\u0001$\u0001$"+
		"\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001&\u0001&\u0001\'\u0001"+
		"\'\u0001\'\u0001\'\u0001\'\u0001\'\u0001(\u0001(\u0001)\u0001)\u0001)"+
		"\u0001*\u0001*\u0001*\u0001+\u0001+\u0001,\u0001,\u0001-\u0001-\u0001"+
		".\u0001.\u0001/\u0001/\u00010\u00010\u00011\u00011\u00011\u00012\u0001"+
		"2\u00013\u00013\u00013\u00014\u00014\u00015\u00015\u00015\u00016\u0001"+
		"6\u00016\u00017\u00017\u00017\u00018\u00018\u00018\u00019\u00019\u0001"+
		"9\u00019\u00019\u0001:\u0001:\u0001:\u0001:\u0001:\u0001:\u0001;\u0001"+
		";\u0001;\u0001;\u0001<\u0001<\u0001<\u0001<\u0001<\u0001<\u0001<\u0001"+
		"<\u0001=\u0001=\u0001=\u0001=\u0001>\u0001>\u0001>\u0001>\u0001>\u0001"+
		">\u0001>\u0001?\u0001?\u0001?\u0001?\u0001?\u0001?\u0001@\u0001@\u0001"+
		"@\u0001@\u0001A\u0001A\u0001A\u0001A\u0001A\u0001B\u0001B\u0001B\u0001"+
		"B\u0001B\u0001B\u0001B\u0001B\u0001B\u0001B\u0001C\u0001C\u0005C\u01b9"+
		"\bC\nC\fC\u01bc\tC\u0001D\u0004D\u01bf\bD\u000bD\fD\u01c0\u0001E\u0004"+
		"E\u01c4\bE\u000bE\fE\u01c5\u0001E\u0001E\u0004E\u01ca\bE\u000bE\fE\u01cb"+
		"\u0001F\u0001F\u0001F\u0001F\u0005F\u01d2\bF\nF\fF\u01d5\tF\u0001F\u0001"+
		"F\u0001G\u0001G\u0001G\u0001G\u0001H\u0001H\u0001H\u0001H\u0005H\u01e1"+
		"\bH\nH\fH\u01e4\tH\u0001H\u0001H\u0001I\u0001I\u0001I\u0001I\u0005I\u01ec"+
		"\bI\nI\fI\u01ef\tI\u0001I\u0001I\u0001I\u0001I\u0001I\u0001\u01ed\u0000"+
		"J\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006"+
		"\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e"+
		"\u001d\u000f\u001f\u0010!\u0011#\u0012%\u0013\'\u0014)\u0015+\u0016-\u0017"+
		"/\u00181\u00193\u001a5\u001b7\u001c9\u001d;\u001e=\u001f? A!C\"E#G$I%"+
		"K&M\'O(Q)S*U+W,Y-[.]/_0a1c2e3g4i5k6m7o8q9s:u;w<y={>}?\u007f@\u0081A\u0083"+
		"B\u0085C\u0087D\u0089E\u008bF\u008dG\u008fH\u0091I\u0093J\u0001\u0000"+
		"\u0006\u0003\u0000AZ__az\u0004\u000009AZ__az\u0001\u000009\u0001\u0000"+
		"\"\"\u0003\u0000\t\n\r\r  \u0002\u0000\n\n\r\r\u01fc\u0000\u0001\u0001"+
		"\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001"+
		"\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000"+
		"\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000"+
		"\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000"+
		"\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000"+
		"\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000"+
		"\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000"+
		"\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000"+
		"\u0000#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000\u0000\u0000\'"+
		"\u0001\u0000\u0000\u0000\u0000)\u0001\u0000\u0000\u0000\u0000+\u0001\u0000"+
		"\u0000\u0000\u0000-\u0001\u0000\u0000\u0000\u0000/\u0001\u0000\u0000\u0000"+
		"\u00001\u0001\u0000\u0000\u0000\u00003\u0001\u0000\u0000\u0000\u00005"+
		"\u0001\u0000\u0000\u0000\u00007\u0001\u0000\u0000\u0000\u00009\u0001\u0000"+
		"\u0000\u0000\u0000;\u0001\u0000\u0000\u0000\u0000=\u0001\u0000\u0000\u0000"+
		"\u0000?\u0001\u0000\u0000\u0000\u0000A\u0001\u0000\u0000\u0000\u0000C"+
		"\u0001\u0000\u0000\u0000\u0000E\u0001\u0000\u0000\u0000\u0000G\u0001\u0000"+
		"\u0000\u0000\u0000I\u0001\u0000\u0000\u0000\u0000K\u0001\u0000\u0000\u0000"+
		"\u0000M\u0001\u0000\u0000\u0000\u0000O\u0001\u0000\u0000\u0000\u0000Q"+
		"\u0001\u0000\u0000\u0000\u0000S\u0001\u0000\u0000\u0000\u0000U\u0001\u0000"+
		"\u0000\u0000\u0000W\u0001\u0000\u0000\u0000\u0000Y\u0001\u0000\u0000\u0000"+
		"\u0000[\u0001\u0000\u0000\u0000\u0000]\u0001\u0000\u0000\u0000\u0000_"+
		"\u0001\u0000\u0000\u0000\u0000a\u0001\u0000\u0000\u0000\u0000c\u0001\u0000"+
		"\u0000\u0000\u0000e\u0001\u0000\u0000\u0000\u0000g\u0001\u0000\u0000\u0000"+
		"\u0000i\u0001\u0000\u0000\u0000\u0000k\u0001\u0000\u0000\u0000\u0000m"+
		"\u0001\u0000\u0000\u0000\u0000o\u0001\u0000\u0000\u0000\u0000q\u0001\u0000"+
		"\u0000\u0000\u0000s\u0001\u0000\u0000\u0000\u0000u\u0001\u0000\u0000\u0000"+
		"\u0000w\u0001\u0000\u0000\u0000\u0000y\u0001\u0000\u0000\u0000\u0000{"+
		"\u0001\u0000\u0000\u0000\u0000}\u0001\u0000\u0000\u0000\u0000\u007f\u0001"+
		"\u0000\u0000\u0000\u0000\u0081\u0001\u0000\u0000\u0000\u0000\u0083\u0001"+
		"\u0000\u0000\u0000\u0000\u0085\u0001\u0000\u0000\u0000\u0000\u0087\u0001"+
		"\u0000\u0000\u0000\u0000\u0089\u0001\u0000\u0000\u0000\u0000\u008b\u0001"+
		"\u0000\u0000\u0000\u0000\u008d\u0001\u0000\u0000\u0000\u0000\u008f\u0001"+
		"\u0000\u0000\u0000\u0000\u0091\u0001\u0000\u0000\u0000\u0000\u0093\u0001"+
		"\u0000\u0000\u0000\u0001\u0095\u0001\u0000\u0000\u0000\u0003\u0098\u0001"+
		"\u0000\u0000\u0000\u0005\u009a\u0001\u0000\u0000\u0000\u0007\u009c\u0001"+
		"\u0000\u0000\u0000\t\u00a1\u0001\u0000\u0000\u0000\u000b\u00a8\u0001\u0000"+
		"\u0000\u0000\r\u00ad\u0001\u0000\u0000\u0000\u000f\u00af\u0001\u0000\u0000"+
		"\u0000\u0011\u00b7\u0001\u0000\u0000\u0000\u0013\u00bd\u0001\u0000\u0000"+
		"\u0000\u0015\u00c1\u0001\u0000\u0000\u0000\u0017\u00c4\u0001\u0000\u0000"+
		"\u0000\u0019\u00ca\u0001\u0000\u0000\u0000\u001b\u00d0\u0001\u0000\u0000"+
		"\u0000\u001d\u00d9\u0001\u0000\u0000\u0000\u001f\u00e0\u0001\u0000\u0000"+
		"\u0000!\u00e4\u0001\u0000\u0000\u0000#\u00e6\u0001\u0000\u0000\u0000%"+
		"\u00e8\u0001\u0000\u0000\u0000\'\u00ea\u0001\u0000\u0000\u0000)\u00ee"+
		"\u0001\u0000\u0000\u0000+\u00f0\u0001\u0000\u0000\u0000-\u00f7\u0001\u0000"+
		"\u0000\u0000/\u00f9\u0001\u0000\u0000\u00001\u00fb\u0001\u0000\u0000\u0000"+
		"3\u0102\u0001\u0000\u0000\u00005\u0105\u0001\u0000\u0000\u00007\u0110"+
		"\u0001\u0000\u0000\u00009\u0112\u0001\u0000\u0000\u0000;\u011c\u0001\u0000"+
		"\u0000\u0000=\u0122\u0001\u0000\u0000\u0000?\u0129\u0001\u0000\u0000\u0000"+
		"A\u012b\u0001\u0000\u0000\u0000C\u0134\u0001\u0000\u0000\u0000E\u0139"+
		"\u0001\u0000\u0000\u0000G\u013e\u0001\u0000\u0000\u0000I\u0141\u0001\u0000"+
		"\u0000\u0000K\u0143\u0001\u0000\u0000\u0000M\u0149\u0001\u0000\u0000\u0000"+
		"O\u014b\u0001\u0000\u0000\u0000Q\u0151\u0001\u0000\u0000\u0000S\u0153"+
		"\u0001\u0000\u0000\u0000U\u0156\u0001\u0000\u0000\u0000W\u0159\u0001\u0000"+
		"\u0000\u0000Y\u015b\u0001\u0000\u0000\u0000[\u015d\u0001\u0000\u0000\u0000"+
		"]\u015f\u0001\u0000\u0000\u0000_\u0161\u0001\u0000\u0000\u0000a\u0163"+
		"\u0001\u0000\u0000\u0000c\u0165\u0001\u0000\u0000\u0000e\u0168\u0001\u0000"+
		"\u0000\u0000g\u016a\u0001\u0000\u0000\u0000i\u016d\u0001\u0000\u0000\u0000"+
		"k\u016f\u0001\u0000\u0000\u0000m\u0172\u0001\u0000\u0000\u0000o\u0175"+
		"\u0001\u0000\u0000\u0000q\u0178\u0001\u0000\u0000\u0000s\u017b\u0001\u0000"+
		"\u0000\u0000u\u0180\u0001\u0000\u0000\u0000w\u0186\u0001\u0000\u0000\u0000"+
		"y\u018a\u0001\u0000\u0000\u0000{\u0192\u0001\u0000\u0000\u0000}\u0196"+
		"\u0001\u0000\u0000\u0000\u007f\u019d\u0001\u0000\u0000\u0000\u0081\u01a3"+
		"\u0001\u0000\u0000\u0000\u0083\u01a7\u0001\u0000\u0000\u0000\u0085\u01ac"+
		"\u0001\u0000\u0000\u0000\u0087\u01b6\u0001\u0000\u0000\u0000\u0089\u01be"+
		"\u0001\u0000\u0000\u0000\u008b\u01c3\u0001\u0000\u0000\u0000\u008d\u01cd"+
		"\u0001\u0000\u0000\u0000\u008f\u01d8\u0001\u0000\u0000\u0000\u0091\u01dc"+
		"\u0001\u0000\u0000\u0000\u0093\u01e7\u0001\u0000\u0000\u0000\u0095\u0096"+
		"\u0005i\u0000\u0000\u0096\u0097\u0005f\u0000\u0000\u0097\u0002\u0001\u0000"+
		"\u0000\u0000\u0098\u0099\u0005{\u0000\u0000\u0099\u0004\u0001\u0000\u0000"+
		"\u0000\u009a\u009b\u0005}\u0000\u0000\u009b\u0006\u0001\u0000\u0000\u0000"+
		"\u009c\u009d\u0005e\u0000\u0000\u009d\u009e\u0005l\u0000\u0000\u009e\u009f"+
		"\u0005s\u0000\u0000\u009f\u00a0\u0005e\u0000\u0000\u00a0\b\u0001\u0000"+
		"\u0000\u0000\u00a1\u00a2\u0005s\u0000\u0000\u00a2\u00a3\u0005w\u0000\u0000"+
		"\u00a3\u00a4\u0005i\u0000\u0000\u00a4\u00a5\u0005t\u0000\u0000\u00a5\u00a6"+
		"\u0005c\u0000\u0000\u00a6\u00a7\u0005h\u0000\u0000\u00a7\n\u0001\u0000"+
		"\u0000\u0000\u00a8\u00a9\u0005c\u0000\u0000\u00a9\u00aa\u0005a\u0000\u0000"+
		"\u00aa\u00ab\u0005s\u0000\u0000\u00ab\u00ac\u0005e\u0000\u0000\u00ac\f"+
		"\u0001\u0000\u0000\u0000\u00ad\u00ae\u0005:\u0000\u0000\u00ae\u000e\u0001"+
		"\u0000\u0000\u0000\u00af\u00b0\u0005d\u0000\u0000\u00b0\u00b1\u0005e\u0000"+
		"\u0000\u00b1\u00b2\u0005f\u0000\u0000\u00b2\u00b3\u0005a\u0000\u0000\u00b3"+
		"\u00b4\u0005u\u0000\u0000\u00b4\u00b5\u0005l\u0000\u0000\u00b5\u00b6\u0005"+
		"t\u0000\u0000\u00b6\u0010\u0001\u0000\u0000\u0000\u00b7\u00b8\u0005w\u0000"+
		"\u0000\u00b8\u00b9\u0005h\u0000\u0000\u00b9\u00ba\u0005i\u0000\u0000\u00ba"+
		"\u00bb\u0005l\u0000\u0000\u00bb\u00bc\u0005e\u0000\u0000\u00bc\u0012\u0001"+
		"\u0000\u0000\u0000\u00bd\u00be\u0005f\u0000\u0000\u00be\u00bf\u0005o\u0000"+
		"\u0000\u00bf\u00c0\u0005r\u0000\u0000\u00c0\u0014\u0001\u0000\u0000\u0000"+
		"\u00c1\u00c2\u0005i\u0000\u0000\u00c2\u00c3\u0005n\u0000\u0000\u00c3\u0016"+
		"\u0001\u0000\u0000\u0000\u00c4\u00c5\u0005g\u0000\u0000\u00c5\u00c6\u0005"+
		"u\u0000\u0000\u00c6\u00c7\u0005a\u0000\u0000\u00c7\u00c8\u0005r\u0000"+
		"\u0000\u00c8\u00c9\u0005d\u0000\u0000\u00c9\u0018\u0001\u0000\u0000\u0000"+
		"\u00ca\u00cb\u0005b\u0000\u0000\u00cb\u00cc\u0005r\u0000\u0000\u00cc\u00cd"+
		"\u0005e\u0000\u0000\u00cd\u00ce\u0005a\u0000\u0000\u00ce\u00cf\u0005k"+
		"\u0000\u0000\u00cf\u001a\u0001\u0000\u0000\u0000\u00d0\u00d1\u0005c\u0000"+
		"\u0000\u00d1\u00d2\u0005o\u0000\u0000\u00d2\u00d3\u0005n\u0000\u0000\u00d3"+
		"\u00d4\u0005t\u0000\u0000\u00d4\u00d5\u0005i\u0000\u0000\u00d5\u00d6\u0005"+
		"n\u0000\u0000\u00d6\u00d7\u0005u\u0000\u0000\u00d7\u00d8\u0005e\u0000"+
		"\u0000\u00d8\u001c\u0001\u0000\u0000\u0000\u00d9\u00da\u0005r\u0000\u0000"+
		"\u00da\u00db\u0005e\u0000\u0000\u00db\u00dc\u0005t\u0000\u0000\u00dc\u00dd"+
		"\u0005u\u0000\u0000\u00dd\u00de\u0005r\u0000\u0000\u00de\u00df\u0005n"+
		"\u0000\u0000\u00df\u001e\u0001\u0000\u0000\u0000\u00e0\u00e1\u0005v\u0000"+
		"\u0000\u00e1\u00e2\u0005a\u0000\u0000\u00e2\u00e3\u0005r\u0000\u0000\u00e3"+
		" \u0001\u0000\u0000\u0000\u00e4\u00e5\u0005[\u0000\u0000\u00e5\"\u0001"+
		"\u0000\u0000\u0000\u00e6\u00e7\u0005]\u0000\u0000\u00e7$\u0001\u0000\u0000"+
		"\u0000\u00e8\u00e9\u0005=\u0000\u0000\u00e9&\u0001\u0000\u0000\u0000\u00ea"+
		"\u00eb\u0005l\u0000\u0000\u00eb\u00ec\u0005e\u0000\u0000\u00ec\u00ed\u0005"+
		"t\u0000\u0000\u00ed(\u0001\u0000\u0000\u0000\u00ee\u00ef\u0005.\u0000"+
		"\u0000\u00ef*\u0001\u0000\u0000\u0000\u00f0\u00f1\u0005a\u0000\u0000\u00f1"+
		"\u00f2\u0005p\u0000\u0000\u00f2\u00f3\u0005p\u0000\u0000\u00f3\u00f4\u0005"+
		"e\u0000\u0000\u00f4\u00f5\u0005n\u0000\u0000\u00f5\u00f6\u0005d\u0000"+
		"\u0000\u00f6,\u0001\u0000\u0000\u0000\u00f7\u00f8\u0005(\u0000\u0000\u00f8"+
		".\u0001\u0000\u0000\u0000\u00f9\u00fa\u0005)\u0000\u0000\u00fa0\u0001"+
		"\u0000\u0000\u0000\u00fb\u00fc\u0005r\u0000\u0000\u00fc\u00fd\u0005e\u0000"+
		"\u0000\u00fd\u00fe\u0005m\u0000\u0000\u00fe\u00ff\u0005o\u0000\u0000\u00ff"+
		"\u0100\u0005v\u0000\u0000\u0100\u0101\u0005e\u0000\u0000\u01012\u0001"+
		"\u0000\u0000\u0000\u0102\u0103\u0005a\u0000\u0000\u0103\u0104\u0005t\u0000"+
		"\u0000\u01044\u0001\u0000\u0000\u0000\u0105\u0106\u0005r\u0000\u0000\u0106"+
		"\u0107\u0005e\u0000\u0000\u0107\u0108\u0005m\u0000\u0000\u0108\u0109\u0005"+
		"o\u0000\u0000\u0109\u010a\u0005v\u0000\u0000\u010a\u010b\u0005e\u0000"+
		"\u0000\u010b\u010c\u0005L\u0000\u0000\u010c\u010d\u0005a\u0000\u0000\u010d"+
		"\u010e\u0005s\u0000\u0000\u010e\u010f\u0005t\u0000\u0000\u010f6\u0001"+
		"\u0000\u0000\u0000\u0110\u0111\u0005,\u0000\u0000\u01118\u0001\u0000\u0000"+
		"\u0000\u0112\u0113\u0005r\u0000\u0000\u0113\u0114\u0005e\u0000\u0000\u0114"+
		"\u0115\u0005p\u0000\u0000\u0115\u0116\u0005e\u0000\u0000\u0116\u0117\u0005"+
		"a\u0000\u0000\u0117\u0118\u0005t\u0000\u0000\u0118\u0119\u0005i\u0000"+
		"\u0000\u0119\u011a\u0005n\u0000\u0000\u011a\u011b\u0005g\u0000\u0000\u011b"+
		":\u0001\u0000\u0000\u0000\u011c\u011d\u0005c\u0000\u0000\u011d\u011e\u0005"+
		"o\u0000\u0000\u011e\u011f\u0005u\u0000\u0000\u011f\u0120\u0005n\u0000"+
		"\u0000\u0120\u0121\u0005t\u0000\u0000\u0121<\u0001\u0000\u0000\u0000\u0122"+
		"\u0123\u0005s\u0000\u0000\u0123\u0124\u0005t\u0000\u0000\u0124\u0125\u0005"+
		"r\u0000\u0000\u0125\u0126\u0005u\u0000\u0000\u0126\u0127\u0005c\u0000"+
		"\u0000\u0127\u0128\u0005t\u0000\u0000\u0128>\u0001\u0000\u0000\u0000\u0129"+
		"\u012a\u0005;\u0000\u0000\u012a@\u0001\u0000\u0000\u0000\u012b\u012c\u0005"+
		"m\u0000\u0000\u012c\u012d\u0005u\u0000\u0000\u012d\u012e\u0005t\u0000"+
		"\u0000\u012e\u012f\u0005a\u0000\u0000\u012f\u0130\u0005t\u0000\u0000\u0130"+
		"\u0131\u0005i\u0000\u0000\u0131\u0132\u0005n\u0000\u0000\u0132\u0133\u0005"+
		"g\u0000\u0000\u0133B\u0001\u0000\u0000\u0000\u0134\u0135\u0005s\u0000"+
		"\u0000\u0135\u0136\u0005e\u0000\u0000\u0136\u0137\u0005l\u0000\u0000\u0137"+
		"\u0138\u0005f\u0000\u0000\u0138D\u0001\u0000\u0000\u0000\u0139\u013a\u0005"+
		"f\u0000\u0000\u013a\u013b\u0005u\u0000\u0000\u013b\u013c\u0005n\u0000"+
		"\u0000\u013c\u013d\u0005c\u0000\u0000\u013dF\u0001\u0000\u0000\u0000\u013e"+
		"\u013f\u0005-\u0000\u0000\u013f\u0140\u0005>\u0000\u0000\u0140H\u0001"+
		"\u0000\u0000\u0000\u0141\u0142\u0005_\u0000\u0000\u0142J\u0001\u0000\u0000"+
		"\u0000\u0143\u0144\u0005i\u0000\u0000\u0144\u0145\u0005n\u0000\u0000\u0145"+
		"\u0146\u0005o\u0000\u0000\u0146\u0147\u0005u\u0000\u0000\u0147\u0148\u0005"+
		"t\u0000\u0000\u0148L\u0001\u0000\u0000\u0000\u0149\u014a\u0005&\u0000"+
		"\u0000\u014aN\u0001\u0000\u0000\u0000\u014b\u014c\u0005p\u0000\u0000\u014c"+
		"\u014d\u0005r\u0000\u0000\u014d\u014e\u0005i\u0000\u0000\u014e\u014f\u0005"+
		"n\u0000\u0000\u014f\u0150\u0005t\u0000\u0000\u0150P\u0001\u0000\u0000"+
		"\u0000\u0151\u0152\u0005?\u0000\u0000\u0152R\u0001\u0000\u0000\u0000\u0153"+
		"\u0154\u0005+\u0000\u0000\u0154\u0155\u0005=\u0000\u0000\u0155T\u0001"+
		"\u0000\u0000\u0000\u0156\u0157\u0005-\u0000\u0000\u0157\u0158\u0005=\u0000"+
		"\u0000\u0158V\u0001\u0000\u0000\u0000\u0159\u015a\u0005-\u0000\u0000\u015a"+
		"X\u0001\u0000\u0000\u0000\u015b\u015c\u0005!\u0000\u0000\u015cZ\u0001"+
		"\u0000\u0000\u0000\u015d\u015e\u0005*\u0000\u0000\u015e\\\u0001\u0000"+
		"\u0000\u0000\u015f\u0160\u0005/\u0000\u0000\u0160^\u0001\u0000\u0000\u0000"+
		"\u0161\u0162\u0005+\u0000\u0000\u0162`\u0001\u0000\u0000\u0000\u0163\u0164"+
		"\u0005%\u0000\u0000\u0164b\u0001\u0000\u0000\u0000\u0165\u0166\u0005>"+
		"\u0000\u0000\u0166\u0167\u0005=\u0000\u0000\u0167d\u0001\u0000\u0000\u0000"+
		"\u0168\u0169\u0005>\u0000\u0000\u0169f\u0001\u0000\u0000\u0000\u016a\u016b"+
		"\u0005<\u0000\u0000\u016b\u016c\u0005=\u0000\u0000\u016ch\u0001\u0000"+
		"\u0000\u0000\u016d\u016e\u0005<\u0000\u0000\u016ej\u0001\u0000\u0000\u0000"+
		"\u016f\u0170\u0005=\u0000\u0000\u0170\u0171\u0005=\u0000\u0000\u0171l"+
		"\u0001\u0000\u0000\u0000\u0172\u0173\u0005!\u0000\u0000\u0173\u0174\u0005"+
		"=\u0000\u0000\u0174n\u0001\u0000\u0000\u0000\u0175\u0176\u0005&\u0000"+
		"\u0000\u0176\u0177\u0005&\u0000\u0000\u0177p\u0001\u0000\u0000\u0000\u0178"+
		"\u0179\u0005|\u0000\u0000\u0179\u017a\u0005|\u0000\u0000\u017ar\u0001"+
		"\u0000\u0000\u0000\u017b\u017c\u0005t\u0000\u0000\u017c\u017d\u0005r\u0000"+
		"\u0000\u017d\u017e\u0005u\u0000\u0000\u017e\u017f\u0005e\u0000\u0000\u017f"+
		"t\u0001\u0000\u0000\u0000\u0180\u0181\u0005f\u0000\u0000\u0181\u0182\u0005"+
		"a\u0000\u0000\u0182\u0183\u0005l\u0000\u0000\u0183\u0184\u0005s\u0000"+
		"\u0000\u0184\u0185\u0005e\u0000\u0000\u0185v\u0001\u0000\u0000\u0000\u0186"+
		"\u0187\u0005n\u0000\u0000\u0187\u0188\u0005i\u0000\u0000\u0188\u0189\u0005"+
		"l\u0000\u0000\u0189x\u0001\u0000\u0000\u0000\u018a\u018b\u0005i\u0000"+
		"\u0000\u018b\u018c\u0005s\u0000\u0000\u018c\u018d\u0005E\u0000\u0000\u018d"+
		"\u018e\u0005m\u0000\u0000\u018e\u018f\u0005p\u0000\u0000\u018f\u0190\u0005"+
		"t\u0000\u0000\u0190\u0191\u0005y\u0000\u0000\u0191z\u0001\u0000\u0000"+
		"\u0000\u0192\u0193\u0005.\u0000\u0000\u0193\u0194\u0005.\u0000\u0000\u0194"+
		"\u0195\u0005.\u0000\u0000\u0195|\u0001\u0000\u0000\u0000\u0196\u0197\u0005"+
		"S\u0000\u0000\u0197\u0198\u0005t\u0000\u0000\u0198\u0199\u0005r\u0000"+
		"\u0000\u0199\u019a\u0005i\u0000\u0000\u019a\u019b\u0005n\u0000\u0000\u019b"+
		"\u019c\u0005g\u0000\u0000\u019c~\u0001\u0000\u0000\u0000\u019d\u019e\u0005"+
		"F\u0000\u0000\u019e\u019f\u0005l\u0000\u0000\u019f\u01a0\u0005o\u0000"+
		"\u0000\u01a0\u01a1\u0005a\u0000\u0000\u01a1\u01a2\u0005t\u0000\u0000\u01a2"+
		"\u0080\u0001\u0000\u0000\u0000\u01a3\u01a4\u0005I\u0000\u0000\u01a4\u01a5"+
		"\u0005n\u0000\u0000\u01a5\u01a6\u0005t\u0000\u0000\u01a6\u0082\u0001\u0000"+
		"\u0000\u0000\u01a7\u01a8\u0005B\u0000\u0000\u01a8\u01a9\u0005o\u0000\u0000"+
		"\u01a9\u01aa\u0005o\u0000\u0000\u01aa\u01ab\u0005l\u0000\u0000\u01ab\u0084"+
		"\u0001\u0000\u0000\u0000\u01ac\u01ad\u0005C\u0000\u0000\u01ad\u01ae\u0005"+
		"h\u0000\u0000\u01ae\u01af\u0005a\u0000\u0000\u01af\u01b0\u0005r\u0000"+
		"\u0000\u01b0\u01b1\u0005a\u0000\u0000\u01b1\u01b2\u0005c\u0000\u0000\u01b2"+
		"\u01b3\u0005t\u0000\u0000\u01b3\u01b4\u0005e\u0000\u0000\u01b4\u01b5\u0005"+
		"r\u0000\u0000\u01b5\u0086\u0001\u0000\u0000\u0000\u01b6\u01ba\u0007\u0000"+
		"\u0000\u0000\u01b7\u01b9\u0007\u0001\u0000\u0000\u01b8\u01b7\u0001\u0000"+
		"\u0000\u0000\u01b9\u01bc\u0001\u0000\u0000\u0000\u01ba\u01b8\u0001\u0000"+
		"\u0000\u0000\u01ba\u01bb\u0001\u0000\u0000\u0000\u01bb\u0088\u0001\u0000"+
		"\u0000\u0000\u01bc\u01ba\u0001\u0000\u0000\u0000\u01bd\u01bf\u0007\u0002"+
		"\u0000\u0000\u01be\u01bd\u0001\u0000\u0000\u0000\u01bf\u01c0\u0001\u0000"+
		"\u0000\u0000\u01c0\u01be\u0001\u0000\u0000\u0000\u01c0\u01c1\u0001\u0000"+
		"\u0000\u0000\u01c1\u008a\u0001\u0000\u0000\u0000\u01c2\u01c4\u0007\u0002"+
		"\u0000\u0000\u01c3\u01c2\u0001\u0000\u0000\u0000\u01c4\u01c5\u0001\u0000"+
		"\u0000\u0000\u01c5\u01c3\u0001\u0000\u0000\u0000\u01c5\u01c6\u0001\u0000"+
		"\u0000\u0000\u01c6\u01c7\u0001\u0000\u0000\u0000\u01c7\u01c9\u0005.\u0000"+
		"\u0000\u01c8\u01ca\u0007\u0002\u0000\u0000\u01c9\u01c8\u0001\u0000\u0000"+
		"\u0000\u01ca\u01cb\u0001\u0000\u0000\u0000\u01cb\u01c9\u0001\u0000\u0000"+
		"\u0000\u01cb\u01cc\u0001\u0000\u0000\u0000\u01cc\u008c\u0001\u0000\u0000"+
		"\u0000\u01cd\u01d3\u0005\"\u0000\u0000\u01ce\u01cf\u0005\\\u0000\u0000"+
		"\u01cf\u01d2\u0005\"\u0000\u0000\u01d0\u01d2\b\u0003\u0000\u0000\u01d1"+
		"\u01ce\u0001\u0000\u0000\u0000\u01d1\u01d0\u0001\u0000\u0000\u0000\u01d2"+
		"\u01d5\u0001\u0000\u0000\u0000\u01d3\u01d1\u0001\u0000\u0000\u0000\u01d3"+
		"\u01d4\u0001\u0000\u0000\u0000\u01d4\u01d6\u0001\u0000\u0000\u0000\u01d5"+
		"\u01d3\u0001\u0000\u0000\u0000\u01d6\u01d7\u0005\"\u0000\u0000\u01d7\u008e"+
		"\u0001\u0000\u0000\u0000\u01d8\u01d9\u0007\u0004\u0000\u0000\u01d9\u01da"+
		"\u0001\u0000\u0000\u0000\u01da\u01db\u0006G\u0000\u0000\u01db\u0090\u0001"+
		"\u0000\u0000\u0000\u01dc\u01dd\u0005/\u0000\u0000\u01dd\u01de\u0005/\u0000"+
		"\u0000\u01de\u01e2\u0001\u0000\u0000\u0000\u01df\u01e1\b\u0005\u0000\u0000"+
		"\u01e0\u01df\u0001\u0000\u0000\u0000\u01e1\u01e4\u0001\u0000\u0000\u0000"+
		"\u01e2\u01e0\u0001\u0000\u0000\u0000\u01e2\u01e3\u0001\u0000\u0000\u0000"+
		"\u01e3\u01e5\u0001\u0000\u0000\u0000\u01e4\u01e2\u0001\u0000\u0000\u0000"+
		"\u01e5\u01e6\u0006H\u0001\u0000\u01e6\u0092\u0001\u0000\u0000\u0000\u01e7"+
		"\u01e8\u0005/\u0000\u0000\u01e8\u01e9\u0005*\u0000\u0000\u01e9\u01ed\u0001"+
		"\u0000\u0000\u0000\u01ea\u01ec\t\u0000\u0000\u0000\u01eb\u01ea\u0001\u0000"+
		"\u0000\u0000\u01ec\u01ef\u0001\u0000\u0000\u0000\u01ed\u01ee\u0001\u0000"+
		"\u0000\u0000\u01ed\u01eb\u0001\u0000\u0000\u0000\u01ee\u01f0\u0001\u0000"+
		"\u0000\u0000\u01ef\u01ed\u0001\u0000\u0000\u0000\u01f0\u01f1\u0005*\u0000"+
		"\u0000\u01f1\u01f2\u0005/\u0000\u0000\u01f2\u01f3\u0001\u0000\u0000\u0000"+
		"\u01f3\u01f4\u0006I\u0000\u0000\u01f4\u0094\u0001\u0000\u0000\u0000\t"+
		"\u0000\u01ba\u01c0\u01c5\u01cb\u01d1\u01d3\u01e2\u01ed\u0002\u0006\u0000"+
		"\u0000\u0000\u0001\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}