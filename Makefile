SERVICE_NAME=messaging
PROJECTS=$(SERVICE_NAME)/...

all:
	@cd src/ && go install $(PROJECTS)

test:
	@cd src/ && go test -coverprofile=coverage.out $(PROJECTS)
	@go tool cover -html=src/coverage.out

# generate XML report in Cobertura format
test.xml:
	@cd src/ && gocov test $(PROJECTS) | gocov-xml > coverage.xml

clean:
	@rm -rf ./bin/$(SERVICE_NAME)

clean-all:

run:
	@./bin/messaging

stop-all:
	@pkill -f "./bin/messaging" &


run-messaging-notifications:
	@./bin/messaging
stop-messaging-notifications:
	@pkill -f "./bin/messaging"