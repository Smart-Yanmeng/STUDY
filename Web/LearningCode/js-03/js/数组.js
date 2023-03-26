var num1 = 1;
var num2 = 2;
var num3 = 3;

// 数组是指一组数据的集合，在数组中可以存放任意类型的元素
var numList = [1, 2, 3, '小明', true, NaN, {name: "小明"}];

// 创建数组
var arrList1 = [1, 2, 3];
var arrList2 = new Array();

// 数组下标，获取元素时如果下标超出数组长度，则会返回 undefined
arrList2[0] = 1;
arrList2[1] = 2;

// 遍历数组
for (var i = 0; i < numList.length; i++) {
    console.log(numList[i]);
}