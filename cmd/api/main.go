package main

import (
	"workout_tracker/config"
)

func main() {
	config.SetupDB()
	config.Seed(config.DB)
}
