# go-circleci

[![CircleCI](https://circleci.com/gh/twogg-git/go-circleci.svg?style=svg)](https://circleci.com/gh/twogg-git/go-circleci)

Simple Go string validator app to test with CircleCI


### Adding the port from the OS
Heroku configurations are fetched from the OS environment variables.
So to fetch the webapp port, you will need to call os.Getenv("PORT").

https://l-lin.github.io/post/2015/2015-01-31-golang-deploy_to_heroku/
