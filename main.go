package main

import (
	"log"
	"os"
	_ "rss/mathers"
	"rss/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Emergency")
}
