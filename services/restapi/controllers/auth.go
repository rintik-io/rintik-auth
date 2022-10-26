package controllers

import (
	"net/http"

	"github.com/fahmyabdul/golibs"
	"github.com/gin-gonic/gin"
	"github.com/rintik-io/rintik-auth/internal/bispro"
	"github.com/rintik-io/rintik-auth/internal/models"
	"github.com/rintik-io/rintik-auth/internal/models/users"
)

// Auth :
type Auth struct{}

// Register godoc
// @Summary      Register New Account
// @Description  Register New Account With JWT Token
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Param		 email 		formData 	string	true	"E-mail"
// @Param		 username 	formData 	string 	true 	"Username"
// @Param		 phone 		formData 	string 	true 	"Phone"
// @Param		 name 		formData 	string 	true 	"Name"
// @Param		 password 	formData 	string 	true 	"Password"
// @Param		 role 		formData 	string 	true 	"Role" Enums(admin,user)
// @Produce      json
// @Success      200  {object}  models.ResponseRestApi
// @Failure      400  {object}  models.ResponseRestApi
// @Failure      404  {object}  models.ResponseRestApi
// @Failure      500  {object}  models.ResponseRestApi
// @Router       /register [post]
func (p *Auth) Register(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		golibs.Log.Printf("| Controller | Auth | Register | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var requestData users.RequestRegister

	err := c.Bind(&requestData)
	if err != nil {
		golibs.Log.Printf("| Controller | Auth | Register | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	responseData, err := (&bispro.BisproRegister{}).Register(&requestData)
	if err != nil {
		golibs.Log.Printf("| Controller | Auth | Register | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    &responseData,
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseRestApi{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    &responseData,
	})
}

// Claims godoc
// @Summary      Claims JWT
// @Description  Claims JWT Token
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Param		 username_email_phone 	formData 	string 	true 	"Username/Email/Phone"
// @Param		 password 	formData 	string 	true 	"Password"
// @Produce      json
// @Success      200  {object}  models.ResponseRestApi
// @Failure      400  {object}  models.ResponseRestApi
// @Failure      404  {object}  models.ResponseRestApi
// @Failure      500  {object}  models.ResponseRestApi
// @Router       /claims [post]
func (p *Auth) Claims(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		golibs.Log.Printf("| Controller | Auth | Claims | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var requestData users.RequestClaims

	err := c.Bind(&requestData)
	if err != nil {
		golibs.Log.Printf("| Controller | Auth | Claims | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	responseData, err := (&bispro.BisproClaims{}).Claims(&requestData)
	if err != nil {
		golibs.Log.Printf("| Controller | Auth | Claims | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    &responseData,
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseRestApi{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    &responseData,
	})
}

// Validate godoc
// @Summary      Validate JWT
// @Description  Validate JWT Token
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Param		 jwt 	formData 	string 	true 	"JWT Token"
// @Produce      json
// @Success      200  {object}  models.ResponseRestApi
// @Failure      400  {object}  models.ResponseRestApi
// @Failure      404  {object}  models.ResponseRestApi
// @Failure      500  {object}  models.ResponseRestApi
// @Router       /validate [post]
func (p *Auth) Validate(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		golibs.Log.Printf("| Controller | Auth | Validate | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: "Unauthorized",
			Data:    nil,
		})
		return
	}

	var requestData users.RequestValidate

	err := c.Bind(&requestData)
	if err != nil {
		golibs.Log.Printf("| Controller | Auth | Validate | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: "Unauthorized",
			Data:    nil,
		})
		return
	}

	responseData, err := (&bispro.BisproValidate{}).Validate(requestData.Jwt)
	if err != nil {
		golibs.Log.Printf("| Controller | Auth | Validate | Failed, error: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ResponseRestApi{
			Code:    http.StatusBadRequest,
			Message: "Unauthorized",
			Data:    &responseData,
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseRestApi{
		Code:    http.StatusOK,
		Message: "Authorized",
		Data:    &responseData,
	})
}
