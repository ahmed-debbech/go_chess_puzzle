var board 
var movesCount = 0

function buildBoard(data){


    orient = 'white'
    if(data.CurrentPlayer == 0){
        orient = 'black'
        $("#who_to_move_white").hide()
        $("#who_to_move_black").show()
    }else{
        orient = 'white'
        $("#who_to_move_white").show()
        $("#who_to_move_black").hide()
    }

    function onDragStart(source, piece, position, orientation){
      event.preventDefault(); 
      if ((orient === 'white' && piece.search(/^w/) === -1) ||
          (orient === 'black' && piece.search(/^b/) === -1)) {
        
        onDrop()
        return false
      }
    }

    last_move_cell_start = ""
    last_move_cell_end = ""

    function onDrop(source, target, piece, newPos, oldPos, orientation){
      document.body.style.overflow = '';

      if(data.BestMoves.length <= movesCount) {$("#status").html("✅ SOLVED"); return 'snapback';}

      if(isRightMove(data.BestMoves[movesCount], source+target)){
        unHighlight(last_move_cell_start, last_move_cell_end)
        last_move_cell_start = source
        last_move_cell_end = target
        highlight(last_move_cell_start, last_move_cell_end)

        if(data.BestMoves.length <= movesCount) return 'snapback'

        updateStatus(1)

        setTimeout(() => {
          unHighlight(last_move_cell_start, last_move_cell_end)
          movesCount++

          if(data.BestMoves.length <= movesCount)  {$("#status").html("✅ SOLVED"); return;}

          computerPlays(adaptMove(data.BestMoves[movesCount]))

          if(data.BestMoves.length <= movesCount)  {$("#status").html("✅ SOLVED"); return;}

        },500)
      }else{
        updateStatus(-1)
        return 'snapback'
      }
    }

    var fen = data.FEN

    var config = {
      orientation: orient,
      position: fen,
      draggable: true,
      dropOffBoard: 'snapback',
      moveSpeed: 'fast',
      snapbackSpeed: 100,
      onDragStart: onDragStart,
      snapSpeed: 100,
      onDrop: onDrop
    }
    board = Chessboard('board1', config)

    setTimeout(() => {
      updateStatus(0)
      let mov = adaptMove(data.BestMoves[movesCount])
      console.log(mov)
      computerPlays(mov)
    }, 500)
}

function resetBoard(){
  board = Chessboard('board1')
  movesCount = 0
  updateStatus(0)
}

function updateStatus(mode){
  if(mode == 0){
    $("#status").html("Play a move..")
    $("#status").css({"color" : ""})
  }
  if(mode == 1){
    $("#status").html("Correct! keep going..")
    $("#status").css({"color" : "#15a51d"})
  }
  if(mode == -1){
    $("#status").html("Wrong! try again..")
    $("#status").css({"color" : "#dc3545"})
  }
}

function highlight(start, end) {
  $("#board1").find('.square-' + start )
  .addClass('highlight-move')
  $("#board1").find('.square-' + end )
  .addClass('highlight-move')
}

function unHighlight(start, end) {
  $("#board1").find('.square-' + start )
  .removeClass('highlight-move')
  $("#board1").find('.square-' + end )
  .removeClass('highlight-move')
}

function adaptMove(move){
  return move.split("").toSpliced(2,0,"-").join("").split('-')
}
function isRightMove(move1, move2){
  return move1 == move2
}

function highlightHint(start, end) {
  $("#board1").find('.square-' + start )
  .addClass('highlight-hint')
  $("#board1").find('.square-' + end )
  .addClass('highlight-hint')
}
function unHighlightHint(start, end) {
  $("#board1").find('.square-' + start )
  .removeClass('highlight-hint')
  $("#board1").find('.square-' + end )
  .removeClass('highlight-hint')
}

function computerPlays(move){
  last_move_cell_start = move[0]
  last_move_cell_end = move[1]
  highlight(last_move_cell_start, last_move_cell_end)
  board.move(move[0]+'-'+move[1])
  movesCount++

}

function hintPuzzle(moves){

  highlightHint(adaptMove(moves[movesCount])[0], adaptMove(moves[movesCount])[1])
  let ss = adaptMove(moves[movesCount])

  setTimeout(() => {
    unHighlightHint(ss[0], ss[1])
  },1500)

}