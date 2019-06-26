# go-circleci

[![CircleCI](https://circleci.com/gh/twogg-git/go-circleci.svg?style=svg)](https://circleci.com/gh/twogg-git/go-circleci)

Simple Go string validator app to test with CircleCI


### Adding the port from the OS

Heroku dynamically assigns your app a port, so you can't set the port to a fixed number. Heroku adds the port to the env, so you can pull it from there

Heroku configurations are fetched from the OS environment variables.
So to fetch the webapp port, you will need to call os.Getenv("PORT").

port := os.Getenv("PORT")


https://l-lin.github.io/post/2015/2015-01-31-golang-deploy_to_heroku/

### App deployed into Heroku

https://go-circle.herokuapp.com/

https://devcenter.heroku.com/articles/go-support#go-versions
https://medium.com/forloop/continuously-deploy-your-golang-binaries-using-circleci-and-heroku-docker-eb27e06d68f2
https://circleci.com/docs/2.0/deployment-integrations/#heroku

```sh
version: 2
jobs:
  build:
    ...
  deploy:
    docker:
      - image: buildpack-deps:trusty
    steps:
      - checkout
      - run:
          name: Deploy Master to Heroku
          command: |
            git push https://heroku:$HEROKU_API_KEY@git.heroku.com/$HEROKU_APP_NAME.git master

workflows:
  version: 2
  build-deploy:
    jobs:
      - build
      - deploy:
          requires:
            - build
          filters:
            branches:
              only: master
 ```
