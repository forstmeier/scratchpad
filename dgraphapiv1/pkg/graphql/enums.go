package graphql

// Sex is the enum defined for Donor fields in the schema.
type Sex string

// Sexes are the enum value options for Sex.
const (
	Male   Sex = "MALE"
	Female Sex = "FEMALE"
)

// Race is the enum defined for Donor fields in the schema.
type Race string

// Races are the enum value options for Race.
const (
	AmericanIndianOrAlaskaNative Race = "AMERICAN_INDIAN_OR_ALASKA_NATIVE"
	Asian                        Race = "ASIAN"
	BlackOrAfricanAmerican       Race = "BLACK_OR_AFRICAN_AMERICAN"
	HispanicOrLatino             Race = "HISPANIC_OR_LATINO"
	White                        Race = "WHITE"
)

// SpecimenType is the enum defined for Specimen fields in the schema.
type SpecimenType string

// Specimen types are the enum value options for SpecimenType.
const (
	Blood SpecimenType = "BLOOD"
)

// Container is the enum defined for Specimen fields in the schema.
type Container string

// Containers are the enum value options for Container.
const (
	Vial Container = "VIAL"
)

// Status is the enum defined for Specimen fields in the schema.
type Status string

// Statuses are the enum value options for Status.
const (
	Destroyed   Status = "DESTROYED"
	Exhausted   Status = "EXHAUSTED"
	InInventory Status = "IN_INVENTORY"
	InTransit   Status = "IN_TRANSIT"
	Lost        Status = "LOST"
	Reserved    Status = "RESERVED"
	Transferred Status = "TRANSFERRED"
)

// BloodType is the enum defined for BloodSpecimen fields in the schema.
type BloodType string

// Blood types are the enum value options for BloodType.
const (
	ONeg  BloodType = "O_NEG"
	OPos  BloodType = "O_POS"
	ANeg  BloodType = "A_NEG"
	APos  BloodType = "A_POS"
	BNeg  BloodType = "B_NEG"
	BPos  BloodType = "B_POS"
	ABNeg BloodType = "AB_NEG"
	ABPos BloodType = "AB_POS"
)
