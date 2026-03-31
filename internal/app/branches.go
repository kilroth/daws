package app

import (
	md "daws/internal/models"
	"daws/internal/utils"
)

func (a *App) GetBranch(projectName string, repoName string, branchName string) (*md.Branch, error) {
	project, err := a.GetProject(projectName)
	if err != nil {
		return nil, utils.Logger.Error("Failed to get project: %v", err)
	}
	branch, err := project.GetBranch(repoName, branchName)
	if err != nil {
		return nil, utils.Logger.Error("Failed to get branch from repo: %v", err)
	}
	return branch, nil
}

func (a *App) NewBranch(projectName string, repoName string, branch *md.Branch) error {
	project, err := a.GetProject(projectName)
	if err != nil {
		return utils.Logger.Error("Failed to get project: %v", err)
	}
	err = project.NewBranch(repoName, branch)
	if err != nil {
		return utils.Logger.Error("Failed to add branch to repo in project: %v", err)
	}
	err = a.SM.SaveProject(project)
	if err != nil {
		return utils.Logger.Error("Failed to save project: %v", err)
	}
	return nil
}

func (a *App) UpdateBranch(projectName string, repoName string, branchName string, branch *md.Branch) error {
	project, err := a.GetProject(projectName)
	if err != nil {
		return utils.Logger.Error("Failed to get project: %v", err)
	}
	err = project.UpdateBranch(repoName, branchName, branch)
	if err != nil {
		return utils.Logger.Error("Failed to update branch in repo in project: %v", err)
	}
	err = a.SM.SaveProject(project)
	if err != nil {
		return utils.Logger.Error("Failed to save project: %v", err)
	}
	return nil
}

func (a *App) DeleteBranch(projectName string, repoName string, branchName string) error {
	project, err := a.GetProject(projectName)
	if err != nil {
		return utils.Logger.Error("Failed to get project: %v", err)
	}
	err = project.DeleteBranch(repoName, branchName)
	if err != nil {
		return utils.Logger.Error("Failed to delete branch from repo in project: %v", err)
	}
	err = a.SM.SaveProject(project)
	if err != nil {
		return utils.Logger.Error("Failed to save project: %v", err)
	}
	return nil
}
