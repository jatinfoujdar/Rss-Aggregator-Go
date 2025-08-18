-- +goose Up
ALTER table users COLOUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')
) 



-- -goose Down
ALTER table users DROP COLUMN api_key;