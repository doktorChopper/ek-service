package models

type User struct {
    ID              int         `json:"id"`
    Name            string      `json:"name"`
    Surname         string      `json:"surname"`
    Email           string      `json:"email"`
    Password        string      `json:"-"`
    Films           []*Film      `json:"fimls"`
    // Login           string      `json:"login"`
}

