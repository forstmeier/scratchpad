package docpars

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/textract"
)

func Test_getValueBlockID(t *testing.T) {
	tests := []struct {
		description string
		selectionID string
		valueBlocks []*textract.Block
		output      string
	}{
		{
			description: "no values returned",
			selectionID: "non_id",
			valueBlocks: []*textract.Block{
				{
					Relationships: []*textract.Relationship{
						{
							Ids: []*string{
								aws.String("selection_id"),
							},
						},
					},
				},
			},
			output: "",
		},
		{
			description: "one value returned",
			selectionID: "selection_id",
			valueBlocks: []*textract.Block{
				{
					Relationships: []*textract.Relationship{
						{
							Ids: []*string{
								aws.String("selection_id"),
							},
						},
					},
					Id: aws.String("value_id"),
				},
			},
			output: "value_id",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if output := getValueBlockID(test.selectionID, test.valueBlocks); output != test.output {
				t.Errorf("incorrect output, received: %s, expected: %s", output, test.output)
			}
		})
	}
}

func Test_getLineBlockID(t *testing.T) {
	tests := []struct {
		description string
		valueID     string
		keyBlocks   []*textract.Block
		output      string
	}{
		{
			description: "no values returned",
			valueID:     "non_id",
			keyBlocks: []*textract.Block{
				{
					Relationships: []*textract.Relationship{
						{
							Ids: []*string{
								aws.String("value_id"),
							},
						},
					},
				},
			},
			output: "",
		},
		{
			description: "one value returned",
			valueID:     "value_id",
			keyBlocks: []*textract.Block{
				{
					Relationships: []*textract.Relationship{
						{
							Ids: []*string{
								aws.String("value_id"),
							},
							Type: aws.String(textract.RelationshipTypeValue),
						},
						{
							Ids: []*string{
								aws.String("line_id"),
							},
							Type: aws.String(textract.RelationshipTypeChild),
						},
					},
				},
			},
			output: "line_id",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if output := getLineBlockID(test.valueID, test.keyBlocks); output != test.output {
				t.Errorf("incorrect output, received: %s, expected: %s", output, test.output)
			}
		})
	}
}

func Test_getSelectionText(t *testing.T) {
	tests := []struct {
		description string
		lineID      string
		lineBlocks  []*textract.Block
		output      string
	}{
		{
			description: "no values returned",
			lineID:      "non_id",
			lineBlocks: []*textract.Block{
				{
					Relationships: []*textract.Relationship{
						{
							Ids: []*string{
								aws.String("line_id"),
							},
						},
					},
				},
			},
			output: "",
		},
		{
			description: "one value returned",
			lineID:      "line_id",
			lineBlocks: []*textract.Block{
				{
					Relationships: []*textract.Relationship{
						{
							Ids: []*string{
								aws.String("line_id"),
							},
							Type: aws.String(textract.RelationshipTypeValue),
						},
					},
					Text: aws.String("example text"),
				},
			},
			output: "example text",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if output := getSelectionText(test.lineID, test.lineBlocks); output != test.output {
				t.Errorf("incorrect output, received: %s, expected: %s", output, test.output)
			}
		})
	}
}

func Test_inRelationshipsIDs(t *testing.T) {
	tests := []struct {
		description     string
		id              string
		relationshipIDs []*string
		output          bool
	}{
		{
			description: "id not in relationship ids",
			id:          "non_id",
			relationshipIDs: []*string{
				aws.String("line_id"),
			},
			output: false,
		},
		{
			description: "id in relationship ids",
			id:          "line_id",
			relationshipIDs: []*string{
				aws.String("line_id"),
			},
			output: true,
		},
	}

	for _, test := range tests {
		if output := inRelationshipsIDs(test.id, test.relationshipIDs); output != test.output {
			t.Errorf("incorrect output, received: %t, expected: %t", output, test.output)
		}
	}
}
