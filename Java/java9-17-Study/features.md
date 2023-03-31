# JDK 9

### 1. 模块化

新增 module-info.java 文件来管理 module

### 2. String 底层存储结构升级

将原来的 String 用 char 数组替换成 type 数组  
java 8：
```java
private final char[]value;
```
java 9：
```java
@Stable
private final byte[]value;
private final byte coder; // 如果使用的是 utf-8 字符则 coder 为 0，如果使用的是 utf-16 字符则 coder 为 1
```
优点：提高内存空间的使用率和减少 JVM 垃圾清理的工作量

### 3. 多版本兼容 jar
扩展 JAR 文件格式以允许多个特定于 Java 发行版的类文件在单个存档中共存

### 4. 接口私有方法
+ 私有静态方法
+ 私有方法  
>接口与抽象类的区别是什么呢?
> + 逐渐取代抽象类的功能特性，Java9接口私有方法则进一步代替了抽象类的职责
> + 接口针对于抽象类最大的优势在于多继承

> 那各自使用的时机要如何抉择呢?
> + 如果我们有多重继承的需求，我们应该考虑使用接口实现，如果有大量模板代码可以优先考虑抽象类，毕竟接口在定义参数的时候局限性比较大

### 5. JShell
是一个命令行工具，具有简化交互的功能

### 6. 增强的 try-with-resources
```java
public static void main(String[] args) {
    String s = "A";

    // Java 7 之前
    ByteArrayInputStream stream1 = new ByteArrayInputStream(s.getBytes());

    try {
        System.out.println(stream1.available());
    } catch (Exception e) {
        e.printStackTrace();
    } finally {
        try {
            stream1.close();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    // Java 7
//        TestDemo testDemo = new TestDemo();

    try (ByteArrayInputStream stream2 = new ByteArrayInputStream(s.getBytes());
         TestDemo testDemo = new TestDemo()) {
    } catch (Exception e) {
        throw new RuntimeException(e);
    }

    // Java 9
    TestDemo testDemo = new TestDemo();

    try(testDemo) {
        testDemo.exec();
    } catch (Exception e) {
        throw new RuntimeException(e);
    }
}
```

### 7. 语法改进
+ 对下划线使用限制
+ 集合增强
```java
// Java 8 - 不可变集合
List<String> list = Collections.unmodifiableList(new ArrayList<>());

// Java 9
List.of("1", "2", "3");
```

### 8. Stream & Optional增强
```java
public static void main(String[] args) {
    List<String> list1 = Stream.of("1", "2", "3").takeWhile("1"::equals).toList();
    System.out.println(list1);

    List<String> list2 = Stream.of("1", "2", "3").dropWhile("1"::equals).toList();
    System.out.println(list2);

    Stream.ofNullable(null).forEach(item -> {
        System.out.println("item" + item);
    });

    test(null);

    // ifPresentOrElse 方法
    Optional.ofNullable(null).ifPresentOrElse(System.out::println, () -> {
        System.out.println("empty list...");
    });

    // or 方法
    System.out.println(Optional.ofNullable(null).or(() -> Optional.of(List.of("4", "5", "6"))));

    // 直接转成流
    System.out.println(Optional.ofNullable("1").stream().toList());
}

public static void test(List<String> list) {
    list = Optional.ofNullable(list).orElse(new ArrayList<>());
    for (String s : list) {
        System.out.println(s);
    }
}
```

# JDK 10

### 1. 局部变量推导
为了省略不必要的局部变量类型的声明
```java
public static void main(String[] args) {
    String s = "1";

    ByteArrayInputStream bytes = new ByteArrayInputStream(s.getBytes());

    var bytesTwo = new ByteArrayInputStream(s.getBytes());

    // 无法使用 var
    Map<String, Map<String, List<String>>> maps = new HashMap<>();

    System.out.println(bytes.available());
    System.out.println(bytesTwo.available());
}
```
该功能并不是一把万能钥匙，我们还需要因地制宜的去使用它，这是一把双刃剑。用的好，可以简化代码，减少复杂度并提高可读性。用错了地方可能会造成反效果
