package repository

import (
	"context"
	"demo-gotenberg/constants"
	"demo-gotenberg/models"
	"demo-gotenberg/request"
	"demo-gotenberg/service/pdf"
	"errors"
	"fmt"
)

type pdfRepository struct {
	client *request.Client
}

func NewPDFRepository(client *request.Client) pdf.IPDFRepository {
	return &pdfRepository{
		client: client,
	}
}

func (p *pdfRepository) GeneratePDFFromURL(ctx context.Context, req *models.PDFFile) error {
	url := fmt.Sprintf("%s/forms/chromium/convert/url", p.client.GetHost())

	payload := map[string]string{
		"url": req.FrontendURL,
	}

	request := p.client.GetRestyClient().R().
		SetMultipartFormData(payload)

	response, err := request.Post(url)
	if err != nil {
		return err
	}

	if response.StatusCode() >= 400 {
		return fmt.Errorf("error generating pdf from url: %s", response.Status())
	}

	if response.StatusCode() == 200 {
		req.ContentType = "application/pdf"
		req.Size = response.RawResponse.ContentLength
		req.Content = response.Body()
	}

	if len(req.Content) == 0 || req.Content == nil {
		return errors.New(constants.ERROR_GENERATE_PDF_CONTENT)
	}

	return nil
}
