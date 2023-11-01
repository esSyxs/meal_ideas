package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	CreateProduceQuery      = `INSERT INTO produce (Name) VALUES ('%s');`
	GetProduceByNameQuery   = `select * from produce WHERE name = '%s';`
	GetProduceByIDQuery     = `select * from produce WHERE ID = '%d';`
	CreateApplienceQuery    = `INSERT INTO appliance (Name) VALUES ('%s');`
	GetApplienceByNameQuery = `select * from appliance WHERE name = '%s';`
	GetApplienceByIDQuery   = `select * from appliance WHERE ID = '%d';`
	GetRecipesByNameQuery   = `select ID from recipes WHERE Name = '%s';`
	GetRecipes              = `SELECT * FROM recipes;`
	GetRecipesProd          = `SELECT * FROM recipe_produce;`
	GetRecipesApl           = `SELECT * FROM recipe_appliance;`

	InsertRecepieQuery = `
	INSERT INTO recipes (Name, Description, Users)
	VALUES ('%s', '%s', '%s');
	`
	InsertRecepieQuery2 = `
 	INSERT INTO recipe_produce (recipe_id, produce_id)
	VALUES %s;
	`

	InsertRecepieQuery3 = `
	INSERT INTO recipe_appliance (recipe_id, appliance_id)
	VALUES %s;
	`
)

func Connect(name, password, url, port, db string) (*sql.DB, error) {
	conn, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", name, password, url, port, db))

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL Database")

	return conn, nil
}
