# simple-mikroserviche
Implements a simple microservice to expose REST APIs for a book store
<br/>Intended to be a tutorial for developing container based service in GoLang


  The following APIs are supported:
  1. Add a book
  2. Get all books
  3. Update book details

## Build steps:
  1. `docker build -t book-store:1.0.0 .` in repo's working directory
  2. `docker run -it -p 8089:8089 book-store:1.0.0` to run docker image on port 8089(by default port 8089 is used in the container image thorugh an env variable). Additional port can be specified for forwarding by using -e option with docker like so: `docker run -it -e "PORT=9090" -p 8089:9090 book-store:1.0.1`
