BEGIN;


CREATE TABLE IF NOT EXISTS employer (
    id                  SERIAL         PRIMARY KEY,
    first_name          TEXT        NOT NULL, 
    last_name           TEXT        NOT NULL UNIQUE,
    email               TEXT        NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
    id                  INT         PRIMARY KEY,
    first_name          TEXT        NOT NULL, 
    last_name           TEXT        NOT NULL UNIQUE,
    email               TEXT        NOT NULL UNIQUE

);


CREATE TABLE IF NOT EXISTS posting (
    id                  INT         PRIMARY KEY,
    title               TEXT        NOT NULL, 
    company_name        TEXT        NOT NULL UNIQUE,
    company_description TEXT        NOT NULL UNIQUE,
    posting_description TEXT        NOT NULL,
    salary              INT         NOT NULL,
    FOREIGN KEY (company_name) REFERENCES jobs.company(company_name)        

);

CREATE TABLE IF NOT EXISTS applicant_info (
    applicant_id                  INT         NOT NULL,
    resume_location     TEXT        NOT NULL,
    company_name         TEXT        NOT NULL ,
    PRIMARY KEY(applicant_id),

    FOREIGN KEY (applicant_id) REFERENCES jobs.users(id)  
);


COMMIT;