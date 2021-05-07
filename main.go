package main

import (
	"aubase/router"
	"flag"
)

func main () {
	port := flag.String("port", "8080", "Listen port.")
	flag.Parse()
	r := router.InitRouter()
	r.Run(":" + *port)
}