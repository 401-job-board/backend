package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fanfit/feed/database"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	CreatePosting(context.Context, Posting) (Posting, error)
	GetAllPosts(context.Context) ([]Posting, error)
	// Other
	GetJobs(context.Context, GetJobsParams) ([]Posting, error)
	Close() error
}

type repository struct {
	queries *Queries
	db      *sql.DB
}

func (repo *repository) GetJobs(ctx context.Context, post GetJobsParams) ([]Posting, error) {
	jobsList, err := repo.queries.GetJobs(ctx, GetJobsParams{
		Title:       post.Title,
		CompanyName: post.CompanyName,
		Salary:      post.Salary,
	})
	if err != nil {
		fmt.Print(err)
	}
	return jobsList, err
}

func (repo *repository) CreatePosting(ctx context.Context, post Posting) (Posting, error) {
	newPosting, err := repo.queries.CreateNewPosting(ctx, CreateNewPostingParams{
		Title:              post.Title,
		CompanyName:        post.CompanyName,
		CompanyDescription: post.CompanyDescription,
		PostingDescription: post.PostingDescription,
		Salary:             post.Salary,
	})
	if err != nil {
		fmt.Print(err)
	}
	return newPosting, err
}

// Gets All workouts
func (repo *repository) GetAllPosts(ctx context.Context) ([]Posting, error) {
	response, err := repo.queries.GetAllPostings(ctx)
	if err != nil {
		fmt.Print(err)
	}

	return response, err
}

// Deletes a workout according to its ID

func (repo *repository) Close() error {
	return repo.db.Close()
}

func NewWorkoutStore(dbURL string) (Repository, error) {
	db, err := database.EstablishConnection(dbURL)
	if err != nil {
		fmt.Println("Error while establishing connection with databse " + err.Error())
	}

	return &repository{
		db:      db,
		queries: New(db),
	}, nil
}
