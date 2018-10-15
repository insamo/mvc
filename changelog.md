# Change Log

## [Unreleased][unreleased]
### Added
- Added nosql(couchdb) datasource and pseudo transaction factory
- Added nosql(couchbase) datasource
- Added sliceHelper package
### Changed
- Changed datasource structure
- Changed NoSql transaction factory
- Changed datasource package structure
### Fixed
- Fixed paginator.CurrentPage zero value
- Fixed database configure logger disable statement
- Fixed query string empty value
- Fixed query string page value calculation


## [0.0.9] - 2018-09-30
### Changed
- Changed nosql db factory to client factory
### Fixed
- Fixed sql query parse with no value

## [0.0.8] - 2018-09-05
### Added
- Added support for nosql(couchdb)
### Removed
- Removed default configs for database

## [0.0.7] - 2018-06-05
### Fixed
- Fixed database configure

## [0.0.6] - 2018-06-05
### Added
- Added scheduler
- Added query params to request logger
### Updated
- Updated iris framework 10.6.6:https://github.com/kataras/iris/blob/master/HISTORY.md#tu-05-june-2018--v1066
### Changed
- Changed database packet location

## [0.0.5] - 2018-06-01
### Added
- Added sql query param constructor
### Changed
- Changed url params, now dividing sql and url params

## [0.0.4] - 2018-05-31
### Removed
- Removed unused utils

## [0.0.3] - 2018-05-31
### Added
- Added env PORT variable support

## [0.0.2] - 2018-05-31
### Added
- Added some utils
- Added repository prepare
- Added disable logging
### Changed
- Changed paginator
- Changed name dialect
- Changed query string

## [0.0.1] - 2018-05-06
### Added
- Based on Iris 10.6.3:https://github.com/kataras/iris/blob/master/HISTORY.md#we-02-may-2018--v1063

# Change Log FAQ

[FAQ LINK](http://keepachangelog.com/)

## [Added] for new features.
## [Changed] for changes in existing functionality.
## [Deprecated] for once-stable features removed in upcoming releases.
## [Removed] for deprecated features removed in this release.
## [Fixed] for any bug fixes.
## [Security] to invite users to upgrade in case of vulnerabilities.
