let waitNum = 1;
let res;


function factorial(fact) {
    waitNum = 1;
    if (fact > 0) {
        if (isNaN(("fact"))) {
            for (let i = fact; i > 0; i--) {
                waitNum = waitNum * i;
            }
        }
        return waitNum;
    } else if (fact = 0) {
        return 1;
    } else { document.write("Geçerli bir sayı giriniz!") }
}

function combination(n, r) {
    if (n > 0 && r >= 0 && n >= r) {
        return factorial(n) / (factorial(n - r) * factorial(r));
    } else { document.write("Geçerli bir sayı giriniz!") }
}

function result() {
    let n = document.getElementById("n").value;
    let r = document.getElementById("r").value;
    document.getElementById("snc").innerHTML = combination(n, r);
}