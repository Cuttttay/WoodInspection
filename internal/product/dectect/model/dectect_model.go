package model

type Defect struct {
	ID      int     `json:"id"`
	ImageID string  `json:"image_id"`
	Label   string  `json:"label"`
	Score   float64 `json:"score"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	Area    int     `json:"area"`
	Level   string  `json:"level"`
}
