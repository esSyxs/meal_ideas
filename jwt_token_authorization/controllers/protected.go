// sign-up login system example taken from https://github.com/VinikaAnthwal/go-jwt and median post
// https://medium.com/@22vinikaanthwal/register-login-api-with-jwt-authentication-in-golang-gin-740633e5707b

package controllers

import (
	"System/food"

	"github.com/gin-gonic/gin"
)

// Profile is a controller function that retrieves the user profile from the database
// based on the email provided in the authorization middleware.
// It returns a 404 status code if the user is not found,
// and a 500 status code if an error occurs while retrieving the user profile.

func Profile(c *gin.Context) {
	// Get the email from the authorization middleware
	email, _ := c.Get("email")
	// Query the database for the user

	user, err := food.GetUser(email.(string))
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}

	// Set the user's password to an empty string
	user.Password = ""
	// Return the user profile with a 200 status code
	c.JSON(200, user)
}
