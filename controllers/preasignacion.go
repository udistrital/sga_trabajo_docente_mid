package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/errorhandler"
)

// PreasignacionController operations for Preasignacion
type PreasignacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *PreasignacionController) URLMapping() {
	c.Mapping("Preasignacion", c.Preasignacion)
	c.Mapping("PreasignacionDocente", c.PreasignacionDocente)
	c.Mapping("Aprobar", c.Aprobar)
}

// Preasignacion ...
// @Title Preasignacion
// @Description Listar todas las preasignaciones
// @Param	vigencia	query 	string	true		"Vigencia de las preasignaciones"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router / [get]
func (c *PreasignacionController) Preasignacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}

// PreasignacionDocente ...
// @Title PreasignacionDocente
// @Description Listar preasignaciones de un docente
// @Param	docente		query 	string	true		"Id docente"
// @Param	vigencia	query 	string	true		"Vigencia de las preasignaciones"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router /docente [get]
func (c *PreasignacionController) PreasignacionDocente() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}

// Aprobar ...
// @Title Aprobar
// @Description Actualizar estado de la aprobación de la preasignación
// @Param   body        body    {}  true        "body Actualizar preasignación plan docente"
// @Success 200 {}
// @Failure 400 The request contains an incorrect data type or an invalid parameter
// @router /aprobar [put]
func (c *PreasignacionController) Aprobar() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}
