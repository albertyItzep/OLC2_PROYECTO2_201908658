function antlr4 { java -Xmx500M -cp ".\antlr-4.13.0-complete.jar;$CLASSPATH" org.antlr.v4.Tool $args }
antlr4 -Dlanguage=Go -o ../Background -package Background -visitor *.g4