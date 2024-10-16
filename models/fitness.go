package fitness

type FitnessData struct {
	UserId string `json:"userId" bson:"userId`
	Date string `json:"date" bson:"date`
	Workouts []Workout `json:"workouts" bson:"workouts`
}

type Workout struct {
	Type string `json:"type" bson:"type"`
	Duration int `json:"duration" bson:"duration`
}