package com.qztc.springboot.web.controller;

import com.qztc.springboot.web.domain.entity.UserEntity;
import com.qztc.springboot.web.service.UserService;
import jakarta.annotation.Resource;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/user")
public class UserController {
    @Resource
    UserService userService;

    @GetMapping
    public List<UserEntity> showUser() {
        List<UserEntity> userEntities = userService.showUser();

        return userEntities;
    }
}
