package models

import (
	"daws/internal/utils"
)

type Repo struct {
	Name            string            `json:"name"`
	Slug            string            `json:"slug"`
	Description     string            `json:"description"`
	Path            string            `json:"path"`
	Branches        map[string]Branch `json:"branches"`
	DefaultBranch   string            `json:"defaultBranch"`
	AWSCredentials  AWSCredentials    `json:"awsCredentials"`
	AWSLocations    []AWSLocation     `json:"awsLocations"`
	DockerImageName string            `json:"dockerImageName"`
	TestScripts     []string          `json:"testScripts"`
	TimeData        TimeData          `json:"timeData"`
	Archived        bool              `json:"archived"`
}

func (r *Repo) Deploy() error {
	var err error
	return err
}

func (r *Repo) RunTests() error {
	var err error
	return err
}

func (r *Repo) NewBranch(branch *Branch) error {
	if branch == nil {
		return utils.Logger.Error("Branch cannot be nil")
	}
	if branch.Name == "" {
		return utils.Logger.Error("Branch name cannot be empty")
	}
	branch.Slug = utils.Slugify(branch.Name)
	if _, exists := r.Branches[branch.Slug]; exists {
		return utils.Logger.Error("Branch with name %s already exists", branch.Slug)
	}
	if r.Branches == nil {
		r.Branches = make(map[string]Branch)
	}
	branch.TimeData.Create()
	r.Branches[branch.Slug] = *branch
	return nil
}

func (r *Repo) GetBranch(branchName string) (*Branch, error) {
	if branchName == "" {
		return nil, utils.Logger.Error("Branch name cannot be empty")
	}
	branch, exists := r.Branches[branchName]
	if !exists {
		return nil, utils.Logger.Error("Could not find branch %s", branchName)
	}
	return &branch, nil
}

func (r *Repo) UpdateBranch(branchName string, branch *Branch) error {
	if branch == nil {
		return utils.Logger.Error("Branch cannot be nil")
	}
	if _, exists := r.Branches[branchName]; !exists {
		return utils.Logger.Error("Branch with name %s does not exist", branchName)
	}
	branch.TimeData.Update()
	branch.Slug = utils.Slugify(branch.Name)
	r.Branches[branchName] = *branch
	return nil
}

func (r *Repo) DeleteBranch(branchName string) error {
	if _, exists := r.Branches[branchName]; !exists {
		return utils.Logger.Error("Branch with name %s does not exist", branchName)
	}
	delete(r.Branches, branchName)
	return nil
}

func (r *Repo) Save_Preprocess() error {
	var err error
	err = r.AWSCredentials.EncryptAll()
	for branchName, branch := range r.Branches {
		err = branch.Save_Preprocess()
		if err != nil {
			return utils.Logger.Error("Failed to preprocess branch %s: %v", branchName, err)
		}
	}
	r.Slug = utils.Slugify(r.Name)
	r.TimeData.Create()
	return err
}
