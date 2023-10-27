package model

type MaterialList struct {
    Materials []Material `json:"materials"`
    Qp        int        `json:"qp"`
}
