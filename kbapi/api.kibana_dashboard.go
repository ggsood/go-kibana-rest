package kbapi

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	basePathKibanaDashboard = "/api/kibana/dashboards" // Base URL to access on Kibana dashboard
)

type KibanaDashboardExport func(listID []string) (map[string]interface{}, error)
type KibanaDashboardImport func(data map[string]interface{}, listExcludeType []string, force bool) error

// newKibanaDashboardExport permit to export Kibana dashboard by its names
func newKibanaDashboardExportFunc(c *resty.Client) KibanaDashboardExport {
	return func(listID []string) (map[string]interface{}, error) {

		if len(listID) == 0 {
			return nil, NewAPIError(600, "You must provide on or more dashboard ID")
		}
		log.Debug("listID: ", listID)

		path := fmt.Sprintf("%s/export", basePathKibanaDashboard)
		query := fmt.Sprintf("dashboard=%s", strings.Join(listID, ","))
		resp, err := c.R().SetQueryString(query).Get(path)
		if err != nil {
			return nil, err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			if resp.StatusCode() == 404 {
				return nil, nil
			} else {
				return nil, NewAPIError(resp.StatusCode(), resp.Status())
			}
		}
		var data map[string]interface{}
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return nil, err
		}
		log.Debug("Data: ", data)

		return data, nil
	}

}

// newKibanaDashboardImportFunc permit to import kibana dashboard
func newKibanaDashboardImportFunc(c *resty.Client) KibanaDashboardImport {
	return func(data map[string]interface{}, listExcludeType []string, force bool) error {

		if data == nil {
			return NewAPIError(600, "You must provide one or more dashboard to import")
		}
		log.Debug("data: ", data)
		log.Debug("List type to exclude: ", listExcludeType)
		log.Debug("Force import: ", force)

		path := fmt.Sprintf("%s/import", basePathKibanaDashboard)
		request := c.R().SetQueryString(fmt.Sprintf("force=%s", force))
		if len(listExcludeType) > 0 {
			request = request.SetQueryString(fmt.Sprintf("exclude=%s", strings.Join(listExcludeType, ",")))
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		resp, err := request.SetBody(jsonData).Post(path)
		if err != nil {
			return err
		}
		log.Debug("Response: ", resp)
		if resp.StatusCode() >= 300 {
			return NewAPIError(resp.StatusCode(), resp.Status())
		}
		var dataResponse map[string]interface{}
		err = json.Unmarshal(resp.Body(), &dataResponse)
		if err != nil {
			return err
		}
		log.Debug("Data response: ", dataResponse)

		// Need to manage error returned in response

		return nil
	}

}