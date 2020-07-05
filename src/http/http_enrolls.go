package http

import (
	"net/http"

	atDomain "github.com/dzikrisyafi/kursusvirtual_gateway-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/services/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

type EnrollsHandler interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type enrollsHandler struct {
	service enrolls.Service
}

func NewEnrollsHandler(service enrolls.Service) EnrollsHandler {
	return &enrollsHandler{
		service: service,
	}
}

func (h *enrollsHandler) Create(c *gin.Context) {
	var request atDomain.Enroll
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	enroll, err := h.service.Create(request, c.Query("access_token"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusCreated("success created enroll", enroll)
	c.JSON(resp.Status(), resp)
}

func (h *enrollsHandler) Update(c *gin.Context) {
	enrollID, err := controller_utils.GetIDInt(c.Param("enroll_id"), "enroll id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var request atDomain.Enroll
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	request.ID = enrollID
	enroll, err := h.service.Update(request, c.Query("access_token"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success updated enroll", enroll)
	c.JSON(resp.Status(), resp)
}

func (h *enrollsHandler) Delete(c *gin.Context) {
	enrollID, idErr := controller_utils.GetIDInt(c.Param("enroll_id"), "enroll id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := h.service.Delete(enrollID, c.Query("access_token")); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted user enroll", "status": http.StatusOK})
}
