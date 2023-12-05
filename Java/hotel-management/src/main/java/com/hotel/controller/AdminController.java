package com.hotel.controller;

import com.hotel.pojo.Admin;
import com.hotel.service.AdminServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("/admin")

public class AdminController {

    @Autowired
            AdminServiceImpl adminService;

    @RequestMapping("/login")
    public String login(String username, int password){

        Admin admin=adminService.findAdmin(username,password );
       if(admin!=null){
           if(admin.getPassword()==password){
               return "index";
           }else {
                return "error";
           }
       }
        return "error";
    }

    @RequestMapping("/updatePwd")
    public String update(Admin admin){
        adminService.updatePwd(admin);
        System.out.println(admin);
        return "suc_a";
    }
}
