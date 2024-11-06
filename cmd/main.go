package main

import (
	"fmt"
	"os"
	"ytlog/internal/application"
	"ytlog/internal/config"
)

func main() {
	cfg := config.Config{}
	err := config.LoadFromJsonFile("./config.json", &cfg)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
	app := application.NewApplication(&cfg)
	app.Run()
}
