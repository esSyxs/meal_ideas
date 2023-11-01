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
	protected.PUT("/profile", updateUser)
	protected.POST("/favourite", addFavourite)
	protected.PUT("/favourite", removeFavourite)

	return r
}

func updateUser(c *gin.Context) {
	u := struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{}

	err := c.ShouldBindJSON(&u)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": fmt.Sprintf("Invalid Inputs %s", err.Error()),
		})
		c.Abort()
		return
	}

	email, found := c.Get("email")
	if !found {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}

	user, err := food.GetUser(email.(string))
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}

	if u.Password != "" {
		err = user.HashPassword(u.Password)
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"Error": "Error Hashing Password",
			})
			c.Abort()
			return
		}
	}

	oldEmail := user.Email

	if u.Email != "" {
		user.Email = u.Email
	}

	if u.Username != "" {
		user.Username = u.Username
	}

	err = food.UpdateUser(user, oldEmail)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Updating User",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"Message": "Sucessfully Updated",
	})
}

func addFavourite(c *gin.Context) {
	rec_id := struct {
		ID uint `json:"id"`
	}{}

	err := c.ShouldBindJSON(&rec_id)
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "Invalid Recipe ID",
		})
		c.Abort()
		return
	}

	rec, err := food.GetRecepie(rec_id.ID)
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "Recipe Not Found",
		})
		c.Abort()
		return
	}

	email, found := c.Get("email")
	if !found {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}

	food.AddUserRecipe(email.(string), *rec)

	c.JSON(200, gin.H{
		"Message": "Sucessfully added favourite recipe",
	})
}

func removeFavourite(c *gin.Context) {
	rec_id := struct {
		ID uint `json:"id"`
	}{}

	err := c.ShouldBindJSON(&rec_id)
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "Invalid Recipe ID",
		})
		c.Abort()
		return
	}

	rec, err := food.GetRecepie(rec_id.ID)
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "Recipe Not Found",
		})
		c.Abort()
		return
	}

	email, found := c.Get("email")
	if !found {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}

	food.RemoveUserRecipe(email.(string), *rec)

	c.JSON(200, gin.H{
		"Message": "Sucessfully removed favourite recipe",
	})
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

	var pIDs, aIDs []uint

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

		pIDs = append(pIDs, uint(i))
	}

	fmt.Println("pIds", pIDs)
	fmt.Println("aIDs", aIDs)

	c.JSON(200, filterRecipes(food.GetRecepies(), pIDs, aIDs, pMatch, aMatch))
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
		produceIDs := map[uint]bool{}
		appliancesIDs := map[uint]bool{}
		for _, p := range r.Produces {
			if p == nil {
				continue
			}

			produceIDs[p.ID] = false
			if inSlice(pIDs, p.ID) {
				tmpP = append(tmpP, p.ID)
			}
		}

		if pMatch && len(r.Produces) != len(tmpP) {
			continue
		}

		var tmp []uint

		for _, a := range r.Appliances {
			if a == nil {
				continue
			}

			appliancesIDs[a.ID] = false
			if inSlice(aIDs, a.ID) {
				tmp = append(tmp, a.ID)
			}
		}

		if aMatch && len(r.Appliances) != len(tmp) {
			continue
		}

		switch true {
		case !aMatch && !pMatch && (len(tmpP) > 0 || len(tmp) > 0):
			var tmp int
			for _, id := range pIDs {
				_, ok := produceIDs[id]
				if ok {
					tmp++
				}

			}

			if len(pIDs) != tmp && len(pIDs) > 0 {
				continue
			}

			tmp = 0

			for _, id := range aIDs {
				_, ok := appliancesIDs[id]
				if ok {
					tmp++
				}
			}

			if len(aIDs) != tmp && len(aIDs) > 0 {
				continue
			}

			out[r.ID] = r
		case aMatch && pMatch && len(tmpP) == len(r.Produces) && len(tmp) == len(r.Appliances):
			out[r.ID] = r
		case aMatch && len(tmp) == len(r.Appliances) && !pMatch:
			var tmp int
			for _, id := range pIDs {
				_, ok := produceIDs[id]
				if ok {
					tmp++
				}

			}

			if len(pIDs) != tmp && len(pIDs) > 0 {
				continue
			}

			out[r.ID] = r
		case pMatch && len(tmpP) == len(r.Produces) && !aMatch:
			var tmp int
			for _, id := range aIDs {
				_, ok := appliancesIDs[id]
				if ok {
					tmp++
				}
			}

			if len(aIDs) != tmp && len(aIDs) > 0 {
				continue
			}

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
