let array = ["elma", "armut", "portakal", "limon", "havuç"]
let secilen = document.getElementById("menu");
let siparis = document.getElementById("order");

let list = "";
for (let i = 0; i < array.length; i++) {
    let noktalama;
    if (i == array.length - 1) {
        noktalama = "";
    } else {
        noktalama = ", ";
    }
    list = list + (array[i] + noktalama);
    secilen.innerHTML = "<br>Menü : " + list;
}

function ekle() {
    let ekli = document.getElementById("inp").value;
    if (ekli == "") { 
        alert("Bilgi giriniz!")
    } else {
        array.push(ekli);
        document.getElementById("inp").value = "";
        list = list + ", " + array[array.length - 1];
        secilen.innerHTML = "<br>Menü : " + list;
    }
}

let rand;
function sec() {
    rand = parseInt((Math.random() * 100) % array.length);
    siparis.innerHTML = "<br>Siparişiniz : " + array[rand];
}
