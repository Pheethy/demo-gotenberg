package pdf

import (
	"context"
	"demo-gotenberg/models"
)

type IPDFUseCase interface {
	GeneratePDFFromURL(ctx context.Context, req *models.PDFFile) error
}
