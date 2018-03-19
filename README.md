# simple-mikroserviche
Implements a simple microservice to expose REST APIs for a book store
<br/>Intended to be a tutorial for developing container based service in GoLang


  The following APIs are supported:
  1. Add a book
  2. Get all books
  3. Update book details

## Build steps:
  1. Run `GOOS=linux GOARCH=amd64 go build` in repos's working directory to create application binary. Note `GOOS=linux` since we are using linux image(This step is required as the binary is used in packaging docker image)
  2. Run `docker build -t book-store:1.0.0 .` in repo's working directory
  2. Use `docker run -it -p 8089:8089 book-store:1.0.0` to run docker image on port 8089(by default port 8089 is used in the container image thorugh an env variable). Additional port can be specified for forwarding by using -e option with docker like so: `docker run -it -e "PORT=9090" -p 8089:9090 book-store:1.0.1`
