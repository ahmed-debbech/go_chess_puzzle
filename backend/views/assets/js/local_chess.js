var board 
var movesCount = 0
var game = null
var fen

function chessJsMove(source, target){
  var move = game.move({
    from: source,
    to: target,
    promotion: 'q' 
  })
  if (move === null) return false

  board.position(game.fen(), false)
  return true
}

function buildBoard(data){

    game = new Chess(data.FEN)
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

        //if(data.BestMoves.length <= movesCount) return 'snapback'
        chessJsMove(source, target)
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

    fen = data.FEN

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
      computerPlays(mov)
    }, 500)
}

function resetBoard(){
  board = Chessboard('board1')
  movesCount = 0
  updateStatus(0)
  game = new Chess(fen)
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
  let x = []
  x.push(move.substring(0,2))
  x.push(move.substring(2,4))
  if(move.length == 5) x.push(move.substring(4,5))
  else x.push("")
  return x
}
function isRightMove(move1, move2){
  let f = adaptMove(move1)
  return (f[0] + f[1]) == move2
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
  //board.move(move[0]+'-'+move[1])
  chessJsMove(move[0], move[1])

  movesCount++
}

function hintPuzzle(moves){

  highlightHint(adaptMove(moves[movesCount])[0], adaptMove(moves[movesCount])[1])
  let ss = adaptMove(moves[movesCount])

  setTimeout(() => {
    unHighlightHint(ss[0], ss[1])
  },1500)

}