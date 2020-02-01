package Contextualizer

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/rrkrish561/relief-call-logger/Caller"
	"github.com/rrkrish561/relief-call-logger/Message"
)

func Contextualize(w http.ResponseWriter, r *http.Request) {
	response := Message.Message{}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		fmt.Fprint(w, "An error has occured")
		return
	}

}
