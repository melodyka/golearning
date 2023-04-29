package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init在main之前调用
func init() {
	log.SetOutput(os.Stdout)
}

// main入口
func main() {
	//使用特定的项搜索
	search.Run("president")
}
