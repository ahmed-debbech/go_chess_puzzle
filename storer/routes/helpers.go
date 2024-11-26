package routes

import (
	"bufio"
	"io"
	"errors"
	"fmt"
	"net/http"
)

func getBody(r io.Reader) ([]byte, error) {
	
	bb := make([]byte, 0)

	reader := bufio.NewReader(r)
	for {
		char, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("[ERROR] could not read request body")
			return nil, err
		}
		bb = append(bb, char)
	}
	return bb, nil
}

func allowedMethod(r *http.Request, method string) error{
	if r.Method != method { 
		fmt.Println("[ERROR] StatusMethodNotAllowed")
		return errors.New("StatusMethodNotAllowed")
	}
	return nil
}