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
)

type servantResponse struct {
    ID                 int                  `json:"id"`
    Name               string               `json:"name"`
    ClassID            int                  `json:"classId"`
    Rarity             int                  `json:"rarity"`
    ExtraAssets        extraAssets          `json:"extraAssets"`
    Skills             []model.Skill        `json:"skills"`
    Appends            []model.Skill        `json:"appendPassive"`
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
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}
