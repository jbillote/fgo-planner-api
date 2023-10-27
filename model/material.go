package model

type Material struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Icon   string `json:"icon"`
    Amount int    `json:"amount"`
}
