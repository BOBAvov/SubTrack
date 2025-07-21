package handler

import (
	"fmt"
	"github.com/BOBAvov/sub_track"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createSubscription(c *gin.Context) {
	var input sub_track.Subscription

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Subscription.Create(input)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}

type AllSubscriptionsRequest struct {
	Data []sub_track.Subscription `json:"data"`
}

func (h *Handler) getAllSubscriptions(c *gin.Context) {
	subs, err := h.services.Subscription.GetAll()
	fmt.Printf("%v", subs)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, AllSubscriptionsRequest{Data: subs})
}

func (h *Handler) getByIdSubscription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, "Invalid Subscription ID")
		return
	}
	sub, err := h.services.Subscription.GetById(id)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sub)

}

func (h *Handler) updateSubscription(c *gin.Context) {
	var input sub_track.SubscriptionUpdate
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, "Invalid Subscription ID")
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Subscription.Update(id, input)

	if err != nil {
		newErrorResponse(c, h.log, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}

func (h *Handler) deleteSubscription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, "Invalid Subscription ID")
		return
	}

	err = h.services.Subscription.Delete(id)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}
