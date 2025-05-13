package main

import "github.com/achmad-dev/simple-ewallet/api/v1/routes"

/*
--- MIT License (c) 2022 achmad
--- See LICENSE for more details
*/

func main() {

	// env path
	envpath := "../.env"

	routes.ServerRouteV1(envpath)
}
