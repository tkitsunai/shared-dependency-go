package main

import (
	"github.com/labstack/echo"
	"shared-dependency/common/config"
)

func main() {
	config.MyConf()
	echo.New()
}
