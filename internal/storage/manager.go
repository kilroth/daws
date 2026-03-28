package storage

import (
	md "daws/internal/models"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type StorageManager struct {
	baseDir string
	dataDir string
}

func slugify(name string) string {
	tmpName := strings.ReplaceAll(name, " ", "-")
	return tmpName // TODO: implement slugification
}

func NewStorageManager(baseDataPath string) *StorageManager {
	dir, err := os.UserConfigDir()
	if err != nil {
		md.Logger.Panic("Failed to get user config directory: %v", err)
	}
	configRoot := filepath.Join(dir, "daws")
	if _, err := os.Stat(configRoot); os.IsNotExist(err) {
		md.Logger.Warn("Config directory does not exist, creating: %s", configRoot)
		err := os.MkdirAll(configRoot, os.ModePerm)
		if err != nil {
			md.Logger.Panic("Failed to create config directory: %v", err)
		}
	}

	pPath := baseDataPath
	if pPath == "" {
		pPath = filepath.Join(configRoot, "projects")
	}
	if _, err := os.Stat(pPath); os.IsNotExist(err) {
		md.Logger.Warn("Data directory does not exist, creating: %s", pPath)
		err := os.MkdirAll(pPath, os.ModePerm)
		if err != nil {
			md.Logger.Panic("Failed to create data directory: %v", err)
		}
	}
	return &StorageManager{baseDir: configRoot, dataDir: pPath}
}

func (s *StorageManager) save(path string, data interface{}) error {
	var err error
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return md.Logger.Error("Failed to marshal data: %v", err)
	}
	err = os.WriteFile(path, bytes, os.ModePerm)
	if err != nil {
		return md.Logger.Error("Failed to write file: %v", err)
	}
	return md.Logger.Success("Data saved to %s", path)
}

func (s *StorageManager) load(path string, target interface{}) error {
	var err error
	bytes, err := os.ReadFile(path)
	if err != nil {
		return md.Logger.Error("Failed to read file: %v", err)
	}
	err = json.Unmarshal(bytes, target)
	if err != nil {
		return md.Logger.Error("Failed to unmarshal data: %v", err)
	}
	return md.Logger.Success("Data loaded from %s", path)
}

func (s *StorageManager) GetDataDir() string {
	return s.dataDir
}

func (s *StorageManager) GetBaseDir() string {
	return s.baseDir
}
