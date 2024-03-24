package main

import "OLO-backend/api_gateway/internal/app"

func main() {
	a, err := app.New()
	if err != nil {
		panic(err)
		return
	}
	a.Start()
}
