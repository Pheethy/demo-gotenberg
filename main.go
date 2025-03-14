package main

import (
	"demo-gotenberg/request"
	"demo-gotenberg/route"
	_pdf_handler "demo-gotenberg/service/pdf/http"
	_pdf_repository "demo-gotenberg/service/pdf/repository"
	_pdf_usecase "demo-gotenberg/service/pdf/usecase"
	"log"
	"time"

	"github.com/Pheethy/psql/helper"

	"github.com/gin-gonic/gin"
)

var gotenbergHost = helper.GetENV("GOTENBERG_HOST", "")

func main() {
	client := request.New(gotenbergHost, false)
	client.GetRestyClient().SetTimeout(30 * time.Second)

	/*Init Repository */
	pdfRepo := _pdf_repository.NewPDFRepository(client)

	/*Init Usecase */
	pdfUs := _pdf_usecase.NewPDFUseCase(pdfRepo)

	/*Init Handler */
	pdfHandler := _pdf_handler.NewPDFHandler(pdfUs)

	/* Init Web Server */
	app := gin.Default()

	/*Init Router */
	route := route.NewRoute(app)
	route.RegisterPDF(pdfHandler)

	/* Start Server On Port 3600 จ้าาาาา*/
	log.Println("🍫 Server demo pdf started on port 3600")
	app.Run(":3600")
}
