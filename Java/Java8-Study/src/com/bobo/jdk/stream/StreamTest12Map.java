package com.bobo.jdk.stream;

import java.util.stream.Stream;

public class StreamTest12Map {
    public static void main(String[] args) {
        Stream.of("1", "2", "3", "4", "5", "6")
                .map(Integer::parseInt)
                .forEach(System.out::println);
    }
}
