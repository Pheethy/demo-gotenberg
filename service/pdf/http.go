package pdf

import "github.com/gin-gonic/gin"

type IPDFHandler interface {
	GenerarePDF(c *gin.Context)
}
