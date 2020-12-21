function volumeUp() {
    window.navigator.vibrate(50);
    return $.post({
        url: '/performAction', data: JSON.stringify({ action: "volumeUp" }), contentType: 'application/json'
    });
}

function volumeDown() {
    window.navigator.vibrate(50);
    return $.post({ url: '/performAction', data: JSON.stringify({ action: "volumeDown" }) , contentType: 'application/json'});
}

function brightnesUp() {
    window.navigator.vibrate(50);
    return $.post({ url: '/performAction', data: JSON.stringify({ action: "brightnessUp" }) , contentType: 'application/json'});
}

function brightnesDown() {
    window.navigator.vibrate(50);
    return $.post({ url: '/performAction', data: JSON.stringify({ action: "brightnessDown" }) , contentType: 'application/json'});
}
function powerOn() {
    window.navigator.vibrate(50);
    return $.post({ url: '/performAction', data: JSON.stringify({ action: "on" }) , contentType: 'application/json'});
}
function powerOff() {
    window.navigator.vibrate(50);
    return $.post({ url: '/performAction', data: JSON.stringify({ action: "off" }) , contentType: 'application/json'});
}

$(document).ready(function () {
});