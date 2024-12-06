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
    function onDrop(){
      document.body.style.overflow = '';
    }
    function onMoveEnd(){
        $("#board1").find('.square-' + squareToHighlight)
        .addClass('highlight-move')
    }

    var fen = data.FEN

    var config = {
      orientation: orient,
      position: fen,
      draggable: true,
      dropOffBoard: 'snapback',
      moveSpeed: 'slow',
      snapbackSpeed: 100,
      onDragStart: onDragStart,
      snapSpeed: 100,
      onDrop: onDrop,
      onMoveEnd: onMoveEnd

    }
    var board = Chessboard('board1', config)

    setTimeout(() => {
      let mov = data.BestMoves[0].split("").toSpliced(2,0,"-").join("")
      console.log(mov)
      board.move(mov)
    }, 500)
}