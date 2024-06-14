package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"regexp"

	"log"
)

type JSONReq struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Version float32     `json:"version"`
}

type JSONServer struct {
	handler interface{}
}

var methodRe = regexp.MustCompile(`\w+`)

func (l *JSONServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != "POST" {
		log.Printf("Unsupported JSON RPC %s %s", r.Method, r.URL.Path)
		return
	}
	jreq := JSONReq{}
	if json.NewDecoder(r.Body).Decode(&jreq) != nil {
		log.Printf("Bad JSON body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error:\": \"Bad JSON body\"}"))
		return
	}
	if !methodRe.MatchString(jreq.Method) {
		log.Printf("Method not valid")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error:\": \"Method not specified!\"}"))
		return
	}
	methodName := jreq.Method
	if method := reflect.ValueOf(l.handler).MethodByName(methodName); method.IsValid() {
		vals := method.Call([]reflect.Value{reflect.ValueOf(any(jreq.Params)), reflect.ValueOf(r)})
		if err := json.NewEncoder(w).Encode(vals[0].Interface()); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"error:\": \"Runtime error!\"}"))
		}
		return
	}
	log.Printf("%s Method not found", methodName)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"error:\": \"Method not found!\"}"))
}
