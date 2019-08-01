package repomanager

import (
	"fmt"
	goget "github.com/hashicorp/go-getter"
	"github.com/pkg/errors"
	"os"
)

type Manager interface {
	GetRepos() []*Repository
	Exists(repo *Repository) bool
	AddRepo(repo *Repository) error
	RemoveRepo(repo *Repository) error
	UpdateAll() error
}

type manager struct {
	opts *ManagerOpts
}

// Creates a new repository Manager
func NewManager(opts *ManagerOpts) Manager {
	if opts == nil {
		panic("manager opts cannot be nil")
	}

	if err := opts.Validate(); err != nil {
		panic("manager opts were not valid")
	}

	return &manager{
		opts: opts,
	}
}

func (m manager) GetRepos() []*Repository {

	cfg := m.opts.Config

	var result []*Repository
	if err := cfg.UnmarshalKey(m.opts.PropName, &result); err != nil {
		panic(err)
	}

	return result
}

func (m manager) Exists(repo *Repository) bool {

	repos := m.GetRepos()
	return repoContainsName(repos, repo.Name)
}

func (m manager) AddRepo(repo *Repository) error {

	cfg := m.opts.Config

	if exists := m.Exists(repo); exists {
		return errors.New("repo already exists")
	}

	repos := m.GetRepos()
	repos = append(repos, repo)

	cfg.Set(m.opts.PropName, &repos)
	if err := cfg.WriteConfig(); err != nil {
		return err
	}

	return nil
}

func (m manager) RemoveRepo(repo *Repository) error {

	cfg := m.opts.Config

	if exists := m.Exists(repo); !exists {
		return errors.New("repo does not exist")
	}

	cacheDir := fmt.Sprintf("%s/%s", m.opts.CacheDir, repo.Name)
	if err := os.RemoveAll(cacheDir); err != nil {
		return err
	}

	repos := m.GetRepos()
	var result []*Repository
	for i,r := range repos {

		if r.Name == repo.Name {
			continue
		}

		result = append(result, repos[i])
	}

	cfg.Set(m.opts.PropName, &result)
	if err := cfg.WriteConfig(); err != nil {
		return err
	}

	return nil
}

func (m manager) UpdateAll() error {
	repos := m.GetRepos()

	for _, r := range repos {
		dst := fmt.Sprintf("%s/%s", m.opts.CacheDir, r.Name)
		url := fmt.Sprintf("git::%s", r.URL)

		if err := goget.Get(dst, url); err != nil {
			return err
		}
	}

	return nil
}
