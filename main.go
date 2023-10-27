package main

import "github.com/jbillote/fgo-planner-api/server"

func main() {
    s := server.NewServer()
    s.Start(":8080")
}
