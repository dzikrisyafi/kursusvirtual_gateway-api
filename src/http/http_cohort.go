package http

import (
	"net/http"

	atDomain "github.com/dzikrisyafi/kursusvirtual_gateway-api/src/domain/cohort"
	"github.com/dzikrisyafi/kursusvirtual_gateway-api/src/services/cohort"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

type CohortHandler interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type cohortHandler struct {
	service cohort.Service
}

func NewCohortHandler(service cohort.Service) CohortHandler {
	return &cohortHandler{
		service: service,
	}
}

func (h *cohortHandler) Create(c *gin.Context) {
	var request atDomain.Cohort
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	cohort, err := h.service.Create(request, c.Query("access_token"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusCreated("success created cohort", cohort)
	c.JSON(resp.Status(), resp)
}

func (h *cohortHandler) Update(c *gin.Context) {
	cohortID, err := controller_utils.GetIDInt(c.Param("cohort_id"), "cohort id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var request atDomain.Cohort
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	request.ID = cohortID
	cohort, err := h.service.Update(request, c.Query("access_token"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success updated cohort", cohort)
	c.JSON(resp.Status(), resp)
}

func (h *cohortHandler) Delete(c *gin.Context) {
	cohortID, err := controller_utils.GetIDInt(c.Param("cohort_id"), "cohort id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := h.service.Delete(cohortID, c.Query("access_token")); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted cohort", "status": http.StatusOK})
}
