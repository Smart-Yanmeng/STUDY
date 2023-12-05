<%@ page contentType="text/html;charset=UTF-8" language="java" %>

<html lang="zh">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <title>后台首页</title>
    <link rel="stylesheet" href="../views/static/common/layui/css/layui.css">
    <link rel="stylesheet" href="../views/static/admin/css/style.css">
    <script src="../views/static/common/layui/layui.js"></script>
    <script src="../views/static/common/jquery-3.3.1.min.js"></script>
    <script src="../views/static/common/vue.min.js"></script>
    <style>
        .right h2{
            margin: 10px 0;
        }
        .right li{
            margin-bottom: 10px;
        }
    </style>
</head>
<body>
<div id="app">
    <!--顶栏-->
    <header>
        <div style="float:left;"><img src="../views/static/admin/img/tubiao.png" width="152" height="55" /></div>
        <div style="float: right">
            <i  class="layui-icon  layui-anim layui-anim-rotate layui-anim-loop" style="font-size: 30px; color: #969baa;">&#xe614;</i>&nbsp; &nbsp; &nbsp;
        </div>
    </header>

    <div class="main" id="">
        <!--左栏-->
        <div class="left">
            <ul class="cl" >
                <!--顶级分类-->
                <li v-for="vo,index in menu" :class="{hidden:vo.hidden}">
                    <a href="javascript:;" :class="{active:vo.active}" @click="onActive(index)">
                        <i class="layui-icon" v-html="vo.icon"></i>
                        <span v-text="vo.name"></span>
                        <i class="layui-icon arrow" v-show="vo.url.length==0">&#xe61a;</i> <i v-show="vo.active" class="layui-icon active">&#xe623;</i>
                    </a>
                    <!--子级分类-->
                    <div v-for="vo2,index2 in vo.list">
                        <a href="javascript:;" :class="{active:vo2.active}" @click="onActive(index,index2)" v-text="vo2.name"></a>
                        <i v-show="vo2.active" class="layui-icon active">&#xe623;</i>
                    </div>
                </li>
            </ul>
        </div>
        <!--右侧-->
        <div class="right">
            <blockquote class="layui-elem-quote">
                    <h2>欢迎登陆温情小筑管理系统</h2>

                    <marquee  vspace="15px">
                        <font face="宋体" size="5px" color="fuchsia" >梦见四季，四季有家，温情小筑，欢迎入住！</font>
                    </marquee>
                <%-- 显示系统当前时间--%>
                <div id="datetime">
                    <script>
                        setInterval("document.getElementById('datetime').innerHTML=new Date().toLocaleString();", 1000);
                    </script>
                </div>

            </blockquote>

            <div class="layui-row">
                <div class="layui-col-md4">
                    <div class="layui-card">
                        <div class="layui-card-header"><h2>系统介绍</h2></div>
                        <div class="layui-card-body" style="height: 250px;">
                            <ul>
                                <li>该是一个轻量级的后台管理系统,基于layui框架</li>
                                <li>风格简洁简单,中小型企业首选</li>
                                <li>您可以通过该面板对如下信息进行管理</li>
                                <li><a href="#" target="_blank" class="layui-btn layui-btn-radius layui-btn-normal">住&nbsp;客&nbsp;信&nbsp;息</a></li>
                                <li><a href="#" target="_blank" class="layui-btn layui-btn-radius">房&nbsp;间&nbsp;信&nbsp;息</a></li>
                                <li><a href="#" target="_blank" class="layui-btn layui-btn-radius layui-btn-warm">会&nbsp;员&nbsp;信&nbsp;息</a></li>
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="layui-col-md4">
                    <div class="layui-card">
                        <div class="layui-card-header"><h2>系统使用指南</h2></div>
                        <div class="layui-card-body" style="height: 250px;">
                            <ul>
                                <li><a href="#">如何访问网站</a></li>
                                <li><a href="#">如何管理信息</a></li>
                                <li><a href="#">后台用户设置(权限)</a></li>
                                <li><a href="#">系统设置</a></li>
                                <li>当前版本: v1.0</li>
                                <li><b>关注官网或github随时下载最新版</b></li>
                                <li>
                                    <a href="http://www.qadmin.net/" target="_blank" class="layui-btn layui-btn-normal">进入官网</a>
                                    <a href="https://github.com/aa24615/qadmin" target="_blank" class="layui-btn layui-btn-normal">github</a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>

                <div class="layui-col-md4">

                    <div class="layui-card">
                        <div class="layui-card-header"><h2>系统特点</h2></div>
                        <div class="layui-card-body" style="height: 250px;">
                            <ul>
                                <li>轻量级</li>
                                <li>简洁直爽</li>
                                <li>基于layui框架,上手快</li>
                                <li>支持jq/vue/layui</li>
                                <li>...</li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    </div>
</div>
<script src="../views/static/admin/js/config.js"></script>
<script src="../views/static/admin/js/script.js"></script>
<script type="text/javascript">

</script>
</body>
</html>
