var lightD = document.getElementById('defaultlight');
var lightR = document.getElementById('redlight');
var lightG = document.getElementById('greenlight');

// TODO：完善此函数 显示红色颜色的灯
function red() {
    setTimeout("lightR.style.display = 'inline', lightD.style.display = 'none', lightG.style.display = 'none'", 3000);
}

// TODO：完善此函数  显示绿色颜色的灯
function green() {
    setTimeout("lightR.style.display = 'none', lightD.style.display = 'none', lightG.style.display = 'inline'", 6000);
}

// TODO：完善此函数
function trafficlights() {
    red();
    green();
}

trafficlights();
