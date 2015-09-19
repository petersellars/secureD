package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samalba/dockerclient"
)

func main() {

	// initiate the docker client
	docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	// set the docker content trust environment variable
	os.Setenv("docker_content_trust", "1")
	fmt.Println("docker_content_trust:", os.Getenv("docker_content_trust"))

	// useful information during hack
	versioninfo, _ := docker.Version()
	log.Println(versioninfo)
	dockerversion := versioninfo.Version
	log.Println(dockerversion)

	info, _ := docker.Info()
	log.Println(info)

	// secureD

	// Create the benchmark container
	// TODO - currently depends on benchmark being available
	
    hostConfig := &dockerclient.HostConfig {
		Binds: []string{"/var/lib:/var/lib",
			"/var/run/docker.sock:/var/run/docker.sock",
			"/usr/lib/systemd:/usr/lib/systemd",
			"/etc:/etc"},
		NetworkMode: "host",
		PidMode: "host",
		CapAdd: []string{"audit_control"},
	}
	
	benchmarkContainerConfig := &dockerclient.ContainerConfig{
		Image: "docker/docker-bench-security:latest",
		AttachStdin: true,
		AttachStdout: true,
		AttachStderr: true,
		OpenStdin: true,
		StdinOnce: true,
		Tty: true,
		Labels: map[string]string{"docker_bench_security":""},
	}
	containerId, err := docker.CreateContainer(benchmarkContainerConfig, "docker-benchmark")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(containerId)

	// Start the benchmark container
	err = docker.StartContainer(containerId, hostConfig)
	if err != nil {
		log.Fatal(err)
	}
}
