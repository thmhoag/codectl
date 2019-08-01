package repomanager

type Repository struct {
	// The repo name
	Name string

	// The URL of the repo
	URL string
}

func repoContainsName(repos []*Repository, name string) bool {
	for _, r := range repos {
		if r.Name == name {
			return true
		}
	}

	return false
}