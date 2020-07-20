package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	tmp, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}
//err = tmp.Execute(os.Stdout, nil)
//if err != nil {
//	log.Fatalln(err)


err = tmp.ExecuteTemplate(os.Stdout, "test.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tmp.ExecuteTemplate(os.Stdout,"test2.gohtml", nil)
	if err != nil{
		log.Fatalln(err)
	}
}