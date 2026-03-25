package handler

import (
	"Document-Service/internal/interfaces"
	"Document-Service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	service interfaces.DocumentService
}

func NewDocumentHandler(service interfaces.DocumentService) *DocumentHandler {
	return &DocumentHandler{service: service}
}

func (handler *DocumentHandler) GetDocuments(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, handler.service.GetAll())
}

func (handler *DocumentHandler) GetDocumentById(context *gin.Context) {
	id := context.Param("id")

	document, err := handler.service.GetByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Document not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, document)
}

func (handler *DocumentHandler) AddDocument(context *gin.Context) {
	var newDocument model.Document

	if err := context.ShouldBindJSON(&newDocument); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	added, err := handler.service.Add(newDocument)
	if err != nil {
		if err.Error() == "Document already exists" {
			context.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusCreated, added)
}

func (handler *DocumentHandler) DeleteDocument(context *gin.Context) {
	id := context.Param("id")

	err := handler.service.Delete(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
