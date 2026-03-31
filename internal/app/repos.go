package app

import (
	md "daws/internal/models"
	"daws/internal/utils"
)

func (a *App) NewRepo(projectName string, repo *md.Repo) error {
	project, err := a.GetProject(projectName)
	if err != nil {
		return utils.Logger.Error("Failed to get project: %v", err)
	}
	err = project.NewRepo(repo)
	if err != nil {
		return utils.Logger.Error("Failed to add repo to project: %v", err)
	}
	err = a.SM.SaveProject(project)
	if err != nil {
		return utils.Logger.Error("Failed to save project: %v", err)
	}
	return nil
}

func (a *App) GetRepo(projectName string, repoName string) (*md.Repo, error) {
	project, err := a.GetProject(projectName)
	if err != nil {
		return nil, utils.Logger.Error("Failed to get project: %v", err)
	}
	repo, err := project.GetRepo(repoName)
	if err != nil {
		return nil, utils.Logger.Error("Failed to get repo from project: %v", err)
	}
	return repo, nil
}

func (a *App) UpdateRepo(projectName string, repoName string, repo *md.Repo) error {
	project, err := a.GetProject(projectName)
	if err != nil {
		return utils.Logger.Error("Failed to get project: %v", err)
	}
	err = project.UpdateRepo(repoName, repo)
	if err != nil {
		return utils.Logger.Error("Failed to update repo in project: %v", err)
	}
	err = a.SM.SaveProject(project)
	if err != nil {
		return utils.Logger.Error("Failed to save project: %v", err)
	}
	return nil
}

func (a *App) DeleteRepo(projectName string, repoName string) error {
	project, err := a.GetProject(projectName)
	if err != nil {
		return utils.Logger.Error("Failed to get project: %v", err)
	}
	err = project.DeleteRepo(repoName)
	if err != nil {
		return utils.Logger.Error("Failed to delete repo from project: %v", err)
	}
	err = a.SM.SaveProject(project)
	if err != nil {
		return utils.Logger.Error("Failed to save project: %v", err)
	}
	return nil
}
