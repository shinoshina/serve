package main

import "goserve/internal/netbase"


func main(){

	e := netbase.NewEngine(3)

	e.Launch()
	e.Start()



}