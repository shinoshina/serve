package test

import "github.com/shinoshina/snet/internal/netbase"


func main(){

	e := netbase.NewEngine(3)

	e.Launch()
	e.Start()



}