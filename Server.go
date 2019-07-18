//go:generate goversioninfo -icon=mkcl.ico
package main

import (
	"DemoProject/server/api"
	"DemoProject/server/api/model"

	"fmt"
	"log"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"github.com/labstack/echo"
)

var APP_ENV = "dev"
var APP_VERSION string
var APP_ISACTIVATION string

// var apicdnServerPort = ":4887"
func main() {
	confighelper.InitViper()
	e := echo.New()
	e.HideBanner = true
	//should be called before initializing api
	fmt.Println("initializeVariable")
	initializeVariable()
	//Bind API
	fmt.Println("before init")

	api.Init(e)
	fmt.Println("after init")

	serverPort := confighelper.GetConfig("serverport")
	logginghelper.LogDebug(serverPort)
	err := e.Start(serverPort)
	if err != nil {
		log.Fatal(err)
	}
}

func initializeVariable() {
	model.DBNAME = confighelper.GetConfig("DBNAME")
}
