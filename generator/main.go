package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ahmed-debbech/go_chess_puzzle/generator/chess"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/logic"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/utils"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/engine"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/config"
	"github.com/ahmed-debbech/go_chess_puzzle/generator/data"
)


func main() {
	
	if len(os.Args) <= 3{
		fmt.Println("NOT ENOUGH ARGS!!")
		os.Exit(1)
	} 

	max, yes := utils.IsNumber(os.Args[1]);
	if yes == false {
		fmt.Println("ERROR: Set the maximum number correctly")
		os.Exit(1)
	}

	if !utils.IsDirectoryExist(os.Args[2]) {
		fmt.Println("ERROR: Not a valid directory to read from.")
		os.Exit(1)
	}

	if !utils.IsDirectoryExist(os.Args[3]) {
		fmt.Println("ERROR: could not find Stockfish exec.")
		os.Exit(1)
	}

	for{
		fmt.Println("Looking for a game match from this directory: ", os.Args[2])

		match_content, id := logic.LookupMatch(os.Args[2], max)
		fmt.Println("[SUCCESS] MATCH FOUND! ID: ", id)

		game := chess.ObjectifyGame(match_content)
		fmt.Println(game.Position().Board().Draw())

		gameWithRandPos, randNum := chess.JumpToRandPosition(game.Clone())
		//gameWithRandPos, randNum := chess.JumpToBeforeCheckmate(game)

		if !chess.IsGameEligible(game, randNum) {continue}

		FEN := chess.GenerateFen(gameWithRandPos)

		fmt.Println(FEN)

		movesNumber := config.BestMovesNumber
		newfen := FEN
		var bestmvs [config.BestMovesNumber]string
		tu := 0
		for chess.IsGameOver(gameWithRandPos) {
			var bestmove string
			if tu % 2 == 0 {
				bestmove = engine.GetBestMove(newfen, 24)
			}else{
				bestmove = engine.GetBestMove(newfen, 4)
			}
			tu++
			if(bestmove == "") {break}
			//bestmvs[i] = bestmove
			newfen = chess.MakeMoveAndFEN(gameWithRandPos, bestmove)
			fmt.Println(gameWithRandPos.Position().Board().Draw())
			if newfen == "" { panic("SYSTEM HALT: because of error happened in generating a new fen") }
		}
		fmt.Println("final play after stockfish")
		fmt.Println(gameWithRandPos.Position().Board().Draw())

		if !chess.IsCheckmate(gameWithRandPos) {continue}

		newfen = chess.NavToMove(config.BestMovesNumber, gameWithRandPos)

		bestmvs = chess.GetFinalBestMoves(gameWithRandPos)

		whosPlaying := chess.DetermineTurn(gameWithRandPos)


		fmt.Println("[SUCCESS] all best ", movesNumber, " moves have been generated. ", bestmvs)

		puzzle := data.Puzzle{
			ID: strconv.Itoa(id),
			FEN: newfen,
			BestMoves: bestmvs,
			GenTime: strconv.Itoa(int(time.Now().UnixNano())),
			CurrentPlayer: whosPlaying,
			SolveCount: 0,
			MatchLink: game.GetTagPair("Site").Value,
			SeenCount: 0,
			FirstSeenTime: "",
		}
		fmt.Println("[SUCCESS] generated puzzle " ,puzzle.String())
		err := utils.SendToStore(puzzle)
		if err != nil {
			fmt.Println(err)	
		}
		time.Sleep(2 * time.Second)
	}
}