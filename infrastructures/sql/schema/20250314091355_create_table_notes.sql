-- +goose Up
-- +goose StatementBegin
CREATE TABLE Notes (
    id varchar(255) NOT NULL,
    title text not null,
    description text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    owner varchar(255) not null,
    PRIMARY KEY (id),
    FOREIGN KEY (owner) REFERENCES Users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table Notes;
-- +goose StatementEnd
