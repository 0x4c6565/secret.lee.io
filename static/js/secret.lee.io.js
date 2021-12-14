var textInput = document.getElementById('text');
var form = document.getElementById('form');

function getText(uuid, encryptionKey) {
    $.ajax({
        url: '/secret/'+uuid,
        timeout: 90000,
        type: 'GET',
        dataType: "json",
        contentType: 'application/json',
        error: function(jqXHR, textStatus, errorThrown) {
            toastr.error(`Failed to retrieve text (HTTP ${jqXHR.status}): ${jqXHR.responseText}`, null, {timeOut: 2000, extendedTimeOut: 1000});
        },
        success: function(data, textStatus, jqXHR) {
            history.pushState("", document.title, "/");
            window.location.hash = "";
            textInput.value = '' + CryptoJS.AES.decrypt(data.text, encryptionKey).toString(CryptoJS.enc.Utf8);
            toastr.success("Text retrieved and burnt", null, {timeOut: 2000, extendedTimeOut: 1000});
        }
    });
}

function postText() {
    PasswordGenerator.length = 24;
    PasswordGenerator.symbols = false;
    var encryptionKey = PasswordGenerator.generate();
    var encryptedText = '' + CryptoJS.AES.encrypt(textInput.value, encryptionKey);

    $.ajax({
        url: '/secret',
        timeout: 90000,
        type: 'POST',
        dataType: "json",
        contentType: 'application/json',
        data: JSON.stringify({
            text: encryptedText
        }),
        error: function(jqXHR, textStatus, errorThrown) {
            toastr.error(`Failed to save text (HTTP ${jqXHR.status}): ${jqXHR.responseText}`, null, {timeOut: 2000, extendedTimeOut: 1000});
        },
        success: function(data, textStatus, jqXHR) {
            history.pushState(null, null, "/"+data.uuid);
            window.location.hash = "#"+encryptionKey;
            toastr.success("Text saved", null, {timeOut: 2000, extendedTimeOut: 1000});
        }
    });
}


form.onsubmit=function(e) {
    e.preventDefault();
    postText();
}

var uuidMatch = window.location.pathname.match(/([a-zA-Z0-9-]+)/);
var pwMatch = window.location.hash.match(/([a-zA-Z0-9]+)/);
if (uuidMatch != null && pwMatch != null) {
    getText(uuidMatch[1], pwMatch[1]);
}

textInput.focus();
