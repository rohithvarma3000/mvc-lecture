# mvc-lecture
- Clone the repo. From the root directory:
```
- go mod vendor
- go mod tidy
```


- MYSQL:
1. `mysql -u root -p` : and enter password
2. Create a new database 'books': `CREATE DATABASE books;`
3. Connect to the database: `USE books;`
4. Import the sql dump file into your database: `mysql -u root -p books < dump.sql`


- Running the server:
1. `go build -o mvc ./cmd/main.go`
2.  Run the binary file: `./mvc`