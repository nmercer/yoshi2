CREATE TABLE IF NOT EXISTS locations(
   location_id serial PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL
);