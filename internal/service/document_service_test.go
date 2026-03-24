package service_test

import (
	"Document-Service/internal/mocks"
	"Document-Service/internal/model"
	"Document-Service/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupService() *service.DocumentService {
	mock := mocks.NewMockDocumentRepository()
	return service.NewDocumentService(mock)
}

func TestGetAll(t *testing.T) {
	svc := setupService()

	documents := svc.GetAll()

	assert.NotNil(t, documents)
	assert.Equal(t, 3, len(documents))
}

func TestGetByID(t *testing.T) {
	svc := setupService()

	// valid id
	document, err := svc.GetByID("1")
	assert.Nil(t, err)
	assert.Equal(t, "1", document.ID)
	assert.Equal(t, "Document One", document.Name)

	// empty id
	document, err = svc.GetByID("")
	assert.NotNil(t, err)
	assert.Nil(t, document)
	assert.Equal(t, "invalid document id", err.Error())

	// non existing id
	document, err = svc.GetByID("999")
	assert.NotNil(t, err)
	assert.Nil(t, document)
}

func TestAdd(t *testing.T) {
	svc := setupService()

	// valid document
	newDocument := model.Document{ID: "4", Name: "Document Four", Description: "description"}
	added, err := svc.Add(newDocument)
	assert.Nil(t, err)
	assert.Equal(t, "4", added.ID)
	assert.Equal(t, "Document Four", added.Name)

	// empty id
	emptyID := model.Document{ID: "", Name: "Document Five", Description: "description"}
	_, err = svc.Add(emptyID)
	assert.NotNil(t, err)
	assert.Equal(t, "Id cannot be empty", err.Error())

	// empty name
	emptyName := model.Document{ID: "5", Name: "", Description: "description"}
	_, err = svc.Add(emptyName)
	assert.NotNil(t, err)
	assert.Equal(t, "Name cannot be empty", err.Error())

	// duplicate id
	duplicate := model.Document{ID: "1", Name: "Duplicate", Description: "description"}
	_, err = svc.Add(duplicate)
	assert.NotNil(t, err)
	assert.Equal(t, "Document already exists", err.Error())
}

func TestDelete(t *testing.T) {
	svc := setupService()

	// valid delete
	err := svc.Delete("1")
	assert.Nil(t, err)

	// verify it was deleted
	documents := svc.GetAll()
	assert.Equal(t, 2, len(documents))

	// empty id
	err = svc.Delete("")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid document id", err.Error())

	// non existing id
	err = svc.Delete("999")
	assert.NotNil(t, err)
	assert.Equal(t, "Document doesn't exists", err.Error())
}
