// Akasha Gateway - API for legacy (web 2.0) applications
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
//
// The Akasha gateway is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// The Akasha gateway is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public
// License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Akasha gateway. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Block represents a single paragraph in an Akasha entry or comment, which
// might be purely text, or might be some media object.
type Block struct {
	Text    string `json:"text,omitempty"`
	Image   string `json:"image,omitempty"`
	Caption string `json:"caption,omitempty"`
}

// parseDraftjs converts a messy draftjs document into a cleaned up and very much
// simplified programatic text dump.
func parseDraftjs(document []byte) ([]Block, error) {
	// Parse the document into our little internal structure
	var doc draftjsDocument
	if err := json.Unmarshal(document, &doc); err != nil {
		return nil, err
	}
	// Convert the fancy draftjs document into a plain text dump
	var blocks []Block
	for _, block := range doc.Blocks {
		// Extract any textual data
		if text := strings.TrimSpace(block.Text); text != "" {
			blocks = append(blocks, Block{Text: text})
		}
		// Extract any embedded media
		if block.Data.Files.source() != "" {
			blocks = append(blocks, Block{
				Image:   "https://ipfs.io/ipfs/" + block.Data.Files.source(),
				Caption: block.Data.Caption,
			})
		}
		// Extract any inlined media
		for _, ref := range block.Refs {
			if obj, ok := doc.Entities[strconv.Itoa(ref.Key)]; ok && obj.Type == "image" {
				if obj.Data.Files.source() != "" {
					blocks = append(blocks, Block{
						Image:   "https://ipfs.io/ipfs/" + obj.Data.Files.source(),
						Caption: obj.Data.Caption,
					})
				}
			}
		}
	}
	return blocks, nil
}

// draftjsDocument represents a draftjs document in its weird and gory format.
// Only those fields have been listed here which are of some use to the API.
type draftjsDocument struct {
	Entities map[string]draftjsEntity `json:"entityMap"`
	Blocks   []struct {
		Text string      `json:"text"`
		Data draftjsData `json:"data"`
		Refs []struct {
			Offset int `json:"offset"`
			Key    int `json:"key"`
		} `json:"entityRanges"`
	} `json:"blocks"`
}

// draftjsEntity is a single complex entity that is can be referenced from within
// one of the text blocks.
type draftjsEntity struct {
	Type string      `json:"type"`
	Data draftjsData `json:"data"`
}

// draftjsData is a multi-purpose media object that can be embedded into a draftjs
// document. We don't care about most of the forms this may take, only a subset
// that makes sense from a programatic API perspective.
type draftjsData struct {
	Caption string `json:"caption"`
	Files   image  `json:"files"`
}
