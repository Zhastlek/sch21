CREATE TABLE waiting_time (
  id INTEGER NOT NULL,
  station_name TEXT NOT NULL,
  FOREIGN KEY("id") REFERENCES "distance_and_time"("id")
);