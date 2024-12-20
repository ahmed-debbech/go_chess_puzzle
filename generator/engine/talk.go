package engine

import(
	"fmt"
	//"os"
	//"io"
	//"strings"

)

func extractBestMove(str string) (string){
	return str[9:13]
}

func GetBestMove(FEN string) (string){

	uci := &UCI{}
	if err := uci.Init(); err != nil{
		return ""
	}

	if err := uci.Start(); err != nil{
		return ""
	}
	fmt.Println("[SUCCESS] started stockfish")

	if err := uci.setPosition(FEN); err != nil{
		return ""
	}
	if err := uci.Go(); err != nil{
		return ""
	}

	bm := uci.GetResultsBestMove()
	fmt.Println("[SUCCESS] got best move (", bm ,")")
	
	if err := uci.Kill(); err != nil{
		return ""
	}

	fmt.Println("[SUCCESS] finished stockfish")
	return extractBestMove(bm)
}