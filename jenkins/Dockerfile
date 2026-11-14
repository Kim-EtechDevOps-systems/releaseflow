FROM jenkins/inbound-agent:latest-jdk21

USER root

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    git golang-go python3 gcc libc6-dev && \
    rm -rf /var/lib/apt/lists/*

USER jenkins