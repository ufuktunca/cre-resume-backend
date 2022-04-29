
generate-mocks:
	mockgen -destination=mocks/mock_user_view.go -source=./internal/user/view.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/user/view UserViewInterface
	mockgen -destination=mocks/mock_user_model.go -source=./internal/user/model.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/user/model UserModelInterface
	mockgen -destination=mocks/mock_jobPost_view.go -source=./internal/job-post/view.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/job-post/view JobPostViewInterface
	mockgen -destination=mocks/mock_jobPost_model.go -source=./internal/job-post/model.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/job-post/model JobPostModelInterface
	mockgen -destination=mocks/mock_cv_view.go -source=./internal/cv/view.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/cv/view CVViewInterface
	mockgen -destination=mocks/mock_cv_model.go -source=./internal/cv/model.go -package mocks github.com/ufuktunca/cre-resume-frontend/internal/cv/model CVModelInterface
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
