package main

import (
	"github.com/tmm6907/Todo/constants"
	"github.com/tmm6907/Todo/server"
)

func main() {
	s := server.New(constants.INTERNAL_PORT, server.Config{ReleaseMode: true})
	s.Run()
}
