package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/config"

	"github.com/bitwormhole/go-wormhole-web-starter/demo"
	"github.com/bitwormhole/go-wormhole-web-starter/web/gin_starter"
)

//go:embed src/main/resources
var resources embed.FS

func try_main() error {

	config := &config.AppConfig{}
	args := os.Args
	config.SetResources(&resources, "src/main/resources")
	demo.Config(config)

	context, err := application.Run(config, args)
	if err != nil {
		return err
	}

	web, err := context.GetComponents().GetComponent("gin-web-server")
	if err == nil {
		err = web.(*gin_starter.GinServerContainer).Run()
	}

	code := application.Exit(context)
	fmt.Println("hello, web-starter. exit with code:", code)
	return nil
}

func main() {
	err := try_main()
	if err != nil {
		panic(err)
	}
}
