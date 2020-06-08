CREATE TABLE IF NOT EXISTS temperatures(
   temperature_id serial PRIMARY KEY,
   temperature NUMERIC,
   location_id integer REFERENCES locations,
   created timestamp default current_timestamp
);