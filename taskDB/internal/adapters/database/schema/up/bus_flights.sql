CREATE TABLE bus_flights (
     id INTEGER NOT NULL UNIQUE,
     departure_city TEXT NOT NULL,
     arrival_city TEXT NOT NULL,
     distance TEXT NOT NULL,
     travel_time TEXT NOT NULL,
     PRIMARY KEY("id" AUTOINCREMENT)
);