package com.hotel.service;

import com.hotel.pojo.Home;

import java.util.ArrayList;

public interface HomeService {

    void addHome(Home home);

    void deleteHomeById(int id);

    void updateHomeById(Home home);

    Home queryHomeById(Integer id);

    ArrayList<Home> queryAllHome();

    Home queryHomeByNum(int num);

    void updateH_TypeById(Home home);
}
