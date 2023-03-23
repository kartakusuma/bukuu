package main

import "bukuu/routers"

func main() {
	PORT := ":8080"

	routers.StartServer().Run(PORT)
}
