package main

import (
	"github.com/labstack/echo/v4"
	"shared-dependency/common/config"
)

func main() {
	config.MyConf()
	echo.New()
}
