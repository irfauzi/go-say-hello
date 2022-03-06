package main

import (
	"fmt"

	"github.com/irfauzi/go-say-hello/say"
	_ "github.com/irfauzi/go-say-hello/say"
	"os"
)

func main()  {


	say.SayHi()
	fmt.Println(os.Args)
	fmt.Println(os.Hostname())
	fmt.Println(os.Getenv("APP_USERNAME"))


}