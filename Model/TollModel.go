package Model

type Toll struct {
	TollId           int     `json:"tollId" bson:"tollId"`
	Administrator    string  `json:"administrator" bson:"administrator"`
	CoorLat          float64 `json:"coor_lat" bson:"coor_lat"`
	CoorLng          float64 `json:"coor_lng" bson:"coor_lng"`
	CranePhoneNumber string  `json:"crane_phone_number" bson:"crane_phone_number"`
	Name             string  `json:"name" bson:"name"`
	Price            float64 `json:"price" bson:"price"`
	Sector           string  `json:"sector" bson:"sector"`
	Territory        string  `json:"territory" bson:"territory"`
	TollPhoneNumber  string  `json:"toll_phone_number" bson:"toll_phone_number"`
}
