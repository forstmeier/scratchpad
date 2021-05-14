package docpars

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/google/uuid"
)

var _ Parser = &Client{}

// Client implements the docpars.Parser.Parse methods using
// AWS Textract.
type Client struct {
	textractClient textractClient
}

type textractClient interface {
	AnalyzeDocument(input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error)
}

// New generates a Client pointer instance with an AWS
// Textract session client.
func New() *Client {
	newSession := session.Must(session.NewSession())
	service := textract.New(newSession)

	return &Client{
		textractClient: service,
	}
}

// Parse implements the docpars.Parser.Parse interface method.
func (c *Client) Parse(ctx context.Context, doc []byte) (*Content, error) {
	content := &Content{
		ID: uuid.NewString(),
	}

	input := &textract.AnalyzeDocumentInput{
		Document: &textract.Document{
			Bytes: doc,
		},
		FeatureTypes: []*string{
			aws.String(textract.FeatureTypeTables),
			aws.String(textract.FeatureTypeForms),
		},
	}

	output, err := c.textractClient.AnalyzeDocument(input)
	if err != nil {
		return nil, &ErrorAnalyzeDocument{err: err}
	}

	selectionBlocks := []*textract.Block{}
	keyBlocks := []*textract.Block{}
	valueBlocks := []*textract.Block{}
	lineBlocks := []*textract.Block{}

	for _, block := range output.Blocks {
		switch *block.BlockType {
		case textract.BlockTypeSelectionElement:
			selectionBlocks = append(selectionBlocks, block)
		case textract.BlockTypeKeyValueSet:
			if *block.EntityTypes[0] == textract.EntityTypeKey {
				keyBlocks = append(keyBlocks, block)
			} else {
				valueBlocks = append(valueBlocks, block)
			}
		case textract.BlockTypeLine:
			lineBlocks = append(lineBlocks, block)
		}
	}

	content.Selections = make([]Selection, len(selectionBlocks))

	for i, selectionBlock := range selectionBlocks {
		selection := Selection{
			ID:         uuid.NewString(),
			Selected:   *selectionBlock.SelectionStatus == textract.SelectionStatusSelected,
			Confidence: *selectionBlock.Confidence,
		}

		valueID := getValueBlockID(*selectionBlock.Id, valueBlocks)
		lineID := getLineBlockID(valueID, keyBlocks)
		selection.Text = getSelectionText(lineID, lineBlocks)

		content.Selections[i] = selection
	}

	return content, nil
}
