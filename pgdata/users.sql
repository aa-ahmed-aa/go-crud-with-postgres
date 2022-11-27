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
