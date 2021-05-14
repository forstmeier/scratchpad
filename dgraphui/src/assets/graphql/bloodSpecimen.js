export const add = `
mutation AddBloodSpecimen($input: [AddBloodSpecimenInput!]!) {
    addBloodSpecimen(input: $input) {
        bloodSpecimen {
            id
        }
    }
}
`;

export const query = `
query QueryBloodSpecimen($bloodSpecimen: BloodSpecimenFilter, $donor: DonorFilter, $consent: ConsentFilter) {
    queryBloodSpecimen(filter: $bloodSpecimen) {
        id
        externalID
        description
        collectionDate
        destructionDate
        type
        container
        status
        donor(filter: $donor) {
            dob
            sex
            race
        }
        consent(filter: $consent) {
            textBody
        }
        bloodType
        volume
    }
}
`;
