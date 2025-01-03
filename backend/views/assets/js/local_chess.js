var board = null
var movesCount = 0
var game = null
var fen
var data
var pieceClickedOn = ''

function chessJsMove(source, target){
  console.log(source + target)
  var move = game.move({
    from: source,
    to: target,
    promotion: 'q' 
  })
  if (move === null) return false

  board.position(game.fen()) 
  return true
}

function showAvilableMoves (square, show) {
  var moves = game.moves({
    square: square,
    verbose: true
  })

  if (moves.length === 0) return

  if(show){
    // highlight the possible squares for this piece
    for (var i = 0; i < moves.length; i++) {
      highlightAvailable(moves[i].to)
    }
  }else{
    // unhighlight the possible squares for this piece
    for (var i = 0; i < moves.length; i++) {
      unHighlightAvailable(moves[i].to)
    }
  }
}

function onDragStart(source, piece, position, orientation){
  event.preventDefault(); 

  if (game.game_over()) return false

  if ((orient === 'white' && piece.search(/^w/) === -1) ||
      (orient === 'black' && piece.search(/^b/) === -1)) {
    
    return false
  }

  pieceClickedOn = source
  showAvilableMoves(source, true)
}

function movePieceAndRevert(source, target){
  board.move(source + "-" + target)
  setTimeout(() => {
    board.move(target + "-" + source)
    board.position(game.fen()) 
  }, 500)
}

function onDrop(source, target, piece, newPos, oldPos, orientation){
  document.body.style.overflow = '';

  pieceClickedOn = ''

  showAvilableMoves(source, false)

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
    if(getDraggingMode() == "0"){
      movePieceAndRevert(source, target)
    }
    updateStatus(-1)
    return 'snapback'
  }
}

function buildBoard(dd){
    data = dd
    game = new Chess(data.FEN)
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


    last_move_cell_start = ""
    last_move_cell_end = ""


    

    fen = data.FEN

    var config = initConfig()
    
    board = Chessboard('board1', config)

    setTimeout(() => {
      updateStatus(0)
      let mov = adaptMove(data.BestMoves[movesCount])
      computerPlays(mov)
    }, 500)
    fnInit()
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
function highlightAvailable(square) {
  $("#board1").find('.square-' + square )
  .addClass('highlight-available')
}
function unHighlightAvailable(square) {
  $("#board1").find('.square-' + square )
  .removeClass('highlight-available')
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


function fnInit(){
  $('.row-5277c').on('click', 'div[class^="square-"]', function() {
    // Your logic when the div is clicked
    console.log("div clicked")
    let squareName = this.id.substring(0,2)
    if(pieceClickedOn != '') {
      onDrop(pieceClickedOn, squareName)
      return
    }
    let s = game.get(squareName)
    s = (s.color + s.type.toUpperCase())
    console.log(this.id)
    for(let i=0; i<=$(this).children().length-1; i++){
      if($(this).children()[i].nodeName == "IMG"){
        console.log(squareName, s)
        if(s == null) return
        onDragStart(squareName, s)
      }
    }
  });
};

function initConfig(){
  var config
  if(getDraggingMode() == "1"){
    config = {
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
  }else{
    config = {
      orientation: orient,
      position: fen,
      draggable: false,
      dropOffBoard: 'snapback',
      moveSpeed: 'fast',
      snapbackSpeed: 100,
      snapSpeed: 100
    }
  }
  return config
}


function boardWithoutOrWithDrag(){

  if (board == null) return

  resetBoard()
  board = Chessboard('board1', initConfig())
  buildBoard(data);
}