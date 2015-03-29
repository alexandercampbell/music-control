

window.onload = function() {
    var bbc = document.getElementById("bbc");
    bbc.onclick = function() {
        var xmlhttp = new XMLHttpRequest;
        xmlhttp.open("POST", "/start-bbc", true);
        xmlhttp.send();
        location.reload();
    }

    var police = document.getElementById("police");
    police.onclick = function() {
        var xmlhttp = new XMLHttpRequest;
        xmlhttp.open("POST", "/start-police", true);
        xmlhttp.send();
        location.reload();
    }
};

