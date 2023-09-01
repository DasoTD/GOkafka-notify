package main

import (
  f "github.com/ambelovsky/gosf"
  "log"
)

func init() {
  // Listen on an endpoint
  f.Listen("echo", func(client *f.Client, request *f.Request) *f.Message {
    return f.NewSuccessMessage(request.Message.Text)
  })
}

func main() {
  // Start the server using a basic configuration
  f.Startup(map[string]interface{}{
    "port": 9999})

	log.Println("connected")
}