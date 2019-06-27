package main

import (
	"github.com/astaxie/beego"
	"os"
	"log"
)

func main() {
	log.Fatal(os.Getwd())
	beego.Run()
}