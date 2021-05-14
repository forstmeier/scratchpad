package docpars

import "github.com/aws/aws-sdk-go/service/textract"

func getValueBlockID(selectionID string, valueBlocks []*textract.Block) string {
	for _, valueBlock := range valueBlocks {
		if inRelationshipsIDs(selectionID, valueBlock.Relationships[0].Ids) {
			return *valueBlock.Id
		}
	}

	return ""
}

func getLineBlockID(valueID string, keyBlocks []*textract.Block) string {
	for _, keyBlock := range keyBlocks {
		for _, valueRelationship := range keyBlock.Relationships {
			if inRelationshipsIDs(valueID, valueRelationship.Ids) {
				for _, childRelationship := range keyBlock.Relationships {
					if *childRelationship.Type == textract.RelationshipTypeChild {
						return *childRelationship.Ids[0]
					}
				}
			}
		}
	}

	return ""
}

func getSelectionText(lineID string, lineBlocks []*textract.Block) string {
	for _, lineBlock := range lineBlocks {
		if inRelationshipsIDs(lineID, lineBlock.Relationships[0].Ids) {
			return *lineBlock.Text
		}
	}

	return ""
}

func inRelationshipsIDs(id string, relationshipsIDs []*string) bool {
	for _, relationshipID := range relationshipsIDs {
		if id == *relationshipID {
			return true
		}
	}

	return false
}
