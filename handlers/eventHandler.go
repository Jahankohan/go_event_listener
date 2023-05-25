package handlers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/jahankohan/go_event_listener/model"
	"github.com/jahankohan/go_event_listener/modules"

	"github.com/jahankohan/go_event_listener/utils"
)

type EventHandler struct {}


func (eh EventHandler) GetAllEvents(c *gin.Context) {
	variables := utils.LoadConfig()
	client, err := ethclient.Dial(variables.AVATestnet.Network)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : "Problem with connecting to Geth",
			"description" : err.Error(),
		})
	}
	handler := modules.ClientHandler{Client: client, DeployedAddress: variables.AVATestnet.DeployedAddress}
	allEvents := handler.PullEvents()
	c.JSON(http.StatusOK, allEvents)
}

