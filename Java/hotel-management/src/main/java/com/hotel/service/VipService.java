package com.hotel.service;

import com.hotel.pojo.Vip;

import java.util.ArrayList;

public interface VipService {

    void addVip(Vip vip);

    void deleteVipById(int id);

    void updateVipById(Vip vip);

    Vip queryVipById(int id);

    ArrayList<Vip> queryAllVip();

    Vip queryVipByPhone(String phone);
}
