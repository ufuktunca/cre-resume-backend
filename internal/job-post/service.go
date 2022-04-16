package jobPost

type JobPostService struct {
	Repository JobPostRepositoryInterface
}

type JobPostServiceInterface interface {
}

func NewJobPostService(repository JobPostRepositoryInterface) JobPostRepositoryInterface {
	return JobPostService{
		Repository: repository,
	}
}
