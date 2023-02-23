package models

// Tweet estructura
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
