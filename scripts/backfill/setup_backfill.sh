#!/bin/bash

function check_docker_installation() {
  echo -n "Checking for Docker installation ... "
  if command -v docker &> /dev/null; then
      echo "DOCKER INSTALLED"
  else
      echo "DOCKER NOT INSTALLED"
      echo "Exiting setup as Docker is needed to proceed, Please install and try again"
      exit 1
  fi
}

function check_requirements() {
  check_docker_installation
}

function create_prometheus_instance() {
  docker run -it --rm -p 9090:9090 -v ${PWD}/configs/prometheus.yml:/etc/prometheus/prometheus.yml \
  -v ${PWD}/../../mimik-data:/mimik-data:ro \
  --name=prometheus \
  prom/prometheus:v2.38.0 \
  --config.file=/etc/prometheus/prometheus.yml \
  --storage.tsdb.path=/prometheus \
  --storage.tsdb.allow-overlapping-blocks \
  --web.console.libraries=/usr/share/prometheus/console_libraries \
  --web.console.templates=/usr/share/prometheus/consoles \
  --storage.tsdb.retention.size=8GB \
  --storage.tsdb.retention.time=0s \
  --web.enable-lifecycle

}

check_requirements
create_prometheus_instance


