package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ekharisma/sendgrid-web-service/internals/controller"
	"github.com/ekharisma/sendgrid-web-service/internals/service"
	"github.com/ekharisma/sendgrid-web-service/internals/static"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "specifiy path of config file")
	flag.Parse()
	if path == "" {
		log.Panicln("Cant get path from args")
		return
	}
	config := static.NewConfig(path)
	emailService := service.NewEmailClient(config)
	sendGridController := controller.NewSendGridController(emailService)
	http.HandleFunc("/email", sendGridController.SendEmail)
	port := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(port, nil))
}
