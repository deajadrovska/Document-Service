package handler_test

import (
	"Document-Service/internal/handler"
	"Document-Service/internal/mocks"
	"Document-Service/internal/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	mock := mocks.NewMockDocumentService()
	h := handler.NewDocumentHandler(mock)

	router := gin.Default()
	router.GET("/documents", h.GetDocuments)
	router.GET("/documents/:id", h.GetDocumentById)
	router.POST("/documents", h.AddDocument)
	router.DELETE("/documents/:id", h.DeleteDocument)

	return router
}

func TestGetDocuments(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/documents", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var documents []model.Document
	err := json.Unmarshal(w.Body.Bytes(), &documents)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(documents))
}

func TestGetDocumentById(t *testing.T) {
	router := setupRouter()

	// valid id
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/documents/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// non existing id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/documents/999", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestAddDocument(t *testing.T) {
	router := setupRouter()

	// valid document
	newDocument := model.Document{ID: "4", Name: "Document Four", Description: "description"}
	body, _ := json.Marshal(newDocument)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// empty name
	emptyName := model.Document{ID: "5", Name: "", Description: "description"}
	body, _ = json.Marshal(emptyName)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// duplicate id
	duplicate := model.Document{ID: "1", Name: "Duplicate", Description: "description"}
	body, _ = json.Marshal(duplicate)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestDeleteDocument(t *testing.T) {
	router := setupRouter()

	// valid delete
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/documents/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// non existing id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/documents/999", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
