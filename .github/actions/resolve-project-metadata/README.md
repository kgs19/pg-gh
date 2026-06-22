# resolve-project-metadata

Composite action that resolves release metadata for this multi-module Gradle repository.

## Scope

This action only computes metadata and writes outputs:

- `version` (`NMS_TAG` format)
- `java_images` (JSON array for matrix `include`)
- `third_party_images` (JSON array of `{name,registry,repository,tag}` objects)

It does **not** build, push, scan, or release anything.

## Inputs

- `version-file` (default: `version.properties`)

## Outputs

- `version`: Example `18.1.1-B12345`
- `java_images`: Example

```json
[
  { "name": "data-access", "registry": "ghcr.io", "repository": "adtran-osa/nc-synca/gnss-data-access", "tag": "18.1.1-B12345" },
  { "name": "gnss-cli-worker", "registry": "ghcr.io", "repository": "adtran-osa/nc-synca/gnss-cli-worker", "tag": "18.1.1-B12345" }
]
```

- `third_party_images`: Example

```json
[
  {
    "name": "kafka",
    "registry": "ghcr.io",
    "repository": "adtran-osa/nc/3rd-party/apache-kafka",
    "tag": "4.2.0-alpine"
  },
  {
    "name": "traefik",
    "registry": "ghcr.io",
    "repository": "adtran-osa/nc/3rd-party/traefik",
    "tag": "v3.6.10"
  }
]
```

## Resolver Interface (Go)

The composite action invokes:

```bash
cd .github/actions/resolve-project-metadata
go run . \
  --version-file ../../../version.properties 
```

Without `--github-output-file`, the resolver prints a JSON document:

```json
{
  "version": "18.1.1-B19301",
  "java_images": [
    { "name": "data-access", "registry": "ghcr.io", "repository": "adtran-osa/nc-synca/gnss-data-access", "tag": "18.1.1-B19301" },
    { "name": "gnss-cli-worker", "registry": "ghcr.io", "repository": "adtran-osa/nc-synca/gnss-cli-worker", "tag": "18.1.1-B19301" }
  ],
  "third_party_images": [
    {
      "name": "kafka",
      "registry": "ghcr.io",
      "repository": "adtran-osa/nc/3rd-party/apache-kafka",
      "tag": "4.2.0-alpine"
    },
    {
      "name": "traefik",
      "registry": "ghcr.io",
      "repository": "adtran-osa/nc/3rd-party/traefik",
      "tag": "v3.6.10"
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


