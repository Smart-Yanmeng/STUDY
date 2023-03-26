# 新时间日期API

## 1. 旧版日期时间的问题

​	在旧版本中JDK对于日期和时间这块的设计是非常差的

```java
public static void main(String[] args) throws ParseException {
        /**
         * 旧版日期时间设计的问题
         */
        // 1、设计不合理
        Date date = new Date(2023, 03, 03);
        System.out.println(date);

        // 2、时间格式化和解析操作是线程不安全的
        SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd");
        for (int i = 0; i < 50; i++) {
            new Thread(() -> {
                try {
                    System.out.println(sdf.parse("2023-03-03"));
                } catch (ParseException e) {
                    e.printStackTrace();
                }
            }).start();
        }
    }
```

1. 设计不合理，在java.util和java.sql的包中都有日期类，java.util.Date同时包含日期和时间的，而java.sql.Date仅仅包含日期，此外用于格式化和解析的类在java.text包下
2. 非线程安全，java.util.Date是非线程安全的，所有的日期类都是可变的，这是java日期类最大的问题之一
3. 时区处理麻烦，日期类并不提供国际化，没有时区支持

## 2. 新日期时间API介绍

JDK8中增加了一套全新的日期时间API，这套API设计合理，是线程安全的。新的日期及时间API位于java.time包中，下面是一些关键类：

+ LocalDate：表示日期，包含年月日，格式为2019-10-16
+ LocalTime：表示时间，包含时分秒，格式为16:38:54.158549300
+ LocalDateTime：表示日期时间，包含年月日，时分秒，格式为2018-09-06T15:33:56.750
+ DateTimeFormatter：日期时间格式化类
+ Instant：时间戳，表示一个特定的时间瞬间
+ Duration：用于计算2个时间（LocalTime，时分秒）的距离
+ Period：用于计算2个日期（LocalDate，年月日）的距离
+ ZonedDateTime：包含时区的时间

Java中使用的历法是ISO 8601日历系统，它是世界民用历法，也就是我们所说的公历。平年有365天，闰年是366天。此外Java8还提供了4套其他历法，分别是：

+ ThaiBuddhistDate：泰国佛教历
+ MingguoDate：中华民国历
+ JapaneseDate：日本历
+ HijrahDate：伊斯兰历

### 2.1 日期时间的常见操作

LocalDate，LocalTime以及LocalDateTime的操作

```java
public static void main(String[] args) {
        /**
         * JDK8 日期时间操作
         */
        // 1.创建指定的日期
        LocalDate date1 = LocalDate.of(2021, 05, 06);
        System.out.println("date1 = " + date1);

        // 2.得到当前日期
        LocalDate now = LocalDate.now();
        System.out.println("date2 = " + now);

        // 3.根据 LocalDate 对象获取对应的日期信息
        System.out.println("Year = " + now.getYear());
        System.out.println("Month = " + now.getMonth().getValue());
        System.out.println("Day = " + now.getDayOfMonth());
        System.out.println("Week = " + now.getDayOfWeek().getValue());

        /**
         * 时间操作
         */
        // 1.得到指定时间
        LocalTime time = LocalTime.of(5, 26, 33, 23145);
        System.out.println(time);

        // 2.获取当前的时间
        LocalTime now1 = LocalTime.now();
        System.out.println(now1);

        // 3.获取时间信息
        System.out.println("Hour = " + now1.getHour());
        System.out.println("Minute = " + now1.getMinute());
        System.out.println("Second = " + now1.getSecond());
        System.out.println("Nano = " + now1.getNano());

        /**
         * 日期时间类型 LocalDateTime
         */
        // 1.获取指定的日期时间
        LocalDateTime dateTime = LocalDateTime.of(2023, 03, 03, 12, 12, 33, 213);
        System.out.println(dateTime);

        // 2.获取当前的日期时间
        LocalDateTime now2 = LocalDateTime.now();
        System.out.println(now2);

        // 3.获取日期时间信息
        System.out.println("Year = " + now2.getYear());
        System.out.println("Month = " + now2.getMonth().getValue());
        System.out.println("Day = " + now2.getDayOfMonth());
        System.out.println("Week = " + now2.getDayOfWeek().getValue());
        System.out.println("Hour = " + now2.getHour());
        System.out.println("Minute = " + now2.getMinute());
        System.out.println("Second = " + now2.getSecond());
        System.out.println("Nano = " + now2.getNano());
    }
```

### 2.2 日期时间的修改和比较

```java
public static void main(String[] args) {
        /**
         * 日期时间的修改
         */
        LocalDateTime now = LocalDateTime.now();
        System.out.println("now = " + now);

        // 修改时间，对日期时间的修改，对已存在的 LocalDate 对象，创建了它的模板
        // 并不会修改原来的信息
        LocalDateTime localDateTime = now.withYear(1998);
        System.out.println("localDateTime = " + localDateTime);
        System.out.println(now.withMonth(10));
        System.out.println(now.withDayOfMonth(6));
        System.out.println(now.withHour(8));
        System.out.println(now.withMinute(15));

        // 在当前日期的基础上，加上或者减去指定的时间
        System.out.println("两天后 = " + now.plusDays(2));
        System.out.println("十年后 = " + now.plusYears(10));
        System.out.println("六个月后 = " + now.plusMonths(6));
        System.out.println("十年前 = " + now.minusYears(10));
        System.out.println("六个月前 = " + now.minusMonths(6));
        System.out.println("七天前 = " + now.minusDays(7));

        /**
         * 日期时间的比较
         */
        LocalDate now1 = LocalDate.now();
        LocalDate date = LocalDate.of(2020, 1, 3);

        // 在 JDK8 中要实现日期的比较
        System.out.println(now.isAfter(date.atStartOfDay())); // true
        System.out.println(now.isBefore(date.atStartOfDay())); // false
        System.out.println(now.isEqual(date.atStartOfDay())); // false
    }
```

注意：在进行日期时间修改的时候，原来的LocalDate对象是不会被修改的，每次操作都是返回一个新的LocalDate对象，所以在多线程场景下是数据安全的。

### 2.3 格式化和解析操作

在JDK8中，我们可以通过`java.time.format.DateTimeFormatter`类可以进行日期的解析和格式化操作

```java
public static void main(String[] args) {
        /**
         * 日期格式化操作
         */
        LocalDateTime now = LocalDateTime.now();

        // 指定格式，使用系统默认的格式
        DateTimeFormatter isoLocalDateTime = DateTimeFormatter.ISO_LOCAL_DATE_TIME;
        // 将日期时间转换为字符串
        String format = now.format(isoLocalDateTime);
        System.out.println("format = " + format);

        // 通过ofPattern 方法来指定特定的格式
        DateTimeFormatter dateTimeFormatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
        String format1 = now.format(dateTimeFormatter);
        System.out.println("format1 = " + format1);

        // 将字符串解析为一个日期时间类型
        LocalDateTime parse = LocalDateTime.parse("1997-05-06 22:45:16", dateTimeFormatter);
        System.out.println("parse = " + parse);
    }
```

