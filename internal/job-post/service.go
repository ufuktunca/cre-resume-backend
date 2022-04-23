package jobPost

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"time"
)

type JobPostService struct {
	Repository JobPostRepositoryInterface
}

type JobPostServiceInterface interface {
	CreateJobPost(jobPost *models.JobPost) (*models.JobPost, error)
	GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error)
}

func NewJobPostService(repository JobPostRepositoryInterface) *JobPostService {
	return &JobPostService{
		Repository: repository,
	}
}

func (s *JobPostService) CreateJobPost(jobPost *models.JobPost) (*models.JobPost, error) {
	jobPost.ID = helpers.GenerateUUID(8)

	jobPost.CreatedAt = time.Now().UTC().Round(time.Second)
	jobPost.UpdatedAt = time.Now().UTC().Round(time.Second)

	err := s.Repository.CreateJobPost(jobPost)
	if err != nil {
		return nil, err
	}

	return jobPost, nil
}

func (s *JobPostService) GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error) {
	return s.Repository.GetJobPosts(jobPostType, category, from, to, sort)
}
