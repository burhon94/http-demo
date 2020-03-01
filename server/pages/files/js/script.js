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
    } else {
        alert("you are spamer!");
    }
}

//////////////////////////////CALCULATOR//////////////////////
const display = document.getElementById("calc-display");
a = "", b = "", c = "", d = "",
    aa = 0, bb = 0, cc = 0, pp = 0,
    e = false, s = false;

function ShowDisplay() {
    if (!e) {
        if (a !== "" || d !== "" || b !== "") {
            display.innerHTML = a + d + b;
        } else if (a === "" && d === "" && b === "") {
            display.innerHTML = "Cleared";
        }

    } else {
        display.innerHTML = "=" + c;
    }
}

function Clear() {
    a = "";
    b = "";
    c = "";
    d = "";
    e = false;
    s = false;
    ShowDisplay();
}

function SetN(n) {
    if (e === true) {
        e = false;
        s = false;
        Clear();
    }
    if (!s) {
        a += n.toString();
    } else {
        b += n.toString();
    }
    ShowDisplay();
}

function SetA(o) {
    if (a === "") {
        a = "0";
    }
    pp = o;
    s = true;
    if (e === true) {
        a = c;
        b = "";
        c = "";
    }
    e = false;
    switch (pp) {
        case 1:
            d = "+";
            break;
        case 2:
            d = "-";
            break;
        case 3:
            d = "*";
            break;
        case 4:
            d = "/";
            break;
        case 5:
            d = "^";
            break;
        case 6:
            d = "root";
            break;
    }
    ShowDisplay();
}

function Result() {
    if (b === "") {
        b = "0";
    }
    if (e === true) {
        a = c;
    }
    switch (pp) {
        case 1:
            c = parseFloat(a) + parseFloat(b);
            break;
        case 2:
            c = parseFloat(a) - parseFloat(b);
            break;
        case 3:
            c = parseFloat(a) * parseFloat(b);
            break;
        case 4:
            c = parseFloat(a) / parseFloat(b);
            break;
        case 5:
            c = Math.pow(parseFloat(a), parseFloat(b));
            break;
        case 6:
            c = Math.sqrt(parseFloat(a));
            break;
    }
    e = true;
    ShowDisplay();
}
