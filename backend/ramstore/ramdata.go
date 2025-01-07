package ramstore

import (
	"fmt"
	"strconv"
	"strings"
    "crypto/sha256"
	"encoding/hex"

	"github.com/ahmed-debbech/go_chess_puzzle/generator/config"
)

type RamStore struct{
	store map[string]string
}

var ramStoreInstance *RamStore = nil

func newRamStore() *RamStore{
	fmt.Println("Creating new RamStore")
	return &RamStore{ store: make(map[string]string) }
}

func GetRamStoreInstance() *RamStore{
	if ramStoreInstance == nil {
		ramStoreInstance = newRamStore()
	}
	return ramStoreInstance
}

func Set(pid string, hash string){
	ramStoreInstance.store[pid] = hash
	fmt.Println(ramStoreInstance)
}

func extractDigits(move string) int{
	p := 1
	for _, c := range move {
		if ('0'<=c) && ('9' >= c) {
			ss, _ := strconv.Atoi(string(c))
			p *= ss
		}
	}
	return p
}

func doHash(toHash string) string {
	hash := sha256.New()
	hash.Write([]byte(toHash))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func Calculate(pid string, bestmove [config.BestMovesNumber]string) string{
	hash := pid
	
	sum := 0
	necessary_moves := make([]string, 0)
	for i:=1; i<=len(bestmove)-1; i+=2 {
		necessary_moves = append(necessary_moves, bestmove[i])
		sum += extractDigits(bestmove[i])
	}
	
	hash += strconv.Itoa(sum)
	hash += strings.Join(necessary_moves, "")
	return doHash(hash)
}