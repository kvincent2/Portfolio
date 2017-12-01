# KVincent's Portfolio <img src="icons/pirate.png" width="50">
===================

Describe me!

## Running
Clone the current repository, and run the following commands:
  ```
  dep ensure
  go run main.go
  ```

## Dependencies
External community dependencies are as follows:
- List your dependencies here
This project uses the Go package management tool [dep](https://github.com/golang/dep) for package versioning. To leverage this tool to install dependencies, run the following command from the project root:
  ```
  dep init
  dep ensure
  ```

## Testing
To run the unit tests on the packages included as a part of this application, run the following command:
  ```
  go test github.com/kvincent2/portfolio/...
  ```
After making changes to any of these packages, please update the testing as needed, and verify by running the above command.

##Running on Docker Container
Pull container from DockerHub using 'docker pull vincentkb0823/portfolio'
Run using 'docker run -p 8080:8080 -e GITHUB_PAT=[github access key] vincentkb0823/portfolio'


