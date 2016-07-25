package api

import (
	"net/http"
	"encoding/json"
	"fmt"
	"bufio"
	"io"
)

type Response struct {
	Status  int
	Message string
	Body    string
}

func GetHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o := Response{Status: 0, Message: "", Body: ""}
		// response
		defer func() {
			w.Header().Set("Content-Type", "application/json")
			result, err := json.Marshal(o);
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintln(w, string(result))
		}()

		switch r.Method {
		case "GET":
			o.Status = 1
			o.Message = "Method is Get"
			o.Body = "Nothing"
			return
		case "POST":
			tmp := ""
			rb := bufio.NewReader(r.Body)
			for {
				s, err := rb.ReadString('\n')
				tmp = tmp + s
				if err == io.EOF {
					break
				}
			}
			o.Body = tmp
		default:
			o.Status = 999
			o.Message = "Not Method"
			return
		}
	}
}