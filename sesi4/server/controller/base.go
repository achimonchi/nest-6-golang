package controller

import (
	"encoding/json"
	"net/http"
	"sesi4/server/view"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(w http.ResponseWriter, payload *view.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(payload.Status)
	json.NewEncoder(w).Encode(payload)
}
func WriteJsonResponseGin(c *gin.Context, payload *view.Response) {
	c.JSON(payload.Status, payload)
}

func WriteErrorJsonResponseGin(c *gin.Context, payload *view.Response) {
	c.AbortWithStatusJSON(payload.Status, payload)
}
