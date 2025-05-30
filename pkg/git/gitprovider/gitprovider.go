/*
Copyright 2023-2025 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gitprovider

import "time"

type GitProviderClient interface {
	// EnsurePaCMergeRequest creates or updates existing (if needed) Pipelines as Code configuration proposal merge request.
	// Returns the merge request web URL.
	// If there is no error and web URL is empty, it means that the merge request is not needed (main branch is up to date).
	EnsurePaCMergeRequest(repoUrl string, data *MergeRequestData) (webUrl string, err error)

	// UndoPaCMergeRequest creates or updates existing Pipelines as Code configuration removal merge request.
	// Returns the merge request web URL.
	// If there is no error and web URL is empty, it means that the merge request is not needed (the configuraton has already been deleted).
	UndoPaCMergeRequest(repoUrl string, data *MergeRequestData) (webUrl string, err error)

	// FindUnmergedPaCMergeRequest searches for existing Pipelines as Code configuration proposal merge request
	FindUnmergedPaCMergeRequest(repoUrl string, data *MergeRequestData) (*MergeRequest, error)

	// SetupPaCWebhook creates Pipelines as Code webhook in the given repository
	SetupPaCWebhook(repoUrl, webhookUrl, webhookSecret string) error

	// DeletePaCWebhook deletes Pipelines as Code webhook in the given repository
	DeletePaCWebhook(repoUrl, webhookUrl string) error

	// GetDefaultBranchWithChecks returns name of default branch in the given repository
	// also performs additional checks if repo exists and hostname is valid and returns more specific error
	GetDefaultBranchWithChecks(repoUrl string) (string, error)

	// DeleteBranch deletes given branch from repository.
	// Returns true if branch was deleted, false if the branch didn't exist.
	DeleteBranch(repoUrl, branchName string) (bool, error)

	// GetBranchSha returns SHA of top commit in the given branch
	GetBranchSha(repoUrl, branchName string) (string, error)

	// DownloadFileContent retrieves file content by full path in the git repository.
	DownloadFileContent(repoUrl, branchName, filePath string) ([]byte, error)

	// IsFileExist check whether given file exists in the given branch of the reposiotry
	IsFileExist(repoUrl, branchName, filePath string) (bool, error)

	// IsRepositoryPublic returns true if the repository could be accessed without authentication
	IsRepositoryPublic(repoUrl string) (bool, error)

	// GetBrowseRepositoryAtShaLink returns web URL of repository state at given SHA
	GetBrowseRepositoryAtShaLink(repoUrl, sha string) string

	// GetConfiguredGitAppName returns configured git application name and id.
	// Not all git providers support applications. Currently only GitHub does.
	GetConfiguredGitAppName() (string, string, error)

	// GetAppUserId get info about application user
	GetAppUserId(userName string) (int64, error)
}

type MergeRequestData struct {
	CommitMessage  string
	SignedOff      bool
	BranchName     string
	BaseBranchName string
	Title          string
	Text           string
	AuthorName     string
	AuthorEmail    string
	Files          []RepositoryFile
}

type RepositoryFile struct {
	FullPath string
	Content  []byte
}

type MergeRequest struct {
	Id        int64
	CreatedAt *time.Time
	WebUrl    string
	Title     string
}
