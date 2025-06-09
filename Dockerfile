FROM jenkins/jenkins:lts

USER root

RUN apt-get update && \
    apt-get install -y curl gnupg lsb-release zip tar wget && \
    apt-get clean

ENV GO_VERSION=1.21.4
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/opt/go
ENV PATH=$PATH:$GOPATH/bin
RUN apt-get update && apt-get install -y sudo
RUN mkdir -p /opt/go && chown -R jenkins:jenkins /opt/go

RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" \
    > /etc/apt/sources.list.d/docker.list && \
    apt-get update && \
    apt-get install -y docker-ce-cli

RUN groupadd docker && usermod -aG docker jenkins

USER jenkins
