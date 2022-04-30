CREATE TABLE distance_and_time (
   id INTEGER NOT NULL,
   distance TEXT NOT NULL,
   FOREIGN KEY("id") REFERENCES "bus_flights"("id")
);