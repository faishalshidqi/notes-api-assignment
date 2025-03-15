package database

import "assignment/domains"

func (u *User) ToDomainsUser() domains.User {
	return domains.User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		FullName:  u.Fullname,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (n *Note) ToDomainsNote() domains.Note {
	return domains.Note{
		ID:        n.ID,
		Title:     n.Title,
		Body:      n.Description,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
		Owner:     n.Owner,
	}
}
