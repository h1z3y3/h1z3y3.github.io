.PHONY: release push

release:
	docker build -f Dockerfile -t h1z3y3/telegram-timeline-bot:latest .

push:
	docker push h1z3y3/telegram-timeline-bot:latest

heroku-push:
	heroku container:push web -a telegram-timeline-robot

heroku-release:
	heroku container:release web -a telegram-timeline-robot