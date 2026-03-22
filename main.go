package main

import (
	"Document-Service/internal/handler"
	"Document-Service/internal/repository"
	"Document-Service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	rep := repository.NewDocumentRepository()
	ser := service.NewDocumentService(rep)
	h := handler.NewDocumentHandler(ser)

	router := gin.Default()
	router.GET("/documents", h.GetDocuments)
	router.GET("/documents/:id", h.GetDocumentById)
	router.POST("/documents", h.AddDocument)
	router.DELETE("/documents/:id", h.DeleteDocument)

	err := router.Run(":7070")
	if err != nil {
		return
	}
}
