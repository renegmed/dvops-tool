package pork

import (
	"fmt"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"

	git "gopkg.in/src-d/go-git.v4"
)

type GHRepo struct {
	RepoDir string
	owner   string
	project string
	repo    *git.Repository
}

func NewGHRepo(respository string) (*GHRepo, error) {
	values := strings.Split(respository, "/")
	if len(values) != 2 {
		return nil, fmt.Errorf("repository must be in format owner/project")
	}
	return &GHRepo{
		owner:   values[0],
		project: values[1],
	}, nil
}

func (g *GHRepo) RepositoryURL() string {
	return fmt.Sprintf("https://github.com/%s/%s.git", g.owner, g.project)
}

// Clone the GitHub repository to local machine
func (g *GHRepo) Clone(dest string) error {
	fullPath := filepath.Join(dest, fmt.Sprintf("%s-%s", g.owner, g.project))
	repo, err := git.PlainClone(fullPath, false, &git.CloneOptions{ // false - normal
		URL: g.RepositoryURL(),
	})
	if err != nil {
		return err
	}
	g.repo = repo
	g.RepoDir = fullPath
	return nil
}

func (g *GHRepo) Checkout(ref string, create bool) error {
	opts := &git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", ref)),
		Create: create,
	}
	if create {
		head, err := g.repo.Head()
		if err != nil {
			return err
		}
		opts.Hash = head.Hash()
	}
	tree, err := g.repo.Worktree()
	if err != nil {
		return err
	}
	return tree.Checkout(opts)
}

func (g *GHRepo) AddUpstream(repository *GHRepo) error {
	_, err := g.repo.CreateRemote(&config.RemoteConfig{
		URLs: []string{repository.RepositoryURL()},
	})
	return err
}
