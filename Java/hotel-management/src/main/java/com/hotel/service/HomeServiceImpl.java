package com.hotel.service;

import com.hotel.dao.HomeMapper;
import com.hotel.pojo.Home;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;

@Service
public class HomeServiceImpl implements HomeService {

    @Autowired
    HomeMapper homeMapper;


    public void addHome(Home home) {
        homeMapper.addHome(home);
    }

    public void deleteHomeById(int id) {
        homeMapper.deleteHomeById(id);
    }

    public void updateHomeById(Home home) {
        homeMapper.updateHomeById(home);
    }

    public Home queryHomeById(Integer id) {
        return homeMapper.queryHomeById(id);
    }

    public ArrayList<Home> queryAllHome() {
        return homeMapper.queryAllHome();
    }

    public Home queryHomeByNum(int num) {
        return homeMapper.queryHomeByNum(num);
    }

    @Override
    public void updateH_TypeById(Home home) {
        homeMapper.updateH_TypeById(home);
    }

}
