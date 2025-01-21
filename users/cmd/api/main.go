package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nimbo1999/lopes_drive/users/internal/app"
	"github.com/nimbo1999/lopes_drive/users/internal/config"
)

func main() {
	config.ReadEnvs()
	env := config.CreateEnvironment()
	usersApp := app.NewApp(env)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", env.Port), usersApp.Handler); err != nil {
		log.Fatalln(err)
	}
}
