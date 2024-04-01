package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/errorhandler"
)

// EspacioAcademicoController operations for Asignacion
type EspacioAcademicoController struct {
	beego.Controller
}

// URLMapping ...
func (c *EspacioAcademicoController) URLMapping() {
	c.Mapping("GrupoEspacioAcademico", c.GrupoEspacioAcademico)
	c.Mapping("GrupoEspacioAcademicoPadre", c.GrupoEspacioAcademicoPadre)
}

// GrupoEspacioAcademico ...
// @Title GrupoEspacioAcademico
// @Description  Lista los grupos de un espacios académico padre por vigencia
// @Param	padre		query 	string	true		"Id del espacio académico padre"
// @Param	vigencia	query 	string	true		"Vigencia del espacio académico"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router /grupo [get]
func (c *EspacioAcademicoController) GrupoEspacioAcademico() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}

// GrupoEspacioAcademicoPadre ...
// @Title GrupoEspacioAcademicoPadre
// @Description Lista los grupos de un espacios académico padre
// @Param	padre		query 	string	true		"Id del espacio académico padre"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router /grupo-padre [get]
func (c *EspacioAcademicoController) GrupoEspacioAcademicoPadre() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}
