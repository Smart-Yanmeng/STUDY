# QZNU - WEBPROJECT

------

## 小组成员

211308121 - 蒋天宇
211308067 - 胡林峰

## 项目简介

### 文件分类

### HTML

HTML文件夹中总共三个文件，分别为

+ index.html - 主页
+ list.html - 栏目页
+ page.html - 内容页

他们之间的连接通过a标签定位实现：

+ 导航栏中的第一个按钮（学校主页）对应主页
+ 导航栏中的第三个按钮（通知通告）对应栏目页
+ 通知通告板块下的第一条新闻对应内容页

### CSS

CSS文件夹中总共四个文件，分别为

+ index.css - 主页内容样式
+ list.css - 栏目页内容样式
+ page.css - 内容页内容样式
+ public.css - 公用内容样式

### JS

JS文件夹中总共一个文件，为

+ main.js - 所有的脚本管理，包括轮播图等

### IMAGE

IMAGE文件夹中保存的是网页中需要的图片文件

## 媒体查询

本项目对屏幕大小做了媒体查询，分别为

+ SCREEN-WIDTH > 1140px

  ```css
  @media screen and (min-width: 1140px) { ... }
  ```

+ CREEN-WIDTH > 0px

  ```css
  @media screen and (min-width: 0) and (max-width: 1139px) { ... }
  ```

