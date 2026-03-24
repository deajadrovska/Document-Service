package mocks

import (
	"Document-Service/internal/model"
	"errors"
)

type MockDocumentRepository struct {
	documents []model.Document
}

func NewMockDocumentRepository() *MockDocumentRepository {
	return &MockDocumentRepository{
		documents: []model.Document{
			{ID: "1", Name: "Document One", Description: "description"},
			{ID: "2", Name: "Document Two", Description: "description"},
			{ID: "3", Name: "Document Three", Description: "description"},
		},
	}
}

func (m *MockDocumentRepository) FindAll() []model.Document {
	return m.documents
}

func (m *MockDocumentRepository) FindByID(id string) (*model.Document, error) {
	for i, doc := range m.documents {
		if doc.ID == id {
			return &m.documents[i], nil
		}
	}
	return nil, errors.New("document not found")
}

func (m *MockDocumentRepository) Save(d model.Document) {
	m.documents = append(m.documents, d)
}

func (m *MockDocumentRepository) Delete(id string) error {
	for i, doc := range m.documents {
		if doc.ID == id {
			m.documents = append(m.documents[:i], m.documents[i+1:]...)
			return nil
		}
	}
	return errors.New("document not found")
}
