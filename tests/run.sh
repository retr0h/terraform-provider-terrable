#!/usr/bin/env sh

NC='\033[0m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'

set -e

for resource in $(ls resources); do
	for scenario in $(ls resources/${resource}/scenarios); do
		echo
		printf "${CYAN}*** Resource: ${resource}${NC}\n"
		printf "${PURPLE}***** Scenario: ${scenario}${NC}\n"
		echo
		(
			cd resources/${resource}/scenarios/${scenario}
			rm -rf ./.terraform ./.terraform.lock.hcl
			terraform init >/dev/nullt

			TF_LOG=$TF_LOG terraform apply -auto-approve
			goss -g goss-apply.yaml validate --color

			TF_LOG=$TF_LOG terraform destroy -auto-approve
			goss -g goss-destroy.yaml validate --color
		)
	done
done
