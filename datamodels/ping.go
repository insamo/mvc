package datamodels

type Ping struct {
	Context string `json:"context"`
	Message string `json:"message" validate:"-"`
	Status  int    `json:"status" validate:"int"`
}
