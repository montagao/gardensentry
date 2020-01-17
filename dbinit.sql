/*
CREATE TABLE events (
  id serial PRIMARY KEY,
  description VARCHAR(200),
  occured TIMESTAMP NOT NULL,
  type VARCHAR(200),
  vidurl VARCHAR(200)
);
*/

insert into events (description, occured, type, vidurl)
values ('tests', now(), 'test', 'ass.com');

