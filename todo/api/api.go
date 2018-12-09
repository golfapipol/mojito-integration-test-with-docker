package api

import (
	"net/http"
	"mojito/todo/repository"

	"github.com/gin-gonic/gin"
)

type API struct {
	Repository repository.Repository
}

func (api API) ViewCountry(context *gin.Context) {
	var countryName string
	countryName, _ = context.GetQuery("name")
	country,_ := api.Repository.FindCountryByName(countryName)
	context.JSON(http.StatusOK, country)
}
