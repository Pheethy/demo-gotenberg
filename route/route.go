package route

import (
	"demo-gotenberg/service/pdf"

	"github.com/gin-gonic/gin"
)

type Route struct {
	e *gin.Engine
}

func NewRoute(e *gin.Engine) *Route {
	return &Route{
		e: e,
	}
}

func (r *Route) RegisterPDF(handler pdf.IPDFHandler) {
	r.e.POST("/v1/pdf", handler.GenerarePDF)
}
