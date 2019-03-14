package main

import (
	"encoding/json"
    "github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type jsonStruct struct {
	Name string `jsonapi:"attr,Name"`
	Text string `jsonapi:"attr,Text"`
}

func main() {
	r := mux.NewRouter()

	reqLimiter := tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Second})

	r.Handle("/api1", tollbooth.LimitFuncHandler(reqLimiter, api1)).Methods("GET")
	r.Handle("/api2", tollbooth.LimitFuncHandler(reqLimiter, api2)).Methods("POST")
	r.Handle("/api3", tollbooth.LimitFuncHandler(reqLimiter, api3)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func api1(w http.ResponseWriter, _ *http.Request) {
	randNum := randomNum(500, 1500)
	randTime := time.Duration(randNum) * time.Millisecond
	time.Sleep(randTime) // emulation of doing something that takes random amount of time (from 500ms to 1.5s)

	w.Write([]byte("done"))
}

func api2(w http.ResponseWriter, r *http.Request) {
	var jsonInterface interface{} // переменная с экземпляром типо interface для преобразования данных из формата JSON

	randNum := randomNum(500, 1500)
	randTime := time.Duration(randNum) * time.Millisecond
	time.Sleep(randTime) // emulation of doing something that takes random amount of time (from 500ms to 1.5s)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpError(w, err)
		return
	}

	err = json.Unmarshal(body, &jsonInterface) // разбор данных в формате JSON и помещение их в переменную с типом interface{}
	if err != nil {
		httpError(w, err)
		return
	}

	jsonData, err := json.MarshalIndent(jsonInterface, " ", "\t")
	if err != nil {
		httpError(w, err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(jsonData)
}

func api3(w http.ResponseWriter, r *http.Request) {
	randNum := randomNum(500, 1500)
	randTime := time.Duration(randNum) * time.Millisecond
	time.Sleep(randTime) // emulation of doing something that takes random amount of time (from 500ms to 1.5s)

	s := new(jsonStruct)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpError(w, err)
		return
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		httpError(w, err)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	err = jsonapi.MarshalPayload(w, s)
	if err != nil {
		httpError(w, err)
		return
	}
}

func httpError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Println(err)
}

func randomNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
