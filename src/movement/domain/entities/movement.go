package entities

type Movement struct {
	ID        int    `json:"id"`
	SensorID  string `json:"sensorId"`
	Timestamp string `json:"timestamp"`
	Motion    bool   `json:"motion"`
}
