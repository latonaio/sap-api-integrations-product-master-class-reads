package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-product-master-class-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) ([]General, error) {
	pm := &responses.General{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	general := make([]General, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		general = append(general, General{
			Product:             data.Product,
			ProductType:         data.ProductType,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
			ProductGroup:        data.ProductGroup,
			BaseUnit:            data.BaseUnit,
			ProductHierarchy:    data.ProductHierarchy,
			ToProductClass:      data.ToProductClass.Deferred.URI,
			ToProductCharc:      data.ToProductCharc.Deferred.URI,
		})
	}

	return general, nil
}

func ConvertToToProductClass(raw []byte, l *logger.Logger) ([]ToProductClass, error) {
	pm := &responses.ToProductClass{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToProductClass. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toProductClass := make([]ToProductClass, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toProductClass = append(toProductClass, ToProductClass{
			Product:               data.Product,
			ClassInternalID:       data.ClassInternalID,
			KeyDate:               data.KeyDate,
			ChangeNumber:          data.ChangeNumber,
			ClassType:             data.ClassType,
			ToClassDetails:        data.ToClassDetails.Deferred.URI,
			ToCharc:               data.ToCharc.Deferred.URI,
		})
	}

	return toProductClass, nil
}

func ConvertToToClassDetails(raw []byte, l *logger.Logger) (*ToClassDetails, error) {
	pm := &responses.ToClassDetails{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToClassDetails. unmarshal error: %w", err)
	}

	return &ToClassDetails{
		ClassInternalID:          pm.D.ClassInternalID,
		ClassType:                pm.D.ClassType,
		ClassTypeName:            pm.D.ClassTypeName,
		Class:                    pm.D.Class,
		ClassStatus:              pm.D.ClassStatus,
		ClassStatusName:          pm.D.ClassStatusName,
		ClassGroup:               pm.D.ClassGroup,
		ClassGroupName:           pm.D.ClassGroupName,
		CreationDate:             pm.D.CreationDate,
		LastChangeDate:           pm.D.LastChangeDate,
		ValidityStartDate:        pm.D.ValidityStartDate,
		ValidityEndDate:          pm.D.ValidityEndDate,
		ClassLastChangedDateTime: pm.D.ClassLastChangedDateTime,
		KeyDate:                  pm.D.KeyDate,
	}, nil
}

func ConvertToToProductCharc(raw []byte, l *logger.Logger) ([]ToProductCharc, error) {
	pm := &responses.ToProductCharc{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToProductCharc. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toProductCharc := make([]ToProductCharc, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toProductCharc = append(toProductCharc, ToProductCharc{
			Product:         data.Product,
			ClassInternalID: data.ClassInternalID,
			CharcInternalID: data.CharcInternalID,
			KeyDate:         data.KeyDate,
			ChangeNumber:    data.ChangeNumber,
			ClassType:       data.ClassType,
		})
	}

	return toProductCharc, nil
}
