# Copyright (c) 2021 John Dewey <john@dewey.ws>
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

FROM debian:buster as tools-builder

ARG TF_VERSION=0.14.7
ARG TF_ZIP=${WORKDIR}/terraform_${TF_VERSION}_linux_amd64.zip
ARG GOSS_VERSION=latest
ARG GOSS_BINARY=goss
ARG WORKDIR=/tmp

WORKDIR ${WORKDIR}

RUN \
	apt-get update \
	&& apt-get install curl unzip -y

RUN \
	curl https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip \
	-o ${TF_ZIP} \
	&& unzip ${TF_ZIP}

RUN \
	curl -L https://github.com/aelsabbahy/goss/releases/${GOSS_VERSION}/download/goss-linux-amd64 \
	-o ${GOSS_BINARY} \
	&& chmod +rx ${GOSS_BINARY}

FROM golang:buster as terrable-builder

ARG WORKDIR=/src

COPY . /${WORKDIR}/
WORKDIR ${WORKDIR}

RUN \
	make build install

# FROM scratch
FROM debian:buster
LABEL maintainer="John Dewey <john@dewey.ws>"

COPY --from=tools-builder /tmp/terraform /usr/local/bin/
COPY --from=tools-builder /tmp/goss /usr/local/bin/
COPY --from=terrable-builder /root/.terraform.d/plugins /root/.terraform.d/plugins

# ENV PATH="/usr/local/bin"
# ENTRYPOINT ["terraform"]
# CMD ["help"]
