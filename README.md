# Url_Shortener
This Go program implements a simple URL shortener using HTTP redirects based on path mappings provided either directly or via YAML configuration
Features
MapHandler: Handles URL redirection based on a predefined map of paths to URLs.
YAMLHandler: Parses a YAML configuration to dynamically create URL redirection mappings.
Usage
Setup:

Ensure you have Go installed on your system.
Install the required YAML package using go get gopkg.in/yaml.v3.
Running the Program:

bash
Copy code
go run main.go handler.go
Endpoints:

/url-godoc: Redirects to https://courses.calhoun.io/courses/cor_gophercises
/yaml-godoc: Redirects to https://www.udemy.com/
/: Returns "Hello, World"
Configuration:
You can configure additional redirects by editing the YAML content in the main.go file under yaml := section.

Dependencies:

gopkg.in/yaml.v3: YAML parsing library for Go.
Files
main.go: Contains the main function and server setup.
handler.go: Defines MapHandler and YAMLHandler functions for URL redirection.
How It Works
MapHandler: Directly maps specified paths to URLs and redirects accordingly.
YAMLHandler: Parses YAML input to dynamically create redirection mappings, allowing for easy configuration without changing code.
Future Improvements
Enhance error handling and logging.
Support for more dynamic configuration sources (e.g., databases).
