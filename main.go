package main
import (
	"fmt"
	"rbac/routers"
)

func main(){
	fmt.Println("hello world")
	r := routers.NewRouters()
	engine := r.InitGin()
	engine.Run(":8080")
}
