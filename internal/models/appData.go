package models

type AppData struct {
	Projects map[string]Project `json:"projects"`
}
