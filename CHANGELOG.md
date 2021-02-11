## 0.3.2

### Changed
- Migrated build images to `circleci/golang:1.15`

## 0.3.1

### Changed
- Remove unused dependencies from `go.mod`

## 0.3.0

### Added
- Deploy package to github via [ghr](https://github.com/tcnksm/ghr)

### Changed
- Return the version without extra text in command line

### Removed
- Remove check for aws-cli, that's not our responsibility

## 0.2.0

### Added
- Add a `version` command
- Add an optional flag, `duration`, to allow setting the duration in seconds for the role to be assumed

## 0.1.0

- Initial release - assume a role temporarily
