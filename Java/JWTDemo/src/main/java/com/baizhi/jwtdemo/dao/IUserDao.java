package com.baizhi.jwtdemo.dao;

import com.baizhi.jwtdemo.entity.UserEntity;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface IUserDao {
    UserEntity login(UserEntity user);
}
