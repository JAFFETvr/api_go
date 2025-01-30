package infraestructure

import (
    "github.com/gin-gonic/gin"
)

func RegisterClientRoutes(router *gin.Engine, createController *CreateClientController, getController *GetClientsController) {
    clients := router.Group("/clients")
    {
        clients.POST("/", createController.Handle)
        clients.GET("/", getController.Handle)
    }
}
