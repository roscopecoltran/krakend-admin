package searx

import (
	"errors"
	"fmt"
)

func ParseGithubRawFile(url string, provider string, patterns map[string]string) (map[string]string, error) {
	keys := make(map[string]string)
	if provider == "" {
		return keys, errors.New("No provider name provided.")
	}

	if url == "" {
		url = fmt.Sprintf("https://raw.githubusercontent.com/ascimoo/searx/blob/%s/searx/engines/%s.py", provider, provider)
	}
	// src_remote_url: https://raw.githubusercontent.com/jibe-b/searx/blob/arxiv/searx/engines/arxiv.py
	return keys, nil
}
