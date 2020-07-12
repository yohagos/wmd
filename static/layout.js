var contentDiv = document.getElementsByClassName("content");
var wrapperDiv = document.getElementsByClassName("wrapper");

function errorComponent() { // benuetze diese methode da im code wo error stattfindet
    this.hideContentComponent()
    var errorDiv = document.createElement("div");
    errorDiv.className = "error"; // clase in css
    var error = document.createElement("div");
    error.innerHTML = "error";
    errorDiv.appendChild(error);
    wrapperDiv[0].appendChild(errorDiv)
}

function hideContentComponent() {
    contentDiv[0].style.display = "none";
}