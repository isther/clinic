package main

import (
	"github.com/isther/clinic/conf"
	"github.com/isther/clinic/internal/routers"
	"github.com/sirupsen/logrus"
)

func main() {
	r := routers.Init()
	logrus.Info("Server listen: ", conf.Server.Listen)
	r.Run(conf.Server.Listen)
}
