package interfaces

import "Document-Service/internal/model"

type DocumentRepository interface {
	FindAll() []model.Document
	FindByID(id string) (*model.Document, error)
	Save(d model.Document)
	Delete(id string) error
}

type DocumentService interface {
	GetAll() []model.Document
	GetByID(id string) (*model.Document, error)
	Add(document model.Document) (model.Document, error)
	Delete(id string) error
}
