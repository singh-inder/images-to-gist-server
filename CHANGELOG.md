# Changelog

All notable changes to this project will be documented in this file. See [commit-and-tag-version](https://github.com/absolute-version/commit-and-tag-version) for commit guidelines.

## [2.0.1](https://github.com/singh-inder/images-to-gist-server/compare/v2.0.0...v2.0.1) (2024-09-30)

### Bug Fixes

- **handlers:** add cache-control header ([80dc3dc](https://github.com/singh-inder/images-to-gist-server/commit/80dc3dc40cb83ed496a7b8605ebb526d08a4515d))
- **docker-compose:** fix container_name ([e4140bc](https://github.com/singh-inder/images-to-gist-server/commit/e4140bce354655b74c0b8fddba7264b6935354ba))

## [2.0.0](https://github.com/singh-inder/images-to-gist-server/compare/v1.3.0...v2.0.0) (2024-09-23)

### ⚠ BREAKING CHANGES

- - migrate to golang

* drop support for resizing

- migrate to golang

- Merge branch 'migrate-go' ([3a6f449](https://github.com/singh-inder/images-to-gist-server/commit/3a6f449dfaf32a9006c4f53f74170488a3232651))

### Features

- **.github:** update publish-img workflow ([993371b](https://github.com/singh-inder/images-to-gist-server/commit/993371b6544ed8a2fa4c5df7103a859f60d9c713))
- add Dockerfile.goserver ([85486c5](https://github.com/singh-inder/images-to-gist-server/commit/85486c5b651e46e525455a4312298d6b9cf90ce2))
- add golang fiber ([b98eb08](https://github.com/singh-inder/images-to-gist-server/commit/b98eb0849752461f77ea4b31b1bdb4a4d078ee7f))

### Bug Fixes

- **handlers:** check response status ([8464997](https://github.com/singh-inder/images-to-gist-server/commit/8464997a1207e6694b0ef5c90366373af6280597))

### Docs

- update readme ([a89837f](https://github.com/singh-inder/images-to-gist-server/commit/a89837f075cc846eae5cdc9ea5212e932b7e0426))

## [1.3.0](https://github.com/singh-inder/images-to-gist-server/compare/v1.2.0...v1.3.0) (2024-08-22)

### Features

- set default DOMAIN to localhost:80 in env var ([dd55d34](https://github.com/singh-inder/images-to-gist-server/commit/dd55d34ca6306bc825260b95f26eb91a2bbbe3f4))

## [1.2.0](https://github.com/singh-inder/images-to-gist-server/compare/v1.1.0...v1.2.0) (2024-08-05)

### Features

- **Dockerfile.nodeserver:** use node:22.4.1-bullseye-slim image ([70a709f](https://github.com/singh-inder/images-to-gist-server/commit/70a709f579be57b4b54412d0c201ea62eb496ffe))
- **docker-compose.yaml:** remove env_file directive ([7012ab4](https://github.com/singh-inder/images-to-gist-server/commit/7012ab4c87e23d7ffcab92caababdf16cde15548))

### Docs

- update readme ([39dc49b](https://github.com/singh-inder/images-to-gist-server/commit/39dc49b9c64fd1ee3e21669958a204959e31c340))

## [1.1.0](https://github.com/singh-inder/images-to-gist-server/compare/v1.0.0...v1.1.0) (2024-07-06)

### Bug Fixes

- **publish-image.yaml:** update workflow ([72ae264](https://github.com/singh-inder/images-to-gist-server/commit/72ae2645ef34b7f505f2fea268e13e613a7d7668))

## [1.0.0](https://github.com/singh-inder/images-to-gist-server/compare/v0.1.0...v1.0.0) (2024-07-06)

### Features

- add Caddy server ([2af90c6](https://github.com/singh-inder/images-to-gist-server/commit/2af90c6b00c702fd57036dccf367cffc5fb90a47))
- add Dockerfile, publish-image workflow ([a831fa8](https://github.com/singh-inder/images-to-gist-server/commit/a831fa8160fca493bc503d2acb8fea6875b9a5d3))
- resize images if query params present ([0870c82](https://github.com/singh-inder/images-to-gist-server/commit/0870c82a23bceee455e7e8373b25528c39784d90))

### Bug Fixes

- **hosts:** update allowed hosts ([1ce776e](https://github.com/singh-inder/images-to-gist-server/commit/1ce776e7b86cf2f5b31f83bafb293c237214bebe))

## 0.1.0 (2024-06-28)

### Features

- stream images to client ([c121e1e](https://github.com/singh-inder/images-to-gist-server/commit/c121e1e7b5f5a8f93d3166a4dd1a79c6465713f0))
