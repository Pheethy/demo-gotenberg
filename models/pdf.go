package models

import (
	"fmt"
	"time"

	"github.com/Pheethy/psql/helper"
)

type PDFFile struct {
	Content     []byte
	ContentType string
	Size        int64
	CreatedAt   *helper.Timestamp

	Filename    string `json:"filename" binding:"required"`
	FrontendURL string `json:"frontend_url" binding:"required"`
}

func (p *PDFFile) SetCreatedAt() {
	time := helper.NewTimestampFromTime(time.Now())
	p.CreatedAt = &time
}

func (p *PDFFile) SetFileType() {
	p.Filename = fmt.Sprintf("%s.pdf", p.Filename)
}
