-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    id varchar(255) NOT NULL,
    username text not null,
    password text not null,
    fullname text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table Users;
-- +goose StatementEnd
