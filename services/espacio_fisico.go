package services

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	requestmanager "github.com/udistrital/utils_oas/requestresponse"
)

func ArbolEspaciosFisicosDependencia(dependencia int64) requestmanager.APIResponse {
	inBog, _ := time.LoadLocation("America/Bogota")
	horaes := time.Now().In(inBog).Format(time.RFC3339)

	resp, err := requestmanager.Get("http://"+beego.AppConfig.String("ProyectoAcademicoService")+
		fmt.Sprintf("proyecto_academico_institucion/%d", dependencia), requestmanager.ParseResonseNoFormat)
	if err != nil {
		logs.Error(err)
		return requestmanager.APIResponseDTO(false, 404, nil, "ProyectoAcademicoService (proyecto_academico_institucion): "+err.Error())
		/* badAns, code := requestmanager.MidResponseFormat("ProyectoAcademicoService (proyecto_academico_institucion)", "GET", false, map[string]interface{}{
			"response": resp,
			"error":    err.Error(),
		})
		c.Ctx.Output.SetStatus(code)
		c.Data["json"] = badAns
		c.ServeJSON()
		return */
	}

	dependencia = int64(resp.(map[string]interface{})["DependenciaId"].(float64))
	if dependencia <= 0 {
		err = fmt.Errorf("no valid Id: %d > 0 = false", dependencia)
		logs.Error(err)
		return requestmanager.APIResponseDTO(false, 404, nil, "GetEspaciosFisicosDependencia (param: dependencia): "+err.Error())
		/* errorAns, statuscode := requestmanager.MidResponseFormat("GetEspaciosFisicosDependencia (param: dependencia)", "GET", false, err.Error())
		c.Ctx.Output.SetStatus(statuscode)
		c.Data["json"] = errorAns
		c.ServeJSON()
		return */
	}

	Salones := map[string][]map[string]interface{}{}
	Edificios := map[string][]map[string]interface{}{}
	Sedes := []map[string]interface{}{}

	resp, err = requestmanager.Get("http://"+beego.AppConfig.String("OikosService")+
		fmt.Sprintf("asignacion_espacio_fisico_dependencia?query=Activo:true,DependenciaId:%d,FechaInicio__lte:%v,FechaFin__gte:%v&fields=EspacioFisicoId&limit=0",
			dependencia, horaes, horaes), requestmanager.ParseResonseNoFormat)
	if err != nil {
		resp, err = requestmanager.Get("http://"+beego.AppConfig.String("OikosService")+"espacio_fisico?query=Nombre:POR%20ASIGNAR,TipoEspacioFisicoId__Id:2", requestmanager.ParseResonseNoFormat)
		if err != nil {
			logs.Error(err)
			return requestmanager.APIResponseDTO(false, 404, nil, "OikosService (espacio_fisico): "+err.Error())
			/* badAns, code := requestmanager.MidResponseFormat("OikosService (espacio_fisico)", "GET", false, map[string]interface{}{
				"response": resp,
				"error":    err.Error(),
			})
			c.Ctx.Output.SetStatus(code)
			c.Data["json"] = badAns
			c.ServeJSON()
			return */
		}
		Idstr := fmt.Sprintf("%v", resp.([]interface{})[0].(map[string]interface{})["Id"])
		Opcion := map[string]interface{}{
			"Id":     resp.([]interface{})[0].(map[string]interface{})["Id"],
			"Nombre": resp.([]interface{})[0].(map[string]interface{})["Nombre"],
		}
		Salones[Idstr] = append(Salones[Idstr], Opcion)
		Edificios[Idstr] = append(Edificios[Idstr], Opcion)
		Sedes = append(Sedes, Opcion)

		return requestmanager.APIResponseDTO(true, 200, map[string]interface{}{
			"Salones":    Salones,
			"Edificios":  Edificios,
			"Sedes":      Sedes,
			"PorAsignar": true,
		})
		/* respuesta, statuscode := requestmanager.MidResponseFormat("GetEspaciosFisicosDependencia", "GET", true, map[string]interface{}{
			"Salones":    Salones,
			"Edificios":  Edificios,
			"Sedes":      Sedes,
			"PorAsignar": true,
		})
		c.Ctx.Output.SetStatus(statuscode)
		c.Data["json"] = respuesta
		c.ServeJSON()
		return */
	}

	for _, EspacioFisico := range resp.([]interface{}) {
		resp, err := requestmanager.Get("http://"+beego.AppConfig.String("OikosService")+
			fmt.Sprintf("espacio_fisico_padre?query=HijoId:%v", EspacioFisico.(map[string]interface{})["EspacioFisicoId"].(map[string]interface{})["Id"]), requestmanager.ParseResonseNoFormat)
		if err == nil {
			tipoEspacio := resp.([]interface{})[0].(map[string]interface{})["PadreId"].(map[string]interface{})["TipoEspacioFisicoId"].(map[string]interface{})["Id"].(float64)
			PadreSalon := fmt.Sprintf("%v", resp.([]interface{})[0].(map[string]interface{})["PadreId"].(map[string]interface{})["Id"])
			for tipoEspacio != 39 {
				resp, err := requestmanager.Get("http://"+beego.AppConfig.String("OikosService")+
					fmt.Sprintf("espacio_fisico_padre?query=HijoId:%v", PadreSalon), requestmanager.ParseResonseNoFormat)
				if err == nil {
					PadreSalon = fmt.Sprintf("%v", resp.([]interface{})[0].(map[string]interface{})["PadreId"].(map[string]interface{})["Id"])
					tipoEspacio = resp.([]interface{})[0].(map[string]interface{})["PadreId"].(map[string]interface{})["TipoEspacioFisicoId"].(map[string]interface{})["Id"].(float64)
				}
			}

			if _, ok := Salones[PadreSalon]; !ok {
				Salones[PadreSalon] = []map[string]interface{}{}
			}
			Salones[PadreSalon] = append(Salones[PadreSalon], map[string]interface{}{
				"Id":                resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Id"],
				"Nombre":            resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Nombre"],
				"Descripcion":       resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Descripcion"],
				"CodigoAbreviacion": resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["CodigoAbreviacion"],
			})

		}
	}

	for PadreSalon := range Salones {
		resp, err := requestmanager.Get("http://"+beego.AppConfig.String("OikosService")+
			fmt.Sprintf("espacio_fisico_padre?query=HijoId:%v", PadreSalon), requestmanager.ParseResonseNoFormat)
		if err == nil {
			PadreEdificio := fmt.Sprintf("%v", resp.([]interface{})[0].(map[string]interface{})["PadreId"].(map[string]interface{})["Id"])
			if _, ok := Edificios[PadreEdificio]; !ok {
				Edificios[PadreEdificio] = []map[string]interface{}{}
			}
			Edificios[PadreEdificio] = append(Edificios[PadreEdificio], map[string]interface{}{
				"Id":                resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Id"],
				"Nombre":            resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Nombre"],
				"Descripcion":       resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Descripcion"],
				"CodigoAbreviacion": resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["CodigoAbreviacion"],
			})
		}
	}

	for PadreEficio := range Edificios {
		resp, err := requestmanager.Get("http://"+beego.AppConfig.String("OikosService")+
			fmt.Sprintf("espacio_fisico_padre?query=HijoId:%v", PadreEficio), requestmanager.ParseResonseNoFormat)
		if err == nil {
			Sedes = append(Sedes, map[string]interface{}{
				"Id":                resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Id"],
				"Nombre":            resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Nombre"],
				"Descripcion":       resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["Descripcion"],
				"CodigoAbreviacion": resp.([]interface{})[0].(map[string]interface{})["HijoId"].(map[string]interface{})["CodigoAbreviacion"],
			})
		}
	}

	return requestmanager.APIResponseDTO(true, 200, map[string]interface{}{
		"Salones":   Salones,
		"Edificios": Edificios,
		"Sedes":     Sedes,
	})

}
