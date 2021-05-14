package docpars

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/textract"
)

func TestNew(t *testing.T) {
	client := New()
	if client == nil {
		t.Error("error creating parser client")
	}
}

type mockTextractClient struct {
	textractClientOutput *textract.AnalyzeDocumentOutput
	textractClientError  error
}

func (m *mockTextractClient) AnalyzeDocument(input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error) {
	return m.textractClientOutput, m.textractClientError
}

func TestParse(t *testing.T) {
	tests := []struct {
		description          string
		textractClientOutput *textract.AnalyzeDocumentOutput
		textractClientError  error
		content              *Content
		error                error
	}{
		{
			description:          "textract client analyze error",
			textractClientOutput: nil,
			textractClientError:  errors.New("mock analyze error"),
			content:              nil,
			error:                &ErrorAnalyzeDocument{},
		},
		{
			description:          "no selections returned",
			textractClientOutput: &textract.AnalyzeDocumentOutput{},
			textractClientError:  nil,
			content: &Content{
				Selections: []Selection{},
			},
			error: nil,
		},
		{
			description: "on selection returned",
			textractClientOutput: &textract.AnalyzeDocumentOutput{
				Blocks: []*textract.Block{
					{
						BlockType:       aws.String(textract.BlockTypeSelectionElement),
						Confidence:      aws.Float64(0.50),
						Id:              aws.String("selection_id"),
						SelectionStatus: aws.String(textract.SelectionStatusSelected),
					},
					{
						BlockType: aws.String(textract.BlockTypeKeyValueSet),
						Id:        aws.String("value_id"),
						Relationships: []*textract.Relationship{
							{
								Type: aws.String(textract.RelationshipTypeChild),
								Ids: []*string{
									aws.String("selection_id"),
								},
							},
						},
						EntityTypes: []*string{
							aws.String(textract.EntityTypeValue),
						},
					},
					{
						BlockType: aws.String(textract.BlockTypeKeyValueSet),
						Relationships: []*textract.Relationship{
							{
								Type: aws.String(textract.RelationshipTypeValue),
								Ids: []*string{
									aws.String("value_id"),
								},
							},
							{
								Type: aws.String(textract.RelationshipTypeChild),
								Ids: []*string{
									aws.String("line_id"),
								},
							},
						},
						EntityTypes: []*string{
							aws.String(textract.EntityTypeKey),
						},
					},
					{
						BlockType: aws.String(textract.BlockTypeLine),
						Text:      aws.String("example text"),
						Relationships: []*textract.Relationship{
							{
								Type: aws.String(textract.RelationshipTypeChild),
								Ids: []*string{
									aws.String("line_id"),
								},
							},
						},
					},
				},
			},
			textractClientError: nil,
			content: &Content{
				Selections: []Selection{
					{
						ID:         "not_used_id",
						Text:       "example text",
						Selected:   true,
						Confidence: float64(0.50),
					},
				},
			},
			error: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := &Client{
				textractClient: &mockTextractClient{
					textractClientOutput: test.textractClientOutput,
					textractClientError:  test.textractClientError,
				},
			}

			ctx := context.Background()
			doc := []byte("content")

			output, err := client.Parse(ctx, doc)

			if err != nil {
				switch test.error.(type) {
				case *ErrorAnalyzeDocument:
					var testError *ErrorAnalyzeDocument
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				default:
					t.Fatalf("unexpected error type: %v", err)
				}
			}

			if err == nil {
				if output.ID == "" {
					t.Errorf("no content id, received: %+v", output)
				}

				if len(output.Selections) != len(test.content.Selections) {
					t.Errorf("unequal selections lengths, received: %d, expected: %d", len(output.Selections), len(test.content.Selections))
				} else {
					for i, receivedSelection := range output.Selections {
						expectedSelection := test.content.Selections[i]

						if receivedSelection.Text != expectedSelection.Text {
							t.Errorf("incorrect selection text, received: %s, expected: %s", receivedSelection.Text, expectedSelection.Text)
						}

						if receivedSelection.Text != expectedSelection.Text {
							t.Errorf("incorrect selection selected, received: %t, expected: %t", receivedSelection.Selected, expectedSelection.Selected)
						}

						if receivedSelection.Text != expectedSelection.Text {
							t.Errorf("incorrect selection confidence, received: %f, expected: %f", receivedSelection.Confidence, expectedSelection.Confidence)
						}
					}
				}
			}
		})
	}
}
