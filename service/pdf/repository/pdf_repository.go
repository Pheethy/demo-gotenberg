package repository

import (
	"context"
	"demo-gotenberg/constants"
	"demo-gotenberg/models"
	"demo-gotenberg/request"
	"demo-gotenberg/service/pdf"
	"errors"
	"fmt"
	"time"
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
	p.client.GetRestyClient().SetTimeout(30 * time.Second)
	url := fmt.Sprintf("%s/forms/chromium/convert/url", p.client.GetHost())

	headers := fmt.Sprintf(`{"Authorization": "%s"}`, req.Token)

	payload := map[string]string{
		"url":               req.FrontendURL,
		"waitDelay":         "2s",
		"extraHttpHeaders":  headers,
		"paperWidth":        "8.27",  // ความกว้าง A4 ในหน่วยนิ้ว (210mm)
		"paperHeight":       "11.7",  // ความสูง A4 ในหน่วยนิ้ว (297mm)
		"preferCssPageSize": "false", // ใช้ค่าที่กำหนดแทนที่จะใช้จาก CSS
		"printBackground":   "true",  // พิมพ์พื้นหลัง
		"landscape":         "false", // แนวตั้ง (portrait) ไม่ใช่แนวนอน
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
