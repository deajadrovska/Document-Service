package service

import (
	"Document-Service/internal/interfaces"
	"Document-Service/internal/model"
	"errors"
)

type DocumentService struct {
	repository interfaces.DocumentRepository
}

func NewDocumentService(repository interfaces.DocumentRepository) *DocumentService {
	return &DocumentService{repository: repository}
}

func (s *DocumentService) GetAll() []model.Document {
	return s.repository.FindAll()
}

func (s *DocumentService) GetByID(id string) (*model.Document, error) {
	if id == " " {
		return nil, errors.New("invalid document id")
	}

	return s.repository.FindByID(id)
}

func (s *DocumentService) Add(document model.Document) (model.Document, error) {
	if document.ID == "" {
		return model.Document{}, errors.New("Id cannot be empty")
	}

	if document.Name == "" {
		return model.Document{}, errors.New("Name cannot be empty")
	}

	existing, _ := s.repository.FindByID(document.ID)
	if existing != nil {
		return model.Document{}, errors.New("Document already exists")
	}

	s.repository.Save(document)
	return document, nil
}

func (s *DocumentService) Delete(id string) error {
	if id == "" {
		return errors.New("invalid document id")
	}

	existing, _ := s.repository.FindByID(id)
	if existing == nil {
		return errors.New("Document doesn't exists")
	}

	return s.repository.Delete(id)
}
