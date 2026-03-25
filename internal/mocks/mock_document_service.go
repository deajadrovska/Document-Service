package mocks

import (
	"Document-Service/internal/model"
	"errors"
)

type MockDocumentService struct {
	documents []model.Document
}

func NewMockDocumentService() *MockDocumentService {
	return &MockDocumentService{
		documents: []model.Document{
			{ID: "1", Name: "Document One", Description: "description"},
			{ID: "2", Name: "Document Two", Description: "description"},
			{ID: "3", Name: "Document Three", Description: "description"},
		},
	}
}

func (m *MockDocumentService) GetAll() []model.Document {
	return m.documents
}

func (m *MockDocumentService) GetByID(id string) (*model.Document, error) {
	for i, doc := range m.documents {
		if doc.ID == id {
			return &m.documents[i], nil
		}
	}
	return nil, errors.New("document not found")
}

func (m *MockDocumentService) Add(document model.Document) (model.Document, error) {
	if document.ID == "" {
		return model.Document{}, errors.New("Id cannot be empty")
	}
	if document.Name == "" {
		return model.Document{}, errors.New("Name cannot be empty")
	}
	existing, _ := m.GetByID(document.ID)
	if existing != nil {
		return model.Document{}, errors.New("Document already exists")
	}
	m.documents = append(m.documents, document)
	return document, nil
}

func (m *MockDocumentService) Delete(id string) error {
	if id == "" {
		return errors.New("invalid document id")
	}
	for i, doc := range m.documents {
		if doc.ID == id {
			m.documents = append(m.documents[:i], m.documents[i+1:]...)
			return nil
		}
	}
	return errors.New("Document doesn't exists")
}
