package models

import (
	"daws/internal/utils"
)

type Project struct {
	Name           string          `json:"name"`
	Slug           string          `json:"slug"`
	Description    string          `json:"description"`
	Repos          map[string]Repo `json:"repos"`
	AWSCredentials AWSCredentials  `json:"awsCredentials"`
	Archived       bool            `json:"archived"`
	TimeData       TimeData        `json:"timeData"`
}

func (p *Project) Deploy() error {
	var err error
	return err
}

func (p *Project) RunTests() error {
	var err error
	return err
}

func (p *Project) Save() error {
	var err error
	// write to file
	return err
}

func (p *Project) Load() (*Project, error) {
	var err error
	return p, err
}

func (p *Project) NewRepo(repo *Repo) error {
	if repo == nil {
		return utils.Logger.Error("Repo cannot be nil")
	}
	if repo.Name == "" {
		return utils.Logger.Error("Repo name cannot be empty")
	}
	if p.Repos == nil {
		p.Repos = make(map[string]Repo)
	}
	repo.Slug = utils.Slugify(repo.Name)
	if _, exists := p.Repos[repo.Slug]; exists {
		return utils.Logger.Error("Repo with name %s already exists", repo.Slug)
	}

	repo.TimeData.Create()
	p.Repos[repo.Slug] = *repo
	return nil
}

func (p *Project) GetRepo(repoName string) (*Repo, error) {
	if repoName == "" {
		return nil, utils.Logger.Error("Repo name cannot be empty")
	}
	repo, exists := p.Repos[repoName]
	if !exists {
		return nil, utils.Logger.Error("Could not find repo %s", repoName)
	}
	return &repo, nil
}

func (p *Project) UpdateRepo(repoName string, repo *Repo) error {
	if repo == nil {
		return utils.Logger.Error("Repo cannot be nil")
	}
	if repo.Name == "" {
		return utils.Logger.Error("Repo name cannot be empty")
	}
	if _, exists := p.Repos[repoName]; !exists {
		return utils.Logger.Error("Repo with name %s does not exist", repoName)
	}
	branches := p.Repos[repoName].Branches
	repo.Branches = branches
	repo.Slug = utils.Slugify(repo.Name)
	repo.TimeData.Update()
	p.Repos[repo.Slug] = *repo
	return nil
}

func (p *Project) DeleteRepo(repoName string) error {
	if repoName == "" {
		return utils.Logger.Error("Repo name cannot be empty")
	}
	if _, exists := p.Repos[repoName]; !exists {
		return utils.Logger.Error("Repo with name %s does not exist", repoName)
	}
	delete(p.Repos, repoName)
	return nil
}

func (p *Project) NewBranch(repoName string, branch *Branch) error {
	if branch == nil {
		return utils.Logger.Error("Branch cannot be nil")
	}
	if branch.Name == "" {
		return utils.Logger.Error("Branch name cannot be empty")
	}
	if repoName == "" {
		return utils.Logger.Error("Repo name cannot be empty")
	}
	repo, exists := p.Repos[repoName]
	if !exists {
		return utils.Logger.Error("Repo with name %s does not exist", repoName)
	}
	return repo.NewBranch(branch)
}

func (p *Project) GetBranch(repoName string, branchName string) (*Branch, error) {
	if branchName == "" {
		return nil, utils.Logger.Error("Branch name cannot be empty")
	}
	if repoName == "" {
		return nil, utils.Logger.Error("Repo name cannot be empty")
	}
	repo, exists := p.Repos[repoName]
	if !exists {
		return nil, utils.Logger.Error("Repo with name %s does not exist", repoName)
	}
	return repo.GetBranch(branchName)
}

func (p *Project) UpdateBranch(repoName string, branchName string, branch *Branch) error {
	if branch == nil {
		return utils.Logger.Error("Branch cannot be nil")
	}
	if branch.Name == "" {
		return utils.Logger.Error("Branch name cannot be empty")
	}
	if repoName == "" {
		return utils.Logger.Error("Repo name cannot be empty")
	}
	repo, exists := p.Repos[repoName]
	if !exists {
		return utils.Logger.Error("Repo with name %s does not exist", repoName)
	}
	return repo.UpdateBranch(branchName, branch)
}

func (p *Project) DeleteBranch(repoName string, branchName string) error {
	if repoName == "" {
		return utils.Logger.Error("Repo name cannot be empty")
	}
	if branchName == "" {
		return utils.Logger.Error("Branch name cannot be empty")
	}

	repo, exists := p.Repos[repoName]
	if !exists {
		return utils.Logger.Error("Repo with name %s does not exist", repoName)
	}
	return repo.DeleteBranch(branchName)
}

func (p *Project) Save_Preprocess() error {
	var err error
	p.AWSCredentials.EncryptAll()
	for repoName, repo := range p.Repos {
		err = repo.Save_Preprocess()
		if err != nil {
			return utils.Logger.Error("Failed to preprocess repo %s: %v", repoName, err)
		}
	}

	p.Slug = utils.Slugify(p.Name)
	p.TimeData.Create()
	return err
}
