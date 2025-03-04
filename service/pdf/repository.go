package pdf

import (
	"context"
	"demo-gotenberg/models"
)

type IPDFRepository interface {
	GeneratePDFFromURL(ctx context.Context, req *models.PDFFile) error
}
