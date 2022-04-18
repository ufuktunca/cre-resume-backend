
generate-mocks:
	mockgen -destination=mocks/mock_user_service.go -source=./internal/user/service.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/user/service UserServiceInterface
	mockgen -destination=mocks/mock_user_repository.go -source=./internal/user/repository.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/user/repository UserRepositoryInterface
	mockgen -destination=mocks/mock_jobPost_service.go -source=./internal/job-post/service.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/job-post/service JobPostServiceInterface
	mockgen -destination=mocks/mock_jobPost_repository.go -source=./internal/job-post/repository.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/job-post/repository JobPostRepositoryInterface
	mockgen -destination=mocks/mock_email.go -source=./internal/email/email.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/email/email EmailInterface
																									
build: clean
	@echo "Building project"
	go build

	
run-local: print-banner build
	@echo "Running with local configuration"
	./favorite -port=8000

run-prod: print-banner build
	@echo "Running with Prod configuration"
	./favorite

test: print-banner test-unit
	$(call print-dash)

test-single: print-banner build
	@echo "Running Go single test"
	go test -v ./... -run ${TEST_NAME} -coverprofile=coverage.out -covermode=atomic && echo "$(boldgreen)Go Tests Passed!$(nc)" || (echo "$(boldred)Go Tests Failed$(nc)" && exit 1)
	$(call print-dash)

test-unit: print-banner build
	@echo "Running Go tests"
	go test -v ./... -coverprofile=coverage.out -covermode=atomic && echo "$(boldgreen)Go Tests Passed!$(nc)" || (echo "$(boldred)Go Tests Failed$(nc)" && exit 1)
	$(call print-dash)
	
test-cdc-provider: print-banner build
	@echo "Running Go Cdc Provider Tests"
	go test ./... -v -tags=cdc_test -run TestProvider && echo "$(boldgreen)Go Tests Passed!$(nc)" || (echo "$(boldred)Go Tests Failed$(nc)" && exit 1)
	$(call print-dash)

clean:
	go clean
	rm -rf coverage.out favorite profile.cov

bold=\033[0;1m
red=\033[0;31m
green=\033[0;32m
boldred=$(bold)$(red)
boldgreen=$(bold)$(green)
nc=\033[0m

print-banner:

define print-dash
	@echo "____________________________________________________"
endef
