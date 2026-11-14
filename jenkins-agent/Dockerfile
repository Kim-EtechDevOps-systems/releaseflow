FROM jenkins/inbound-agent:latest-jdk21

USER root

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    git python3 curl ca-certificates golang-go && \
    rm -rf /var/lib/apt/lists/*

USER jenkins
