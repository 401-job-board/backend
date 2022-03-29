package service

import (
	"context"

	"github.com/fanfit/feed/models/posts/repository"
)

// Service receives commands from handlers and forwards them to the repository
type Service interface {
	// CREATES

	GetPosts(context.Context) ([]repository.Posting, error)
	CreatePosting(context.Context, repository.Posting) (repository.Posting, error)
	GetJobs(context.Context, repository.GetJobsParams) ([]repository.Posting, error)
}

type service struct {
	repository repository.Repository
}

// New creates a service instance with the repository passed
func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (service *service) CreatePosting(ctx context.Context, inputPosting repository.Posting) (repository.Posting, error) {
	return service.repository.CreatePosting(ctx, inputPosting)
}

func (service *service) GetPosts(ctx context.Context) ([]repository.Posting, error) {
	return service.repository.GetAllPosts(ctx)
}

func (service *service) GetJobs(ctx context.Context, jobData repository.GetJobsParams) ([]repository.Posting, error) {
	return service.repository.GetJobs(ctx, jobData)
}
