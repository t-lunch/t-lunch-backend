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


CREATE TABLE IF NOT EXISTS histories (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    lunch_id BIGINT NOT NULL,
    is_liked BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (lunch_id) REFERENCES lunches(id)
);
