package models

type Project struct {
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Repos        map[string]Repo `json:"repos"`
	AWSAccount   string          `json:"awsAccount"`
	AWSLocations []AWSLocation   `json:"awsLocations"`
	LastDeploy   string          `json:"lastDeploy"`
	LastBuild    string          `json:"lastBuild"`
	LastTest     string          `json:"lastTest"`
	Archived     bool            `json:"archived"`
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
		return Logger.Error("Repo cannot be nil")
	}
	if repo.Name == "" {
		return Logger.Error("Repo name cannot be empty")
	}
	if p.Repos == nil {
		p.Repos = make(map[string]Repo)
	}
	if _, exists := p.Repos[repo.Name]; exists {
		return Logger.Error("Repo with name %s already exists", repo.Name)
	}
	newRepo := Repo{
		Name:        repo.Name,
		Description: repo.Description,
	}
	p.Repos[repo.Name] = newRepo
	return nil
}

func (p *Project) GetRepo(repoName string) (*Repo, error) {
	if repoName == "" {
		return nil, Logger.Error("Repo name cannot be empty")
	}
	repo, exists := p.Repos[repoName]
	if !exists {
		return nil, Logger.Error("Could not find repo %s", repoName)
	}
	return &repo, nil
}

func (p *Project) UpdateRepo(repoName string, repo *Repo) error {
	if repo == nil {
		return Logger.Error("Repo cannot be nil")
	}
	if repo.Name == "" {
		return Logger.Error("Repo name cannot be empty")
	}
	if _, exists := p.Repos[repoName]; !exists {
		return Logger.Error("Repo with name %s does not exist", repoName)
	}
	branches := p.Repos[repoName].Branches
	repo.Branches = branches
	p.Repos[repoName] = *repo
	return nil
}

func (p *Project) DeleteRepo(repoName string) error {
	if repoName == "" {
		return Logger.Error("Repo name cannot be empty")
	}
	if _, exists := p.Repos[repoName]; !exists {
		return Logger.Error("Repo with name %s does not exist", repoName)
	}
	delete(p.Repos, repoName)
	return nil
}

func (p *Project) NewBranch(repoName string, branch *Branch) error {
	if branch == nil {
		return Logger.Error("Branch cannot be nil")
	}
	if branch.Name == "" {
		return Logger.Error("Branch name cannot be empty")
	}
	if repoName == "" {
		return Logger.Error("Repo name cannot be empty")
	}
	repo, exists := p.Repos[repoName]
	if !exists {
		return Logger.Error("Repo with name %s does not exist", repoName)
	}
	return repo.NewBranch(branch)
}

func (p *Project) GetBranch(repoName string, branchName string) (*Branch, error) {
	if branchName == "" {
		return nil, Logger.Error("Branch name cannot be empty")
	}
	if repoName == "" {
		return nil, Logger.Error("Repo name cannot be empty")
	}
	repo, exists := p.Repos[repoName]
	if !exists {
		return nil, Logger.Error("Repo with name %s does not exist", repoName)
	}
	return repo.GetBranch(branchName)
}

func (p *Project) UpdateBranch(repoName string, branchName string, branch *Branch) error {
	if branch == nil {
		return Logger.Error("Branch cannot be nil")
	}
	if branch.Name == "" {
		return Logger.Error("Branch name cannot be empty")
	}
	if repoName == "" {
		return Logger.Error("Repo name cannot be empty")
	}
	repo, exists := p.Repos[repoName]
	if !exists {
		return Logger.Error("Repo with name %s does not exist", repoName)
	}
	return repo.UpdateBranch(branchName, branch)
}

func (p *Project) DeleteBranch(repoName string, branchName string) error {
	if repoName == "" {
		return Logger.Error("Repo name cannot be empty")
	}
	if branchName == "" {
		return Logger.Error("Branch name cannot be empty")
	}

	repo, exists := p.Repos[repoName]
	if !exists {
		return Logger.Error("Repo with name %s does not exist", repoName)
	}
	return repo.DeleteBranch(branchName)
}
