# go-count

[![Release](https://github.com/elmariofredo/go-count/actions/workflows/release.yml/badge.svg)](https://github.com/elmariofredo/go-count/actions/workflows/release.yml)
[![Coverage Status](https://coveralls.io/repos/github/elmariofredo/go-count/badge.svg?branch=main)](https://coveralls.io/github/elmariofredo/go-count?branch=main)

I'm going to build counter service with following features. Main purpose is to learn how to build go service from sratch and find best way how to deliver it to the production.

TODO:

- [X] HTTP endpoint response with increasing number on every call
- [X] Build CI ( Inspired by [vojtechmares/goreleaser-live](https://github.com/vojtechmares/goreleaser-live/blob/master/.goreleaser.yml) )
- [X] Tests 👻
- [X] Reset counter request
- [X] JSON logging with verbosity level
- [X] Database backend to preserve counter during restart
- [ ] Metrics
- [ ] Automated deployment/rollout to Kubernetes
- [ ] Reasonable test coverage
- [ ] Help flag listing vars ( https://github.com/spf13/pflag )
- [ ] Use models
- [ ] Check conformance with https://github.com/golang-standards/project-layout
- [ ] Add online versioned documentation e.g. https://docusaurus.io/docs

## Versioning

I'm using André Staltz's [ComVer](https://staltz.com/i-wont-use-semver-patch-versions-anymore.html) as I prefer backward compatibility tracking over major.minor.patch decission hell.
