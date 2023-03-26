package com.qztc.springboot.web.dao;

import com.qztc.springboot.web.domain.entity.UserEntity;
import org.apache.ibatis.annotations.Mapper;

import java.util.List;

@Mapper
public interface IUserDao {
    List<UserEntity> queryUserList();
}
