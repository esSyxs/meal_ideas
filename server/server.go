package server

import (
	logger "System/Log"
	"System/food"
	"System/jwt_token_authorization/controllers"
	"System/jwt_token_authorization/middlewares"
	"fmt"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const idName = "id"

func Start(port string) {
	r := setupRouter()

	r.Run(fmt.Sprintf(":%s", port))
}

// setupRouter sets up the router and adds the routes.
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})
	// Create a new group for the API
	api := r.Group("/api")

	public := api.Group("/public")
	public.POST("/login", controllers.Login)
	public.POST("/signup", controllers.Signup)

	public.GET(fmt.Sprintf("/recipes/:%s", idName), recpieGet)
	public.GET("/recipes", recpiesGet)

	protected := api.Group("/protected").Use(middlewares.Authz())
	protected.GET("/profile", controllers.Profile)

	// Return the router
	return r
}

func recpiesGet(c *gin.Context) {
	logger.Default.Println("got recepies request")
	c.JSON(200, food.GetRecepies())
}

func recpieGet(c *gin.Context) {
	id := c.Param(idName)
	logger.Default.Printf("got recepie request with id: %s\n", id)

	i, err := strconv.Atoi(id)
	if err != nil {
		logger.Default.Printf("failed to parse recepie id: %s, err: %s\n", id, err.Error())
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		c.Abort()
		return
	}

	rec, err := food.GetRecepie(uint(i))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(200, rec)
}
