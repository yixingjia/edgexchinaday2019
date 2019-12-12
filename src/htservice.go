package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var data = make(map[string]string)

type (
	reading struct {
		Name string `json:"name"`
		Value string `json:"value"`
		Origin  int64    `json:"origin"`
	}

	SensorData struct {
		Device string `json:"device"`
		Origin  int64    `json:"origin"`
		Readings []reading `json:"readings"`
	}
)

func main(){
	port := ":9000"
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("./src/web/static/js/"))))
	http.HandleFunc("/",handler)
	http.HandleFunc("/receiveDataPost", handlePostJson)
	http.HandleFunc("/ajax", OnAjax)
	log.Println("Http server listened at port 9000")
	err := http.ListenAndServe(port,nil)
	if err != nil{
		log.Println(err)	
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	t ,err := template.ParseFiles("src/web/static/index.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w,nil)
}

func handlePostJson(w http.ResponseWriter, r *http.Request) {
	result := SensorData{}
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(con, &result)
	formatTimeStr:=time.Unix(time.Now().Unix(),0).Format("03:04:05")
	data["time"] = formatTimeStr
	for i := range result.Readings {
		if result.Readings[i].Name == "temperature" {
			data["dataT"] = result.Readings[i].Value
		}
		if result.Readings[i].Name == "humidity" {
			data["dataH"] = result.Readings[i].Value
		}
	}
}

func OnAjax(res http.ResponseWriter, req *http.Request) {
	macon,_ :=json.Marshal(data)
	mString :=string(macon)
	io.WriteString(res, mString)
}
