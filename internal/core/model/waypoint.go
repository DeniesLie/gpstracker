package model

type Waypoint struct {
	ID        uint64
	Lat       float64
	LatHem    string
	Long      float64
	LongHem   string
	Timestamp int64
	TrackID   uint
	Track     Track
}
