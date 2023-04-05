package interfaces

type TrackInterface struct {
	Path              string `json:"path" validate:"required"`
	TrackNameFallback string `json:"track_name_fallback" validate:"required"`
}
