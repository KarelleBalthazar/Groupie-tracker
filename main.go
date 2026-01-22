package main

import "groupie_tracker/internal/app"

func main() {
	if err := app.Init(); err != nil {
		panic(err)
	}
}
