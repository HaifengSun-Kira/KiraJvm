Java8 与 Java9  由于模块化，jre有了较大的改变。

Java9 中  没有了rt.jar (runtime)

目前的手段是在  kira-jvm  docker中，运行。

debian 中 jre 路径 /usr/lib/jvm/java-8-oracle/jre

完成第三章，对于classfile的解析之后，如果想要解析标准的class file。例如`java.lang.String`，可以运行如下命令：
`./gojvm -Xjre "/usr/lib/jvm/java-8-oracle/jre" java.lang.String`
如果想要解析自定义的java类，首先需要将java文件编译到指定目录。
例如，在`KiraJvm/java/`目录下，使用如下命令，可以将`KiraJvm/java/src/gojvm/ClassFileTest.java`文件编译到`KiraJvm/java/class/`目录下。（gojvm为包目录）
`javac src/gojvm/ClassFileTest.java -d class`
然后，在`KiraJvm/bin`目录下，运行如下命令，解析自定义类的class file:
`./gojvm -Xjre "/usr/lib/jvm/java-8-oracle/jre" -classpath /home/kira/KiraJvm/java/class/ gojvm.ClassFileTest`
