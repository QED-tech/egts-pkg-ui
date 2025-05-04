.PHONY: roles
IMAGE_TAG=$(shell git log | head -n 1 | tail -c 7)

roles:
	ansible-galaxy install -r requirements.yml
server:
	ansible-playbook -i inventory.yml server.yml -b --ask-become-pass
deploy:
	ansible-playbook -i inventory.yml deploy.yml -b --ask-become-pass -e IMAGE_TAG=${IMAGE_TAG}