package service

import "github.com/DeniesLie/gpstracker/internal/core/model"

type TrackRepoMock struct {
	GetAllResult    []model.Track
	GetAllError     error
	GetByIdResult   *model.Track
	GetByIdError    error
	GetByNameResult *model.Track
	GetByNameError  error
	AddError        error
	UpdateError     error
	DeleteError     error
}

func (r *TrackRepoMock) GetAll() (tracks []model.Track, err error) {
	return r.GetAllResult, r.GetAllError
}

func (r *TrackRepoMock) GetById(id uint) (track *model.Track, err error) {
	return r.GetByIdResult, r.GetByIdError
}

func (r *TrackRepoMock) GetByName(name string) (t *model.Track, err error) {
	return r.GetByNameResult, r.GetByNameError
}

func (r *TrackRepoMock) Add(t *model.Track) error {
	return r.AddError
}

func (r *TrackRepoMock) Update(t *model.Track) error {
	return r.UpdateError
}

func (r *TrackRepoMock) Delete(id uint) error {
	return r.DeleteError
}
