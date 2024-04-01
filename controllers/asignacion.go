package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/errorhandler"
)

// AsignacionController operations for Asignacion
type AsignacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *AsignacionController) URLMapping() {
	c.Mapping("Asignacion", c.Asignacion)
	c.Mapping("AsignacionDocente", c.AsignacionDocente)
}

// Asignacion ...
// @Title Asignacion
// @Description Listar todas las asignaciones de la vigencia determinada
// @Param	vigencia	query 	string	true		"Vigencia de las asignaciones"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router / [get]
func (c *AsignacionController) Asignacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}

// AsignacionDocente ...
// @Title AsignacionDocente
// @Description Listar todas las asignaciones de la vigencia determinada de un docente
// @Param	docente		query 	string	true		"Id docente"
// @Param	vigencia	query 	string	true		"Vigencia de las asignaciones"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router /docente [get]
func (c *AsignacionController) AsignacionDocente() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}
