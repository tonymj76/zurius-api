package models

//Place to be search
type Place struct {
	Input, Fields, LocationBias, Point, Center, SouthWest, NorthEast string
	Radius                                                           int
}

//Places list of place
type Places []*Place
