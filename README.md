# Golang + CircleCI + Heroku
Simple Go string flip code call by a webpage deployed in Heroku by CircleCI. [![CircleCI](https://circleci.com/gh/twogg-git/go-circleci.svg?style=svg)](https://circleci.com/gh/twogg-git/go-circleci)

## Content 

This repo contains the following source files:
- strings.go : a very small golang code to flip a given string  
- strings_test.go: unit test cases in golang for CircleCI to run
- main.go: main class with the endpoints and the setters for the html template 
- temp.html: the actual webpage that we are going to visualize

Heroku requieres two files to be able to deploy in their servers:
- requirements.txt: that at ths point are none, so its a blank file
- Godeps/Godeps.json: the descriptor of the Go Version to use

To test locally, we added a Docker deployment
- Dockerfile: small Dockerfile to be able to build an image and run a container for our webpage

Finally, CircleCI config file
- config.yml: here we setup the actual build code in Circle and the futher deployment in Heroku.

## Local Docker Deployment

To build locally the image
```sh
docker build -t webtest:v1 .
```

To run the container
```sh
docker run --name webtest -p 8181:8080 -t webtest:v1
```

Test the webpage
```sh
http:localhost:8181
```

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
