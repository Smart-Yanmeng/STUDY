package com.hotel.controller;

import com.hotel.pojo.Vip;
import com.hotel.service.VipServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.ModelAndView;

import java.util.ArrayList;
import java.util.List;

@Controller
@RequestMapping("/vip")
public class VipController {

    @Autowired
            VipServiceImpl vipService;

    @RequestMapping("/add")
    public ModelAndView add(Vip vip){
        ModelAndView mv = new ModelAndView();
        vipService.addVip(vip);
        mv.setViewName("suc_v");
        return mv;
    }

    @RequestMapping("/delete")
    public String delete(int id){
        vipService.deleteVipById(id);
        return "redirect:/vip/list";
    }

    @RequestMapping("/list")
    public ModelAndView list(){
        ModelAndView mv = new ModelAndView();
        List<Vip> vipList=vipService.queryAllVip();
        mv.addObject("list",vipList);
        mv.setViewName("vip_list");
        return mv;
    }

    @RequestMapping("/update1")
    public ModelAndView update1(int id){
        ModelAndView mv = new ModelAndView();
        Vip vip = vipService.queryVipById(id);
        mv.addObject("v",vip);
        mv.setViewName("vip_update");
        return mv;
    }

    @RequestMapping("/update2")
    public String update2(Vip v){
        vipService.updateVipById(v);
        return ("redirect:/vip/list");
    }

    @RequestMapping("/find")
    public ModelAndView find(String findByPhone ){
        ModelAndView mv = new ModelAndView();
        Vip vip = vipService.queryVipByPhone(findByPhone);
        List<Vip> vipList=new ArrayList<Vip>();
        vipList.add(vip);
        if (vip==null){
            vipList=vipService.queryAllVip();
            mv.addObject("error","未查询出结果");
        }
        mv.addObject("list",vipList);
        mv.setViewName("vip_list");
        return mv;
    }
}

