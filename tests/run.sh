#!/usr/bin/env sh

set -ex

(
	cd user
	rm -rf ./.terraform ./.terraform.lock.hcl
	terraform init

	terraform apply -auto-approve
	goss -g goss-apply.yaml validate

	terraform destroy -auto-approve
	goss -g goss-destroy.yaml validate
)
