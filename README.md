# Gotodos

# How to run the program

- Initially, download and install golang

  - link: https://go.dev/dl/

- Next thing is postgresql, download and install it too

  - link: https://go.dev/dl/

- Run "sudo -u postgres psql"
  - You would get postgres cli
  - Now, you can create new database with:
    - "CREATE DATABASE "YOUR_DATABASE_NAME;"
  - Change user password if you don't know the current user password with:
    - "ALTER USER user_name WITH PASSWORD 'new_password';"
  - You can check the username if you don't know the current username with:
    - "\l" List all the database details available in current schema
- We are almost done setting database and tables so create environment variables in your system as:

  - db_user:"YOUR_DB_USERNAME"
  - db_pass:"USER_PASSWORD"
  - db_name:"DATABASENAME"
  - db_host:"localhost"

- Before running the program, check the required packages are installed with running:
  - "go mod tidy"
- Its time to run our program run the server with:
  - "go run server.go"

You can check the interactive Graphql playground in your localhost at port 8080
