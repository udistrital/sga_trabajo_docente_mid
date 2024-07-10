package services

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/sga_trabajo_docente_mid/utils"
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
