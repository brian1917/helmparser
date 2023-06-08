# Helm Parser
helmparser is a small Go script that parses the Illumio Helm template YAML file into individual YAML documents.

## Run using executable
1. Download the latest executable from the release section of this repository: https://github.com/brian1917/helmparser/releases.
2. Run `helmparser illumio.yaml` (the `illumio.yaml` file is the YAML output from `helm template`)

## Run from source
1. Have Go install.
2. Run `go run main.go illumio.yaml` (the `illumio.yaml` file is the YAML output from `helm template`)

