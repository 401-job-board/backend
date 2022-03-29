CREATE SCHEMA IF NOT EXISTS jobs;

CREATE TABLE IF NOT EXISTS jobs.posting (
    id                  INT         PRIMARY KEY,
    title               TEXT        NOT NULL, 
    company_name        TEXT        NOT NULL UNIQUE,
    company_description TEXT        NOT NULL UNIQUE,
    posting_description TEXT        NOT NULL,
    salary              INT         NOT NULL
);