
function giveH() {
    let edge_1 = parseInt(document.getElementById("edg1").value);
    let edge_2 = parseInt(document.getElementById("edg2").value);

    if (edge_1 >= 0 || edge_2 >= 0) {
        document.querySelector("h2").innerHTML = "RESULT :  " + Math.sqrt(Math.pow(edge_1, 2) + Math.pow(edge_2, 2));
    } else {
        document.querySelector("h2").innerHTML = "An edge must be a positive integer! Pythogoras bones are aching..."
    }
}



