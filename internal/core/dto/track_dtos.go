package dto

type TrackGet struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type TrackInfo struct {
	ID                uint    `json:"trackId"`
	Name              string  `json:"trackName"`
	State             string  `json:"state"`
	TotalDistanceMtrs float64 `json:"totalDistanceMeters"`
	AverageSpeedMps   float64 `json:"averageSpeedMps"`
}

type TrackPost struct {
	Name string `json:"name"`
}
