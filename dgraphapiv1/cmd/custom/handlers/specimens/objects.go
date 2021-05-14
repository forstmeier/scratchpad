package specimens

// AddBloodSpecimenUpdateRequest is a request received from the @custom
// directive from Dgraph sent to the custom server to add a blood
// specimen update.
type AddBloodSpecimenUpdateRequest struct {
	OrgID           string  `json:"orgID" valid:"required"`
	BloodSpecimenID string  `json:"bloodSpecimenID" valid:"required"`
	BloodType       string  `json:"bloodType"`
	Container       string  `json:"container"`
	Volume          float64 `json:"volume"`
	Description     string  `json:"description"`
}

// Response holds the HTTP response body from Dgraph.
type Response struct {
	Body Body `json:"data"`
}

// Body holds the Dgraph response data.
type Body struct {
	Data Data `json:"data"`
}

// Data holds the Dgraph response BloodSpecimen(s) and
// BloodSpecimenUpdate(s) data.
type Data struct {
	BloodSpecimen       []BloodSpecimen       `json:"bloodSpecimen"`
	BloodSpecimenUpdate []BloodSpecimenUpdate `json:"bloodSpecimenUpdate"`
}

// BloodSpecimen holds the Dgraph response BloodSpecimen data.
type BloodSpecimen struct {
	ID string `json:"id"`
}

// BloodSpecimenUpdate holds the Dgraph response BloodSpecimenUpdate data.
type BloodSpecimenUpdate struct {
	ID string `json:"id"`
}
