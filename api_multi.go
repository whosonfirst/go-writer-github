package writer

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	wof_writer "github.com/whosonfirst/go-writer"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

const GITHUBAPI_MULTI_SCHEME string = "githubapi-multi"

// base_ is the thing a PR is being created "against"
// pr_ is the thing where the PR is being created

type GitHubAPIMultiWriter struct {
	wof_writer.Writer
	base_owner         string
	base_repo          string
	base_branch        string
	commit_owner       string
	commit_repo        string
	commit_branch      string
	commit_author      string
	commit_email       string
	commit_title       string
	commit_description string
	commit_entries     []github.TreeEntry
	commit_ensure_repo bool
	prefix             string
	client             *github.Client
	user               *github.User
	logger             *log.Logger
}

func init() {

	ctx := context.Background()
	wof_writer.RegisterWriter(ctx, GITHUBAPI_MULTI_SCHEME, NewGitHubAPIMultiWriter)
}

func NewGitHubAPIMultiWriter(ctx context.Context, uri string) (wof_writer.Writer, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	base_owner := u.Host

	path := strings.TrimLeft(u.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) != 1 {
		return nil, errors.New("Invalid path")
	}

	base_repo := parts[0]
	base_branch := DEFAULT_BRANCH

	q := u.Query()

	token := q.Get("access_token")

	prefix := q.Get("prefix")
	branch := q.Get("branch")

	if token == "" {
		return nil, errors.New("Missing access token")
	}

	if branch != "" {
		base_branch = branch
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	users := client.Users
	user, _, err := users.Get(ctx, "")

	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve user for token, %w", err)
	}

	commit_owner := base_owner
	commit_repo := base_repo

	commit_branch := q.Get("pr-branch")

	if commit_branch == "" {
		commit_branch = base_branch
	}

	commit_title := q.Get("commit-title")

	if commit_title == "" {
		return nil, fmt.Errorf("Invalid title")
	}

	commit_description := q.Get("commit-description")

	if commit_title == "" {
		return nil, fmt.Errorf("Invalid description")
	}

	commit_author := q.Get("commit-author")

	if commit_author == "" {
		commit_author = user.GetName()
	}

	if commit_author == "" {
		return nil, fmt.Errorf("Invalid author")
	}

	commit_email := q.Get("commit-email")

	if commit_email == "" {
		commit_email = user.GetEmail()
	}

	if commit_email == "" {
		return nil, fmt.Errorf("Invalid email address")
	}

	commit_entries := []github.TreeEntry{}

	logger := log.Default()

	wr := &GitHubAPIMultiWriter{
		client:             client,
		user:               user,
		base_owner:         base_owner,
		base_repo:          base_repo,
		base_branch:        base_branch,
		commit_owner:       commit_owner,
		commit_repo:        commit_repo,
		commit_branch:      commit_branch,
		commit_author:      commit_author,
		commit_email:       commit_email,
		commit_title:       commit_title,
		commit_description: commit_description,
		commit_entries:     commit_entries,
		prefix:             prefix,
		logger:             logger,
	}

	return wr, nil
}

func (wr *GitHubAPIMultiWriter) Write(ctx context.Context, uri string, r io.ReadSeeker) (int64, error) {

	// Something something something account for cases with a bazillion commits and not keeping
	// everything in memory until we call Close(). One option would be to keep a local map of io.ReadSeeker
	// instances but then we will just have filehandle exhaustion problems. Add option to write to
	// disk or something like a SQLite database (allowing a custom DSN to determine whether to write to
	// disk or memory) ?

	body, err := io.ReadAll(r)

	if err != nil {
		return 0, err
	}

	wr_uri := wr.WriterURI(ctx, uri)

	e := github.TreeEntry{
		Path:    github.String(wr_uri),
		Type:    github.String("blob"),
		Content: github.String(string(body)),
		Mode:    github.String("100644"),
	}

	wr.commit_entries = append(wr.commit_entries, e)

	return 0, nil
}

func (wr *GitHubAPIMultiWriter) Close(ctx context.Context) error {

	if len(wr.commit_entries) == 0 {
		return nil
	}

	ref, err := wr.getRef(ctx)

	if err != nil {

		if err != nil {
			return fmt.Errorf("Failed to get ref, %w", err)
		}
	}

	tree, _, err := wr.client.Git.CreateTree(ctx, wr.commit_owner, wr.commit_repo, *ref.Object.SHA, wr.commit_entries)

	if err != nil {
		return fmt.Errorf("Failed to create tree, %w", err)
	}

	err = wr.pushCommit(ctx, ref, tree)

	if err != nil {
		return fmt.Errorf("Failed to push commit, %w", err)
	}

	return nil
}

func (wr *GitHubAPIMultiWriter) WriterURI(ctx context.Context, key string) string {

	uri := key

	if wr.prefix != "" {
		uri = filepath.Join(wr.prefix, key)
	}

	return uri
}

func (wr *GitHubAPIMultiWriter) getRef(ctx context.Context) (*github.Reference, error) {

	base_branch := fmt.Sprintf("refs/heads/%s", wr.base_branch)
	commit_branch := fmt.Sprintf("refs/heads/%s", wr.commit_branch)

	commit_ref, _, _ := wr.client.Git.GetRef(ctx, wr.commit_owner, wr.commit_repo, commit_branch)

	if commit_ref != nil {
		return commit_ref, nil
	}

	base_ref, _, err := wr.client.Git.GetRef(ctx, wr.commit_owner, wr.commit_repo, base_branch)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve base branch '%s' for %s/%s, %w", base_branch, wr.commit_owner, wr.commit_repo, err)
	}

	new_ref := &github.Reference{Ref: github.String(commit_branch), Object: &github.GitObject{SHA: base_ref.Object.SHA}}

	commit_ref, _, err = wr.client.Git.CreateRef(ctx, wr.commit_owner, wr.commit_repo, new_ref)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new ref, %w")
	}

	return commit_ref, err
}

// pushCommit creates the commit in the given reference using the given tree.
func (wr *GitHubAPIMultiWriter) pushCommit(ctx context.Context, ref *github.Reference, tree *github.Tree) error {

	// Get the parent commit to attach the commit to.

	parent, _, err := wr.client.Repositories.GetCommit(ctx, wr.commit_owner, wr.commit_repo, *ref.Object.SHA)

	if err != nil {
		return fmt.Errorf("Failed to determine parent commit, %w", err)
	}

	// This is not always populated, but is needed.
	parent.Commit.SHA = parent.SHA

	// Create the commit using the tree.
	date := time.Now()

	author := &github.CommitAuthor{
		Date:  &date,
		Name:  &wr.commit_author,
		Email: &wr.commit_email,
	}

	parents := []github.Commit{
		*parent.Commit,
	}

	commit := &github.Commit{
		Author:  author,
		Message: &wr.commit_description,
		Tree:    tree,
		Parents: parents,
	}

	newCommit, _, err := wr.client.Git.CreateCommit(ctx, wr.commit_owner, wr.commit_repo, commit)

	if err != nil {
		return fmt.Errorf("Failed to create commit, %w", err)
	}

	// Attach the commit to the main branch.
	ref.Object.SHA = newCommit.SHA

	_, _, err = wr.client.Git.UpdateRef(ctx, wr.commit_owner, wr.commit_repo, ref, false)

	if err != nil {
		return fmt.Errorf("Failed to update ref, %w", err)
	}

	return nil
}
