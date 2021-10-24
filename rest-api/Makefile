mock:
	mockery --dir service --all --case underscore
	mockery --dir repository --all --case underscore
	mockery --dir adapter/push --all --case underscore
test:
	go clean -testcache
	go test ./...
