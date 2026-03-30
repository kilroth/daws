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

func isWailsDev() bool {
	_, err := os.Stat("frontend")
	return !os.IsNotExist(err)
}

func getBasePath() string {
	if isWailsDev() {
		md.Logger.Info("Running in development mode, using current directory for storage")
		basePath, _ := os.Getwd()
		return filepath.Join(basePath, ".dev_data")
	}
	ex, err := os.Executable()
	if err != nil {
		md.Logger.Panic("Failed to get executable path: %v", err)
	}
	exPath := filepath.Dir(ex)
	return filepath.Join(exPath, "data")
}

func isReadOnly(dir string) bool {
	testFile := filepath.Join(dir, ".write_test")

	// Attempt to create/write a tiny file
	err := os.WriteFile(testFile, []byte("ok"), 0644)
	if err != nil {
		return true // It's read-only
	}

	// Clean up if it worked
	os.Remove(testFile)
	return false
}

func slugify(name string) string {
	tmpName := strings.ReplaceAll(name, " ", "-")
	return tmpName // TODO: implement slugification
}

func NewStorageManager(baseDataPath string) *StorageManager {
	/* Old idea
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
	*/

	basePath := getBasePath()

	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		md.Logger.Warn("Working directory does not exist, creating: %s", basePath)
		err := os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			md.Logger.Panic("Failed to create data directory: %v", err)
		}
	}

	if isReadOnly(basePath) {
		md.Logger.Panic("Working directory is read-only, cannot continue: ")
	}

	configRoot := basePath

	pPath := baseDataPath
	if pPath == "" {
		pPath = filepath.Join(basePath, "projects")
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
