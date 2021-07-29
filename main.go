package main

import (
	"github.com/raff/godet"
)

func main() {
	remote, _ := godet.Connect("localhost:9222", false)
	remote.SetCookie(godet.Cookie{Name: "Foo", Value: "Bar"})
}
