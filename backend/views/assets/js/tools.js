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