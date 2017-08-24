package gohelpencoding

import "encoding/base64"

// Base64DecodedBatch represents the decoded batch
type Base64DecodedBatch struct {
	dataCollection [][]byte
	iteration      int
}

// NewBase64StdDecodeBatch decodes the batch with the standard encoding
func NewBase64StdDecodeBatch(dataCollection ...string) (decodedBatch *Base64DecodedBatch, err error) {
	decodedBatch = &Base64DecodedBatch{}
	decodedBatch.dataCollection = make([][]byte, len(dataCollection))
	for index, encodedData := range dataCollection {
		decodedBatch.dataCollection[index], err = base64.StdEncoding.DecodeString(encodedData)
		if err != nil {
			return nil, err
		}
	}
	return
}

// Next returns the next entry in the decoded batch collection
func (decodedBatch *Base64DecodedBatch) Next() (data []byte) {
	data = decodedBatch.dataCollection[decodedBatch.iteration]
	decodedBatch.iteration++
	return
}
