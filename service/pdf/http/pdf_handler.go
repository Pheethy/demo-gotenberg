package http

import (
	"demo-gotenberg/constants"
	"demo-gotenberg/models"
	"demo-gotenberg/service/pdf"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type pdfHandler struct {
	pdfUs pdf.IPDFUseCase
}

func NewPDFHandler(pdfUs pdf.IPDFUseCase) pdf.IPDFHandler {
	return &pdfHandler{
		pdfUs: pdfUs,
	}
}

func (p *pdfHandler) GenerarePDF(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(models.PDFFile)
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := p.pdfUs.GeneratePDFFromURL(ctx, req); err != nil {
		if ok := strings.Contains(err.Error(), constants.ERROR_GENERATE_PDF_CONTENT); ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	/* Set Meta Data จ้าาาา */
	req.SetCreatedAt()
	req.SetFileType()

	/* เขียน File ให้ Frontend */
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", req.Filename))
	c.Header("Content-Type", req.ContentType)
	c.Header("Content-Length", fmt.Sprintf("%d", req.Size))
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	c.Data(http.StatusOK, req.ContentType, req.Content)
}
