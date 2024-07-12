CREATE TABLE users (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  email text NOT NULL,
  password text NOT NULL
);

CREATE TABLE masters (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  userID integer,
  name text NOT NULL,
  category text NOT NULL,
  point integer
);

CREATE TABLE tasks (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  userID integer,
  masterID integer,
  person text NOT NULL,
  date date
);