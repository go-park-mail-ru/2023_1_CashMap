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

func easyjson1c045807DecodeDepecheInternalEntities(in *jlexer.Lexer, out *GroupManagement) {
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
		case "link":
			out.Link = string(in.String())
		case "role":
			out.Role = string(in.String())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
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
func easyjson1c045807EncodeDepecheInternalEntities(out *jwriter.Writer, in GroupManagement) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"link\":"
		out.RawString(prefix[1:])
		out.String(string(in.Link))
	}
	{
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		out.String(string(in.Role))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GroupManagement) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1c045807EncodeDepecheInternalEntities(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GroupManagement) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1c045807EncodeDepecheInternalEntities(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GroupManagement) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1c045807DecodeDepecheInternalEntities(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GroupManagement) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1c045807DecodeDepecheInternalEntities(l, v)
}
func easyjson1c045807DecodeDepecheInternalEntities1(in *jlexer.Lexer, out *Group) {
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
		case "link":
			out.Link = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "info":
			out.Info = string(in.String())
		case "privacy":
			out.Privacy = string(in.String())
		case "avatar_url":
			out.Avatar = string(in.String())
		case "avg_avatar_color":
			out.AvgAvatarColor = string(in.String())
		case "subscribers":
			out.MembersCount = int(in.Int())
		case "creation_date":
			out.CreationDate = string(in.String())
		case "hide_owner":
			out.HideOwner = bool(in.Bool())
		case "is_deleted":
			out.IsDeleted = bool(in.Bool())
		case "management":
			if in.IsNull() {
				in.Skip()
				out.Management = nil
			} else {
				in.Delim('[')
				if out.Management == nil {
					if !in.IsDelim(']') {
						out.Management = make([]GroupManagement, 0, 0)
					} else {
						out.Management = []GroupManagement{}
					}
				} else {
					out.Management = (out.Management)[:0]
				}
				for !in.IsDelim(']') {
					var v1 GroupManagement
					(v1).UnmarshalEasyJSON(in)
					out.Management = append(out.Management, v1)
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
func easyjson1c045807EncodeDepecheInternalEntities1(out *jwriter.Writer, in Group) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"link\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Link))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"info\":"
		out.RawString(prefix)
		out.String(string(in.Info))
	}
	{
		const prefix string = ",\"privacy\":"
		out.RawString(prefix)
		out.String(string(in.Privacy))
	}
	{
		const prefix string = ",\"avatar_url\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"avg_avatar_color\":"
		out.RawString(prefix)
		out.String(string(in.AvgAvatarColor))
	}
	{
		const prefix string = ",\"subscribers\":"
		out.RawString(prefix)
		out.Int(int(in.MembersCount))
	}
	{
		const prefix string = ",\"creation_date\":"
		out.RawString(prefix)
		out.String(string(in.CreationDate))
	}
	{
		const prefix string = ",\"hide_owner\":"
		out.RawString(prefix)
		out.Bool(bool(in.HideOwner))
	}
	{
		const prefix string = ",\"is_deleted\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsDeleted))
	}
	{
		const prefix string = ",\"management\":"
		out.RawString(prefix)
		if in.Management == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Management {
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
func (v Group) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1c045807EncodeDepecheInternalEntities1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Group) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1c045807EncodeDepecheInternalEntities1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Group) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1c045807DecodeDepecheInternalEntities1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Group) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1c045807DecodeDepecheInternalEntities1(l, v)
}
