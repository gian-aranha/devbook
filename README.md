# DevBook 

Devbook is a project that simulates a social network. It has two main parts, a Rest API and a WEB interface. 

First, to be able to test all the functionalities, you need to configure the database to already have some data. For that I created two files, `sql.sql` (creates the database and it's tables) and `data.sql` (add some data to the tables that were just created). Both files can be found in `devbook/api/sql`

The database used in this project is MySQL, so you'll have to install it. You also need to create an user and a password for this user. Don't forget to change the first two lines of the file `devbook/api/.env`, because, otherwise the database won't work.

```bash
DB_USER=<your-username>
DB_PASSWORD=<your-password>
```

To run the API is simple, you can go to `/devbook/api/src/` directory and run the command bellow.

```bash
go run main.go
```

Or, if you prefer you can build the binaries and execute the project.

```bash
    # build
    go build .

    # execute
    ./api
```

For now, the API can be tested from Postman, Insomnia or any program you prefer. The WEB interface (front-end) is being developed, and will be available as soon as possible.
