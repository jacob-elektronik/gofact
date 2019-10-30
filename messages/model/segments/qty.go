package segments

type Quantity struct {
	QuantityDetails QuantityDetails
}

type QuantityDetails struct {
	QuantityTypeCodeQualifier string
	Quantity                  string
	MeasurementUnitCode       string
}
