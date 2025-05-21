CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    tg TEXT NOT NULL UNIQUE,
    office TEXT NOT NULL,
    emoji TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS lunches (
    id BIGSERIAL PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    time TIMESTAMP NOT NULL,
    place TEXT NOT NULL,
    description TEXT,
    participants BIGINT[] NOT NULL,
    number_of_participants BIGINT NOT NULL,
    liked_by BIGINT[] NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES users(id)
);

CREATE OR REPLACE FUNCTION check_participants_exist()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM unnest(NEW.participants) AS p
        WHERE p NOT IN (SELECT id FROM users)
    ) THEN
        RAISE EXCEPTION 'One or more participants do not exist in the users table';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_check_participants
BEFORE INSERT OR UPDATE ON lunches
FOR EACH ROW
EXECUTE FUNCTION check_participants_exist();

-- Function to update number_of_participants based on participants array length
CREATE OR REPLACE FUNCTION update_number_of_participants()
RETURNS TRIGGER AS $$
BEGIN
    NEW.number_of_participants := cardinality(NEW.participants);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call the update function before insert or update
CREATE TRIGGER trg_update_number_of_participants
BEFORE INSERT OR UPDATE ON lunches
FOR EACH ROW
EXECUTE FUNCTION update_number_of_participants();

-- Function to check for unique participants in the array
CREATE OR REPLACE FUNCTION check_unique_participants()
RETURNS TRIGGER AS $$
BEGIN
    IF (
        SELECT COUNT(*)
        FROM unnest(NEW.participants) AS p
        GROUP BY p
        HAVING COUNT(*) > 1
    ) > 0 THEN
        RAISE EXCEPTION 'Duplicate userIDs found in participants array';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call the uniqueness check function before insert or update
CREATE TRIGGER trg_check_unique_participants
BEFORE INSERT OR UPDATE ON lunches
FOR EACH ROW
EXECUTE FUNCTION check_unique_participants();


CREATE OR REPLACE FUNCTION check_liked_by_exist()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM unnest(NEW.liked_by) AS p
        WHERE p NOT IN (SELECT id FROM users)
    ) THEN
        RAISE EXCEPTION 'One or more liked_by do not exist in the users table';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_check_liked_by
BEFORE INSERT OR UPDATE ON lunches
FOR EACH ROW
EXECUTE FUNCTION check_liked_by_exist();

-- Function to check for unique liked_by in the array
CREATE OR REPLACE FUNCTION check_unique_liked_by()
RETURNS TRIGGER AS $$
BEGIN
    IF (
        SELECT COUNT(*)
        FROM unnest(NEW.liked_by) AS p
        GROUP BY p
        HAVING COUNT(*) > 1
    ) > 0 THEN
        RAISE EXCEPTION 'Duplicate userIDs found in liked_by array';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call the uniqueness check function before insert or update
CREATE TRIGGER trg_check_unique_liked_by
BEFORE INSERT OR UPDATE ON lunches
FOR EACH ROW
EXECUTE FUNCTION check_unique_liked_by();
