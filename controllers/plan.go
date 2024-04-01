package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/errorhandler"
)

// PlanController operations for Asignacion
type PlanController struct {
	beego.Controller
}

// URLMapping ...
func (c *PlanController) URLMapping() {
	c.Mapping("DefinePlanTrabajoDocente", c.DefinePlanTrabajoDocente)
	c.Mapping("PlanTrabajoDocenteAsignacion", c.PlanTrabajoDocenteAsignacion)
	c.Mapping("CopiarPlanTrabajoDocente", c.CopiarPlanTrabajoDocente)
	c.Mapping("PlanPreaprobado", c.PlanPreaprobado)
}

// DefinePlanTrabajoDocente ...
// @Title DefinePlanTrabajoDocente
// @Description Actualiza la información de los planes de trabajo
// @Param   body        body    {}  true        "body plan a actualizar"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router / [put]
func (c *PlanController) DefinePlanTrabajoDocente() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}

// PlanTrabajoDocenteAsignacion ...
// @Title PlanTrabajoDocenteAsignacion
// @Description Traer la información de las asignaciones de un docente en la vigencia determinada
// @Param	docente		query 	string	true		"Id docente"
// @Param	vigencia	query 	string	true		"Vigencia de las asignaciones"
// @Param	vinculacion	query 	string	true		"Id vinculacion"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router / [get]
func (c *PlanController) PlanTrabajoDocenteAsignacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}

// CopiarPlanTrabajoDocente ...
// @Title CopiarPlanTrabajoDocente
// @Description Copia de plan de trabajo docente de una vigencia anterior
// @Param   body        body    {}  true        "body plan a copiar {docente,vigenciaAnterior,vigencia,vinculacion,carga(1,2) all number}"
// @Success 201 {}
// @Failure 400 The request contains an incorrect data type or an invalid parameter
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router /copiar [post]
func (c *PlanController) CopiarPlanTrabajoDocente() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}

// PlanPreaprobado ...
// @Title PlanPreaprobado
// @Description Listar planes que han sido aprobados en asignar ptd
// @Param	vigencia		query 	int	true	"Id periodo"
// @Param	proyecto		query 	int	true	"Id proyecto"
// @Success 200 {}
// @Failure 404 he request contains an incorrect parameter or no record exist
// @router /preaprobado [get]
func (c *PlanController) PlanPreaprobado() {
	defer errorhandler.HandlePanic(&c.Controller)

	c.ServeJSON()
}
