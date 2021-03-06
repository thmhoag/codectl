project_name: {{.RepoName}}
env:
  - GO111MODULE=on
  - CGO_ENABLED=0
before:
  hooks:
    - go mod download
builds:
  - main: cmd/{{.RepoName}}/main.go
    binary: {{.RepoName}}
    ldflags:
      - -s -w -X main.semver=v{{`{{.Version}}`}} -X main.commit={{`{{.Commit}}`}} -X main.built={{`{{.Date}}`}}
    goos:
      - darwin
      - windows
      - linux
    goarch:
      - amd64
archives:
- format_overrides:
    - goos: windows
      format: zip
  name_template: {{`'{{ .Binary }}-{{ .Os }}-{{ .Arch }}{{ if eq .Os "windows"}}.exe{{end}}'`}}
  # only the binary
  files:
    - none*
checksum:
  name_template: 'checksums.txt'
snapshot:
  name_template: {{`"{{ .Tag }}-next"`}}
changelog:
  sort: asc
  filters:
    exclude:
    - '^docs:'
    - '^test:'
release:
  github:
    owner: {{.OrgName}}
    name: {{.RepoName}}
  prerelease: auto
