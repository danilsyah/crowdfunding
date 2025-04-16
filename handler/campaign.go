package handler

import (
	"crowdfunding/campaign"
	"crowdfunding/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// analisa flow
// tangkap parameter di handler
// handler panggil service
// service yang mementukan repository mana yang di panggil
// repository yang di panggil sesuai dengan parameter yang di terima : FindAll, FindByUserID
// di repository panggil database

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// GetCampaigns godoc
// @Summary Get Campaigns
// @Description Get all campaigns or campaigns by user ID
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
