package com.baizhi.jwtdemo.service;

import com.baizhi.jwtdemo.dao.IUserDao;
import com.baizhi.jwtdemo.entity.UserEntity;
import jakarta.annotation.Resource;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserService {
    @Resource
    private IUserDao userDao;

    public UserEntity login(UserEntity user) {
        return userDao.login(user);
    }
}
