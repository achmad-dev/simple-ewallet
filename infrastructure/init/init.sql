--- This file is used to create the tables and insert the initial data

--- for storing the users
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    username text NOT NULL UNIQUE,
    password text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);