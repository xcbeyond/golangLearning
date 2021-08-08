package main

import (
	"fmt"
	"golangLearning/framework/gin-demo/route/route-split/split-v1/routes"
)

func main() {
	route := routes.SetupRouter()
	if err := route.Run(); err != nil {
		fmt.Printf("startup server failed,err: %v", err)
	}
}
