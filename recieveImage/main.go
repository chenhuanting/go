package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"encoding/base64"
)

func recieveImage(w http.ResponseWriter , r *http.Request){
	r.ParseForm()
	fmt.Println("path",r.URL.Path)
	imagebuff := r.Form["image"]
	formatTimeStr:=time.Now().Format("2006-01-02-15-04-05")
	fmt.Println(formatTimeStr)   //打印结果：2017-04-11 13:30:39	
	dataTimeString := "static/"+string(formatTimeStr)+".png"
	fmt.Println(dataTimeString)
	imageString := string(imagebuff[0])
	//fmt.Println(imageString)
	base2Img(imageString,dataTimeString)
	fmt.Fprintf(w,"submit sucess!")
}
func base2Img(imageString string,filename string){
	ddd, _ := base64.StdEncoding.DecodeString(imageString) 
	err :=ioutil.WriteFile(filename, ddd, 0666) 
	fmt.Println(err)
}
func main(){
	fmt.Println(" \t welcome to use recivev image http server \t")
	fmt.Println("\t\t\t\t\t深圳市德沃仪器\t")
	http.HandleFunc("/image/upload",recieveImage)
	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}