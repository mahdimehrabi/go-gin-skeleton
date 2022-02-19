package controllers

import (
	"boilerplate-api/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UtilityController struct {
	logger infrastructure.Logger
	env    infrastructure.Env
}

func NewUtilityController(logger infrastructure.Logger,
	env infrastructure.Env,
) UtilityController {
	return UtilityController{
		logger: logger,
		env:    env,
	}
}

// Response -> response for the util scope
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    string      `json:"data"`
	Path    string      `json:"path"`
	Value   interface{} `json:"attributes"`
}

const storageURL string = "https://storage.googleapis.com/"

// FileUploadHandler -> handles file upload
func (uc UtilityController) FileUploadHandler(ctx *gin.Context) {
	response := &Response{
		Success: true,
		Message: "Uploaded Successfully",
		Data:    "file",
		Path:    "/file",
	}
	ctx.JSON(http.StatusOK, response)
}
