package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zignd/turbo-octo-robot/model"
)

var accounts map[string]*model.Account

func init() {
	accounts = make(map[string]*model.Account)
}

// PostAccountHandler is the handler function for the POST /account route
func PostAccountHandler(w http.ResponseWriter, req *http.Request) {
	defer (func() {
		if req.Body != nil {
			req.Body.Close()
		}
	})()

	fmt.Println("[POST /account]")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, fmt.Errorf("failed to read the request body: %w", err).Error())
		return
	}

	account := &model.Account{}
	if err := json.Unmarshal(body, account); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, fmt.Errorf("failed to parse the request body: %w", err).Error())
		return
	}

	if accounts[account.ID] != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "There's already an account with provided ID, %s", account.ID)
		return
	}

	fmt.Printf("[POST /account] Created an account %+v\n", account)
	accounts[account.ID] = account
}
