package api

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nano2nano/valorant_tips/internal/cloud"
	"github.com/nano2nano/valorant_tips/internal/image"
	"github.com/nano2nano/valorant_tips/internal/model"
	"github.com/olahol/go-imageupload"
	"gorm.io/gorm"
)

func Status() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "api is working")
	}
}

func GetAgent() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		tx := c.Get("Tx").(*gorm.DB)

		agent := new(model.Agent)
		if err := agent.Load(tx, uint(id)); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, agent)
	}
}

func GetAgents() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)

		agents := new(model.Agents)
		if err := agents.Load(tx); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, agents)
	}
}

func GetAbility() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		tx := c.Get("Tx").(*gorm.DB)

		ability := new(model.Ability)
		if err := ability.Load(tx, uint(id)); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, ability)
	}
}

func GetAbilities() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)

		abilities := new(model.Abilities)
		if err := abilities.Load(tx); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, abilities)
	}
}

func GetMap() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		tx := c.Get("Tx").(*gorm.DB)

		m := new(model.Map)
		if err := m.Load(tx, uint(id)); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, m)
	}
}

func PostMap() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		m := model.Map{Name: c.FormValue("name")}
		tx := c.Get("Tx").(*gorm.DB)
		if err := m.Save(tx); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, m)
	}
}

func GetMaps() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)

		m := new(model.Maps)
		if err := m.Load(tx); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, m)
	}
}

func GetTip() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		tx := c.Get("Tx").(*gorm.DB)

		t := new(model.Tip)
		if err := t.Load(tx, uint(id)); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, t)
	}
}

func PostTip() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		img, err := imageupload.Process(c.Request(), "stand_img")
		if err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}
		f_name_stand, err := image.SaveImage(img)
		if err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}

		img, err = imageupload.Process(c.Request(), "aim_img")
		if err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}
		f_name_aim, err := image.SaveImage(img)
		if err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}

		side_id, err := strconv.ParseInt(c.FormValue("side_id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		map_id, err := strconv.ParseInt(c.FormValue("map_id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		ability_id, err := strconv.ParseInt(c.FormValue("ability_id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		title := c.FormValue("title")
		description := c.FormValue("description")
		t := model.Tip{StandingPosition: f_name_stand, AimPosition: f_name_aim, SideID: uint(side_id), MapID: uint(map_id), AbilityID: uint(ability_id), Description: description, Title: title}
		tx := c.Get("Tx").(*gorm.DB)

		if err := t.Save(tx); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, t)
	}
}

func GetTips() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)

		cond := new(model.Tip)
		map_id := c.QueryParam("map_id")
		if len(map_id) != 0 {
			map_id, err := strconv.ParseInt(map_id, 0, 64)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid map_id.")
			}
			cond.MapID = uint(map_id)
		}

		side_id := c.QueryParam("side_id")
		if len(side_id) != 0 {
			side_id, err := strconv.ParseInt(side_id, 0, 64)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid side_id.")
			}
			cond.SideID = uint(side_id)
		}

		ability_id := c.QueryParam("ability_id")
		if len(ability_id) != 0 {
			ability_id, err := strconv.ParseInt(ability_id, 0, 64)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid ability_id.")
			}
			cond.AbilityID = uint(ability_id)
		}

		tips := new(model.Tips)
		if err := tx.Where(cond).Preload("Side").Preload("Map").Preload("Ability").Find(tips).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, tips)
	}
}

func GetSides() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)

		sides := new(model.Sides)
		if err := sides.Load(tx); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, sides)
	}
}

func PostImg() echo.HandlerFunc {
	return func(c echo.Context) error {
		img, err := imageupload.Process(c.Request(), "file")
		if img.ContentType != "image/jpeg" {
			return c.String(http.StatusBadRequest, "only 'png' image")
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		thumb, err := imageupload.ThumbnailPNG(img, 896, 504)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		f_name := fmt.Sprintf("%s.jpeg", time.Now().Format("20060102150405"))
		if err := cloud.Upload(f_name, bytes.NewReader(thumb.Data)); err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}

		return c.String(http.StatusOK, f_name)
	}
}

func GetImg() echo.HandlerFunc {
	return func(c echo.Context) error {
		f_name := c.Param("name")
		bs, err := cloud.Download(f_name)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		i := &imageupload.Image{
			Filename:    f_name,
			ContentType: "image/jpeg",
			Data:        bs,
			Size:        len(bs),
		}
		i.Write(c.Response().Writer)

		return nil
	}
}
