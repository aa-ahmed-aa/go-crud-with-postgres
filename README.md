# GOWT
Sample crud web application project using Golang(http, templates, os, sql), Bootstrap 4, DataTables, MySQL, Docker.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software

* Golang, preferably the latest version (1.16).
* MySQL Database
* Docker (optional)

### Installing

1. Clone this repository

```
git clone https://github.com/le4ndro/gowt.git
cd gowt
```

3. Create database on Postgress and insert some data

```
CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  first_name TEXT,
  last_name TEXT,
  birthday DATE,
  gender INT,
  email TEXT,
  address TEXT
);

INSERT INTO "customers" ("first_name", "last_name", "birthday", "gender", "email", "address")
VALUES ('first1', 'last1', '1995-01-01', '0', '1@1.com', '1 tallin estonia');

INSERT INTO "customers" ("first_name", "last_name", "birthday", "gender", "email", "address")
VALUES ('first2', 'last2', '1995-02-02', '1', '2@2.com', '2 tallin estonia');

INSERT INTO "customers" ("first_name", "last_name", "birthday", "gender", "email", "address")
VALUES ('first3', 'last3', '1995-03-03', '0', '3@3.com', '3 tallin estonia');
```
