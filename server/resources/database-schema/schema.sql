CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(50),
    role       VARCHAR(50),
    manager_id INT REFERENCES users (id),
    plant_id   INT REFERENCES plants (id)
);

CREATE TABLE plants
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE shifts
(
    id              SERIAL PRIMARY KEY,
    user_id         INT REFERENCES users (id),
    plant_id        INT REFERENCES plants (id),
    machine_details TEXT,
    notes           TEXT,
    shift_time      TIMESTAMP
);
