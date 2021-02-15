package main

import (
	"github.com/kenshin579/analyzing-import-cycle-example-go/p1"
)

func main() {
	pp1 := p1.PP1{}
	pp1.HelloFromP2Side()
}
