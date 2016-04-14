package main

import (
	"github.com/Reyurnible/Sudoku-server/route"
)
// 初期化
func init() {
}

func main() {
	router := route.Init()
	router.Run(":8080")
}
