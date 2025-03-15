-- name: CreateUser :exec
insert into users (id, username, password, fullname, created_at, updated_at) values(?, ?, ?, ?, current_timestamp, current_timestamp);

-- name: GetByUsername :one
select * from users where username = ?;

-- name: GetUserByID :one
select * from users where id = ?;
