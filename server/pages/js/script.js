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

function ShowTime() {
    let ctime = new Date();
    alert("Day: " + ctime.getDate() + ", Month: " + ctime.getMonth() + ", Year: " + ctime.getFullYear() + ", Hours: " + ctime.getHours() + ", Minutes: " + ctime.getMinutes());
}

function ShowMessage(msg) {
    alert(msg)
}

function SetWord() {
    let sName, tName;
    sName = prompt("Your name: ", "NoName");
    tName = confirm("Your name is: " + sName + " ?");
    if (tName === true) {
        alert("Welcome: " + sName + "!!!");
    }else {
        alert("you are spamer!");
    }
}