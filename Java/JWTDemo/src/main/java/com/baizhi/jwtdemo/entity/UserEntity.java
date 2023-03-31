package com.baizhi.jwtdemo.entity;

//public record UserEntity(Integer id, String username, String password) {
//}

import lombok.Data;

@Data
public class UserEntity {

    private Integer id;

    private String username;

    private String password;
}