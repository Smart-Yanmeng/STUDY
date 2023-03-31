//package com.baizhi.jwtdemo;
//
//import com.auth0.jwt.JWT;
//import com.auth0.jwt.JWTVerifier;
//import com.auth0.jwt.algorithms.Algorithm;
//import com.auth0.jwt.interfaces.DecodedJWT;
//import org.junit.Test;
//
//import java.util.Calendar;
//import java.util.HashMap;
//
//public class TestDemo {
//    @Test
//    public void contextLoads() {
//        HashMap<String, Object> map = new HashMap<>();
//
//        Calendar instance = Calendar.getInstance();
//        instance.add(Calendar.MINUTE, 30);
//
//        String token = JWT.create()
//                .withHeader(map) // Header
//                .withClaim("userId", 1) // Payload
//                .withClaim("username", "York")
//                .withExpiresAt(instance.getTime()) // 指定令牌过期时间
//                .sign(Algorithm.HMAC256("SHaRwaSS675%^$%&YORK44$"));// Signature
//
//        System.out.println(token);
//    }
//
//    @Test
//    public void test() {
//        // 创建验证对象
//        JWTVerifier jwtVerifier = JWT.require(Algorithm.HMAC256("SHaRwaSS675%^$%&YORK44$")).build();
//
//        DecodedJWT verify = jwtVerifier.verify("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2ODAxNTk5MDYsInVzZXJJZCI6MSwidXNlcm5hbWUiOiJZb3JrIn0.3d0vOHTZvUR77RUbzv36Xu5AmgH6nCLmP1dbW4DkiLI");
//
//        System.out.println(verify.getClaim("userId").asInt());
//        System.out.println(verify.getClaim("username").asString());
//    }
//}
