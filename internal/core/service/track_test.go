//go:build unit_test

package service

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/DeniesLie/gpstracker/internal/core/mapper"
	"github.com/DeniesLie/gpstracker/internal/core/model"
	"github.com/DeniesLie/gpstracker/internal/core/model/enum"
	"github.com/DeniesLie/gpstracker/internal/core/validation"
	"github.com/DeniesLie/gpstracker/internal/helpers"
	"github.com/pkg/errors"
)

func TestGetAll(t *testing.T) {
	notEmptyRepoResult := getTracks(5)
	notEmptyResult := mapper.MapSlice(notEmptyRepoResult, mapper.ToTrackDto)
	repoGetAllError := errors.New("db error")
	wrappedError := errors.Wrap(repoGetAllError, "failed at TrackService.GetAll(), some error occurred in repo")

	testCases := []struct {
		description       string
		trackGetAllResult []model.Track
		trackGetAllError  error
		expectedResult    []dto.TrackGet
		expectedError     error
	}{
		{
			description:       "get entries from repo",
			trackGetAllResult: notEmptyRepoResult,
			expectedResult:    notEmptyResult,
		},
		{
			description:       "get empty slice from repo",
			trackGetAllResult: []model.Track{},
			expectedResult:    []dto.TrackGet{},
		},
		{
			description:      "get error from repo",
			trackGetAllError: repoGetAllError,
			expectedError:    wrappedError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			trackRepoMock := &TrackRepoMock{
				GetAllResult: tc.trackGetAllResult,
				GetAllError:  tc.trackGetAllError,
			}
			waypointRepoMock := &WaypointRepoMock{}

			sut := NewTrackService(trackRepoMock, waypointRepoMock)
			res, err := sut.GetAll()

			if !helpers.ErrorsEqual(err, tc.expectedError) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedError, err)
			}

			if !reflect.DeepEqual(tc.expectedResult, res) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedResult, res)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	newValidTrack := dto.TrackPost{Name: "NewTrack"}
	emptyNameTrack := dto.TrackPost{Name: ""}

	existingTrack := getTrack(3)

	repoError := errors.New("repo error")
	wrappedRepoError := errors.Wrap(repoError, "failed at TrackService.Create(), some error occurred in repo")
	responseTrackDto := dto.TrackGet{ID: 0, Name: "NewTrack", State: enum.TrackActive.String()}

	testCases := []struct {
		description         string
		trackToAdd          dto.TrackPost
		repoGetByNameResult *model.Track
		repoGetByNameError  error
		repoAddError        error
		expectedResult      *dto.TrackGet
		expectedError       error
	}{
		{
			description:    "create track",
			trackToAdd:     newValidTrack,
			expectedResult: &responseTrackDto,
		},
		{
			description:   "invalid track data",
			trackToAdd:    emptyNameTrack,
			expectedError: validation.ValidationError{},
		},
		{
			description:         "case name is taken",
			trackToAdd:          newValidTrack,
			repoGetByNameResult: &existingTrack,
			expectedError:       BusinessError{Message: "Case name is taken, try another name"},
		},
		{
			description:        "get case by name returns error from repo",
			trackToAdd:         newValidTrack,
			repoGetByNameError: repoError,
			expectedError:      wrappedRepoError,
		},
		{
			description:   "add function returned error from repo",
			trackToAdd:    newValidTrack,
			repoAddError:  repoError,
			expectedError: wrappedRepoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			trackRepoMock := &TrackRepoMock{
				GetByNameResult: tc.repoGetByNameResult,
				GetByNameError:  tc.repoGetByNameError,
				AddError:        tc.repoAddError,
			}
			waypointRepoMock := &WaypointRepoMock{}

			sut := NewTrackService(trackRepoMock, waypointRepoMock)
			res, err := sut.Create(tc.trackToAdd)

			if !helpers.ErrorsEqual(err, tc.expectedError) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedError, err)
			}

			if !reflect.DeepEqual(tc.expectedResult, res) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedResult, res)
			}
		})
	}
}

func TestComplete(t *testing.T) {
	track := getTrack(1)
	var trackId uint = 1
	repoError := errors.New("repo error")
	wrappedRepoError := errors.Wrap(repoError, "failed at TrackService.Complete(), some error occurred in repo")

	testCases := []struct {
		description       string
		repoGetByIdResult *model.Track
		repoGetByIdError  error
		repoUpdateError   error
		expectedError     error
	}{
		{
			description:       "complete track",
			repoGetByIdResult: &track,
		},
		{
			description:      "error returned from repo's GetById method",
			repoGetByIdError: repoError,
			expectedError:    wrappedRepoError,
		},
		{
			description:       "track was not found",
			repoGetByIdResult: nil,
			expectedError:     NotFoundError{},
		},
		{
			description:       "error returned from repo's Update method",
			repoGetByIdResult: &track,
			repoUpdateError:   repoError,
			expectedError:     wrappedRepoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			trackRepoMock := &TrackRepoMock{
				GetByIdResult: tc.repoGetByIdResult,
				GetByIdError:  tc.repoGetByIdError,
				UpdateError:   tc.repoUpdateError,
			}
			waypointRepoMock := &WaypointRepoMock{}

			sut := NewTrackService(trackRepoMock, waypointRepoMock)
			err := sut.Complete(trackId)

			if !helpers.ErrorsEqual(err, tc.expectedError) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedError, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	track := getTrack(1)
	var trackId uint = 1
	repoError := errors.New("repo error")
	wrappedRepoError := errors.Wrap(repoError, "failed at TrackService.Delete(), some error occurred in repo")

	testCases := []struct {
		description       string
		repoGetByIdResult *model.Track
		repoGetByIdError  error
		repoDeleteError   error
		expectedError     error
	}{
		{
			description:       "delete track",
			repoGetByIdResult: &track,
		},
		{
			description:      "error returned from repo's GetById method",
			repoGetByIdError: repoError,
			expectedError:    wrappedRepoError,
		},
		{
			description:       "track was not found",
			repoGetByIdResult: nil,
			expectedError:     NotFoundError{},
		},
		{
			description:       "error returned from repo's Delete method",
			repoGetByIdResult: &track,
			repoDeleteError:   repoError,
			expectedError:     wrappedRepoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			trackRepoMock := &TrackRepoMock{
				GetByIdResult: tc.repoGetByIdResult,
				GetByIdError:  tc.repoGetByIdError,
				DeleteError:   tc.repoDeleteError,
			}
			waypointRepoMock := &WaypointRepoMock{}

			sut := NewTrackService(trackRepoMock, waypointRepoMock)
			err := sut.Delete(trackId)

			if !helpers.ErrorsEqual(err, tc.expectedError) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedError, err)
			}
		})
	}
}

func getTrack(trackId uint) (track model.Track) {
	track = model.Track{}
	track.ID = trackId
	track.Name = "TestTrack_" + strconv.FormatUint(uint64(trackId), 10)
	track.State = enum.TrackCreated
	track.CreatedAt = time.Now()
	return
}

func getTracks(size int) (tracks []model.Track) {
	for i := 1; i <= size; i++ {
		tracks = append(tracks, getTrack(uint(i)))
	}
	return tracks
}
