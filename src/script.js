"use strict";
function alertBox()
{
    alert("Здесь можно найти пасхалку. Поищи странности");
}

async function showHelpLabel() {
    let inv = "Click on me<br>"
    let text = inv;
    for (let i = 0; i < 3; i++){
    let myPromise = new Promise(function(resolve) 
        {
            setTimeout(function() {resolve(text);}, 5000);
        });
    document.getElementById("easter_egg_help").innerHTML = await myPromise;
    text += inv;
    }
}
showHelpLabel()
    .then((response) => 
        {console.log("Всё прошло нормально");})
    .catch((error)   => 
        {console.error(error)});
