vet:
	go vet ./...
test: vet
	ginkgo run -p -vv --race --trace --cover ./optional/
