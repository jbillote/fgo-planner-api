package main

import (
    "github.com/jbillote/fgo-planner-api/server"
    "os"
)

func main() {
    s := server.NewServer()
    s.Start(os.Getenv("PORT"))
}
