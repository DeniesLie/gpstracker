package repo

import (
	"github.com/DeniesLie/gpstracker/internal/core/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TrackRepo struct {
	db *gorm.DB
}

func NewTrackRepo(db *gorm.DB) *TrackRepo {
	return &TrackRepo{db}
}

func (r *TrackRepo) GetAll() (t []model.Track, err error) {
	t = make([]model.Track, 0)
	result := r.db.Find(&t)
	if result.Error != nil {
		err = errors.Wrap(result.Error, "failed at TrackRepo.GetAll(), some db error occurred")
	}
	return
}

func (r *TrackRepo) GetById(id uint) (t *model.Track, err error) {
	result := r.db.First(&t, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		err = errors.Wrap(result.Error, "failed at TrackRepo.GetById(), some db error occurred")
	}
	return
}

func (r *TrackRepo) GetByName(name string) (t *model.Track, err error) {
	result := r.db.Where(&model.Track{Name: name}).First(&t)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		err = errors.Wrap(result.Error, "failed at TrackRepo.GetByName(), some db error occurred")
	}
	return
}

func (r *TrackRepo) Add(t *model.Track) (err error) {
	result := r.db.Create(t)
	if result.Error != nil {
		err = errors.Wrap(result.Error, "failed at TrackRepo.Add(), some db error occurred")
	}
	return
}

func (r *TrackRepo) Update(t *model.Track) (err error) {
	result := r.db.Save(t)
	if result.Error != nil {
		err = errors.Wrap(result.Error, "failed at TrackRepo.Update(), some db error occurred")
	}
	return
}

func (r *TrackRepo) Delete(id uint) (err error) {
	result := r.db.Delete(&model.Track{}, id)
	if result.Error != nil {
		err = errors.Wrap(result.Error, "failed at TrackRepo.Delete(), some db error occurred")
	}
	return
}
