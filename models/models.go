package models

//Place to be search
type Place struct {
	Input, InputType, Fields, LocationBias, Point, Center Radius string
}

type Places []*Place