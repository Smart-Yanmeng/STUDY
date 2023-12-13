package com.hotel.service;

import com.hotel.dao.VipMapper;
import com.hotel.pojo.Vip;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;

@Service
public class VipServiceImpl implements VipService {

    @Autowired
    VipMapper vipMapper;

    public void addVip(Vip vip) {
        vipMapper.addVip(vip);
    }

    public void deleteVipById(int id) {
        vipMapper.deleteVipById(id);
    }

    public void updateVipById(Vip vip) {
        vipMapper.updateVipById(vip);
    }

    public Vip queryVipById(int id) {
        return vipMapper.queryVipById(id);
    }

    public ArrayList<Vip> queryAllVip() {
        return vipMapper.queryAllVip();
    }

    public Vip queryVipByPhone(String phone) {
        return vipMapper.queryVipByPhone(phone);
    }
}
