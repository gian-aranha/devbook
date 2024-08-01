# devbook

Devbook is a project that simulates a social network. It has two main parts, an Rest API that runs the program logic, and an CLI that makes possible the execution of all the functions of the project.

First, to be able to test all the functionalities, you need to configure the database to already have some data. For that I create two files, `sql.sql` (creates the database and it's tables) and `data.sql` (add some data to the tables that were just created). Both files can be found in `devbook/api/sql`

The database used in this project is MySQL, so you'll have to install it. You also need to create an user and a password for this user. Don't forget to change the first to lines of the file `devbook/api/.env`, because, otherwise the database won't work.

```bash
DB_USER=<your-username>
DB_PASSWORD=<your-password>
```

To run the API is simple, you can go to `/devbook/api/src/` directory and run the command bellow.

```bash
go run main.go
```

For now the API can be tested from Postman, Insomnia or any program you prefer. The CLI is not ready yet, but will be developed the as quickly as possible.
