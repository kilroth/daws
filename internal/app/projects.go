package app

import (
	"os"
	"path/filepath"

	md "daws/internal/models"
)

func (a *App) CheckProjectsDir() error {
	projectsDir := filepath.Join(a.SM.GetDataDir(), "projects")
	if _, err := os.Stat(projectsDir); os.IsNotExist(err) {
		md.Logger.Warn("Projects directory does not exist, creating: %s", projectsDir)
		err := os.MkdirAll(projectsDir, os.ModePerm)
		if err != nil {
			md.Logger.Panic("Failed to create projects directory: %v", err)
		}
	}
	return nil
}

func (a *App) DataLoad() (md.AppData, error) {
	var data md.AppData
	var err error

	projects, err := a.SM.GetAllProjects()
	if err != nil {
		return data, md.Logger.Error("Failed to load projects: %v", err)
	}
	data.Projects = projects
	return data, nil
}

func (a *App) NewProject(project *md.Project) error {
	if project == nil {
		return md.Logger.Error("Project cannot be nil")
	}
	if project.Name == "" {
		return md.Logger.Error("Project name cannot be empty")
	}
	err := a.SM.SaveProject(project)
	if err != nil {
		return md.Logger.Error("Failed to save project: %v", err)
	}
	return nil
}

func (a *App) GetProject(projectName string) (*md.Project, error) {
	if projectName == "" {
		return nil, md.Logger.Error("Project name cannot be empty")
	}
	project, exists := a.Data.Projects[projectName]
	if !exists {
		return nil, md.Logger.Error("Project not found: %s", projectName)
	}
	return &project, nil
}

func (a *App) UpdateProject(projectName string, project *md.Project) error {
	if project == nil {
		return md.Logger.Error("Project cannot be nil")
	}
	if project.Name == "" {
		return md.Logger.Error("Project name cannot be empty")
	}
	if _, exists := a.Data.Projects[projectName]; !exists {
		return md.Logger.Error("Project not found: %s", projectName)
	}
	err := a.SM.SaveProject(project)
	if err != nil {
		return md.Logger.Error("Failed to save project: %v", err)
	}
	return nil
}

func (a *App) DeleteProject(projectName string) error {
	if projectName == "" {
		return md.Logger.Error("Project name cannot be empty")
	}
	if _, exists := a.Data.Projects[projectName]; !exists {
		return md.Logger.Error("Project not found: %s", projectName)
	}
	err := a.SM.DeleteProject(projectName)
	if err != nil {
		return md.Logger.Error("Failed to delete project: %v", err)
	}
	delete(a.Data.Projects, projectName)
	return nil
}
