package jobPost

import (
	"cre-resume-backend/internal/cv"
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"cre-resume-backend/internal/user"
	"errors"
	"time"
)

type JobPostView struct {
	Model     JobPostModelInterface
	UserModel user.UserModelInterface
	CVModel   cv.CVModelInterface
}

type JobPostViewInterface interface {
	CreateJobPost(jobPost *models.JobPost, userID string) (*models.JobPost, error)
	GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error)
	ApplyJobPost(jobPostDTO *models.ApplyJobPostDTO, userID, jobID string) error
	GetUserJobPosts(userEmail string, postType string) (*[]models.JobPost, error)
	GetUserAppliedJobs(userId string) (*[]models.JobPost, error)
	GetJobApplies(jobId string) ([]models.CV, error)
	DeleteJobPost(jobId string) error
}

func NewJobPostView(model JobPostModelInterface, userModel user.UserModelInterface, cvModel cv.CVModelInterface) *JobPostView {
	return &JobPostView{
		Model:     model,
		UserModel: userModel,
		CVModel:   cvModel,
	}
}

func (s *JobPostView) CreateJobPost(jobPost *models.JobPost, userID string) (*models.JobPost, error) {
	jobPost.ID = helpers.GenerateUUID(8)

	jobPost.OwnerID = userID
	jobPost.CreatedAt = time.Now().UTC().Round(time.Second)
	jobPost.UpdatedAt = time.Now().UTC().Round(time.Second)

	err := s.Model.CreateJobPost(jobPost)
	if err != nil {
		return nil, err
	}

	return jobPost, nil
}

func (s *JobPostView) GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error) {
	return s.Model.GetJobPosts(jobPostType, category, from, to, sort)
}

func (s *JobPostView) ApplyJobPost(jobPostDTO *models.ApplyJobPostDTO, userID, jobID string) error {
	_, err := s.Model.GetJobApplyWithUserIDAndJobID(userID, jobID)
	if err == nil {
		return errors.New("you cannot apply to same job")
	}

	jobPost, err := s.Model.GetJobPostByID(jobID)
	if err != nil {
		return err
	}

	applyJobPost := &models.ApplyJobPost{
		ID:          helpers.GenerateUUID(8),
		JobPostID:   jobPost.ID,
		CVID:        jobPostDTO.CVID,
		ApplierID:   userID,
		PostOwnerID: jobPost.OwnerID,
	}

	return s.Model.CreateApplyJobPost(applyJobPost)
}

func (s *JobPostView) GetUserJobPosts(userID string, postType string) (*[]models.JobPost, error) {
	_, err := s.UserModel.GetUserByEmail(userID)
	if err != nil {
		return nil, err
	}

	return s.Model.GetJobPostsWithUserID(userID, postType)
}

func (s *JobPostView) GetUserAppliedJobs(userId string) (*[]models.JobPost, error) {

	applies, err := s.Model.GetUserApplies(userId)
	if err != nil {
		return nil, err
	}

	jobs := []models.JobPost{}
	for _, apply := range applies {
		job, err := s.Model.GetJobPostByID(apply.JobPostID)
		if err != nil {
			return nil, err
		}
		job.CVID = apply.CVID
		jobs = append(jobs, *job)
	}

	return &jobs, nil
}

func (s *JobPostView) GetJobApplies(jobId string) ([]models.CV, error) {

	applies, err := s.Model.GetJobApplies(jobId)
	if err != nil {
		return nil, err
	}

	cvs := []models.CV{}
	for _, apply := range *applies {
		cv, err := s.CVModel.GetCV(apply.CVID)
		if err != nil {
			return nil, err
		}
		cvs = append(cvs, *cv)
	}

	return cvs, nil
}

func (s *JobPostView) DeleteJobPost(jobId string) error {
	return s.Model.DeleteJobPost(jobId)
}
