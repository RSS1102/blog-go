package main

import (
	Init "blog-go/Config"
	"blog-go/Routers"
)

func main() {
	Init.DBInit()
	Routers.RouterInit()
}
