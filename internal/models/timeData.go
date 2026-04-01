package models

import "time"

type TimeData struct {
	LastBuild  time.Time `json:"lastBuild"`
	LastTest   time.Time `json:"lastTest"`
	LastDeploy time.Time `json:"lastDeploy"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (t *TimeData) Create() {
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	t.UpdatedAt = time.Now()
}

func (t *TimeData) Update() {
	t.UpdatedAt = time.Now()
}

func (t *TimeData) Set(key string) {
	switch key {
	case "build":
		t.LastBuild = time.Now()
		return
	case "test":
		t.LastTest = time.Now()
		return
	case "deploy":
		t.LastDeploy = time.Now()
		return
	case "create":
		t.Create()
	case "update":
		t.UpdatedAt = time.Now()
		return
	default:
		return
	}
}
