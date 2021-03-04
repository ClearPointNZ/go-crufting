default: test

test:
	@go test ./... --cover

mocks:
	@mkdir -p internal/cruft/mock
	@counterfeiter -o internal/cruft/mock/crufter.go internal/cruft Crufter
