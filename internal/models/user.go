package models

type User struct {
    ID              int64       `json:"id"`
    Name            string      `json:"name"`
    Surname         string      `json:"surname"`
    Email           string      `json:"email"`
    Password        string      `json:"password"`
    // Films           []*Film     `json:"fimls"`
    // Login           string      `json:"login"`
}

