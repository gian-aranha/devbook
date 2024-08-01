# devbook

Devbook is a project that simulates a social network. It has two main parts, an Rest API that runs the program logic, and an CLI that makes possible the execute all the functions from the project.

First, to be able to test all the functionalities, you need to configure the database to already have some data. For that I create two files, `sql.sql` (creates the database and it's tables) and `data.sql` (add some data to the tables that were just created). Both files can be found in `devbook/api/sql`

The database that's used in this project is MySQL, so you'll have to install it. Also you need to create an user and a password for this user. Don't forget to change the first to lines of the file `devbook/api/.env`, because, otherwise the database won't work.

```bash
DB_USER=<your-username>
DB_PASSWORD=<your-password>
```

To run the API is simple, you can go to `/devbook/api/src/` directory and build the binaries from the application. For that just run the command bellow.

```bash
go run main.go
```
