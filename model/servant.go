package model

type Servant struct {
    ID                 int            `json:"id"`
    Name               string         `json:"name"`
    ClassIcon          string         `json:"classIcon"`
    Icon               string         `json:"icon"`
    Portraits          []string       `json:"portraits,omitempty"`
    Skills             []Skill        `json:"skills,omitempty"`
    Appends            []Skill        `json:"appends,omitempty"`
    AscensionMaterials []MaterialList `json:"ascensionMaterials,omitempty"`
    SkillMaterials     []MaterialList `json:"skillMaterials,omitempty"`
    AppendMaterials    []MaterialList `json:"appendMaterials,omitempty"`
}
