CREATE TABLE users (
    Id varchar(256) NOT NULL PRIMARY KEY,
    Data json NOT NULL,
    AuthMethod varchar(64) NOT NULL,
    AuthData json NOT NULL
);