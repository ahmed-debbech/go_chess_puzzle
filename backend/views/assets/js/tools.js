var createdMoves = ""


function requestNew(obj){
    $("#loading_block").show()
    document.getElementById('board1').innerHTML = '';

    $.ajax({
        url: '/load',
        type: "GET",
        success: function (data) {
            console.log(data)
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

        },

        error: function (error) {
            console.log(`Error ${error}`);
        }
    });
}

function getUuid(){
    console.log("log")
    if(localStorage.getItem("chess_uuid") == null){
        let uuid = "10000000-1000-4000-8000-100000000000".replace(/[018]/g, c =>
            (+c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> +c / 4).toString(16)
        );

        localStorage.setItem("chess_uuid", uuid)
        return uuid
    }else{
        return localStorage.getItem("chess_uuid")
    }
}

function addCreatedMove(move){
    createdMoves += move
    console.log(createdMoves)
}
function resetCreatedMoves(){
    createdMoves = ""
}

function calculateHash(){

}

function countSum(){

}

function solved(id){
    
}
