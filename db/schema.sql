-- Operators
-- Schema for postgresql store
-- -------------------------

-- Operators table
CREATE TABLE IF NOT EXISTS operators (
    id          TEXT NOT NULL PRIMARY KEY,
    name        TEXT NOT NULL, 
    valid       BOOLEAN NOT NULL,
    created     TIMESTAMP NOT NULL     
);