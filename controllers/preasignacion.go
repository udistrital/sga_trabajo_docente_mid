package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sga_plan_trabajo_docente_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	requestmanager "github.com/udistrital/utils_oas/requestresponse"
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

	vigencia := c.GetString("vigencia")

	if vigencia == "" {
		logs.Error(vigencia)
		c.Data["json"] = requestmanager.APIResponseDTO(false, 400, nil, "Error: Parametro(s) Valores nó validos")
	} else {
		resultado := services.ListaPreasignacion(vigencia)
		c.Data["json"] = resultado
		c.Ctx.Output.SetStatus(resultado.Status)
	}

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

	docente := c.GetString("docente")
	vigencia := c.GetString("vigencia")

	if docente == "" || vigencia == "" {
		logs.Error(docente, vigencia)
		c.Data["json"] = requestmanager.APIResponseDTO(false, 400, nil, "Error: Parametro(s) Valores nó validos")
	} else {
		resultado := services.ListaPreasignacionDocente(docente, vigencia)
		c.Data["json"] = resultado
		c.Ctx.Output.SetStatus(resultado.Status)
	}

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

	var body map[string]interface{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = requestmanager.APIResponseDTO(false, 400, nil, "Error: Parametro(s) nó valido(s) o faltante(s)")
		c.Ctx.Output.SetStatus(400)
	}

	params := []string{"preasignaciones", "no-preasignaciones", "docente"}
	errParam := false
	for _, param := range params {
		if _, ok := body[param]; !ok {
			errParam = true
			logs.Error("No existe el parametro %s", param)
			break
		}
	}
	if errParam {
		c.Data["json"] = requestmanager.APIResponseDTO(false, 400, nil, "Error: Parametro(s) nó valido(s) o faltante(s)")
		c.Ctx.Output.SetStatus(400)
	} else {
		resultado := services.DefinePreasignacion(body)
		c.Data["json"] = resultado
		c.Ctx.Output.SetStatus(resultado.Status)
	}

	c.ServeJSON()
}
