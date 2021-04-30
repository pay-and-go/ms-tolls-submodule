package Model

type TollsInRoute struct {
	Route string     `json:"route" bson:"route"`
	Tolls []int   `json:"tolls" bson:"tolls"`
}