package main

import (
	"encoding/json"
	"net/http"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	itemList := FileDaoHandle.GetItemList(1, 10)
	res, err := json.Marshal(itemList)
	if err != nil {
		w.Write([]byte("Internal Error"))
	}

	w.Write(res)
}

func newHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

}

func updateHandler(w http.ResponseWriter, r *http.Request) {

}
