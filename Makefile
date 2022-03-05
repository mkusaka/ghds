.PHONY: setup
setup:
	go install github.com/goreleaser/goreleaser@latest

TAG=v0.0.1
MESSAGE=release

.PHONY: tag
tag:
	git tag -a $(TAG) -m "$(MESSAGE)"
	git push origin $(TAG)

.PHONY: release
release:
	goreleaser release
