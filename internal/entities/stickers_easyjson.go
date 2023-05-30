// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entities

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson687b2e5cDecodeDepecheInternalEntities(in *jlexer.Lexer, out *StickerPack) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint(in.Uint())
		case "title":
			out.Title = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "depeche_authored":
			out.DepecheAuthored = bool(in.Bool())
		case "cover":
			out.Cover = string(in.String())
		case "creation_date":
			out.CreationDate = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "stickers":
			if in.IsNull() {
				in.Skip()
				out.Stickers = nil
			} else {
				in.Delim('[')
				if out.Stickers == nil {
					if !in.IsDelim(']') {
						out.Stickers = make([]Sticker, 0, 2)
					} else {
						out.Stickers = []Sticker{}
					}
				} else {
					out.Stickers = (out.Stickers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Sticker
					(v1).UnmarshalEasyJSON(in)
					out.Stickers = append(out.Stickers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson687b2e5cEncodeDepecheInternalEntities(out *jwriter.Writer, in StickerPack) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"depeche_authored\":"
		out.RawString(prefix)
		out.Bool(bool(in.DepecheAuthored))
	}
	{
		const prefix string = ",\"cover\":"
		out.RawString(prefix)
		out.String(string(in.Cover))
	}
	{
		const prefix string = ",\"creation_date\":"
		out.RawString(prefix)
		out.String(string(in.CreationDate))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"stickers\":"
		out.RawString(prefix)
		if in.Stickers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Stickers {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StickerPack) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson687b2e5cEncodeDepecheInternalEntities(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StickerPack) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson687b2e5cEncodeDepecheInternalEntities(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StickerPack) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson687b2e5cDecodeDepecheInternalEntities(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StickerPack) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson687b2e5cDecodeDepecheInternalEntities(l, v)
}
func easyjson687b2e5cDecodeDepecheInternalEntities1(in *jlexer.Lexer, out *Sticker) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint(in.Uint())
		case "url":
			out.Url = string(in.String())
		case "stickerpack_id":
			out.StickerpackID = uint(in.Uint())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson687b2e5cEncodeDepecheInternalEntities1(out *jwriter.Writer, in Sticker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.Url))
	}
	{
		const prefix string = ",\"stickerpack_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.StickerpackID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Sticker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson687b2e5cEncodeDepecheInternalEntities1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Sticker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson687b2e5cEncodeDepecheInternalEntities1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Sticker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson687b2e5cDecodeDepecheInternalEntities1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Sticker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson687b2e5cDecodeDepecheInternalEntities1(l, v)
}