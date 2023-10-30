package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/jbillote/fgo-planner-api/constant"
	"github.com/jbillote/fgo-planner-api/model"
	"github.com/labstack/echo/v4"
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
	servantRes, err := atlasServantSearch(query)
	if err != nil {
		c.Logger().Fatal(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var servants []model.Servant
	for _, s := range servantRes {
		servants = append(servants, model.Servant{
			ID:        s.ID,
			Name:      s.Name,
			ClassIcon: fmt.Sprintf(constant.AtlasAcademyClassIcon, classIconFilename(s.Rarity, s.ClassID)),
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
	s, err := atlasGetServant(id)
	if err != nil {
		c.Logger().Fatal(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	portraits := make([]string, 0, len(s.ExtraAssets.CharacterGraph.Ascension))
	keys := getSortedKeys(s.ExtraAssets.CharacterGraph.Ascension)
	for _, v := range keys {
		portraits = append(portraits, s.ExtraAssets.CharacterGraph.Ascension[v])
	}

	skills := processSkills(s.Skills)
	appends := processAppends(s.Appends)
	ascensionMaterials := processMaterialList(s.AscensionMaterials)
	skillMaterials := processMaterialList(s.SkillMaterials)
	appendMaterials := processMaterialList(s.AppendMaterials)

	servant := model.Servant{
		ID:                 s.ID,
		Name:               s.Name,
		ClassIcon:          fmt.Sprintf(constant.AtlasAcademyClassIcon, classIconFilename(s.Rarity, s.ClassID)),
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

func atlasServantSearch(query string) ([]servantResponse, error) {
	url := fmt.Sprintf(constant.AtlasAcademySearch, query)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var servantRes []servantResponse
	err = json.Unmarshal(body, &servantRes)
	if err != nil {
		return nil, err
	}

	return servantRes, nil
}

func atlasGetServant(cid int) (servantResponse, error) {
	var servantRes servantResponse
	url := fmt.Sprintf(constant.AtlasAcademyServantInfo, cid)

	resp, err := http.Get(url)
	if err != nil {
		return servantRes, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return servantRes, err
	}

	err = json.Unmarshal(body, &servantRes)
	if err != nil {
		return servantRes, err
	}

	return servantRes, nil
}

func getSortedKeys[V any](s map[string]V) []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func processMaterialList(ml map[string]materials) []model.MaterialList {
	m := make([]model.MaterialList, 0, len(ml))
	keys := getSortedKeys(ml)
	for _, v := range keys {
		items := make([]model.Material, 0, len(ml[v].Items))
		for _, i := range ml[v].Items {
			items = append(items, model.Material{
				ID:     i.Details.ID,
				Name:   i.Details.Name,
				Icon:   i.Details.Icon,
				Amount: i.Amount,
			})
		}

		m = append(m, model.MaterialList{
			Materials: items,
			QP:        ml[v].QP,
		})
	}

	return m
}

func processSkills(ss []model.Skill) []model.Skill {
	skills := make([]model.Skill, 0, len(ss))
	for _, v := range ss {
		skills = append(skills, model.Skill{
			Name: v.Name,
			Icon: v.Icon,
		})
	}

	return skills
}

func processAppends(as []appendPassive) []model.Skill {
	appends := make([]model.Skill, 0, len(as))
	for _, v := range as {
		appends = append(appends, model.Skill{
			Name: v.Skill.Name,
			Icon: v.Skill.Icon,
		})
	}

	return appends
}

func classIconFilename(r int, cid int) string {
	return fmt.Sprintf("%d_%d", int(math.Max(1, math.Min(3, float64(r)))), cid)
}
