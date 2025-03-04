package usecase

import (
	"context"
	"demo-gotenberg/models"
	"demo-gotenberg/service/pdf"
)

type pdfUseCase struct {
	pdfRepo pdf.IPDFRepository
}

func NewPDFUseCase(pdfRepo pdf.IPDFRepository) pdf.IPDFUseCase {
	return &pdfUseCase{
		pdfRepo: pdfRepo,
	}
}

func (p *pdfUseCase) GeneratePDFFromURL(ctx context.Context, req *models.PDFFile) error {
	return p.pdfRepo.GeneratePDFFromURL(ctx, req)
}
