# URL SHORT
This is a url shortner api built using golang and [golang-gin](https://github.com/gin-gonic/gin), along with MongoDB.

### Requirements :
. Golang
. [Golang-Gin](https://github.com/gin-gonic/gin)
. MongoDB

### Environment variables :
- URL_MONGO : mongodb connection url / uri
    

### Run :
- Create the vendor files for the project
    ```bash
    $ make vendor
    ```
- Build the compiled source of the project
    ```bash
    $ make build
    ```
    - build takes an optional argument OS : target OS to be built to [ windows | linux | darwin ...]
        ```
        $ make -e OS=windows build
        ```
    Build outputs to a folder named bin/url_short with the executable.
    
- Run in debug mode
    ```
    $ go run cmd/main.go
    ```
- Run in release mode
    ```
    $ go run cmd/main.go -mode release
    ```