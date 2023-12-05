package com.hotel.pojo;

public class Vip {

    private int Id;
    private String Name;
    private String Sex;
    private long Card;
    private long Phone;
    private String V_Type;
    private String StartTime;
    private String EndTime;

    public Vip() {
    }

    public Vip(int id, String name, String sex, long card, long phone, String v_Type, String startTime, String endTime) {
        Id = id;
        Name = name;
        Sex = sex;
        Card = card;
        Phone = phone;
        V_Type = v_Type;
        StartTime = startTime;
        EndTime = endTime;
    }

    public int getId() {
        return Id;
    }

    public void setId(int id) {
        Id = id;
    }

    public String getName() {
        return Name;
    }

    public void setName(String name) {
        Name = name;
    }

    public String getSex() {
        return Sex;
    }

    public void setSex(String sex) {
        Sex = sex;
    }

    public long getCard() {
        return Card;
    }

    public void setCard(long card) {
        Card = card;
    }

    public long getPhone() {
        return Phone;
    }

    public void setPhone(long phone) {
        Phone = phone;
    }

    public String getV_Type() {
        return V_Type;
    }

    public void setV_Type(String v_Type) {
        V_Type = v_Type;
    }

    public String getStartTime() {
        return StartTime;
    }

    public void setStartTime(String startTime) {
        StartTime = startTime;
    }

    public String getEndTime() {
        return EndTime;
    }

    public void setEndTime(String endTime) {
        EndTime = endTime;
    }

    @Override
    public String toString() {
        return "Vip{" +
                "Id=" + Id +
                ", Name='" + Name + '\'' +
                ", Sex='" + Sex + '\'' +
                ", Card=" + Card +
                ", Phone=" + Phone +
                ", V_Type='" + V_Type + '\'' +
                ", StartTime='" + StartTime + '\'' +
                ", EndTime='" + EndTime + '\'' +
                '}';
    }
}
