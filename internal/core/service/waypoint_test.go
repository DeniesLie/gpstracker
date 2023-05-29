//go:build unit_test

package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/DeniesLie/gpstracker/internal/core/dto"
	"github.com/DeniesLie/gpstracker/internal/core/mapper"
	"github.com/DeniesLie/gpstracker/internal/core/model"
	"github.com/DeniesLie/gpstracker/internal/core/model/enum"
	"github.com/DeniesLie/gpstracker/internal/helpers"
	"github.com/pkg/errors"
)

func TestGetByTrackId(t *testing.T) {
	track := getTrack(1)
	waypoints := getWaypoints(track)
	waypointsGetDto := mapper.MapSlice(waypoints, mapper.ToWaypointGetDto)
	repoError := errors.New("repo error")
	wrappedRepoError := errors.Wrap(repoError, "failed at WaypointService.GetByTrackId(), some error occurred in repo")

	testCases := []struct {
		description            string
		trackRepoGetByIdResult *model.Track
		trackRepoGetByIdError  error
		repoGetByTrackIdResult []model.Waypoint
		repoGetByTrackIdError  error
		expectedResult         []dto.WaypointGet
		expectedError          error
	}{
		{
			description:            "get waypoints by track id",
			trackRepoGetByIdResult: &track,
			repoGetByTrackIdResult: waypoints,
			expectedResult:         waypointsGetDto,
		},
		{
			description:            "get empty list of waypoints by track id",
			trackRepoGetByIdResult: &track,
			repoGetByTrackIdResult: []model.Waypoint{},
			expectedResult:         []dto.WaypointGet{},
		},
		{
			description:           "get track by id returns error",
			trackRepoGetByIdError: repoError,
			expectedError:         wrappedRepoError,
		},
		{
			description:            "get waypoints by track id returned error",
			trackRepoGetByIdResult: &track,
			repoGetByTrackIdError:  repoError,
			expectedError:          wrappedRepoError,
		},
		{
			description:            "track was not found",
			trackRepoGetByIdResult: nil,
			expectedError:          NotFoundError{Resource: "Track"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			trackRepoMock := &TrackRepoMock{
				GetByIdResult: tc.trackRepoGetByIdResult,
				GetByIdError:  tc.trackRepoGetByIdError,
			}
			waypointRepoMock := &WaypointRepoMock{
				GetByTrackIdResult: tc.repoGetByTrackIdResult,
				GetByTrackIdError:  tc.repoGetByTrackIdError,
			}

			sut := NewWaypointService(waypointRepoMock, trackRepoMock)

			res, err := sut.GetByTrackId(track.ID)

			if !helpers.ErrorsEqual(err, tc.expectedError) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedError, err)
			}

			if !reflect.DeepEqual(tc.expectedResult, res) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedResult, res)
			}
		})
	}
}

func TestAddBatch(t *testing.T) {
	track := getTrack(1)

	completedTrack := getTrack(1)
	completedTrack.State = enum.TrackCompleted

	waypointsPostDto := getWaypointsPostDto(track.ID)
	waypoints := getWaypoints(track)
	waypointsGetDto := mapper.MapSlice(waypoints, mapper.ToWaypointGetDto)

	repoError := errors.New("repo error")
	wrappedRepoError := errors.Wrap(repoError, "failed at WaypointService.AddBatch(), some error occurred in repo")

	testCases := []struct {
		description            string
		trackRepoGetByIdResult *model.Track
		trackRepoGetByIdError  error
		repoAddRangeError      error
		expectedResult         []dto.WaypointGet
		expectedError          error
	}{
		{
			description:            "add waypoints",
			trackRepoGetByIdResult: &track,
			expectedResult:         waypointsGetDto,
		},
		{
			description:           "get track by id returned error",
			trackRepoGetByIdError: repoError,
			expectedError:         wrappedRepoError,
		},
		{
			description:            "track was not found",
			trackRepoGetByIdResult: nil,
			expectedError:          NotFoundError{Resource: "Track"},
		},
		{
			description:            "track is completed",
			trackRepoGetByIdResult: &completedTrack,
			expectedError:          BusinessError{Message: "can't add waypoints to completed tracks"},
		},
		{
			description:            "add range returned error",
			trackRepoGetByIdResult: &track,
			repoAddRangeError:      repoError,
			expectedError:          wrappedRepoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			trackRepoMock := &TrackRepoMock{
				GetByIdResult: tc.trackRepoGetByIdResult,
				GetByIdError:  tc.trackRepoGetByIdError,
			}
			waypointRepoMock := &WaypointRepoMock{
				AddRangeError: tc.repoAddRangeError,
			}

			sut := NewWaypointService(waypointRepoMock, trackRepoMock)

			err := sut.AddBatch(waypointsPostDto)

			if !helpers.ErrorsEqual(err, tc.expectedError) {
				t.Fatalf("%s. expected '%v', got '%v'", tc.description, tc.expectedError, err)
			}
		})
	}
}

func getWaypoints(track model.Track) []model.Waypoint {
	return []model.Waypoint{
		{
			ID:        1,
			Lat:       1.0,
			LatHem:    "E",
			Long:      1.0,
			LongHem:   "N",
			Timestamp: time.Now().UnixMilli(),
			TrackID:   track.ID,
			Track:     track,
		},
		{
			ID:        2,
			Lat:       1.1,
			LatHem:    "E",
			Long:      1.1,
			LongHem:   "N",
			Timestamp: time.Now().UnixMilli(),
			TrackID:   track.ID,
			Track:     track,
		},
		{
			ID:        3,
			Lat:       1.1,
			LatHem:    "E",
			Long:      1.1,
			LongHem:   "N",
			Timestamp: time.Now().UnixMilli(),
			TrackID:   track.ID,
			Track:     track,
		},
	}
}

func getWaypointsPostDto(trackId uint) []dto.WaypointPost {
	return []dto.WaypointPost{
		{
			TrackID: trackId,
			Lat:     1.0,
			LatHem:  "S",
			Long:    1.0,
			LongHem: "W",
			Time:    time.Now().UnixMilli(),
		},
		{
			TrackID: trackId,
			Lat:     1.1,
			LatHem:  "S",
			Long:    1.1,
			LongHem: "W",
			Time:    time.Now().UnixMilli(),
		},
		{
			TrackID: trackId,
			Lat:     1.2,
			LatHem:  "S",
			Long:    1.2,
			LongHem: "W",
			Time:    time.Now().UnixMilli(),
		},
	}
}
