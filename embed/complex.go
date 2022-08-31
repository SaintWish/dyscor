// Credit to Necroforger https://gist.github.com/Necroforger/8b0b70b1a69fa7828b8ad6387ebb3835
// Modified to allow self hosted images.
package embed

import (
  	"io"

  	"github.com/bwmarrin/discordgo"
)

type ComplexEmbed struct {
	*discordgo.MessageSend
}

func NewComplexEmbed() *ComplexEmbed {
  	obj := &ComplexEmbed{&discordgo.MessageSend{}}
  	obj.Embed = &discordgo.MessageEmbed{}

	return obj
}

//SetTitle ...
func (e *ComplexEmbed) SetTitle(name string) *ComplexEmbed {
	e.Embed.Title = name
	return e
}

//SetDescription [desc]
func (e *ComplexEmbed) SetDescription(description string) *ComplexEmbed {
	if len(description) > 2048 {
		description = description[:2048]
	}
	e.Embed.Description = description
	return e
}

//AddField [name] [value]
func (e *ComplexEmbed) AddField(name, value string) *ComplexEmbed {
	fields := make([]*discordgo.MessageEmbedField, 0)

	if len(name) > EmbedLimitFieldName {
		name = name[:EmbedLimitFieldName]
	}

	if len(value) > EmbedLimitFieldValue {
		i := EmbedLimitFieldValue
		extended := false
		for i = EmbedLimitFieldValue; i < len(value); {
			if i != EmbedLimitFieldValue && extended == false {
				name += " (extended)"
				extended = true
			}
			if value[i] == []byte(" ")[0] || value[i] == []byte("\n")[0] || value[i] == []byte("-")[0] {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:  name,
					Value: value[i-EmbedLimitFieldValue : i],
				})
			} else {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:  name,
					Value: value[i-EmbedLimitFieldValue:i-1] + "-",
				})
				i--
			}

			if (i + EmbedLimitFieldValue) < len(value) {
				i += EmbedLimitFieldValue
			} else {
				break
			}
		}
		if i < len(value) {
			name += " (extended)"
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  name,
				Value: value[i:],
			})
		}
	} else {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  name,
			Value: value,
		})
	}

	e.Embed.Fields = append(e.Embed.Fields, fields...)
	return e
}

//SetFooter [Text] [iconURL]
func (e *ComplexEmbed) SetFooter(args ...string) *ComplexEmbed {
	iconURL := ""
	text := ""
	proxyURL := ""

	switch {
	case len(args) > 2:
		proxyURL = args[2]
		fallthrough
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		text = args[0]
	case len(args) == 0:
		return e
	}

	e.Embed.Footer = &discordgo.MessageEmbedFooter{
		IconURL:      iconURL,
		Text:         text,
		ProxyIconURL: proxyURL,
	}

	return e
}

//SetImage ...
func (e *ComplexEmbed) SetImage(args ...string) *ComplexEmbed {
	var URL string
	var proxyURL string

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		URL = args[0]
	}
	if len(args) > 1 {
		proxyURL = args[1]
	}

	e.Embed.Image = &discordgo.MessageEmbedImage{
		URL:      URL,
		ProxyURL: proxyURL,
	}

	return e
}

//SetLocalImage ...
func (e *ComplexEmbed) SetLocalImage(buffer io.Reader, filename string) *ComplexEmbed {
	e.Embed.Image = &discordgo.MessageEmbedImage{
		URL: "attachment://"+filename,
	}

	e.Files = []*discordgo.File{
		{
			Name: filename,
			Reader: buffer,
		},
	}

	return e
}

//SetThumbnail ...
func (e *ComplexEmbed) SetThumbnail(args ...string) *ComplexEmbed {
	var URL string
	var proxyURL string

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		URL = args[0]
	}
	if len(args) > 1 {
		proxyURL = args[1]
	}
	e.Embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

//SetAuthor ...
func (e *ComplexEmbed) SetAuthor(args ...string) *ComplexEmbed {
	var (
		name     string
		iconURL  string
		URL      string
		proxyURL string
	)

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		name = args[0]
	}
	if len(args) > 1 {
		iconURL = args[1]
	}
	if len(args) > 2 {
		URL = args[2]
	}
	if len(args) > 3 {
		proxyURL = args[3]
	}

	e.Embed.Author = &discordgo.MessageEmbedAuthor{
		Name:         name,
		IconURL:      iconURL,
		URL:          URL,
		ProxyIconURL: proxyURL,
	}

	return e
}

//SetURL ...
func (e *ComplexEmbed) SetURL(URL string) *ComplexEmbed {
	e.Embed.URL = URL
	return e
}

//SetColor ...
func (e *ComplexEmbed) SetColor(clr int) *ComplexEmbed {
	e.Embed.Color = clr
	return e
}

// InlineAllFields sets all fields in the embed to be inline
func (e *ComplexEmbed) InlineAllFields() *ComplexEmbed {
	for _, v := range e.Embed.Fields {
		v.Inline = true
	}
	return e
}

// Truncate truncates any embed value over the character limit.
func (e *ComplexEmbed) Truncate() *ComplexEmbed {
	e.TruncateDescription()
	e.TruncateFields()
	e.TruncateFooter()
	e.TruncateTitle()
	return e
}

// TruncateFields truncates fields that are too long
func (e *ComplexEmbed) TruncateFields() *ComplexEmbed {
	if len(e.Embed.Fields) > 25 {
		e.Embed.Fields = e.Embed.Fields[:EmbedLimitField]
	}

	for _, v := range e.Embed.Fields {
		if len(v.Name) > EmbedLimitFieldName {
			v.Name = v.Name[:EmbedLimitFieldName]
		}

		if len(v.Value) > EmbedLimitFieldValue {
			v.Value = v.Value[:EmbedLimitFieldValue]
		}
	}

	return e
}

// TruncateDescription ...
func (e *ComplexEmbed) TruncateDescription() *ComplexEmbed {
	if len(e.Embed.Description) > EmbedLimitDescription {
		e.Embed.Description = e.Embed.Description[:EmbedLimitDescription]
	}

	return e
}

// TruncateTitle ...
func (e *ComplexEmbed) TruncateTitle() *ComplexEmbed {
	if len(e.Embed.Title) > EmbedLimitTitle {
		e.Embed.Title = e.Embed.Title[:EmbedLimitTitle]
	}

	return e
}

// TruncateFooter ...
func (e *ComplexEmbed) TruncateFooter() *ComplexEmbed {
	if e.Embed.Footer != nil && len(e.Embed.Footer.Text) > EmbedLimitFooter {
		e.Embed.Footer.Text = e.Embed.Footer.Text[:EmbedLimitFooter]
	}

	return e
}