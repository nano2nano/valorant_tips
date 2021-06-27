package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nano2nano/valorant_tips/internal/model"
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
		sp_path := c.FormValue("sp_path")
		ap_path := c.FormValue("ap_path")
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
		t := model.Tip{StandingPosition: sp_path, AimPosition: ap_path, SideID: uint(side_id), MapID: uint(map_id), AbilityID: uint(ability_id)}
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
