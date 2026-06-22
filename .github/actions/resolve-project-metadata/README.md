# resolve-project-metadata

Composite action that resolves release metadata for this multi-module Gradle repository.

## Scope

This action only computes metadata and writes outputs:

- `version` (`NMS_TAG` format)
- `java_images` (JSON array for matrix `include`)
- `third_party_images` (JSON array of `{name,registry,repository,tag}` objects)

It does **not** build, push, scan, or release anything.

## Inputs
- `github-output-file` (optional, default: none) - if provided, the action writes outputs to this file in GitHub Actions format.
- `version-file` (default: `version.properties`)


## Resolver Interface (Go)

The composite action invokes:

```bash
cd .github/actions/resolve-project-metadata
go run . \
  --version-file ../../../version.properties 
```

## Outputs
Without `--github-output-file`, the resolver prints a JSON document:

```json
{
  "version": "latest",
  "java_images": [
    {
      "name": "iot-collector",
      "registry": "docker.io",
      "repository": "dgs19/iot-collector",
      "tag": "latest"
    },
    {
      "name": "iot-collector-ui",
      "registry": "docker.io",
      "repository": "dgs19/iot-collector-ui",
      "tag": "latest"
    }
  ],
  "third_party_images": [
    {
      "name": "traefik",
      "registry": "docker.io",
      "repository": "traefik",
      "tag": "v3.7.5"
    }
  ]
}
```

## init modules and dependencies
```shell
# cd in to this directory
go mod init synca/resolver
go mod tidy
go get github.com/joho/godotenv
```


