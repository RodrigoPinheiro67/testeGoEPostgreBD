# Go integration with PostgreSQL DB

Creation of a Go system that captures user data such as full name, email and telephone number and sends them to a PostgreSQL database through an HTML form.

>## Instructions

#### Requirements
You will need to use a PostgreSQL driver package. Let's use the pq package to interact with PostgreSQL;

```sh
go get github.com/lib/pq
```
- Make sure you have a PostgreSQL server running and create an appropriate database and table.

- Be sure to replace the values ​​in the host, port, user, password, and dbname constants with the values ​​corresponding to your PostgreSQL environment.

This example uses the PostgreSQL database to store user data. Make sure you create the users table in your database:


```sh
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    phone VARCHAR(20)
);
```

#### Clone the project locally:

```sh
$ git clone github.com/RodrigoPinheiro67/testeGoEPostgreBD.git
```
