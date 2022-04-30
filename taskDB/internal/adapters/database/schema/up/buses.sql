CREATE TABLE buses (
   id INTEGER NOT NULL UNIQUE,
   bus_model TEXT NOT NULL,
   number_bus TEXT NOT NULL,
   flight_id INTEGER,
   FOREIGN KEY("flight_id") REFERENCES "bus_flights"("id")
);