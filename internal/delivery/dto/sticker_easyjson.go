// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

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

func easyjsonA68a6153DecodeDepecheInternalDeliveryDto(in *jlexer.Lexer, out *UploadStickerPack) {
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
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "cover":
			out.Cover = string(in.String())
		case "stickers":
			if in.IsNull() {
				in.Skip()
				out.Stickers = nil
			} else {
				in.Delim('[')
				if out.Stickers == nil {
					if !in.IsDelim(']') {
						out.Stickers = make([]string, 0, 4)
					} else {
						out.Stickers = []string{}
					}
				} else {
					out.Stickers = (out.Stickers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
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
func easyjsonA68a6153EncodeDepecheInternalDeliveryDto(out *jwriter.Writer, in UploadStickerPack) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"cover\":"
		out.RawString(prefix)
		out.String(string(in.Cover))
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
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UploadStickerPack) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA68a6153EncodeDepecheInternalDeliveryDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UploadStickerPack) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA68a6153EncodeDepecheInternalDeliveryDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UploadStickerPack) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA68a6153DecodeDepecheInternalDeliveryDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UploadStickerPack) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA68a6153DecodeDepecheInternalDeliveryDto(l, v)
}
func easyjsonA68a6153DecodeDepecheInternalDeliveryDto1(in *jlexer.Lexer, out *AddStickerPack) {
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
		case "stickerpack_id":
			out.StickerPackId = uint(in.Uint())
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
func easyjsonA68a6153EncodeDepecheInternalDeliveryDto1(out *jwriter.Writer, in AddStickerPack) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"stickerpack_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.StickerPackId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddStickerPack) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA68a6153EncodeDepecheInternalDeliveryDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddStickerPack) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA68a6153EncodeDepecheInternalDeliveryDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddStickerPack) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA68a6153DecodeDepecheInternalDeliveryDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddStickerPack) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA68a6153DecodeDepecheInternalDeliveryDto1(l, v)
}