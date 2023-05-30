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

func easyjson4086215fDecodeDepecheInternalDeliveryDto(in *jlexer.Lexer, out *NewMessageDTO) {
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
			out.ChatId = uint(in.Uint())
		case "message_content_type":
			out.ContentType = string(in.String())
		case "sticker_id":
			if in.IsNull() {
				in.Skip()
				out.StickerID = nil
			} else {
				if out.StickerID == nil {
					out.StickerID = new(uint)
				}
				*out.StickerID = uint(in.Uint())
			}
		case "text_content":
			out.Text = string(in.String())
		case "reply_to":
			if in.IsNull() {
				in.Skip()
				out.ReplyTo = nil
			} else {
				if out.ReplyTo == nil {
					out.ReplyTo = new(uint)
				}
				*out.ReplyTo = uint(in.Uint())
			}
		case "attachments":
			if in.IsNull() {
				in.Skip()
				out.Attachments = nil
			} else {
				in.Delim('[')
				if out.Attachments == nil {
					if !in.IsDelim(']') {
						out.Attachments = make([]string, 0, 4)
					} else {
						out.Attachments = []string{}
					}
				} else {
					out.Attachments = (out.Attachments)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Attachments = append(out.Attachments, v1)
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
func easyjson4086215fEncodeDepecheInternalDeliveryDto(out *jwriter.Writer, in NewMessageDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"chat_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.ChatId))
	}
	{
		const prefix string = ",\"message_content_type\":"
		out.RawString(prefix)
		out.String(string(in.ContentType))
	}
	{
		const prefix string = ",\"sticker_id\":"
		out.RawString(prefix)
		if in.StickerID == nil {
			out.RawString("null")
		} else {
			out.Uint(uint(*in.StickerID))
		}
	}
	{
		const prefix string = ",\"text_content\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"reply_to\":"
		out.RawString(prefix)
		if in.ReplyTo == nil {
			out.RawString("null")
		} else {
			out.Uint(uint(*in.ReplyTo))
		}
	}
	{
		const prefix string = ",\"attachments\":"
		out.RawString(prefix)
		if in.Attachments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Attachments {
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
func (v NewMessageDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeDepecheInternalDeliveryDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NewMessageDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeDepecheInternalDeliveryDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NewMessageDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeDepecheInternalDeliveryDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NewMessageDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeDepecheInternalDeliveryDto(l, v)
}
func easyjson4086215fDecodeDepecheInternalDeliveryDto1(in *jlexer.Lexer, out *HasDialogDTO) {
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
		case "user_link":
			if in.IsNull() {
				in.Skip()
				out.UserLink = nil
			} else {
				if out.UserLink == nil {
					out.UserLink = new(string)
				}
				*out.UserLink = string(in.String())
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
func easyjson4086215fEncodeDepecheInternalDeliveryDto1(out *jwriter.Writer, in HasDialogDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_link\":"
		out.RawString(prefix[1:])
		if in.UserLink == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.UserLink))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HasDialogDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeDepecheInternalDeliveryDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HasDialogDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeDepecheInternalDeliveryDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HasDialogDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeDepecheInternalDeliveryDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HasDialogDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeDepecheInternalDeliveryDto1(l, v)
}
func easyjson4086215fDecodeDepecheInternalDeliveryDto2(in *jlexer.Lexer, out *GetMessagesDTO) {
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
			if in.IsNull() {
				in.Skip()
				out.ChatID = nil
			} else {
				if out.ChatID == nil {
					out.ChatID = new(uint)
				}
				*out.ChatID = uint(in.Uint())
			}
		case "batch_size":
			if in.IsNull() {
				in.Skip()
				out.BatchSize = nil
			} else {
				if out.BatchSize == nil {
					out.BatchSize = new(uint)
				}
				*out.BatchSize = uint(in.Uint())
			}
		case "last_msg_date":
			if in.IsNull() {
				in.Skip()
				out.LastMessageDate = nil
			} else {
				if out.LastMessageDate == nil {
					out.LastMessageDate = new(string)
				}
				*out.LastMessageDate = string(in.String())
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
func easyjson4086215fEncodeDepecheInternalDeliveryDto2(out *jwriter.Writer, in GetMessagesDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"chat_id\":"
		out.RawString(prefix[1:])
		if in.ChatID == nil {
			out.RawString("null")
		} else {
			out.Uint(uint(*in.ChatID))
		}
	}
	{
		const prefix string = ",\"batch_size\":"
		out.RawString(prefix)
		if in.BatchSize == nil {
			out.RawString("null")
		} else {
			out.Uint(uint(*in.BatchSize))
		}
	}
	{
		const prefix string = ",\"last_msg_date\":"
		out.RawString(prefix)
		if in.LastMessageDate == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.LastMessageDate))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetMessagesDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeDepecheInternalDeliveryDto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMessagesDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeDepecheInternalDeliveryDto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMessagesDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeDepecheInternalDeliveryDto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMessagesDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeDepecheInternalDeliveryDto2(l, v)
}
func easyjson4086215fDecodeDepecheInternalDeliveryDto3(in *jlexer.Lexer, out *GetChatsDTO) {
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
		case "offset":
			if in.IsNull() {
				in.Skip()
				out.Offset = nil
			} else {
				if out.Offset == nil {
					out.Offset = new(uint)
				}
				*out.Offset = uint(in.Uint())
			}
		case "batch_size":
			if in.IsNull() {
				in.Skip()
				out.BatchSize = nil
			} else {
				if out.BatchSize == nil {
					out.BatchSize = new(uint)
				}
				*out.BatchSize = uint(in.Uint())
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
func easyjson4086215fEncodeDepecheInternalDeliveryDto3(out *jwriter.Writer, in GetChatsDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"offset\":"
		out.RawString(prefix[1:])
		if in.Offset == nil {
			out.RawString("null")
		} else {
			out.Uint(uint(*in.Offset))
		}
	}
	{
		const prefix string = ",\"batch_size\":"
		out.RawString(prefix)
		if in.BatchSize == nil {
			out.RawString("null")
		} else {
			out.Uint(uint(*in.BatchSize))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetChatsDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeDepecheInternalDeliveryDto3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetChatsDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeDepecheInternalDeliveryDto3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetChatsDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeDepecheInternalDeliveryDto3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetChatsDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeDepecheInternalDeliveryDto3(l, v)
}
func easyjson4086215fDecodeDepecheInternalDeliveryDto4(in *jlexer.Lexer, out *CreateChatDTO) {
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
		case "user_links":
			if in.IsNull() {
				in.Skip()
				out.UserLinks = nil
			} else {
				in.Delim('[')
				if out.UserLinks == nil {
					if !in.IsDelim(']') {
						out.UserLinks = make([]string, 0, 4)
					} else {
						out.UserLinks = []string{}
					}
				} else {
					out.UserLinks = (out.UserLinks)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.UserLinks = append(out.UserLinks, v4)
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
func easyjson4086215fEncodeDepecheInternalDeliveryDto4(out *jwriter.Writer, in CreateChatDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_links\":"
		out.RawString(prefix[1:])
		if in.UserLinks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.UserLinks {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateChatDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeDepecheInternalDeliveryDto4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateChatDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeDepecheInternalDeliveryDto4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateChatDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeDepecheInternalDeliveryDto4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateChatDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeDepecheInternalDeliveryDto4(l, v)
}