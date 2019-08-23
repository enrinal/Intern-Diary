# Simple Order System Management

## Description
This is an example of implementation of Clean Architecture in Go (Golang) projects. But with some case :)

Rule of Clean Architecture by Uncle Bob

* Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
* Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
* Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
* Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
* Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

This project has 4 Domain layer :

* Models Layer
* Repository Layer
* Usecase Layer
* Delivery Layer

#### The Diagram
![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

#### The Case 
* There are 3 roles for users, ordinary buyers, subscription buyers, and sellers
* Regular buyers can only buy a maximum of 5 times a day
* Subscribers can buy more than 5 times a day
* Buyers can buy goods and buy a maximum of 10 items per product
* Buyers can enter into the Basket before processing the purchase
* The buyer processes the purchase of the order
* Buyer will see pending purchase status
* The seller will process the order and the purchase status is processed
* Send goods and purchase status to be sent
* Buyers can see the status of the purchase whether it is already in process, has been sent or is still pending
* Buyers can change the status to be accepted or automatically changed if it passes 1 day.

This project using Mysql and MongoDb

### How To Run This Project
>  Make Sure you have run the simpleorder.sql in your mysql

```bash
# Clone into YOUR $GOPATH/src
git clone https://gitlab.warungpintar.co/enrinal/intern-diary.git

#move to project
cd Simple-Order-System-Management-Golang

# Test the code
make test

# Build Project
go build

# Run Project
./simple-order
```
