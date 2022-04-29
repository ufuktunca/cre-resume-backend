package jobPost

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"cre-resume-backend/internal/user"
	"time"
)

type JobPostView struct {
	Model     JobPostModelInterface
	UserModel user.UserModelInterface
}

type JobPostViewInterface interface {
	CreateJobPost(jobPost *models.JobPost, ownerEmail string) (*models.JobPost, error)
	GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error)
	ApplyJobPost(jobPostDTO *models.ApplyJobPostDTO, applierEmail, jobID string) error
	GetUserJobPosts(userEmail string, postType string) (*[]models.JobPost, error)
}

func NewJobPostView(model JobPostModelInterface, userModel user.UserModelInterface) *JobPostView {
	return &JobPostView{
		Model:     model,
		UserModel: userModel,
	}
}

func (s *JobPostView) CreateJobPost(jobPost *models.JobPost, ownerEmail string) (*models.JobPost, error) {
	jobPost.ID = helpers.GenerateUUID(8)

	owner, err := s.UserModel.GetUserByEmail(ownerEmail)
	if err != nil {
		return nil, err
	}

	jobPost.OwnerID = owner.UserID
	jobPost.CreatedAt = time.Now().UTC().Round(time.Second)
	jobPost.UpdatedAt = time.Now().UTC().Round(time.Second)

	err = s.Model.CreateJobPost(jobPost)
	if err != nil {
		return nil, err
	}

	return jobPost, nil
}

func (s *JobPostView) GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error) {
	return s.Model.GetJobPosts(jobPostType, category, from, to, sort)
}

func (s *JobPostView) ApplyJobPost(jobPostDTO *models.ApplyJobPostDTO, applierEmail, jobID string) error {
	applierUser, err := s.UserModel.GetUserByEmail(applierEmail)
	if err != nil {
		return err
	}

	jobPost, err := s.Model.GetJobPostByID(jobID)
	if err != nil {
		return err
	}

	applyJobPost := &models.ApplyJobPost{
		ID:          helpers.GenerateUUID(8),
		JobPostID:   jobPost.ID,
		CVID:        jobPostDTO.CVID,
		ApplierID:   applierUser.UserID,
		PostOwnerID: jobPost.OwnerID,
	}

	return s.Model.CreateApplyJobPost(applyJobPost)
}

func (s *JobPostView) GetUserJobPosts(userEmail string, postType string) (*[]models.JobPost, error) {
	user, err := s.UserModel.GetUserByEmail(userEmail)
	if err != nil {
		return nil, err
	}

	return s.Model.GetJobPostsWithUserID(user.UserID, postType)
}
