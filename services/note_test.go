package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/driver005/gateway/utils"
	"github.com/google/uuid"
)

func TestNewNoteService(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *NoteService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoteService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *NoteService
		args args
		want *NoteService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteService_Retrieve(t *testing.T) {
	type args struct {
		id     uuid.UUID
		config *sql.Options
	}
	tests := []struct {
		name  string
		s     *NoteService
		args  args
		want  *models.Note
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Retrieve(tt.args.id, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.Retrieve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NoteService.Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNoteService_List(t *testing.T) {
	type args struct {
		selector *types.FilterableNote
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *NoteService
		args  args
		want  []models.Note
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.List(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.List() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NoteService.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNoteService_ListAndCount(t *testing.T) {
	type args struct {
		selector *types.FilterableNote
		config   *sql.Options
	}
	tests := []struct {
		name  string
		s     *NoteService
		args  args
		want  []models.Note
		want1 *int64
		want2 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.ListAndCount(tt.args.selector, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.ListAndCount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NoteService.ListAndCount() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("NoteService.ListAndCount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestNoteService_Create(t *testing.T) {
	type args struct {
		data   *types.CreateNoteInput
		config map[string]interface{}
	}
	tests := []struct {
		name  string
		s     *NoteService
		args  args
		want  *models.Note
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Create(tt.args.data, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NoteService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNoteService_Update(t *testing.T) {
	type args struct {
		id   uuid.UUID
		data *types.UpdateNoteInput
	}
	tests := []struct {
		name  string
		s     *NoteService
		args  args
		want  *models.Note
		want1 *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Update(tt.args.id, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NoteService.Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNoteService_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name string
		s    *NoteService
		args args
		want *utils.ApplictaionError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
