CREATE SCHEMA IF NOT EXISTS jobs;

CREATE TABLE IF NOT EXISTS jobs.users (
    id                  SERIAL         PRIMARY KEY,
    first_name          TEXT        NOT NULL, 
    last_name           TEXT        NOT NULL UNIQUE,
    email               TEXT        NOT NULL UNIQUE,

);
CREATE TABLE IF NOT EXISTS jobs.employer (
    id                  INT         PRIMARY KEY,
    title               TEXT        NOT NULL, 
    company_name        TEXT        NOT NULL UNIQUE,
    company_description TEXT        NOT NULL UNIQUE,
);


CREATE TABLE IF NOT EXISTS jobs.posting (
    id                  INT         PRIMARY KEY,
    title               TEXT        NOT NULL, 
    company_name        TEXT        NOT NULL UNIQUE,
    company_description TEXT        NOT NULL UNIQUE,
    posting_description TEXT        NOT NULL,
    salary              INT         NOT NULL,
    FOREIGN KEY (company_name) REFERENCES jobs.company(company_name)        

);

CREATE TABLE IF NOT EXISTS jobs.applicant_info (
    applicant_id                  INT         NOT NULL,
    resume_location     TEXT        NOT NULL,
    company_name         TEXT        NOT NULL ,
    PRIMARY KEY(applicant_id),

    FOREIGN KEY (applicant_id) REFERENCES jobs.users(id)         

);
