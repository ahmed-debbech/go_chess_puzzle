var createdMoves = ""


function requestNew(obj){
    $("#loading_block").show()
    document.getElementById('board1').innerHTML = '';

    $.ajax({
        url: '/load',
        type: "GET",
        success: function (data) {
            if (data.ID == undefined) return

            $("#puzzle_id").html(data.ID)
            $("#gen_time").html(moment.unix(data.GenTime.substring(0, data.GenTime.length - 9)).format('DD/MM/YYYY'))
            $("#seen_times").html(data.SeenCount)
            $("#sol_times").html(data.SolveCount)

            obj.url_original = data.MatchLink

            resetBoard()
            buildBoard(data);
            obj.dataMatch = data
            $("#loading_block").hide()
            seen(data.ID)

        },

        error: function (error) {
            console.log(`Error ${error}`);
        }
    });
}

function seen(pid){
    $.ajax({
        url: '/seen?pid='+pid,
        type: "GET",
        success: function (data) {
            
        },

        error: function (error) {
            console.log(`Error ${error}`);
        }
    });
}

function setCookie(cname, cvalue, exdays) {
    const d = new Date();
    d.setTime(d.getTime() + (exdays*24*60*60*1000));
    let expires = "expires="+ d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}
function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for(let i = 0; i <ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
        c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
        }
    }
    return "";
}

function getUuid(){
    if(getCookie("chess_uuid") == ""){
        let uuid = "10000000-1000-4000-8000-100000000000".replace(/[018]/g, c =>
            (+c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> +c / 4).toString(16)
        );

        setCookie("chess_uuid", uuid, 365)
        return uuid
    }else{
        cookie_uuid = getCookie("chess_uuid")
        if(cookie_uuid != ""){
            return cookie_uuid
        }
        return null
    }
}

function addCreatedMove(move){
    createdMoves += move
}
function resetCreatedMoves(){
    createdMoves = ""
}

function calculateHash(content_hash){
    return sha256(content_hash)
}

function countSum(movesToCount){
    let sum = 0
    let numbersOnly = "";
    for(let i=0; i<=movesToCount.length-1; i++){
        if((movesToCount[i] <= '9') && (movesToCount[i] >= '0')){
            numbersOnly += movesToCount[i]
        }
    }
    for(let i=0; i<=numbersOnly.length-1; i+=2){
        let x = parseInt(numbersOnly[i]) * parseInt(numbersOnly[i+1])
        sum += x
    }
    return sum
}

function solved(id){
    let content_hash = id
    content_hash += countSum(createdMoves)
    content_hash += createdMoves
    let hash = calculateHash(content_hash)
    $.ajax({
        url: '/solved?pid=' + id + '&h=' + hash,
        type: "GET",
        success: function (data) {
            console.log(data)
        },

        error: function (error) {
            console.log(`Error ${error}`);
        }
    });
}
