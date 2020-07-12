let currentValue = "";
var selected = "";

function chooseSelectInfo(elm) {
    var select2 = document.getElementById('sel2');
    select2.style.visibility = 'visible';

    currentValue = elm.value;

    var dev = ["Development Concepts", "Programming Language", "Agile Software Development"]
    var db = ["SQL", "NoSQL", "Database Architecture"]
    var net = ["Routing Concepts", "OSI Model", "Router, Switches and more"];

    if (currentValue === "Software Development")
        createOptions(dev, select2)
    else if (currentValue === "Database Technologies")
        createOptions(db, select2)
    else if (currentValue === "Network Technologies")
        createOptions(net, select2)
}

function createOptions(options, elm) {
    elm.innerHTML = "<option selected disabled>Sub categories</option>";

    options.forEach(function (optionValue) {
        var opt = document.createElement('option');
        opt.value = optionValue;
        opt.innerHTML = optionValue;
        elm.add(opt)
    })
}

function GetSelected() {
    selected = document.getElementById('sel2');
    var result = selected.options[selected.selectedIndex].value;
    var cate = currentValue + ";" + result;

    document.getElementById('c_hidden').value = cate;
}