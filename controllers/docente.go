package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/errorhandler"
)

// DocenteController operations for Asignacion
type DocenteController struct {
	beego.Controller
}

// URLMapping ...
func (c *DocenteController) URLMapping() {
	c.Mapping("DocumentoDocenteVinculacion", c.DocumentoDocenteVinculacion)
	c.Mapping("NombreDocenteVinculacion", c.NombreDocenteVinculacion)
}

// DocumentoDocenteVinculacion ...
// @Title DocumentoDocenteVinculacion
// @Description Listar los docentes de acuerdo a la vinculacion y su documento
// @Param	documento		query 	string	true		"Documento docente"
// @Param	vinculacion		query 	int	true			"Id tipo de vinculación"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router /documento [get]
func (c *DocenteController) DocumentoDocenteVinculacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}

// NombreDocenteVinculacion ...
// @Title NombreDocenteVinculacion
// @Description Listar los docentes de acuerdo a la vinculacion y su nombre
// @Param	nombre			query 	string	true		"Nombre docente"
// @Param	vinculacion		query	int	true			"Id tipo de vinculación"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router /nombre [get]
func (c *DocenteController) NombreDocenteVinculacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}
