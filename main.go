package main

import (
	"AUBase/router"
	"AUBase/util"
	"flag"
	"fmt"
)

func main () {
	port := flag.String("port", "8080", "Listen port.")
	flag.Parse()
	r := router.InitRouter()
	fmt.Println(util.PasswordEncrypt("123456"))
	r.Run(":" + *port)
}