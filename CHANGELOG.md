# Changelog

## [0.1.15](https://github.com/Wielewout/arc-cleaner/compare/v0.1.14...v0.1.15) (2024-10-26)


### Bug Fixes

* **deps:** update github.com/actions/actions-runner-controller digest to 96d1bbc ([#22](https://github.com/Wielewout/arc-cleaner/issues/22)) ([a068d35](https://github.com/Wielewout/arc-cleaner/commit/a068d352a8cd983d601ac5861d06839656b023a9))
* **deps:** update golang.org/x/exp digest to f66d83c ([#24](https://github.com/Wielewout/arc-cleaner/issues/24)) ([3146eac](https://github.com/Wielewout/arc-cleaner/commit/3146eac81aa72488d739b48ba3a381748f59df38))
* **deps:** update kubernetes packages to v0.31.2 ([#28](https://github.com/Wielewout/arc-cleaner/issues/28)) ([a999468](https://github.com/Wielewout/arc-cleaner/commit/a999468fe8c4b8ea42b766a414a9cc3b823d6edf))
* **deps:** update module github.com/docker/docker to v27.3.1+incompatible ([#36](https://github.com/Wielewout/arc-cleaner/issues/36)) ([a4355ae](https://github.com/Wielewout/arc-cleaner/commit/a4355ae94575616756bce0e7b72651cbae43a6a2))
* **deps:** update module github.com/mittwald/go-helm-client to v0.12.14 ([#29](https://github.com/Wielewout/arc-cleaner/issues/29)) ([b66f4b1](https://github.com/Wielewout/arc-cleaner/commit/b66f4b126896403dd7fa17373395e2e5853be289))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.19.1 ([#30](https://github.com/Wielewout/arc-cleaner/issues/30)) ([4052ad1](https://github.com/Wielewout/arc-cleaner/commit/4052ad1aca5fc3cf3bbdae630a25bc7ae94cc41d))
* **deps:** update testcontainers-go monorepo to v0.34.0 ([#37](https://github.com/Wielewout/arc-cleaner/issues/37)) ([3379586](https://github.com/Wielewout/arc-cleaner/commit/337958679c544b4e301485a0b6ef817a18807455))

## [0.1.14](https://github.com/Wielewout/arc-cleaner/compare/v0.1.13...v0.1.14) (2024-09-15)


### Bug Fixes

* fix release artifact summary in pipeline ([036ac93](https://github.com/Wielewout/arc-cleaner/commit/036ac9305320eac91d5da1057eaa2895a3a84231))

## [0.1.13](https://github.com/Wielewout/arc-cleaner/compare/v0.1.12...v0.1.13) (2024-09-14)


### Bug Fixes

* fix missing newline escaping ([cfab4ed](https://github.com/Wielewout/arc-cleaner/commit/cfab4edd1258676189713648ccd282a2340f7479))

## [0.1.12](https://github.com/Wielewout/arc-cleaner/compare/v0.1.11...v0.1.12) (2024-09-14)


### Bug Fixes

* ensure .cr-index directory exists ([1858c65](https://github.com/Wielewout/arc-cleaner/commit/1858c65d19dba0cd91b285190d49341e1c3c16c9))

## [0.1.11](https://github.com/Wielewout/arc-cleaner/compare/v0.1.10...v0.1.11) (2024-09-14)


### Bug Fixes

* fix chart release name template ([bd9594e](https://github.com/Wielewout/arc-cleaner/commit/bd9594e98d8ba0eb8a6a19cca88bcd581497a490))

## [0.1.10](https://github.com/Wielewout/arc-cleaner/compare/v0.1.9...v0.1.10) (2024-09-14)


### Bug Fixes

* skip chart release steps if not releasing ([1dbb995](https://github.com/Wielewout/arc-cleaner/commit/1dbb995e2c2d5067b619edf7e4d4b663cd91d945))

## [0.1.9](https://github.com/Wielewout/arc-cleaner/compare/v0.1.8...v0.1.9) (2024-09-14)


### Bug Fixes

* fix repository name for chart release ([2fc3106](https://github.com/Wielewout/arc-cleaner/commit/2fc310681b015fe50377d0f78b9e64f9e1d50bd3))
* skip release if acceptance is skipped ([d25eef8](https://github.com/Wielewout/arc-cleaner/commit/d25eef83e5219485f050d2fd58394af6cebfd87a))

## [0.1.8](https://github.com/Wielewout/arc-cleaner/compare/v0.1.7...v0.1.8) (2024-09-14)


### Bug Fixes

* setup git user without bash for chart release ([a18619e](https://github.com/Wielewout/arc-cleaner/commit/a18619ec8c8d202d8c3840dbb3f3f07b45c992bc))

## [0.1.7](https://github.com/Wielewout/arc-cleaner/compare/v0.1.6...v0.1.7) (2024-09-14)


### Bug Fixes

* checkout full repo for chart release ([b7de741](https://github.com/Wielewout/arc-cleaner/commit/b7de741a213e18bdeb716888561b0d109d15a738))
* setup git user for chart release ([2bf81eb](https://github.com/Wielewout/arc-cleaner/commit/2bf81eb5b5a584f54b700915edb547193f561bb6))

## [0.1.6](https://github.com/Wielewout/arc-cleaner/compare/v0.1.5...v0.1.6) (2024-09-14)


### Bug Fixes

* make workdir safe after checkout in release step of pipeline ([5f9acb6](https://github.com/Wielewout/arc-cleaner/commit/5f9acb679c1d19340d8ad45174db7cfabc63d1ef))

## [0.1.5](https://github.com/Wielewout/arc-cleaner/compare/v0.1.4...v0.1.5) (2024-09-14)


### Bug Fixes

* fix chart asset path ([8f51e60](https://github.com/Wielewout/arc-cleaner/commit/8f51e60cfc839b25560c77602527c38797616182))
* fix missing line escape for chart release ([8df45e8](https://github.com/Wielewout/arc-cleaner/commit/8df45e88b4ae69caea760430db649e7d3bde4b93))

## [0.1.4](https://github.com/Wielewout/arc-cleaner/compare/v0.1.3...v0.1.4) (2024-09-14)


### Bug Fixes

* fix major release version containing a v prefix ([309cf9e](https://github.com/Wielewout/arc-cleaner/commit/309cf9ec9ed012df5f6a774ba5f2f2f6a67ee938))

## [0.1.3](https://github.com/Wielewout/arc-cleaner/compare/v0.1.2...v0.1.3) (2024-09-14)


### Bug Fixes

* run release in devcontainer ([02ca7f2](https://github.com/Wielewout/arc-cleaner/commit/02ca7f26cb3c5eebc9cb47b90c58eceabbe87a4a))

## [0.1.2](https://github.com/Wielewout/arc-cleaner/compare/v0.1.1...v0.1.2) (2024-09-14)


### Bug Fixes

* fix container image tags in pipeline artifacts ([3bbc996](https://github.com/Wielewout/arc-cleaner/commit/3bbc996c05ce9d486981c29fa8be6cfb5e86d9d0))
* strip v from release name in pipeline artifacts ([649bfe3](https://github.com/Wielewout/arc-cleaner/commit/649bfe3ac6cdf690fe180a9e854fd2d776451f78))

## [0.1.1](https://github.com/Wielewout/arc-cleaner/compare/v0.1.0...v0.1.1) (2024-09-14)


### Bug Fixes

* fix release binary download in pipeline ([5c038ff](https://github.com/Wielewout/arc-cleaner/commit/5c038ffd8ce90993d2f5d76972beab3699b05bbb))

## 0.1.0 (2024-09-14)


### Features

* add app entrypoint with --config option ([2243bfc](https://github.com/Wielewout/arc-cleaner/commit/2243bfc0bf9bafcccb237529856ebf6b90b2c4bc))
* add context with cancel on os signals ([882cbb3](https://github.com/Wielewout/arc-cleaner/commit/882cbb3d668f8f9606f1e80e2d80d9529d259c23))
* add dryrun to config ([c50cc83](https://github.com/Wielewout/arc-cleaner/commit/c50cc83c23a26fccf6ce9271eec0c6c0aacf2689))
* add helm chart ([f11f122](https://github.com/Wielewout/arc-cleaner/commit/f11f1229a2ab6d46165ddf12006e1dfc70c7caf0))
* add k8s client with github arc scheme ([e18dfc7](https://github.com/Wielewout/arc-cleaner/commit/e18dfc71eca562c05ed569f251bef8610e8f1c4d))
* add liveness and readiness http endpoints ([f11f122](https://github.com/Wielewout/arc-cleaner/commit/f11f1229a2ab6d46165ddf12006e1dfc70c7caf0))
* add log level to config ([ff49114](https://github.com/Wielewout/arc-cleaner/commit/ff49114b92d2f26be373d277768529085fb3fa49))
* add namespace to config ([8aefaf4](https://github.com/Wielewout/arc-cleaner/commit/8aefaf4e85fbdacfe9bcf1707c27c9419ce9a582))
* add version and commit in startup log ([a95bcd1](https://github.com/Wielewout/arc-cleaner/commit/a95bcd1861be781b42a8138bc6c4b269a5ab9126))
* compress binary ([33781ff](https://github.com/Wielewout/arc-cleaner/commit/33781ffa099e0062505890775b54917f2396bbf4))
* delete workflow pod if runner pod pending ([964efd0](https://github.com/Wielewout/arc-cleaner/commit/964efd031d9b9b98fd48fab5d224a4a477501b8f))
* delete workflow pod without ephemeral runner ([01eeefb](https://github.com/Wielewout/arc-cleaner/commit/01eeefb8a07e6392caeb565d1ecef356d0b4a46f))
* make clean up period configurable ([21f7944](https://github.com/Wielewout/arc-cleaner/commit/21f79440ed4f735b05d44e831ceb1338213a46ec))
* periodically clean up ephemeral runners ([583a469](https://github.com/Wielewout/arc-cleaner/commit/583a469bc647f186befa4eaa4d9af1fcca95f001))


### Bug Fixes

* fix logging with context ([61adaa4](https://github.com/Wielewout/arc-cleaner/commit/61adaa4b3cbb30fd29552cdc4df112a6831135d6))
* stop app when kubernetes client creation fails ([c0f2bb6](https://github.com/Wielewout/arc-cleaner/commit/c0f2bb63bef03116baaa07193a67cd77eb1e60e3))
* work on ephemeral runner iso set ([8b39e45](https://github.com/Wielewout/arc-cleaner/commit/8b39e458f7c3c1f7094e88e480ca36e9a9924bc0))


### Miscellaneous Chores

* release please as 0.1.0 ([81733ec](https://github.com/Wielewout/arc-cleaner/commit/81733ec367278bd971b1965f2352bdc41a268174))
