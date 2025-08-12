-- +goose Up
CREATE TABLE users{
    id UUID PRIMARY KEY,
    created_at Timestamp NOT NULL ,
    updated_at Timestamp NOT NULL,
    name text NOT NULL,
}


-- +goose Down

DRoP TABLE users;