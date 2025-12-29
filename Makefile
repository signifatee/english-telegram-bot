.PHONY: buildbot
buildbot:
	go build -o ./.bin/bot cmd/bot/main.go

.PHONY:
runbot: buildbot
	./.bin/bot

.PHONY: buildapi
buildapi:
	go build -o ./.bin/api cmd/api/main.go

.PHONY:
runapi: buildapi
	./.bin/api

.PHONY:
migrate:

.DEFAULT_GOAL := build

