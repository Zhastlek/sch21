CREATE TABLE time_flights (
      id INTEGER NOT NULL,
      start_flight TEXT not null,
      FOREIGN KEY("id") REFERENCES "bus_flights"("id")
);