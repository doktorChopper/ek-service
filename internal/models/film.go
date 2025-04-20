package models

type Film struct {
    ID      int64   `json:"id"`
    Name    string  `json:"name"`
    Genre   string  `json:"genre"`
    Review  string  `json:"review"`
    Rate    int64   `json:"rate"`
    UserID  int64   `json:"user_id"`
}
