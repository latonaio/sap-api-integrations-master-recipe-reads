package sap_api_output_formatter

type MasterRecipe struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	MasterRecipe  string `json:"business_partner_code"`
	Deleted       bool   `json:"deleted"`
}

type Header struct {
	MasterRecipeGroup           string `json:"MasterRecipeGroup"`
	MasterRecipe                string `json:"MasterRecipe"`
	MasterRecipeInternalVersion string `json:"MasterRecipeInternalVersion"`
	IsMarkedForDeletion         bool   `json:"IsMarkedForDeletion"`
	BillOfOperationsDesc        string `json:"BillOfOperationsDesc"`
	LongTextLanguageCode        string `json:"LongTextLanguageCode"`
	PlainLongText               string `json:"PlainLongText"`
	Plant                       string `json:"Plant"`
	BillOfOperationsStatus      string `json:"BillOfOperationsStatus"`
	BillOfOperationsUsage       string `json:"BillOfOperationsUsage"`
	ResponsiblePlannerGroup     string `json:"ResponsiblePlannerGroup"`
	BillOfOperationsProfile     string `json:"BillOfOperationsProfile"`
	MinimumLotSizeQuantity      string `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity      string `json:"MaximumLotSizeQuantity"`
	BillOfOperationsUnit        string `json:"BillOfOperationsUnit"`
	CreationDate                string `json:"CreationDate"`
	LastChangeDate              string `json:"LastChangeDate"`
	ValidityStartDate           string `json:"ValidityStartDate"`
	ValidityEndDate             string `json:"ValidityEndDate"`
	ChangeNumber                string `json:"ChangeNumber"`
	ChangedDateTime             string `json:"ChangedDateTime"`
	ToMaterialAssignment        string `json:"to_MatlAssgmt"`
	ToOperation                 string `json:"to_Operation"`
}

type ToMaterialAssignment struct {
	Product                        string `json:"Product"`
	Plant                          string `json:"Plant"`
	MasterRecipeGroup              string `json:"MasterRecipeGroup"`
	MasterRecipe                   string `json:"MasterRecipe"`
	MasterRecipeMaterialAssignment string `json:"MasterRecipeMaterialAssignment"`
	MstrRcpMatlAssgmtIntVersion    string `json:"MstrRcpMatlAssgmtIntVersion"`
	CreationDate                   string `json:"CreationDate"`
	LastChangeDate                 string `json:"LastChangeDate"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	ChangeNumber                   string `json:"ChangeNumber"`
	ChangedDateTime                string `json:"ChangedDateTime"`
}

type ToOperation struct {
	MasterRecipeGroup             string `json:"MasterRecipeGroup"`
	MasterRecipe                  string `json:"MasterRecipe"`
	MasterRecipeOperationIntID    string `json:"MasterRecipeOperationIntID"`
	MstrRcpOperationIntVersion    string `json:"MstrRcpOperationIntVersion"`
	Operation                     string `json:"Operation"`
	OperationText                 string `json:"OperationText"`
	LongTextLanguageCode          string `json:"LongTextLanguageCode"`
	PlainLongText                 string `json:"PlainLongText"`
	WorkCenterTypeCode            string `json:"WorkCenterTypeCode"`
	WorkCenterInternalID          string `json:"WorkCenterInternalID"`
	OperationStandardTextCode     string `json:"OperationStandardTextCode"`
	Plant                         string `json:"Plant"`
	OperationControlProfile       string `json:"OperationControlProfile"`
	OperationReferenceQuantity    string `json:"OperationReferenceQuantity"`
	OperationUnit                 string `json:"OperationUnit"`
	OpQtyToBaseQtyNmrtr           string `json:"OpQtyToBaseQtyNmrtr"`
	OpQtyToBaseQtyDnmntr          string `json:"OpQtyToBaseQtyDnmntr"`
	NumberOfTimeTickets           string `json:"NumberOfTimeTickets"`
	NumberOfConfirmationSlips     string `json:"NumberOfConfirmationSlips"`
	OperationCostingRelevancyType string `json:"OperationCostingRelevancyType"`
	BusinessProcess               string `json:"BusinessProcess"`
	OperationSetupType            string `json:"OperationSetupType"`
	OperationSetupGroupCategory   string `json:"OperationSetupGroupCategory"`
	OperationSetupGroup           string `json:"OperationSetupGroup"`
	CapacityCategoryCode          string `json:"CapacityCategoryCode"`
	CreationDate                  string `json:"CreationDate"`
	LastChangeDate                string `json:"LastChangeDate"`
	ValidityStartDate             string `json:"ValidityStartDate"`
	ValidityEndDate               string `json:"ValidityEndDate"`
	ChangeNumber                  string `json:"ChangeNumber"`
	ChangedDateTime               string `json:"ChangedDateTime"`
	ToPhase                       string `json:"to_Phase"`
}

type ToPhase struct {
	MasterRecipeGroup              string `json:"MasterRecipeGroup"`
	MasterRecipe                   string `json:"MasterRecipe"`
	MasterRecipeOperationIntID     string `json:"MasterRecipeOperationIntID"`
	MstrRcpSuperiorOpIntVersion    string `json:"MstrRcpSuperiorOpIntVersion"`
	MstrRcpOperationIntVersion     string `json:"MstrRcpOperationIntVersion"`
	SuperiorOperationInternalID    string `json:"SuperiorOperationInternalID"`
	Operation                      string `json:"Operation"`
	OperationText                  string `json:"OperationText"`
	LongTextLanguageCode           string `json:"LongTextLanguageCode"`
	PlainLongText                  string `json:"PlainLongText"`
	WorkCenterTypeCode             string `json:"WorkCenterTypeCode"`
	WorkCenterInternalID           string `json:"WorkCenterInternalID"`
	Plant                          string `json:"Plant"`
	OperationControlProfile        string `json:"OperationControlProfile"`
	ControlRecipeDestination       string `json:"ControlRecipeDestination"`
	OperationStandardTextCode      string `json:"OperationStandardTextCode"`
	OperationReferenceQuantity     string `json:"OperationReferenceQuantity"`
	OperationUnit                  string `json:"OperationUnit"`
	OpQtyToBaseQtyNmrtr            string `json:"OpQtyToBaseQtyNmrtr"`
	OpQtyToBaseQtyDnmntr           string `json:"OpQtyToBaseQtyDnmntr"`
	StandardWorkFormulaParam1      string `json:"StandardWorkFormulaParam1"`
	StandardWorkFormulaParamName1  string `json:"StandardWorkFormulaParamName1"`
	StandardWorkQuantity1          string `json:"StandardWorkQuantity1"`
	StandardWorkQuantityUnit1      string `json:"StandardWorkQuantityUnit1"`
	CostCtrActivityType1           string `json:"CostCtrActivityType1"`
	StandardWorkFormulaParam2      string `json:"StandardWorkFormulaParam2"`
	StandardWorkFormulaParamName2  string `json:"StandardWorkFormulaParamName2"`
	StandardWorkQuantity2          string `json:"StandardWorkQuantity2"`
	StandardWorkQuantityUnit2      string `json:"StandardWorkQuantityUnit2"`
	CostCtrActivityType2           string `json:"CostCtrActivityType2"`
	StandardWorkFormulaParam3      string `json:"StandardWorkFormulaParam3"`
	StandardWorkFormulaParamName3  string `json:"StandardWorkFormulaParamName3"`
	StandardWorkQuantity3          string `json:"StandardWorkQuantity3"`
	StandardWorkQuantityUnit3      string `json:"StandardWorkQuantityUnit3"`
	CostCtrActivityType3           string `json:"CostCtrActivityType3"`
	StandardWorkFormulaParam4      string `json:"StandardWorkFormulaParam4"`
	StandardWorkFormulaParamName4  string `json:"StandardWorkFormulaParamName4"`
	StandardWorkQuantity4          string `json:"StandardWorkQuantity4"`
	StandardWorkQuantityUnit4      string `json:"StandardWorkQuantityUnit4"`
	CostCtrActivityType4           string `json:"CostCtrActivityType4"`
	StandardWorkFormulaParam5      string `json:"StandardWorkFormulaParam5"`
	StandardWorkFormulaParamName5  string `json:"StandardWorkFormulaParamName5"`
	StandardWorkQuantity5          string `json:"StandardWorkQuantity5"`
	StandardWorkQuantityUnit5      string `json:"StandardWorkQuantityUnit5"`
	CostCtrActivityType5           string `json:"CostCtrActivityType5"`
	StandardWorkFormulaParam6      string `json:"StandardWorkFormulaParam6"`
	StandardWorkFormulaParamName6  string `json:"StandardWorkFormulaParamName6"`
	StandardWorkQuantity6          string `json:"StandardWorkQuantity6"`
	StandardWorkQuantityUnit6      string `json:"StandardWorkQuantityUnit6"`
	CostCtrActivityType6           string `json:"CostCtrActivityType6"`
	NumberOfTimeTickets            string `json:"NumberOfTimeTickets"`
	NumberOfConfirmationSlips      string `json:"NumberOfConfirmationSlips"`
	OperationCostingRelevancyType  string `json:"OperationCostingRelevancyType"`
	BusinessProcess                string `json:"BusinessProcess"`
	OperationSetupType             string `json:"OperationSetupType"`
	OperationSetupGroupCategory    string `json:"OperationSetupGroupCategory"`
	OperationSetupGroup            string `json:"OperationSetupGroup"`
	CapacityCategoryCode           string `json:"CapacityCategoryCode"`
	OpIsExtlyProcdWithSubcontrg    bool   `json:"OpIsExtlyProcdWithSubcontrg"`
	InspectionLotType              string `json:"InspectionLotType"`
	PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
	PurchasingOrganization         string `json:"PurchasingOrganization"`
	PurchaseContract               string `json:"PurchaseContract"`
	PurchaseContractItem           string `json:"PurchaseContractItem"`
	PurchasingInfoRecdAddlGrpgName string `json:"PurchasingInfoRecdAddlGrpgName"`
	PlannedDeliveryDuration        string `json:"PlannedDeliveryDuration"`
	MaterialGroup                  string `json:"MaterialGroup"`
	PurchasingGroup                string `json:"PurchasingGroup"`
	Supplier                       string `json:"Supplier"`
	NumberOfOperationPriceUnits    string `json:"NumberOfOperationPriceUnits"`
	CostElement                    string `json:"CostElement"`
	OpExternalProcessingPrice      string `json:"OpExternalProcessingPrice"`
	OpExternalProcessingCurrency   string `json:"OpExternalProcessingCurrency"`
	CreationDate                   string `json:"CreationDate"`
	LastChangeDate                 string `json:"LastChangeDate"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	ChangeNumber                   string `json:"ChangeNumber"`
	ChangedDateTime                string `json:"ChangedDateTime"`
	ToPhaseRelationship            string `json:"to_PhseRelshp"`
}

type ToPhaseRelationship struct {
	PrdcssrMasterRecipeGroup      string `json:"PrdcssrMasterRecipeGroup"`
	PrdcssrMasterRecipe           string `json:"PrdcssrMasterRecipe"`
	PrdcssrMstrRcpOpInternalID    string `json:"PrdcssrMstrRcpOpInternalID"`
	SuccssrMasterRecipeGroup      string `json:"SuccssrMasterRecipeGroup"`
	SuccssrMasterRecipe           string `json:"SuccssrMasterRecipe"`
	SuccssrMstrRcpOpInternalID    string `json:"SuccssrMstrRcpOpInternalID"`
	MasterRecipeRelationshipType  string `json:"MasterRecipeRelationshipType"`
	MaxTimeIntvlIsUsedForSchedg   bool   `json:"MaxTimeIntvlIsUsedForSchedg"`
	MstrRcpRelationshipIntVersion string `json:"MstrRcpRelationshipIntVersion"`
	TimeIntvlBtwnRelshp           string `json:"TimeIntvlBtwnRelshp"`
	MaxTimeIntvlBtwnRelshp        string `json:"MaxTimeIntvlBtwnRelshp"`
	TimeIntvlBtwnRelshpUnit       string `json:"TimeIntvlBtwnRelshpUnit"`
	FactoryCalendar               string `json:"FactoryCalendar"`
	WorkCenterInternalID          string `json:"WorkCenterInternalID"`
	Plant                         string `json:"Plant"`
	CreationDate                  string `json:"CreationDate"`
	LastChangeDate                string `json:"LastChangeDate"`
	ValidityStartDate             string `json:"ValidityStartDate"`
	ValidityEndDate               string `json:"ValidityEndDate"`
	ChangeNumber                  string `json:"ChangeNumber"`
	ChangedDateTime               string `json:"ChangedDateTime"`
}
