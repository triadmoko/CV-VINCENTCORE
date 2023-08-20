package handler

import (
	"net/http"
	"vincentcore_interview/model/userModel"
	"vincentcore_interview/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func (h *handler) Register(c *gin.Context) {
	req := userModel.RequestUser{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseApi("invalid request body", http.StatusBadRequest, false, nil))
		c.Abort()
		return
	}
	if errs := req.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, helpers.ResponseApi("validation register", http.StatusBadRequest, false, errs))
		c.Abort()
		return
	}
	user, status, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(status, helpers.ResponseApi(err.Error(), status, false, nil))
		c.Abort()
		return
	}
	c.JSON(status, helpers.ResponseApi("success register", status, true, user))
	c.Abort()
}
