package main

import (
	"os"
	"log"
	"fmt"

	"github.com/samalba/dockerclient"
)

func main() {
	
	// Initiate the Docker Client
	docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)

	// Set the Docker Content Trust environment variable
	os.Setenv("DOCKER_CONTENT_TRUST","1")
	fmt.Println("DOCKER_CONTENT_TRUST:",os.Getenv("DOCKER_CONTENT_TRUST"))

	// Useful Information during hack
    versionInfo, _ := docker.Version()
	log.Println(versionInfo)
	dockerVersion := versionInfo.Version
	log.Println(dockerVersion)

	info, _ := docker.Info()
	log.Println(info)
}
