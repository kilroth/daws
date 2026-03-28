package models

type Repo struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Path            string            `json:"path"`
	Branches        map[string]Branch `json:"branches"`
	DefaultBranch   string            `json:"defaultBranch"`
	AWSAccount      string            `json:"awsAccount"`
	AWSLocations    []AWSLocation     `json:"awsLocations"`
	DockerImageName string            `json:"dockerImageName"`
	TestScripts     []string          `json:"testScripts"`
	LastBuild       string            `json:"lastBuild"`
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
		return Logger.Error("Branch cannot be nil")
	}
	if branch.Name == "" {
		return Logger.Error("Branch name cannot be empty")
	}
	if _, exists := r.Branches[branch.Name]; exists {
		return Logger.Error("Branch with name %s already exists", branch.Name)
	}
	if r.Branches == nil {
		r.Branches = make(map[string]Branch)
	}

	r.Branches[branch.Name] = *branch
	return nil
}

func (r *Repo) GetBranch(branchName string) (*Branch, error) {
	if branchName == "" {
		return nil, Logger.Error("Branch name cannot be empty")
	}
	branch, exists := r.Branches[branchName]
	if !exists {
		return nil, Logger.Error("Could not find branch %s", branchName)
	}
	return &branch, nil
}

func (r *Repo) UpdateBranch(branchName string, branch *Branch) error {
	if branch == nil {
		return Logger.Error("Branch cannot be nil")
	}
	if _, exists := r.Branches[branchName]; !exists {
		return Logger.Error("Branch with name %s does not exist", branchName)
	}
	r.Branches[branchName] = *branch
	return nil
}

func (r *Repo) DeleteBranch(branchName string) error {
	if _, exists := r.Branches[branchName]; !exists {
		return Logger.Error("Branch with name %s does not exist", branchName)
	}
	delete(r.Branches, branchName)
	return nil
}
