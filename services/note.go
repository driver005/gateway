package services

import (
	"context"
	"reflect"

	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
	"github.com/icza/gox/gox"
)

type NoteService struct {
	ctx context.Context
	r   Registry
}

func NewNoteService(
	r Registry,
) *NoteService {
	return &NoteService{
		context.Background(),
		r,
	}
}

func (s *NoteService) SetContext(context context.Context) *NoteService {
	s.ctx = context
	return s
}

func (s *NoteService) Retrieve(id uuid.UUID, config *sql.Options) (*models.Note, *utils.ApplictaionError) {
	var res *models.Note
	query := sql.BuildQuery(&models.Note{Model: core.Model{Id: id}}, config)
	if err := s.r.NoteRepository().FindOne(s.ctx, res, query); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *NoteService) List(selector *types.FilterableNote, config *sql.Options) ([]models.Note, *utils.ApplictaionError) {
	result, _, err := s.ListAndCount(selector, config)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *NoteService) ListAndCount(selector *types.FilterableNote, config *sql.Options) ([]models.Note, *int64, *utils.ApplictaionError) {
	if reflect.DeepEqual(config, &sql.Options{}) {
		config.Skip = gox.NewInt(0)
		config.Take = gox.NewInt(50)
		config.Order = gox.NewString("created_at DESC")
	}

	var res []models.Note

	query := sql.BuildQuery(selector, config)

	count, err := s.r.NoteRepository().FindAndCount(s.ctx, res, query)
	if err != nil {
		return nil, nil, err
	}
	return res, count, nil
}

func (s *NoteService) Create(data *types.CreateNoteInput, config map[string]interface{}) (*models.Note, *utils.ApplictaionError) {
	model := &models.Note{
		Model: core.Model{
			Metadata: data.Metadata,
		},
		Value:        data.Value,
		ResourceType: data.ResourceType,
		ResourceId:   uuid.NullUUID{UUID: data.ResourceId},
		AuthorId:     uuid.NullUUID{UUID: data.AuthorId},
		Author:       data.Author,
	}

	model.Metadata = utils.MergeMaps(model.Metadata, config)

	// s.eventBus_.withTransaction(manager).emit(NoteService.Events.CREATED, map[string]interface{}{"id": result.id})

	if err := s.r.NoteRepository().Save(s.ctx, model); err != nil {
		return nil, err
	}
	return model, nil
}

func (s *NoteService) Update(id uuid.UUID, data *types.UpdateNoteInput) (*models.Note, *utils.ApplictaionError) {
	note, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return nil, err
	}

	note.Value = data.Value
	// s.eventBus_.withTransaction(manager).emit(NoteService.Events.UPDATED, map[string]interface{}{"id": result.id})

	if err := s.r.NoteRepository().Save(s.ctx, note); err != nil {
		return nil, err
	}
	return note, nil
}

func (s *NoteService) Delete(id uuid.UUID) *utils.ApplictaionError {
	note, err := s.Retrieve(id, &sql.Options{})
	if err != nil {
		return err
	}

	if note == nil {
		return nil
	}

	if err := s.r.NoteRepository().SoftRemove(s.ctx, note); err != nil {
		return err
	}
	// s.eventBus_.withTransaction(manager).emit(NoteService.Events.DELETED, map[string]interface{}{"id": noteId})

	return nil
}
