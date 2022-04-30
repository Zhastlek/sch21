CREATE TABLE bus_drivers (
     id INTEGER NOT NULL,
     drver_name TEXT NOT NULL,
     FOREIGN KEY("id") REFERENCES "buses"("id")
);