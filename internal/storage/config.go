package storage

import (
	md "daws/internal/models"
	"path/filepath"
)

func (s *StorageManager) SaveConfig(config md.AppConfig) error {
	path := filepath.Join(s.baseDir, "config.json")
	return s.save(path, config)
}

func (s *StorageManager) LoadConfig(config *md.AppConfig) error {
	path := filepath.Join(s.baseDir, "config.json")
	return s.load(path, config)
}
