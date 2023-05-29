package enum

type TrackState int8

const (
	TrackCreated TrackState = iota
	TrackActive
	TrackCompleted
)

func (s TrackState) String() string {
	switch s {
	case TrackCreated:
		return "Created"
	case TrackActive:
		return "Active"
	case TrackCompleted:
		return "Completed"
	}
	return "NotSet"
}
