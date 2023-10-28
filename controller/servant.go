package controller

import (
    "encoding/json"
    "fmt"
    "github.com/jbillote/fgo-planner-api/constant"
    "github.com/jbillote/fgo-planner-api/model"
    "github.com/labstack/echo/v4"
    "io"
    "math"
    "net/http"
    "sort"
    "strconv"
)

type servantResponse struct {
    ID                 int                  `json:"id"`
    Name               string               `json:"name"`
    ClassID            int                  `json:"classId"`
    Rarity             int                  `json:"rarity"`
    ExtraAssets        extraAssets          `json:"extraAssets"`
    Skills             []model.Skill        `json:"skills"`
    Appends            []appendPassive      `json:"appendPassive"`
    AscensionMaterials map[string]materials `json:"ascensionMaterials"`
    SkillMaterials     map[string]materials `json:"skillMaterials"`
    AppendMaterials    map[string]materials `json:"appendSkillMaterials"`
}

type extraAssets struct {
    CharacterGraph characterImages `json:"charaGraph"`
    Faces          characterImages `json:"faces"`
}

type characterImages struct {
    Ascension map[string]string `json:"ascension"`
}

type appendPassive struct {
    Skill model.Skill `json:"skill"`
}

type materials struct {
    Items []item `json:"items"`
    QP    int    `json:"qp"`
}

type item struct {
    Details itemDetails `json:"item"`
    Amount  int         `json:"amount"`
}

type itemDetails struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Icon string `json:"icon"`
}

func SearchServant(c echo.Context) (err error) {
    query := c.QueryParam("query")
    uri := fmt.Sprintf(constant.AtlasAcademySearch, query)

    resp, err := http.Get(uri)
    if err != nil {
        c.Logger().Fatal(err)
        return c.String(http.StatusInternalServerError, err.Error())
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        c.Logger().Fatal(err)
        return c.String(http.StatusInternalServerError, err.Error())
    }

    var servantRes []servantResponse
    err = json.Unmarshal(body, &servantRes)
    if err != nil {
        c.Logger().Fatal(err)
        return c.String(http.StatusInternalServerError, err.Error())
    }

    var servants []model.Servant
    for _, s := range servantRes {
        servants = append(servants, model.Servant{
            ID:        s.ID,
            Name:      s.Name,
            ClassIcon: fmt.Sprintf(constant.AtlasAcademyClassIcon, int(math.Min(3, float64(s.Rarity))), s.ClassID),
            Icon:      s.ExtraAssets.Faces.Ascension["1"],
        })
    }

    return c.JSON(http.StatusOK, servants)
}

func GetServant(c echo.Context) (err error) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.Logger().Fatal(err)
        return c.String(http.StatusInternalServerError, err.Error())
    }
    uri := fmt.Sprintf(constant.AtlasAcademyServantInfo, id)

    resp, err := http.Get(uri)
    if err != nil {
        c.Logger().Fatal(err)
        return c.String(http.StatusInternalServerError, err.Error())
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        c.Logger().Fatal(err)
        return c.String(http.StatusInternalServerError, err.Error())
    }

    var s servantResponse
    err = json.Unmarshal(body, &s)
    if err != nil {
        c.Logger().Fatal(err)
        return c.String(http.StatusInternalServerError, err.Error())
    }

    portraits := make([]string, 0, len(s.ExtraAssets.CharacterGraph.Ascension))
    keys := make([]string, 0, len(s.ExtraAssets.CharacterGraph.Ascension))
    for k := range s.ExtraAssets.CharacterGraph.Ascension {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, v := range keys {
        portraits = append(portraits, s.ExtraAssets.CharacterGraph.Ascension[v])
    }

    skills := make([]model.Skill, 0, len(s.Skills))
    for _, v := range s.Skills {
        skills = append(skills, model.Skill{
            Name: v.Name,
            Icon: v.Icon,
        })
    }

    appends := make([]model.Skill, 0, len(s.Appends))
    for _, v := range s.Appends {
        appends = append(appends, model.Skill{
            Name: v.Skill.Name,
            Icon: v.Skill.Icon,
        })
    }

    ascensionMaterials := make([]model.MaterialList, 0, len(s.AscensionMaterials))
    keys = make([]string, 0, len(s.AscensionMaterials))
    for k := range s.AscensionMaterials {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, v := range keys {
        items := make([]model.Material, 0, len(s.AscensionMaterials[v].Items))
        for _, i := range s.AscensionMaterials[v].Items {
            items = append(items, model.Material{
                ID:     i.Details.ID,
                Name:   i.Details.Name,
                Icon:   i.Details.Icon,
                Amount: i.Amount,
            })
        }

        ascensionMaterials = append(ascensionMaterials, model.MaterialList{
            Materials: items,
            QP:        s.AscensionMaterials[v].QP,
        })
    }

    skillMaterials := make([]model.MaterialList, 0, len(s.SkillMaterials))
    keys = make([]string, 0, len(s.SkillMaterials))
    for k := range s.SkillMaterials {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, v := range keys {
        items := make([]model.Material, 0, len(s.SkillMaterials[v].Items))
        for _, i := range s.SkillMaterials[v].Items {
            items = append(items, model.Material{
                ID:     i.Details.ID,
                Name:   i.Details.Name,
                Icon:   i.Details.Icon,
                Amount: i.Amount,
            })
        }

        skillMaterials = append(skillMaterials, model.MaterialList{
            Materials: items,
            QP:        s.SkillMaterials[v].QP,
        })
    }

    appendMaterials := make([]model.MaterialList, 0, len(s.AppendMaterials))
    keys = make([]string, 0, len(s.AppendMaterials))
    for k := range s.SkillMaterials {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, v := range keys {
        items := make([]model.Material, 0, len(s.AppendMaterials[v].Items))
        for _, i := range s.AppendMaterials[v].Items {
            items = append(items, model.Material{
                ID:     i.Details.ID,
                Name:   i.Details.Name,
                Icon:   i.Details.Icon,
                Amount: i.Amount,
            })
        }

        appendMaterials = append(appendMaterials, model.MaterialList{
            Materials: items,
            QP:        s.AppendMaterials[v].QP,
        })
    }

    servant := model.Servant{
        ID:                 s.ID,
        Name:               s.Name,
        ClassIcon:          fmt.Sprintf(constant.AtlasAcademyClassIcon, int(math.Min(3, float64(s.Rarity))), s.ClassID),
        Icon:               s.ExtraAssets.Faces.Ascension["1"],
        Portraits:          portraits,
        Skills:             skills,
        Appends:            appends,
        AscensionMaterials: ascensionMaterials,
        SkillMaterials:     skillMaterials,
        AppendMaterials:    appendMaterials,
    }

    return c.JSON(http.StatusOK, servant)
}
