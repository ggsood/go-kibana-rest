package kbapi

import (
	"github.com/go-resty/resty/v2"
)

// API handle the API specification
type API struct {
	KibanaSpaces         *KibanaSpacesAPI
	KibanaRoleManagement *KibanaRoleManagementAPI
	KibanaDashboard      *KibanaDashboardAPI
	KibanaSavedObject    *KibanaSavedObjectAPI
	KibanaStatus         *KibanaStatusAPI
}

// KibanaSpacesAPI handle the spaces API
type KibanaSpacesAPI struct {
	Get              KibanaSpaceGet
	List             KibanaSpaceList
	Create           KibanaSpaceCreate
	Delete           KibanaSpaceDelete
	Update           KibanaSpaceUpdate
	CopySavedObjects KibanaSpaceCopySavedObjects
}

// KibanaRoleManagementAPI handle the role management API
type KibanaRoleManagementAPI struct {
	Get            KibanaRoleManagementGet
	List           KibanaRoleManagementList
	CreateOrUpdate KibanaRoleManagementCreateOrUpdate
	Delete         KibanaRoleManagementDelete
}

// KibanaDashboardAPI handle the dashboard API
type KibanaDashboardAPI struct {
	Export KibanaDashboardExport
	Import KibanaDashboardImport
}

// KibanaSavedObjectAPI handle the saved object API
type KibanaSavedObjectAPI struct {
	Get    KibanaSavedObjectGet
	Find   KibanaSavedObjectFind
	Create KibanaSavedObjectCreate
	Update KibanaSavedObjectUpdate
	Delete KibanaSavedObjectDelete
	Import KibanaSavedObjectImport
	Export KibanaSavedObjectExport
}

// KibanaStatusAPI handle the status API
type KibanaStatusAPI struct {
	Get KibanaStatusGet
}

// New initialise the API implementation
func New(c *resty.Client) *API {
	return &API{
		KibanaSpaces: &KibanaSpacesAPI{
			Get:              newKibanaSpaceGetFunc(c),
			List:             newKibanaSpaceListFunc(c),
			Create:           newKibanaSpaceCreateFunc(c),
			Update:           newKibanaSpaceUpdateFunc(c),
			Delete:           newKibanaSpaceDeleteFunc(c),
			CopySavedObjects: newKibanaSpaceCopySavedObjectsFunc(c),
		},
		KibanaRoleManagement: &KibanaRoleManagementAPI{
			Get:            newKibanaRoleManagementGetFunc(c),
			List:           newKibanaRoleManagementListFunc(c),
			CreateOrUpdate: newKibanaRoleManagementCreateOrUpdateFunc(c),
			Delete:         newKibanaRoleManagementDeleteFunc(c),
		},
		KibanaDashboard: &KibanaDashboardAPI{
			Export: newKibanaDashboardExportFunc(c),
			Import: newKibanaDashboardImportFunc(c),
		},
		KibanaSavedObject: &KibanaSavedObjectAPI{
			Get:    newKibanaSavedObjectGetFunc(c),
			Find:   newKibanaSavedObjectFindFunc(c),
			Create: newKibanaSavedObjectCreateFunc(c),
			Update: newKibanaSavedObjectUpdateFunc(c),
			Delete: newKibanaSavedObjectDeleteFunc(c),
			Import: newKibanaSavedObjectImportFunc(c),
			Export: newKibanaSavedObjectExportFunc(c),
		},
		KibanaStatus: &KibanaStatusAPI{
			Get: newKibanaStatusGetFunc(c),
		},
	}
}
