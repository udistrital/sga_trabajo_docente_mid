package services

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sga_plan_trabajo_docente_mid/utils"
	request "github.com/udistrital/utils_oas/request"
	requestmanager "github.com/udistrital/utils_oas/requestresponse"
)

// DocumentoDocenteVinculacion ...
func ListaDocentesxDocumentoVinculacion(documento string, vinculacion int64) requestmanager.APIResponse {
	resVinculacion := []interface{}{}
	resDocumento := []interface{}{}
	response := []interface{}{}

	if errDocumento := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"datos_identificacion?query=Activo:true,Numero:"+documento+"&fields=TerceroId", &resDocumento); errDocumento == nil {
		if fmt.Sprintf("%v", resDocumento) != "[map[]]" {
			for _, documentoGet := range resDocumento {
				if errVinculacion := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"vinculacion?query=Activo:true,"+
					"TipoVinculacionId:"+fmt.Sprintf("%d", vinculacion)+",TerceroPrincipalId.Id:"+fmt.Sprintf("%v", documentoGet.(map[string]interface{})["TerceroId"].(map[string]interface{})["Id"])+"&fields=TerceroPrincipalId", &resVinculacion); errVinculacion == nil {
					if fmt.Sprintf("%v", resVinculacion) != "[map[]]" {
						response = append(response, map[string]interface{}{
							"Nombre":    utils.Capitalize(resVinculacion[0].(map[string]interface{})["TerceroPrincipalId"].(map[string]interface{})["NombreCompleto"].(string)),
							"Documento": documento,
							"Id":        resVinculacion[0].(map[string]interface{})["TerceroPrincipalId"].(map[string]interface{})["Id"]})
						/* c.Ctx.Output.SetStatus(200)
						c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Query successful", "Data": response} */
					} else {
						return requestmanager.APIResponseDTO(false, 404, nil, "No se encontraron registros de docente")
						/* c.Ctx.Output.SetStatus(404)
						c.Data["json"] = map[string]interface{}{"Success": false, "Status": "404", "Message": "No se encontraron registros de docente"} */
					}
				}
			}
		} else {
			return requestmanager.APIResponseDTO(false, 404, nil, "No se encontraron registros de docente")
			/* c.Ctx.Output.SetStatus(404)
			c.Data["json"] = map[string]interface{}{"Success": false, "Status": "404", "Message": "No se encontraron registros de docente"} */
		}
	} else {
		logs.Error(errDocumento)
		return requestmanager.APIResponseDTO(false, 404, nil, "No se encontraron registros de docentes")
		/* c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "404", "Message": "No se encontraron registros de docentes"} */
	}

	return requestmanager.APIResponseDTO(true, 200, response)
}

// NombreDocenteVinculacion ...
func ListaDocentesxNombreVinculacion(nombre string, vinculacion int64) requestmanager.APIResponse {
	resVinculacion := []interface{}{}
	resDocumento := []interface{}{}
	response := []interface{}{}

	if errVinculacion := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"vinculacion?limit=0&query=TipoVinculacionId__in:"+fmt.Sprintf("%d", vinculacion)+
		",Activo:true,TerceroPrincipalId.NombreCompleto__icontains:"+nombre+"&fields=TerceroPrincipalId", &resVinculacion); errVinculacion == nil {
		if fmt.Sprintf("%v", resVinculacion) != "[map[]]" {
			var tercerosIds string
			for _, vinculacion := range resVinculacion {
				tercerosIds += fmt.Sprintf("%v", vinculacion.(map[string]interface{})["TerceroPrincipalId"].(map[string]interface{})["Id"]) + "|"
			}
			tercerosIds = tercerosIds[:len(tercerosIds)-1]

			if errDocumento := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"datos_identificacion?query=Activo:true,TerceroId__in:"+tercerosIds+"&fields=Numero,TerceroId", &resDocumento); errDocumento == nil {
				for _, vinculacion := range resVinculacion {
					for indexDocumento, documento := range resDocumento {

						if vinculacion.(map[string]interface{})["TerceroPrincipalId"].(map[string]interface{})["Id"] == documento.(map[string]interface{})["TerceroId"].(map[string]interface{})["Id"] {
							response = append(response, map[string]interface{}{
								"Nombre":    utils.Capitalize(vinculacion.(map[string]interface{})["TerceroPrincipalId"].(map[string]interface{})["NombreCompleto"].(string)),
								"Documento": resDocumento[0].(map[string]interface{})["Numero"],
								"Id":        vinculacion.(map[string]interface{})["TerceroPrincipalId"].(map[string]interface{})["Id"]})
							resDocumento = append(resDocumento[:indexDocumento], resDocumento[indexDocumento+1:]...)
							break
						}
					}
				}
			} else {
				logs.Error(errDocumento)
				return requestmanager.APIResponseDTO(false, 404, nil, "No se encontraron registros de docentes")
				/* c.Ctx.Output.SetStatus(404)
				c.Data["json"] = map[string]interface{}{"Success": false, "Status": "404", "Message": "No se encontraron registros de docentes"} */
			}
			return requestmanager.APIResponseDTO(true, 200, response)
			/* c.Ctx.Output.SetStatus(200)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Query successful", "Data": response} */
		} else {
			return requestmanager.APIResponseDTO(false, 404, nil, "No se encontraron registros de docentes")
			/* c.Ctx.Output.SetStatus(404)
			c.Data["json"] = map[string]interface{}{"Success": false, "Status": "404", "Message": "No se encontraron registros de docentes"} */
		}
	} else {
		logs.Error(errVinculacion)
		return requestmanager.APIResponseDTO(false, 404, nil, "No se encontraron registros de docentes")
		/* c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "404", "Message": "No se encontraron registros de docentes"} */
	}
}

// Busca el docente por documento y sus vinculaciones asociadas de tipo docente
func BuscarDocenteConVinculaciones(documento string) (response requestmanager.APIResponse) {
	// 0. Inicialización de variables
	var docente map[string]interface{}
	response = requestmanager.APIResponseDTO(false, 404, docente, "No se encontraron registros de docentes")
	idsVinculacionTipoDocente := []int64{293, 294, 296, 297, 298, 299}

	// 1. Busca los terceros relacionados al documento
	resTercero := []interface{}{}
	errResTercero := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"datos_identificacion?query=Activo:true,Numero:"+documento, &resTercero)
	if errResTercero != nil {
		logs.Error(errResTercero)
		response = requestmanager.APIResponseDTO(false, 404, nil, "No se encontraron registros de docentes")
		return response
	}
	if fmt.Sprint(resTercero) == "[map[]]" {
		response = requestmanager.APIResponseDTO(false, 404, nil, "No se encontraron registros de docentes")
		return response
	}

	for _, tercero := range resTercero {
		// 1.1 Obtener el id del tercero
		terceroId := fmt.Sprintf("%v", tercero.(map[string]interface{})["TerceroId"].(map[string]interface{})["Id"])

		// 2. Buscar las vinculaciones del tercero por id del tercero
		resVinculacion := []interface{}{}
		errResVinculacion := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"vinculacion?query=Activo:true,tercero_principal_id:"+terceroId+"&fields=TipoVinculacionId", &resVinculacion)
		if errResVinculacion != nil {
			logs.Error(errResVinculacion)
			continue // Continuar con el siguiente tercero
		}
		if fmt.Sprint(resVinculacion) == "[map[]]" {
			continue // Continuar con el siguiente tercero
		}

		// 3. Buscar el tipo de vinculacion docente
		var docenteVinculacionesIds []int64
		for _, vinculacion := range resVinculacion {
			vinculacionId := fmt.Sprintf("%v", vinculacion.(map[string]interface{})["TipoVinculacionId"])

			// 4. Verificar si la vinculacion es de tipo docente
			for _, idVinculacionTipoDocente := range idsVinculacionTipoDocente {
				if vinculacionId == fmt.Sprint(idVinculacionTipoDocente) {
					docenteVinculacionesIds = append(docenteVinculacionesIds, idVinculacionTipoDocente)
				}
			}
		}

		// 5. Si el tercero tiene vinculaciones de tipo docente
		if len(docenteVinculacionesIds) > 0 {
			docente = map[string]interface{}{
				"Nombre":        utils.Capitalize(tercero.(map[string]interface{})["TerceroId"].(map[string]interface{})["NombreCompleto"].(string)),
				"Documento":     documento,
				"Id":            tercero.(map[string]interface{})["TerceroId"].(map[string]interface{})["Id"],
				"Vinculaciones": docenteVinculacionesIds,
			}
			response = requestmanager.APIResponseDTO(true, 200, docente, "Registro de docente encontrado")
			return response // Retorna el docente encontrado
		}
	}

	return response // Retorna la respuesta por defecto si no se encuentra ningún docente con vinculaciones
}
