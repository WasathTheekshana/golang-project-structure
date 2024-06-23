#!/bin/bash

CONTAINER_NAME="golang-proj-struct"

# Function to log and print messages
log_message() {
    echo "$1"
    echo "$(date) - $1" >> /var/log/docker_script.log
}

# Check if the container exists (whether it's running or stopped)
if [ "$(docker ps -a -q -f name=${CONTAINER_NAME})" ]; then
    # Start the container if it's not running
    if [ ! "$(docker ps -q -f name=${CONTAINER_NAME})" ]; then
        docker start ${CONTAINER_NAME}
        if [ $? -eq 0 ]; then
            log_message "Started existing container ${CONTAINER_NAME}."
        else
            log_message "Failed to start existing container ${CONTAINER_NAME}."
            docker logs ${CONTAINER_NAME} >> /var/log/docker_script.log 2>&1
        fi
    else
        log_message "Container ${CONTAINER_NAME} is already running."
    fi
else
    # Run a new container if it doesn't exist
    docker run -d --name ${CONTAINER_NAME} -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=golang_proj -p 5432:5432 -v golang-proj-struct:/var/lib/postgresql/data postgres:latest
    if [ $? -eq 0 ]; then
        log_message "Created and started new container ${CONTAINER_NAME}."
    else
        log_message "Failed to create and start new container ${CONTAINER_NAME}."
        docker logs ${CONTAINER_NAME} >> /var/log/docker_script.log 2>&1
    fi
fi
