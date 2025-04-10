package models

type User struct {
    ID              int64       `json:"id"`
    Name            string      `json:"name"`
    Surname         string      `json:"surname"`
    Email           string      `json:"email"`
    // Films           []*Film     `json:"fimls"`
    // Login           string      `json:"login"`
    // HashedPassword  string      `json:"hashed_password"`
}

type UserCred struct {
    Login           string      `json:"login"`
    Password        string      `json:"password"`
}

