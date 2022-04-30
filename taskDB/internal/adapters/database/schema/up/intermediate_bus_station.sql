CREATE TABLE intermediate_bus_station (
      id INTEGER NOT NULL,
      station TEXT NOT NULL,
      number_station INTEGER,
      arrival_time TEXT NOT NULL,
      departure_time TEXT NOT NULL,
      FOREIGN KEY("id") REFERENCES "bus_flights"("id")
);