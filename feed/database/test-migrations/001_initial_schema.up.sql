BEGIN;

DROP TABLE IF EXISTS activity_type CASCADE;
DROP TABLE IF EXISTS workout_type CASCADE;
DROP TABLE IF EXISTS tier CASCADE;
DROP TABLE IF EXISTS workouts CASCADE;


-- Holds the 5 main activity types for all workouts: 
-- Aerobic, Strength, Athletics (Sports), Flexibility, Balance
CREATE TABLE activity_type (
    id          INT         PRIMARY KEY,
    activity_type_description    TEXT    NOT NULL UNIQUE
);

-- Each one is part of an activity type
-- Example: the workout type HIIT and Core are both part of Strength
CREATE TABLE workout_type (
    id      INT     PRIMARY KEY,
    workout_type_description     TEXT   NOT NULL UNIQUE,
    activity_type_description   TEXT    NOT NULL,
    FOREIGN KEY (activity_type_description) REFERENCES  activity_type(activity_type_description)         
);

-- Desccirbes what level of premium subscription a creator's content will be under
-- For now will just hold 'Premium'
CREATE TABLE tier (
    id      INT         PRIMARY KEY,
    tier_description    TEXT        NOT NULL UNIQUE
);

-- Holds the workout objects and all of the data associated with it 
-- including Creator details of who created the workout
CREATE TABLE workouts (
    -- Workout Specific Fields
    id                              SERIAL              PRIMARY KEY,
    workout_name                    TEXT                NOT NULL,
    workout_type                    TEXT                NOT NULL,
    tier                            TEXT                NOT NULL,
    duration                        INT                 NOT NULL,
    created_date                    TIMESTAMP           NOT NULL DEFAULT Now(),
    video_path                      TEXT                NOT NULL,
    workout_description             TEXT                NOT NULL,
    preview_image                   TEXT                NOT NULL,
    difficulty                      TEXT                NOT NULL,
        -- Creator Specific Fields
    email                           TEXT                NOT NULL,
    username                        TEXT                NOT NULL,
    first_name                      TEXT                NOT NULL,
    last_name                       TEXT                NOT NULL,
    creator_profile_picture         TEXT                NOT NULL
);




INSERT INTO activity_type(id, activity_type_description) VALUES(1, 'Aerobic');
INSERT INTO activity_type(id, activity_type_description) VALUES(2, 'Strength');
INSERT INTO activity_type(id, activity_type_description) VALUES(3, 'Sports');
INSERT INTO activity_type(id, activity_type_description) VALUES(4, 'Flexibility');
INSERT INTO activity_type(id, activity_type_description) VALUES(5, 'Balance');

COMMIT;