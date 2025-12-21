package dto

type DetectResponse struct {
	Image   string    `json:"image"`
	Defect  DetectOne `json:"defect"`
	Verdict string    `json:"verdict"`
}
