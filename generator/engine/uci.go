package engine

import (
	"fmt"
	"os"
	"os/exec"
	"errors"
	"bufio"
	"strings"
	"strconv"
)

type UCI struct{
	Name string
	Process *exec.Cmd
	Stdout *bufio.Reader
	Stdin *bufio.Writer
	FoundBM bool
}

func (uci *UCI) Init() error{

	stockfish := os.Args[3]
	fmt.Println("init ", stockfish)

	uci.Process = exec.Command(stockfish)
	uci.Name = "Stockfish 17" 
	uci.FoundBM = false

	stdout, err := uci.Process.StdoutPipe()
	if err != nil {
		fmt.Println("[ERROR] could not set up stdout pipe")
		return errors.New("[ERROR] could not set up stdout pipe");
	}
	stdoutReader := bufio.NewReader(stdout)
	uci.Stdout = stdoutReader

	stdin, err := uci.Process.StdinPipe()
	if err != nil {
		fmt.Println("[ERROR] could not set up stdin pipe")
		return errors.New("[ERROR] could not set up stdin pipe");
	}
	stdinWriter := bufio.NewWriter(stdin)
	uci.Stdin = stdinWriter

	fmt.Println("[SUCCESS] init engine")
	return nil
}

func (uci *UCI) Start() error{
	
	if err := uci.Process.Start(); err != nil {
		fmt.Println("[SUCCESS] init engine")
		return errors.New("[ERROR] could not start stockfish")
	}
	return nil
}

func (uci *UCI) setPosition(FEN string) error {
	_, err := uci.Stdin.WriteString("position fen "+FEN+"\n")
	if err != nil {
		fmt.Println("[ERROR] while writing position")
		return errors.New("[ERROR] while writing position")
	}
	fmt.Println("[SUCCESS] wrote position to SF: ", string("position fen "+FEN+"\n"))

	if err = uci.Stdin.Flush(); err != nil {
		fmt.Println("[ERROR] could not flush SF STDIN")
		return errors.New("[ERROR] could not flush SF STDIN")
	}

	return nil
}

func (uci *UCI) Go(level int) error {
	//nm, err := uci.Stdin.WriteString("go movetime "+ config.MoveTimeEngine +"\n")
	
	sslevel := strconv.Itoa(level)
	_, err := uci.Stdin.WriteString("go depth "+ sslevel +"\n")
	if err != nil {
		fmt.Println("[ERROR] when GO command issued")
		return errors.New("[ERROR] when GO command issued")
	}
	fmt.Println("[SUCCESS] wrote position to SF: ", string("go depth "+ sslevel +"\n"))

	if err = uci.Stdin.Flush(); err != nil {
		fmt.Println("[ERROR] could not flush SF STDIN")
		return errors.New("[ERROR] could not flush SF STDIN")
	}
	fmt.Println("[SUCCESS] started searching for bestmove...")
	return nil
}

func (uci * UCI) GetResultsBestMove() string{

	bestmove := ""
	for{
		n, _, _ := uci.Stdout.ReadLine();
		//fmt.Println("READING: ", string(n))	
		if strings.HasPrefix(string(n), "bestmove") {
			fmt.Println("BESTMOVE IS: ", string(n))	
			bestmove = string(n)
			uci.FoundBM = true
			break
		}
	}
	return bestmove
}

func (uci *UCI) WaitUntilEnd() error{
	err := uci.Process.Wait()
	if err != nil {
		fmt.Println("[ERROR] could not wait or end stockfish")
		return errors.New("[ERROR] could not wait or end stockfish")
	}
	return nil
} 
func (uci *UCI) Kill() error{
	if uci.FoundBM == true {
		err := uci.Process.Process.Kill()
		return err
	}
	return nil
} 