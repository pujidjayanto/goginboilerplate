package main

import "github.com/pujidjayanto/goginboilerplate/pkg/log"

func main() {
	log.Init()
	defer log.SyncLogger()

	env, err := loadEnvironment()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.ConfigureLogger(env.Server.Env)
}
