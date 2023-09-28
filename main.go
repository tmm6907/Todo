package main

import (
	"github.com/tmm6907/Todo/constants"
	"github.com/tmm6907/Todo/server"
)

func main() {
	router := server.New(constants.INTERNAL_PORT)
	router.Run()
}
