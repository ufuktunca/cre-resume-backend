package jobPost

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"cre-resume-backend/internal/user"
	"time"
)

type JobPostService struct {
	Repository     JobPostRepositoryInterface
	UserRepository user.UserRepositoryInterface
}

type JobPostServiceInterface interface {
	CreateJobPost(jobPost *models.JobPost, ownerEmail string) (*models.JobPost, error)
	GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error)
}

func NewJobPostService(repository JobPostRepositoryInterface, userRepository user.UserRepositoryInterface) *JobPostService {
	return &JobPostService{
		Repository:     repository,
		UserRepository: userRepository,
	}
}

func (s *JobPostService) CreateJobPost(jobPost *models.JobPost, ownerEmail string) (*models.JobPost, error) {
	jobPost.ID = helpers.GenerateUUID(8)

	owner, err := s.UserRepository.GetUserByEmail(ownerEmail)
	if err != nil {
		return nil, err
	}

	jobPost.OwnerID = owner.UserID
	jobPost.CreatedAt = time.Now().UTC().Round(time.Second)
	jobPost.UpdatedAt = time.Now().UTC().Round(time.Second)

	err = s.Repository.CreateJobPost(jobPost)
	if err != nil {
		return nil, err
	}

	return jobPost, nil
}

func (s *JobPostService) GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error) {
	return s.Repository.GetJobPosts(jobPostType, category, from, to, sort)
}
