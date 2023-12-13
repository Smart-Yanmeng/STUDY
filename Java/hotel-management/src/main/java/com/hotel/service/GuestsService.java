package com.hotel.service;

import com.hotel.pojo.Guests;

import java.util.ArrayList;

public interface GuestsService {

    void addGuests(Guests guests);

    void deleteGuestsById(int id);

    void updateGuestsById(Guests guests);

    Guests queryGuestsById(int id);

    ArrayList<Guests> queryAllGuests();

    Guests queryGuestsByPhone(String phone);
}
