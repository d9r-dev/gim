package main

import (
	"strings"
)

type Piece struct {
	// Links to the previous and next textAreaSpan objects as indices into the
	// [TextArea.spans] slice. The sentinel spans (index 0 and 1) have -1 as
	// their previous or next links, respectively.
	previous, next int

	// The start index and the length of the text segment this span represents.
	// If "length" is negative, the span represents a substring of
	// [TextArea.initialText] and the actual length is its absolute value. If it
	// is positive, the span represents a substring of [TextArea.editText]. For
	// the sentinel spans (index 0 and 1), both values will be 0. Others will
	// never have a zero length.
	offset, length int
}

type TextPieceTable struct {
	originBuffer string
	addBuffer    strings.Builder
	pieces       []Piece
	length       int
}

func NewTextEditor() *TextPieceTable {
	t := &TextPieceTable{
		pieces: make([]Piece, 2),
	}
	t.pieces[0] = Piece{previous: -1, next: 1}
	t.pieces[1] = Piece{previous: 0, next: -1}

	return t
}

func (t *TextPieceTable) SetText(text string) *TextPieceTable {
	t.pieces = t.pieces[:2]
	t.originBuffer = text
	t.length = len(text)

	if len(text) > 0 {
		t.pieces = append(t.pieces, Piece{
			previous: 0,
			next:     1,
			offset:   0,
			length:   -len(text),
		})
		t.pieces[0].next = 2
		t.pieces[1].previous = 2
	} else {
		t.pieces[0].next = 1
		t.pieces[1].previous = 0
	}

	return t
}

func (t *TextPieceTable) GetText() string {
	if t.length == 0 {
		return ""
	}

	var text strings.Builder
	pieceIndex := t.pieces[0].next
	for pieceIndex != 1 {
		piece := &t.pieces[pieceIndex]
		if piece.length < 0 {
			text.WriteString(t.originBuffer[piece.offset : piece.offset-piece.length])
		} else {
			text.WriteString(t.addBuffer.String()[piece.offset : piece.offset+piece.length])
		}
		pieceIndex = t.pieces[pieceIndex].next
	}

	return text.String()
}
