package user

type Achievement struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Level       uint    `json:"level"`
	Progress    float32 `json:"progress"`
}
