package com.hotel.service;

import com.hotel.dao.GuestsMapper;
import com.hotel.pojo.Guests;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;

@Service
public class GuestsServiceImpl implements GuestsService {

    @Autowired
    GuestsMapper guestsMapper;

    public void addGuests(Guests guests) {
        guestsMapper.addGuests(guests);
    }

    public void deleteGuestsById(int id) {
        guestsMapper.deleteGuestsById(id);
    }

    public void updateGuestsById(Guests guests) {
        guestsMapper.updateGuestsById(guests);
    }

    public Guests queryGuestsById(int id) {
        return guestsMapper.queryGuestsById(id);
    }

    public ArrayList<Guests> queryAllGuests() {
        return guestsMapper.queryAllGuests();
    }

    public Guests queryGuestsByPhone(String phone) {
        return guestsMapper.queryGuestsByPhone(phone);
    }


}
