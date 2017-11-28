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

##Contributing
1. Create a Jira Ticket for tracking your work and the reason for the change
2. Create a named feature branch with the Jira ticket included (like `PAAS-1337-add_component_x`)
3. Write your change
4. Write tests for your change (if applicable)
5. Run the tests, ensuring they all pass
6. Bump the version in `CHANGELOG.md` with the new version and what you changed.
7. Commit your changes with your Jira ticket in the commit message (like `git commit -m 'PAAS-1337' added component x and spruced up docs on y`)
8. Submit a Pull Request using Github

##License and Authors
Authors: paas-api@gannett.com
_Copyright (c) 2016 Gannett Co., Inc, All Rights Reserved._
