package main

import (
	Init "blog-go/Config"
	"blog-go/routers"
)

func main() {
	Init.DBInit()
	routers.RouterInit()
}
