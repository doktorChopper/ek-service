package models

type User struct {
    ID              int64       `json:"id"`
    Name            string      `json:"name"`
    Surname         string      `json:"surname"`
    Email           string      `json:"email"`
    // Films           []*Film     `json:"fimls"`
}
