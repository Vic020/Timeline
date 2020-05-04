package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New Request", r.URL.Path)

	sugar := false

	p := r.URL.Query().Get("p")
	if p != "" && SugarCounter < 3 {
		if p == NewPostSugar {
			sugar = true
		} else {
			log.Println("wrong sugar", p, "counter is ", SugarCounter)
			atomic.AddInt32(&SugarCounter, 1)
		}
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 100
	}

	itemList := FileDaoHandle.GetItemList(page, limit)

	res := []Item{}

	for i := len(itemList) - 1; i >= 0; i-- {
		v := itemList[i]
		in, _ := strconv.ParseInt(v.Timestamp, 10, 64)
		v.Timestamp = time.Unix(in, 0).UTC().String()

		res = append(res, v)
	}

	GetInstance().Render(w, "index.html", map[string]interface{}{
		"items":  res,
		"slogan": Slogan,
		"sugar":  sugar,
		"p":      p,
	})

	//res, err := json.Marshal(itemList)
	//if err != nil {
	//	w.WriteHeader(500)
	//	return
	//}
	//
	//w.Write(res)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New Request", r.URL.Path)
	if r.Method == "POST" {

		// sugar check
		p := r.URL.Query().Get("p")
		if p != NewPostSugar && SugarCounter < 3 {
			log.Println("wrong sugar", p, "counter is ", SugarCounter)
			atomic.AddInt32(&SugarCounter, 1)
			w.WriteHeader(404)
			return
		}

		// normal process
		r.ParseForm()

		//name := r.FormValue("name")
		//avatar := r.FormValue("avatar")
		//timestamp := r.FormValue("timestamp")

		name := "Vic Yu"
		avatar := "https://vicyu.com/img/avatar.jpeg"
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

		content := r.FormValue("content")
		if content == "" {
			w.WriteHeader(200)
			w.Write([]byte("content is nil"))
		}
		item := FileDaoHandle.AddItem(name, avatar, timestamp, content)

		res, err := json.Marshal(item)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(res)

	} else {
		w.WriteHeader(404)
		return
	}

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	return
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	return
}
