# Golang + CircleCI + Heroku
Simple Go string flip code call by a webpage deployed in Heroku by CircleCI. [![CircleCI](https://circleci.com/gh/twogg-git/go-circleci.svg?style=svg)](https://circleci.com/gh/twogg-git/go-circleci)

Here is the live deployment in Heroku: **https://go-circle.herokuapp.com/** 

## Requirements
- GitHub account (https://github.com), so you can build your app in CircleCI servers. 
- Heroku Account https://www.heroku.com, to deploy the webpage in Heroku servers.
- CircleCI connection (https://circleci.com), provide CircleCI access to your repository in GitHub.

## Docker Deployment
To run the container
```sh
docker run --name go-circle -p 8181:8080 -t twogghub/go-circle:v1

# Then test the webpage going in yuor browser to:
http:localhost:8181
```

## Local building and deployment
To build locally the image, go to the folder's project then run:
```sh
docker build -t go-circle .
docker run --name go-circle -p 8181:8080 -t go-circle
```

## Source Code 
This repo contains the following source files:
- strings.go : a very small golang code to flip a given string  
- strings_test.go: unit test cases in golang for CircleCI to run
```sh
// flip reverses all the characters on the give string s
func Flip(s string) string {
	if len(s) <= 1 {
		return s
	}
	return s[len(s)-1:] + Flip(s[:len(s)-1])
}
```
- main.go: main class with the endpoints and the setters for the html template 
```sh
  # the endpoints available 
  http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		...
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		...
	})

	http.HandleFunc("/flip", func(w http.ResponseWriter, r *http.Request) {
		...
	})
```
- temp.html: the actual webpage that we are going to visualize
```sh
# loading the info as a parameter from the golang code
<div class="content">
	Current IP: {{.ServerIP}}<br>
	Deployed: {{.Deployed}}<br>
	Version: {{.Version}}
</div>
...
# calling flip's endpoint 
<script>
    function flipFunc() {
		var callTo = "http://{{.ServerIP}}:{{.Port}}/flip?text=" + toflip.value;
		var xhttp = new XMLHttpRequest();
	    xhttp.onreadystatechange = function() {
	         if (this.readyState == 4 && this.status == 200) {
				document.getElementById('fliResult').innerHTML = this.responseText;
	         }
	    };
	    xhttp.open("GET", callTo, true);
	   	xhttp.setRequestHeader("Content-type", "text/plain");
	    xhttp.send();
 }
</script>
```

Heroku requieres two files to be able to deploy in their servers:
- requirements.txt: that at ths point are none, so its a blank file
- Godeps/Godeps.json: the descriptor of the Go Version to use
```sh
{
	"ImportPath": "github.com/freeformz/go-heroku-example",
	"GoVersion": "go1.4.2",
  ...
}
```

To test locally, we added a Docker deployment
- Dockerfile: small Dockerfile to be able to build an image and run a container for our webpage. 
```sh
  FROM golang:1.10-alpine3.7 as builder
  WORKDIR /go/src/go-circleci/main
  COPY . .
  RUN go get -d ./... && go build -o main .

  FROM alpine:3.8
  RUN apk --no-cache add ca-certificates
  WORKDIR /root/
  COPY --from=builder /go/src/go-circleci/main .

  EXPOSE 8080
  ENTRYPOINT ./main
```

Finally, CircleCI config file
- config.yml: here we setup the actual build code in Circle and the futher deployment in Heroku.
```sh
# Including heere the orb for Heroku
orbs:
  heroku: circleci/heroku@1.0.0
  ...
    # Calling remotly the github repo
    working_directory: /go/src/github.com/twogg-git/go-circleci
    steps:
      - checkout
      - run: go get -v -t -d ./...
      - run: go test -v ./...
  ...
        # And deploying in Heroku
        - run:
            name: Deploy Master to Heroku
            command: |
              git push https://heroku:$HEROKU_KEY@git.heroku.com/go-circle.git master
 ...
```

### Adding the port from the OS

Heroku dynamically assigns your app a port, so you can't set the port to a fixed number. Heroku adds the port to the env, so you can pull it from there. Heroku configurations are fetched from the OS environment variables. So to fetch the webapp port, you will need to call **os.Getenv("PORT")**.
```sh
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = PORT
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return port
}
```
Source: https://l-lin.github.io/post/2015/2015-01-31-golang-deploy_to_heroku/

## App deployed into Heroku

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
