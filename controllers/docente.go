package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sga_plan_trabajo_docente_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	requestmanager "github.com/udistrital/utils_oas/requestresponse"
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

	documento := c.GetString("documento")
	vinculacion, errvin := c.GetInt64("vinculacion")

	if errvin != nil {
		logs.Error(errvin)
		c.Data["json"] = requestmanager.APIResponseDTO(false, 400, nil, "Error: Parámetro(s) no válido(s) o faltante(s)")
		c.Ctx.Output.SetStatus(400)
	} else if documento == "" || vinculacion <= 0 {
		logs.Error(documento, vinculacion)
		c.Data["json"] = requestmanager.APIResponseDTO(false, 400, nil, "Error: Parámetro(s) con valores no válidos")
		c.Ctx.Output.SetStatus(400)
	} else {
		resultado := services.ListaDocentesxDocumentoVinculacion(documento, vinculacion)
		c.Data["json"] = resultado
		c.Ctx.Output.SetStatus(resultado.Status)
	}

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

	nombre := c.GetString("nombre")
	vinculacion, errvin := c.GetInt64("vinculacion")

	if errvin != nil {
		logs.Error(errvin)
		c.Data["json"] = requestmanager.APIResponseDTO(false, 400, nil, "Error: Parámetro(s) no válido(s) o faltante(s)")
		c.Ctx.Output.SetStatus(400)
	} else if nombre == "" || vinculacion <= 0 {
		logs.Error(nombre, vinculacion)
		c.Data["json"] = requestmanager.APIResponseDTO(false, 400, nil, "Error: Parámetro(s) con valores no válidos")
		c.Ctx.Output.SetStatus(400)
	} else {
		resultado := services.ListaDocentesxNombreVinculacion(nombre, vinculacion)
		c.Data["json"] = resultado
		c.Ctx.Output.SetStatus(resultado.Status)
	}

	c.ServeJSON()
}
