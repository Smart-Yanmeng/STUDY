package com.hotel.dao;

import com.hotel.pojo.Vip;
import org.springframework.stereotype.Component;

import java.util.ArrayList;

@Component
public interface VipMapper {

    int addVip(Vip vip);

    int deleteVipById(int id);

    int updateVipById(Vip vip);

    Vip queryVipById(int id);

    ArrayList<Vip> queryAllVip();

    Vip queryVipByPhone(String phone);
}
