CREATE TABLE driving_period (
    id INTEGER NOT NULL,
    begining_period TIMESTAMP NOT NULL,
    end_of_period TIMESTAMP NOT NULL,
    FOREIGN KEY("id") REFERENCES "bus_flights"("id")
);