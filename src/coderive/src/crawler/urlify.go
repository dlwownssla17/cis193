package crawler

import "fmt"

// GitHubUrlify builds the GitHub URL given the content-specific URL.
func GitHubUrlify(partURL string) string {
	if len(partURL) > 0 && partURL[0] == '/' {
		partURL = partURL[1:]
	}
	return fmt.Sprintf("%s/%s", GitHubURL, partURL)
}

// GitHubUrlifyWithParams builds the GitHub URL given the content-specific parameters.
func GitHubUrlifyWithParams(username, repositoryName, branchName, filePath string, isFile bool) string {
	contentType := "tree"
	if isFile {
		contentType = "blob"
	}
	return GitHubUrlify(fmt.Sprintf("%s/%s/%s/%s/%s", username, repositoryName, contentType, branchName, filePath))
}

// GitHubRepositoryUrlifyWithParams builds the GitHub URL for the repository given its parameters.
func GitHubRepositoryUrlifyWithParams(username, repositoryName string) string {
	return GitHubUrlify(fmt.Sprintf("%s/%s", username, repositoryName))
}

// GitHubRawContentUrlify builds the GitHub Raw Content URL given the content-specific URL. (assumes public repository)
func GitHubRawContentUrlify(partURL string) string {
	if len(partURL) > 0 && partURL[0] == '/' {
		partURL = partURL[1:]
	}
	return fmt.Sprintf("%s/%s", GitHubRawContentURL, partURL)
}

// GitHubRawContentUrlifyWithParams builds the GitHub Raw Content URL given the content-specific parameters. (assumes public repository)
func GitHubRawContentUrlifyWithParams(username, repositoryName, branchName, filePath string) string {
	return GitHubRawContentUrlify(fmt.Sprintf("%s/%s/%s/%s", username, repositoryName, branchName, filePath))
}