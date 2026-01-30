package camt

import (
	"bytes"
	"encoding/xml"

	"golang.org/x/net/html/charset"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type camt052FileReader struct {
	xmlDecoder *xml.Decoder
}


type camt053FileReader struct {
	xmlDecoder *xml.Decoder
}


func (r *camt052FileReader) read(ctx core.Context) (*camt052File, error) {
	file := &camt052File{}

	err := r.xmlDecoder.Decode(&file)

	if err != nil {
		return nil, err
	}

	return file, nil
}


func (r *camt053FileReader) read(ctx core.Context) (*camt053File, error) {
	file := &camt053File{}

	err := r.xmlDecoder.Decode(&file)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func createNewCamt052FileReader(data []byte) (*camt052FileReader, error) {
	if len(data) > 5 && data[0] == 0x3C && data[1] == 0x3F && data[2] == 0x78 && data[3] == 0x6D && data[4] == 0x6C { 
		xmlDecoder := xml.NewDecoder(bytes.NewReader(data))
		xmlDecoder.CharsetReader = charset.NewReaderLabel

		return &camt052FileReader{
			xmlDecoder: xmlDecoder,
		}, nil
	}

	return nil, errs.ErrInvalidXmlFile
}

func createNewCamt053FileReader(data []byte) (*camt053FileReader, error) {
	if len(data) > 5 && data[0] == 0x3C && data[1] == 0x3F && data[2] == 0x78 && data[3] == 0x6D && data[4] == 0x6C { 
		xmlDecoder := xml.NewDecoder(bytes.NewReader(data))
		xmlDecoder.CharsetReader = charset.NewReaderLabel

		return &camt053FileReader{
			xmlDecoder: xmlDecoder,
		}, nil
	}

	return nil, errs.ErrInvalidXmlFile
}
