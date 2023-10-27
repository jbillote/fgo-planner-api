package model

type Servant struct {
    ID                 string         `json:"id"`
    Name               string         `json:"name"`
    ClassIcon          string         `json:"classIcon"`
    Icon               string         `json:"icon"`
    Portraits          []string       `json:"portraits"`
    Skills             []Skill        `json:"skills"`
    Appends            []Skill        `json:"appends"`
    AscensionMaterials []MaterialList `json:"ascensionMaterials"`
    SkillMaterials     []MaterialList `json:"skillMaterials"`
    AppendMaterials    []MaterialList `json:"appendMaterials"`
}
