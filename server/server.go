package server

import (
	logger "System/Log"
	"System/food"
	"System/jwt_token_authorization/controllers"
	"System/jwt_token_authorization/middlewares"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	idName = "id"

	//query_string_keys
	produceID      = "produce_id"
	produceMatch   = "produce_match_strict"
	applianceID    = "appliance_id"
	applianceMatch = "appliance_match_strict"
)

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

	var pMatch, aMatch bool
	var err error

	query := c.Request.URL.Query()
	produceIDs := query[produceID]
	applianceIDs := query[applianceID]

	if len(produceIDs) == 0 && len(applianceIDs) == 0 {
		c.JSON(200, food.GetRecepies())
		c.Abort()
		return
	}

	produceMatchString := query.Get(produceMatch)
	applianceMatchString := query.Get(applianceMatch)

	if produceMatchString != "" {
		pMatch, err = strconv.ParseBool(produceMatchString)
		if err != nil {
			c.JSON(400, gin.H{
				"Error": fmt.Errorf("failed to parse %s, %s", produceMatchString, err.Error()),
			})
			c.Abort()
			return
		}
	}

	if applianceMatchString != "" {
		aMatch, err = strconv.ParseBool(applianceMatchString)
		if err != nil {
			c.JSON(400, gin.H{
				"Error": fmt.Errorf("failed to parse %s, %s", applianceMatchString, err.Error()),
			})
			c.Abort()
			return
		}
	}

	log.Printf("got values: produce ids: %+v, applance ids: %+v, strict produce: %v, strict applance: %v\n", produceIDs, applianceIDs, produceMatchString, applianceMatchString)

	var rIDs, aIDs []uint

	for _, id := range applianceIDs {
		i, err := strconv.Atoi(id)
		if err != nil {
			logger.Default.Printf("failed to parse appliance id: %s, err: %s\n", id, err.Error())
			c.JSON(400, gin.H{
				"Error": err.Error(),
			})
			c.Abort()
			return
		}

		aIDs = append(aIDs, uint(i))
	}

	for _, id := range produceIDs {
		i, err := strconv.Atoi(id)
		if err != nil {
			logger.Default.Printf("failed to parse produce id: %s, err: %s\n", id, err.Error())
			c.JSON(400, gin.H{
				"Error": err.Error(),
			})
			c.Abort()
			return
		}

		rIDs = append(rIDs, uint(i))
	}

	c.JSON(200, filterRecipes(food.GetRecepies(), rIDs, aIDs, pMatch, aMatch))
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

// database should handle this
func filterRecipes(all map[uint]*food.Recepie, pIDs, aIDs []uint, pMatch, aMatch bool) map[uint]*food.Recepie {

	out := map[uint]*food.Recepie{}

	for _, r := range all {
		var tmpP []uint
		for _, p := range r.Produces {
			if inSlice(pIDs, p.ID) {
				tmpP = append(tmpP, p.ID)
			}
		}

		if pMatch && len(r.Produces) != len(tmpP) {
			continue
		}

		var tmp []uint

		for _, a := range r.Appliances {
			if inSlice(aIDs, a.ID) {
				tmp = append(tmp, a.ID)
			}
		}

		if aMatch && len(r.Appliances) != len(tmp) {
			continue
		}

		switch true {
		case !aMatch && !pMatch && (len(tmpP) > 0 || len(tmp) > 0):

			out[r.ID] = r
		case aMatch && pMatch && len(tmpP) == len(r.Produces) && len(tmp) == len(r.Appliances):
			out[r.ID] = r
		case aMatch && len(tmp) == len(r.Appliances) && !pMatch:
			out[r.ID] = r
		case pMatch && len(tmpP) == len(r.Produces) && !aMatch:
			out[r.ID] = r
		}
	}

	return out
}

func inSlice(in []uint, match uint) bool {
	for _, v := range in {
		if v == match {
			return true
		}
	}

	return false
}
