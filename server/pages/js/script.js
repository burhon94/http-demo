//////////////////////////////BACKGROUND ANIMATION//////////////////////
setInterval('AnimBg()', 10);
let anim = 1;

function AnimBg() {
    anim -= 0.3;
    document.body.style.backgroundPosition = -(anim / 9) + "% 0%," + anim + "% 0%";
}

//////////////////////////////DROP DOWN IMAGE//////////////////////
function imgDrop(element) {
    let style;
    element = element.getElementsByTagName("img")[0].style;
    style = element.display;
    if (style === "none") {
        element.display = "block"
    } else {
        element.display = "none"
    }
}