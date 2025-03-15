-- +goose Up
-- +goose StatementBegin
CREATE TABLE notes (
    id varchar(255) NOT NULL,
    title text not null,
    description text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    owner varchar(255) not null,
    PRIMARY KEY (id),
    FOREIGN KEY (owner) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table notes;
-- +goose StatementEnd
