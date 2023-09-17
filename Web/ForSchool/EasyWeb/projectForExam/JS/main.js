const flag = 0;

function rotation(flag) {
    let imgArr = document.querySelectorAll("#container .campus-news .content-box .rotation img")
    let imgSelectorArr = document.querySelectorAll("#container .campus-news .content-box .rotation .image-select a")

    init(flag++)
    let timer = setInterval(() => {
        init(flag)

        imgArr[flag % 6].className = "";
        imgSelectorArr[flag % 6].className = "opa-70 back-ff0000-color";

        flag++;
    }, 5000)

    imgSelectorArr.forEach((index, value) => {
        index.onclick = () => {
            clearInterval(timer);

            rotation(value)
        }
    })

    function init(flag) {
        imgArr.forEach((index) => {
            index.className = "blocked";
        })
        imgSelectorArr.forEach((index) => {
            index.className = "opa-70 back-6c6d6e-color"
        })

        imgArr[flag % 6].className = "";
        imgSelectorArr[flag % 6].className = "opa-70 back-ff0000-color";
    }
}

function school() {
    let speed = 50;
    let rolling = document.querySelector("#container .mod-third .rolling");

    rolling.innerHTML += rolling.innerHTML;
    let timer1 = setInterval(marquee, speed);
    rolling.onmouseover = function () {
        clearInterval(timer1);
    }
    rolling.onmouseout = function () {
        timer1 = setInterval(marquee, speed);
    };

    function marquee() {
        if (rolling.scrollLeft > 1224) {
            rolling.scrollLeft = 0;
        } else {
            rolling.scrollLeft++;
            console.log(rolling.scrollLeft)
        }
    }
}

school();
rotation(flag);