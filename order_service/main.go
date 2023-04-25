package main

import "github.com/A-Siam/super_delivery/order_service/src/controller"

func main() {
    server := controller.CreateServer("", ":3000", 5)
    controller.Start(server)
    defer server.Close()
}
