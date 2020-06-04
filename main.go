package main

import (
	"log"

	"github.com/astaxie/beego"

	_ "github.com/YouDad/localhost/modules"
	_ "github.com/YouDad/localhost/routers"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ltime)
	log.SetPrefix("[INFO]: ")

	beego.Run()
}
