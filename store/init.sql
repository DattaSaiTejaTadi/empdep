-- Ensure the pgcrypto extension is enabled for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create the department table with UUID and unique name
CREATE TABLE public.department (
    id UUID  PRIMARY KEY,
    name character varying(255) NOT NULL UNIQUE,
    floor integer
);

-- Create the employee table referencing the department table's UUID id
CREATE TABLE public.employee (
    id UUID PRIMARY KEY,
    name character varying(255) NOT NULL,
    dob date NOT NULL,
    major character varying(255),
    department UUID REFERENCES public.department(id)
);
