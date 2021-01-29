let timed;
let kont = false;
function ackapa() {
    if (kont == false) {
        timed = parseInt(document.getElementById("time").value);
    }

    if (timed == 0) {
        let cho = document.getElementById("resm");
        if (cho.style.display == "") {
            cho.style.display = "none";
        } else {
            cho.style.display = "";
        }
    }

    let yazı = document.getElementById("don");
    yazı.innerHTML = (timed / 1000) + "saniye kaldı..."

    timed = timed - 1000;
    kont = true;
    if (timed >= 0) {
        setTimeout(ackapa, 1000);
    } else {
        kont = false;
    }
}
