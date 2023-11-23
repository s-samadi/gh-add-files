package cmd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/cli/go-gh/v2/pkg/api"
)

func Test_getRepos(t *testing.T) {
	type args struct {
		client       *api.RESTClient
		Organization string
	}

	var emptyRepos []Repository

	tests := []struct {
		name    string
		args    args
		want    []Repository
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the organization is valid and has repos
		// 2. When the organization is valid and has no repos
		// 3. When the organization is invalid

		// Test case 1
		{
			name: "When the organization is valid and has repos",
			args: args{
				Organization: "paradisisland",
			},
			want: []Repository{
				{
					FullName:      "paradisisland/maria",
					Name:          "maria",
					DefaultBranch: "main",
				},
				{
					FullName:      "paradisisland/rose",
					Name:          "rose",
					DefaultBranch: "main",
				},
				{
					FullName:      "paradisisland/sheena",
					Name:          "sheena",
					DefaultBranch: "main",
				},
				{
					FullName:      "paradisisland/titanforest",
					Name:          "titanforest",
					DefaultBranch: "main",
				},
			},
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the organization is valid and has no repos",
			args: args{
				Organization: "sandora-desert",
			},
			want:    emptyRepos,
			wantErr: false,
		},

		// Test case 3
		{
			name: "When the organization is invalid",
			args: args{
				Organization: "atotallyrealorgname",
			},
			want:    emptyRepos,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRepos(tt.args.Organization)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRepo(t *testing.T) {
	type args struct {
		RepositoryName string
	}
	tests := []struct {
		name    string
		args    args
		want    Repository
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the repository is valid
		// 2. When the repository is invalid

		// Test case 1
		{
			name: "When the repository is valid",
			args: args{
				RepositoryName: "paradisisland/maria",
			},
			want: Repository{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the repository is invalid",
			args: args{
				RepositoryName: "paradisisland/marley",
			},
			want:    Repository{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRepo(tt.args.RepositoryName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_callApi(t *testing.T) {
	type args struct {
		requestPath string
		parseType   interface{}
		method      HttpMethod
		body        []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the request path is a valid repo path
		// 2. When the request path is an invalid repo path
		// 3. When the request path is a valid organization path
		// 4. When the request path is a valid organization path with a next page
		// 5. When the request path is an invalid organization path

		// Test case 1
		{
			name: "When the request path is valid",
			args: args{
				requestPath: "repos/paradisisland/maria",
				parseType:   Repository{},
				method:      GET,
			},
			want:    200,
			want1:   "",
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the request path is invalid",
			args: args{
				requestPath: "repos/paradisisland/marley",
				parseType:   Repository{},
				method:      GET,
			},
			want:    404,
			want1:   "",
			wantErr: true,
		},

		// Test case 3
		{
			name: "When the request path is a valid organization path",
			args: args{
				requestPath: "orgs/paradisisland/repos",
				parseType:   []Repository{},
				method:      GET,
			},
			want:    200,
			want1:   "",
			wantErr: false,
		},

		// Test case 4
		{
			name: "When the request path is a valid organization path with a next page",
			args: args{
				requestPath: "orgs/ansible/repos",
				parseType:   []Repository{},
				method:      GET,
			},
			want:    200,
			want1:   "<https://api.github.com/organizations/1507452/repos?page=2>; rel=\"next\", <https://api.github.com/organizations/1507452/repos?page=9>; rel=\"last\"",
			wantErr: false,
		},
		// Test case 5
		{
			name: "When the request path is an invalid organization path",
			args: args{
				requestPath: "orgs/atotallyrealorgname/repos",
				parseType:   []Repository{},
				method:      GET,
			},
			want:    404,
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := callApi(tt.args.requestPath, tt.args.parseType, tt.args.method, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("callApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("callApi() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("callApi() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_GetCodeqlLanguages(t *testing.T) {
	type fields struct {
		FullName      string
		Name          string
		DefaultBranch string
	}
	var codeqlLanguages []string

	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the repository has at least one codeql supported language
		// 2. When the repository has no codeql supported languages
		// 3. When the repository is invalid

		// Test case 1
		{
			name: "When the repository has at least one codeql supported language",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			want: []string{
				"Go",
				"Java",
				"JavaScript",
				"Python",
			},
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the repository has no codeql supported languages",
			fields: fields{
				FullName:      "paradisisland/titanforest",
				Name:          "titanforest",
				DefaultBranch: "main",
			},
			want:    codeqlLanguages,
			wantErr: false,
		},

		// Test case 3
		{
			name: "When the repository is invalid",
			fields: fields{
				FullName:      "paradisisland/marley",
				Name:          "marley",
				DefaultBranch: "main",
			},
			want:    codeqlLanguages,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				FullName:      tt.fields.FullName,
				Name:          tt.fields.Name,
				DefaultBranch: tt.fields.DefaultBranch,
			}
			got, err := repo.GetCodeqlLanguages()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetCodeqlLanguages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetCodeqlLanguages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNextPage(t *testing.T) {
	type args struct {
		nextPageLink string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When input contains a next page link
		// 2. When input does not contain a next page link

		// Test case 1
		{
			name: "When input contains a next page link",
			args: args{
				nextPageLink: "<https://api.github.com/organizations/1507452/repos?page=2>; rel=\"next\", <https://api.github.com/organizations/1507452/repos?page=9>; rel=\"last\"",
			},
			want:  "https://api.github.com/organizations/1507452/repos?page=2",
			want1: true,
		},

		// Test case 2
		{
			name: "When input does not contain a next page link",
			args: args{
				nextPageLink: "",
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findNextPage(tt.args.nextPageLink)
			if got != tt.want {
				t.Errorf("findNextPage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findNextPage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_createBranchForRepo(t *testing.T) {
	type fields struct {
		FullName      string
		Name          string
		DefaultBranch string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the repository does not have the branch codescanningworkflow
		// 2. When the repository has the branch codescanningworkflow
		// 3. When the repository is invalid

		// Test case 1
		{
			name: "When the repository does not have the branch codescanningworkflow",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			want:    "refs/heads/gh-cli/codescanningworkflow",
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the repository has the branch codescanningworkflow",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			want:    "",
			wantErr: true,
		},

		// Test case 3
		{
			name: "When the repository is invalid",
			fields: fields{
				FullName:      "paradisisland/marley",
				Name:          "marley",
				DefaultBranch: "main",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				FullName:      tt.fields.FullName,
				Name:          tt.fields.Name,
				DefaultBranch: tt.fields.DefaultBranch,
			}
			got, err := repo.createBranchForRepo()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.createBranchForRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.createBranchForRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_doesCodeqlWorkflowExist(t *testing.T) {
	type fields struct {
		FullName      string
		Name          string
		DefaultBranch string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the repository does not have the CodeQL workflow file
		// 2. When the repository has the CodeQL workflow file
		// 3. When the repository is invalid

		// Test case 1
		{
			name: "When the repository does not have the CodeQL workflow file",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			want:    false,
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the repository has the CodeQL workflow file",
			fields: fields{
				FullName:      "paradisisland/sheena",
				Name:          "sheena",
				DefaultBranch: "main",
			},
			want:    true,
			wantErr: false,
		},

		// Test case 3
		{
			name: "When the repository is invalid",
			fields: fields{
				FullName:      "paradisisland/marley",
				Name:          "marley",
				DefaultBranch: "main",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				FullName:      tt.fields.FullName,
				Name:          tt.fields.Name,
				DefaultBranch: tt.fields.DefaultBranch,
			}
			got, err := repo.doesCodeqlWorkflowExist()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.doesCodeqlWorkflowExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.doesCodeqlWorkflowExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_createWorkflowFile(t *testing.T) {
	type fields struct {
		FullName      string
		Name          string
		DefaultBranch string
	}
	type args struct {
		WorkflowFile string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the repository does not have the CodeQL workflow file
		// 2. When the repository has the CodeQL workflow file
		// 3. When the repository is invalid

		// Test case 1
		{
			name: "When the repository does not have the CodeQL workflow file",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			args: args{
				WorkflowFile: "../examples/codeql.yml",
			},
			want:    "codeql.yml",
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the repository has the CodeQL workflow file",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			args: args{
				WorkflowFile: "../examples/codeql.yml",
			},
			want:    "",
			wantErr: true,
		},

		// Test case 3
		{
			name: "When the repository is invalid",
			fields: fields{
				FullName:      "paradisisland/marley",
				Name:          "marley",
				DefaultBranch: "main",
			},
			args: args{
				WorkflowFile: "../examples/codeql.yml",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				FullName:      tt.fields.FullName,
				Name:          tt.fields.Name,
				DefaultBranch: tt.fields.DefaultBranch,
			}
			got, err := repo.createWorkflowFile(tt.args.WorkflowFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.createWorkflowFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.createWorkflowFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_raisePullRequest(t *testing.T) {
	type fields struct {
		FullName      string
		Name          string
		DefaultBranch string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the repository does not have a pull request
		// 2. When the repository has a pull request
		// 3. When the repository is invalid

		// Test case 1
		{
			name: "When the repository does not have a pull request",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			want:    "https://github.com/paradisisland/maria/pull/",
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the repository has a pull request",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			want:    "",
			wantErr: true,
		},

		// Test case 3
		{
			name: "When the repository is invalid",
			fields: fields{
				FullName:      "paradisisland/marley",
				Name:          "marley",
				DefaultBranch: "main",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				FullName:      tt.fields.FullName,
				Name:          tt.fields.Name,
				DefaultBranch: tt.fields.DefaultBranch,
			}
			got, err := repo.raisePullRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.raisePullRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && tt.want == "" {
				t.Errorf("Repository.raisePullRequest() = %v, want %v", got, tt.want)
			} else if !strings.Contains(got, tt.want) && tt.want != "" {
				t.Errorf("Repository.raisePullRequest() PR LINK = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_deleteBranch(t *testing.T) {
	type fields struct {
		FullName      string
		Name          string
		DefaultBranch string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		// Write test cases for the following scenarios:
		// 1. When the repository has the codescanningworkflow branch
		// 2. When the repository does not have the codescanningworkflow branch
		// 3. When the repository is invalid

		// Test case 1
		{
			name: "When the repository has the codescanningworkflow branch",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			wantErr: false,
		},

		// Test case 2
		{
			name: "When the repository does not have the codescanningworkflow branch",
			fields: fields{
				FullName:      "paradisisland/maria",
				Name:          "maria",
				DefaultBranch: "main",
			},
			wantErr: true,
		},

		// Test case 3
		{
			name: "When the repository is invalid",
			fields: fields{
				FullName:      "paradisisland/marley",
				Name:          "marley",
				DefaultBranch: "main",
			},
			wantErr: true,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				FullName:      tt.fields.FullName,
				Name:          tt.fields.Name,
				DefaultBranch: tt.fields.DefaultBranch,
			}
			if err := repo.deleteBranch(); (err != nil) != tt.wantErr {
				t.Errorf("Repository.deleteBranch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
