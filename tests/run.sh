#!/usr/bin/env sh

NC='\033[0m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'

TERRAFORM_FLAGS="-auto-approve"
GOSS_FLAGS="--color"

set -e

for resource in $(ls resources); do
	for scenario in $(ls resources/${resource}/scenarios); do
		echo
		printf "${CYAN}*** Resource: ${resource}${NC}\n"
		printf "${PURPLE}***** Scenario: ${scenario}${NC}\n"
		echo
		(
			cd resources/${resource}/scenarios/${scenario}
			rm -rf ./.terraform ./.terraform.lock.hcl terraform.tfstate*
			terraform init >/dev/nullt

			TF_LOG=$TF_LOG terraform apply ${TERRAFORM_FLAGS}
			goss -g goss-apply.yaml validate ${GOSS_FLAGS}

			TF_LOG=$TF_LOG terraform destroy ${TERRAFORM_FLAGS}
			goss -g goss-destroy.yaml validate ${GOSS_FLAGS}
		)
	done
done
