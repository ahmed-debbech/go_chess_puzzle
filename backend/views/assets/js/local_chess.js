var board 
var movesCount = 0

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


function buildBoard(data){


    orient = 'white'
    if(data.CurrentPlayer == 1){
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

      if(isRightMove(data.BestMoves[movesCount], source+target)){
        unHighlight(last_move_cell_start, last_move_cell_end)
        last_move_cell_start = source
        last_move_cell_end = target
        highlight(last_move_cell_start, last_move_cell_end)

        movesCount++
        computerPlays(adaptMove(data.BestMoves[movesCount]))
      }else{
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
      let mov = adaptMove(data.BestMoves[movesCount])
      console.log(mov)
      computerPlays(mov)
    }, 500)
}

function adaptMove(move){
  return move.split("").toSpliced(2,0,"-").join("").split('-')
}
function isRightMove(move1, move2){
  return move1 == move2
}
function computerPlays(move){
  unHighlight(last_move_cell_start, last_move_cell_end)

  last_move_cell_start = move[0]
  last_move_cell_end = move[1]
  highlight(last_move_cell_start, last_move_cell_end)
  board.move(move[0]+'-'+move[1])
  movesCount++

}
