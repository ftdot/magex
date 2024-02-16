package uitext

import (
	"image/color"

	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/mmath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var defaultDPI float64 = 72

type iTextTransform interface {
	SetUIText(*UIText)
}

////

type UIText struct {
	Transform interfaces.ITransform
	Text      string
	Font      *sfnt.Font
	Options   *interfaces.FontOptions
	fontFace  font.Face

	ID string
}

func New(tf interfaces.ITransform, text string, f *sfnt.Font, fontSize float64) (*UIText, error) {
	uit := &UIText{
		Transform: tf,
		Text:      text,
		Font:      f,
		Options:   &interfaces.FontOptions{FontColor: color.RGBA{255, 255, 255, 255}, FontSize: fontSize, DPI: defaultDPI},
		fontFace:  nil,
		ID:        utils.GenerateComponentID(),
	}
	err := uit.updateFontFace()
	if err != nil {
		return nil, err
	}
	if tt, ok := tf.(iTextTransform); ok {
		tt.SetUIText(uit)
	}
	return uit, nil
}

////

func (t *UIText) GetFontFace() font.Face {
	return t.fontFace
}

func (t *UIText) updateFontFace() error {
	fontFace, err := opentype.NewFace(t.Font, &opentype.FaceOptions{Size: t.Options.FontSize, DPI: t.Options.DPI, Hinting: font.HintingFull})
	if err != nil {
		return err
	}
	t.fontFace = fontFace
	return nil
}

func (t *UIText) SetColor(fontColor color.Color) {
	t.Options.FontColor = fontColor
}

func (t *UIText) SetDPI(dpi float64) error {
	t.Options.DPI = dpi
	return t.updateFontFace()
}

func (t *UIText) SetSize(fontSize float64) error {
	t.Options.FontSize = fontSize
	return t.updateFontFace()
}

func (t *UIText) GetOptions() *interfaces.FontOptions {
	return t.Options
}

func (t *UIText) SetOptions(fontOptions *interfaces.FontOptions) error {
	t.Options = fontOptions
	return t.updateFontFace()
}

////

func (t *UIText) DrawUIPriority() float64 {
	return t.Transform.GetLayer()
}

func (t *UIText) DrawUI(gb interfaces.IGameBase, screen *ebiten.Image) {

	var opts ebiten.DrawImageOptions

	opts.GeoM.Scale(t.Transform.GetScale().X, t.Transform.GetScale().Y)
	opts.GeoM.Rotate(t.Transform.GetRotation() * mmath.RadiansMeasurement90)
	opts.GeoM.Translate(t.Transform.GetPosition().X, t.Transform.GetPosition().Y)

	opts.ColorScale.ScaleWithColor(t.Options.FontColor)

	text.DrawWithOptions(screen, t.Text, t.fontFace, &opts)
}

////

func (t *UIText) GetID() string {
	return t.ID
}
