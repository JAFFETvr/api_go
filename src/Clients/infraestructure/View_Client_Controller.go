package infraestructure

import (
    "github.com/gin-gonic/gin"
    "demo/src/Clients/applications"
    "net/http"
)

type GetClientsController struct {
    UseCase *applications.GetClientsUseCase
}

func NewGetClientsController(useCase *applications.GetClientsUseCase) *GetClientsController {
    return &GetClientsController{UseCase: useCase}
}

func (c *GetClientsController) Handle(ctx *gin.Context) {
    clients, err := c.UseCase.Execute()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, clients)
}