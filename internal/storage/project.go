package storage

import (
	md "daws/internal/models"
	"daws/internal/utils"
	"fmt"
	"os"
	"path/filepath"
)

func (s *StorageManager) SaveProject(project *md.Project) error {
	filename := fmt.Sprintf("%s.json", utils.Slugify(project.Name))
	path := filepath.Join(s.dataDir, filename)
	project.Save_Preprocess()
	return s.save(path, project)
}

func (s *StorageManager) LoadProject(projectName string) (*md.Project, error) {
	var project md.Project
	filename := fmt.Sprintf("%s.json", utils.Slugify(projectName))
	err := s.load(filename, &project)
	return &project, err
}

func (s *StorageManager) DeleteProject(projectName string) error {
	filename := fmt.Sprintf("%s.json", utils.Slugify(projectName))
	path := filepath.Join(s.dataDir, filename)
	err := os.Remove(path)
	if err != nil {
		return utils.Logger.Error("Failed to delete project file: %v", err)
	}
	return utils.Logger.Success("Project deleted: %s", projectName)
}

func (s *StorageManager) GetAllProjects() (map[string]md.Project, error) {
	projects := make(map[string]md.Project)
	projectsDir := s.GetDataDir()
	files, err := os.ReadDir(projectsDir)
	if err != nil {
		return projects, utils.Logger.Error("Failed to read projects directory: %v", err)
	}
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}
		var project md.Project
		err := s.load(filepath.Join(projectsDir, file.Name()), &project)
		if err != nil {
			utils.Logger.Warn("Failed to load project from file %s: %v", file.Name(), err)
			continue
		}
		projects[project.Slug] = project
	}
	return projects, nil
}
