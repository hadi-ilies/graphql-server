# graphql-server
  A Little graphql-server written in golang using mongoDB

# Usage
  ## Run mongoDB container
  ```sh
    $> docker-compose up
  ```
  ## Export env vars
  ```sh
    $> export $(cat .env | xargs)
  ```

  ## Run Golang server
  ```sh
    $> go run server.go
  ```
