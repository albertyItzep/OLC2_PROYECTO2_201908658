function antlr4 { java -jar .\antlr-4.13.0-complete.jar $args }

antlr4 -o TestRig Control.g4
javac -cp ".\antlr-4.13.0-complete.jar" -g -Xlint  TestRig/Control*.java
java -cp ".\antlr-4.13.0-complete.jar;./TestRig" org.antlr.v4.gui.TestRig Control prog -gui .\..\entrada.txt