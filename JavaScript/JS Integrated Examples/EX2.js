let rand = Math.floor(Math.random() * 100);
let hak = 3;

function karsilastir() {
    var tahmin = parseInt(document.getElementById("inp").value)
    if (tahmin == rand) {
        document.write("Tebrikler! OMG, kahin :D ")
    } else {
        hak--;
        document.querySelector("h2").innerHTML = "Yanlış tahmin XXX <br> Belki bir şans daha..."
    }

    if (hak == 0) {
        document.write("Oyunu kaybettiniz! Doğru cevap: " + rand)
    }
}