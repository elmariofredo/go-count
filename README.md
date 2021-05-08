# go-count

[![Release](https://github.com/elmariofredo/go-count/actions/workflows/release.yml/badge.svg)](https://github.com/elmariofredo/go-count/actions/workflows/release.yml)

I'm going to build counter service with following features. Main purpose is to learn how to build go service from sratch and find best way how to deliver it to the production.

TODO:

- [X] HTTP endpoint response with increasing number on every call
- [X] Build CI ( Inspired by https://github.com/vojtechmares/goreleaser-live/blob/master/.goreleaser.yml )
- [ ] Tests 👻
- [X] Reset counter request
- [X] JSON logging with verbosity level
- [ ] Automated deployment/rollout to Kubernetes
- [ ] Database backend to preserve counter during restart

## Versioning

I'm using André Staltz's [ComVer](https://staltz.com/i-wont-use-semver-patch-versions-anymore.html) as I prefer backward compatibility tracking over major.minor.patch decission hell.

