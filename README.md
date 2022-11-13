# Gotodos

# How to run the program

- In the begining download and install golang
  link: https://go.dev/dl/
- Next thing is postgresql, download and install it too
  link: https://go.dev/dl/
- Run "sudo -u postgres psql"
  - You would get postgres cli
  - Now, you can create new database with:
    "CREATE DATABASE "YOUR_DATABASE_NAME;"
  - Change user password if you don't know the current user password with:
    "ALTER USER user_name WITH PASSWORD 'new_password';"
  - You can check the username if you don't know the current username with:
    "\l"
  - Now, Its time for creating our tables, run the following queries:
    "CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    password VARCHAR ( 50 ) NOT NULL,
    );"
    CREATE TABLE todos(
    id serial PRIMARY KEY,
    title VARCHAR(25) NOT NULL,
    description VARCHAR(255),
    completed BOOLEAN NOT NULL,
    user_id int NOT NULL,
    CONSTRAINT fk_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    );
- We are almost done setting database and tables so create environment variables in your system as:
  - db_user:"YOUR_DB_USERNAME"
    db_pass:"USER_PASSWORD"
    db_name:"DATABASENAME"
    db_host:"localhost"
- Before running the program, check the required packages are installed with running:
  "go mod tidy"
- Its time to run our program run the server with:
  "go run server.go"

You can check the interactive Graphql playground in your localhost at port 8080
