package main

import (
	"log"

	"githiub.com/qodirtok/go-rest-login/cmd"
	"github.com/joho/godotenv"
)

func goDotEnvVariable() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	// return os.Getenv()
}


func main()  {
	goDotEnvVariable()

	// fmt.Printf("------ %s ------", dotenv)

	cmd.Cmd()
}


