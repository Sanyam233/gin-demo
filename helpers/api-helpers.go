package helpers

import "github.com/gin-gonic/gin"

func CreateResponse(results interface{}, errors interface{}) gin.H {
	if results == nil {
		results = []gin.H{}
	} else if !IsSliceOrArray(results) {
		results = []interface{}{results}
	}

	if errors == nil {
		errors = []gin.H{}
	} else if !IsSliceOrArray(errors) {
		errors = []interface{}{errors}
	}

	return gin.H{
		"results": results,
		"errors":  errors,
	}
}
