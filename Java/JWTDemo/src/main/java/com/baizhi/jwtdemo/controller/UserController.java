package com.baizhi.jwtdemo.controller;

import com.baizhi.jwtdemo.entity.UserEntity;
import com.baizhi.jwtdemo.service.UserService;
import jakarta.annotation.Resource;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/user")
public class UserController {
    @Resource
    private UserService userService;

    @GetMapping("/list")
    public UserEntity login(UserEntity user) {
        UserEntity userEntity = userService.login(user);

        return userEntity;
    }
}
