package repository

import (
	"Document-Service/internal/model"
	"errors"
	"sync"
)

type DocumentRepository struct {
	documents []model.Document
	mu        sync.Mutex
}

func NewDocumentRepository() *DocumentRepository {
	return &DocumentRepository{
		documents: []model.Document{
			{ID: "1", Name: "Document One", Description: "description"},
			{ID: "2", Name: "Document Two", Description: "description"},
			{ID: "3", Name: "Document Three", Description: "description"},
		},
	}
}

func (r *DocumentRepository) FindAll() []model.Document {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.documents
}

func (r *DocumentRepository) FindByID(id string) (*model.Document, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, doc := range r.documents {
		if doc.ID == id {
			return &r.documents[i], nil
		}
	}
	return nil, errors.New("document not found")
}

func (r *DocumentRepository) Save(document model.Document) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.documents = append(r.documents, document)
}

func (r *DocumentRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, doc := range r.documents {
		if doc.ID == id {
			r.documents = append(r.documents[:i], r.documents[i+1:]...)
			return nil
		}
	}
	return errors.New("document not found")
}
