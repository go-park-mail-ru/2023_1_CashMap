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

func easyjson9b8f5552DecodeDepecheInternalEntities(in *jlexer.Lexer, out *Chat) {
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
		case "chat_id":
			out.ChatID = uint(in.Uint())
		case "members":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]*UserInfo, 0, 8)
					} else {
						out.Users = []*UserInfo{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *UserInfo
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(UserInfo)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Users = append(out.Users, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "read":
			out.Read = bool(in.Bool())
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
func easyjson9b8f5552EncodeDepecheInternalEntities(out *jwriter.Writer, in Chat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"chat_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ChatID))
	}
	{
		const prefix string = ",\"members\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Users {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"read\":"
		out.RawString(prefix)
		out.Bool(bool(in.Read))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Chat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b8f5552EncodeDepecheInternalEntities(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Chat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b8f5552EncodeDepecheInternalEntities(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Chat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b8f5552DecodeDepecheInternalEntities(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Chat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b8f5552DecodeDepecheInternalEntities(l, v)
}
