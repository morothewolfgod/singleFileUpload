package main

import(
	"./app"
	"os"
	"github.com/joho/godotenv"
)

var port string

func init(){
	err := godotenv.Load("config.ini")
	if err != nil {
	  panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {

	s := app.NewServer()
	s.Init(port)
	s.Start()
	
}