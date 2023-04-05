package dto

type WaypointGet struct {
	ID      uint64  `json:"id"`
	TrackID uint    `json:"trackId"`
	Lat     float64 `json:"lat"`
	LatHem  string  `json:"latHem"`
	Long    float64 `json:"long"`
	LongHem string  `json:"longHem"`
	Time    int64   `json:"time"`
}

type WaypointPost struct {
	TrackID uint    `json:"trackId"`
	Lat     float64 `json:"lat"`
	LatHem  string  `json:"latHem"`
	Long    float64 `json:"long"`
	LongHem string  `json:"longHem"`
	Time    int64   `json:"time"`
}
