package com.qztc.springboot.web.service;

import com.qztc.springboot.web.dao.IUserDao;
import com.qztc.springboot.web.domain.entity.UserEntity;
import jakarta.annotation.Resource;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserService {
    @Resource
    IUserDao userDao;

    public List<UserEntity> showUser() {
        List<UserEntity> userEntities = userDao.queryUserList();

        return userEntities;
    }
}
