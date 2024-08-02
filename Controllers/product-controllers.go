package controllers

import (
	"fmt"
	"gin-demo/helpers"
	"gin-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductParams struct {
	Product models.Product `json:"product" binding:"required"`
}

type CreateProductReq struct {
	Params CreateProductParams `json:"params"`
}

func GetAllProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.CreateResponse(models.Products, nil))
}

func GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, product := range models.Products {
		if product.ID == id {
			ctx.JSON(http.StatusOK, helpers.CreateResponse(product, nil))
			return
		}
	}

	ctx.JSON(http.StatusNotFound, helpers.CreateResponse(nil, gin.H{
		"status":  http.StatusNotFound,
		"message": fmt.Sprintf("product with id %s doesnot exist", id),
	}))
}

func CreateProduct(ctx *gin.Context) {
	var body CreateProductReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(404, helpers.CreateResponse(nil, models.CustomError(err.Error(), 404)))
		return
	}

	fmt.Printf("%v", body.Params.Product.Seller.FirstName)
	ctx.JSON(201, helpers.CreateResponse(nil, nil))

}
