package main

import (
	"fmt"
	"log"

	"api-starter/database"
	_"api-starter/docs"
	i "api-starter/internal"
	"api-starter/service"
)

//	@title			projectname API
//	@version		0.1
//	@description	description
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	TODO
//	@contact.email	TODO

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@servers	/projectname/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	dbConfig, err := i.NewDatabaseConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app, err := i.NewApp()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.NewDatabasePool(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Close()
	router := service.InitRouter(app, db)

	err = router.Run(app.Port())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is running on port: ", app.Port())
}
