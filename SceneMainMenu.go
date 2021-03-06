package OpenDiablo2

import (
	"image/color"

	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/hajimehoshi/ebiten"
)

type MainMenu struct {
	Engine              *Engine
	TrademarkBackground Sprite
	Background          Sprite
	DiabloLogoLeft      Sprite
	DiabloLogoRight     Sprite
	DiabloLogoLeftBack  Sprite
	DiabloLogoRightBack Sprite
	CopyrightLabel      *UILabel
	CopyrightLabel2     *UILabel
	ShowTrademarkScreen bool
}

func CreateMainMenu(engine *Engine) *MainMenu {
	result := &MainMenu{
		Engine:              engine,
		ShowTrademarkScreen: true,
	}

	return result
}

func (v *MainMenu) Load() {
	go func() {
		loadStep := 1.0 / 8.0
		v.Engine.LoadingProgress = 0
		{
			v.CopyrightLabel = CreateUILabel(v.Engine, ResourcePaths.FontFormal12, "static")
			v.CopyrightLabel.Alignment = UILabelAlignCenter
			v.CopyrightLabel.SetText("Diablo 2 is © Copyright 2000-2016 Blizzard Entertainment")
			v.CopyrightLabel.ColorMod = color.RGBA{188, 168, 140, 255}
			v.CopyrightLabel.MoveTo(400, 500)
			v.Engine.LoadingProgress += loadStep
		}
		{
			v.CopyrightLabel2 = CreateUILabel(v.Engine, ResourcePaths.FontFormal12, "static")
			v.CopyrightLabel2.Alignment = UILabelAlignCenter
			v.CopyrightLabel2.SetText("All Rights Reserved.")
			v.CopyrightLabel2.ColorMod = color.RGBA{188, 168, 140, 255}
			v.CopyrightLabel2.MoveTo(400, 525)
			v.Engine.LoadingProgress += loadStep
		}
		{
			v.Background = v.Engine.LoadSprite(ResourcePaths.GameSelectScreen, v.Engine.Palettes["sky"])
			v.Background.MoveTo(0, 0)
			v.Engine.LoadingProgress += loadStep
		}
		{
			v.TrademarkBackground = v.Engine.LoadSprite(ResourcePaths.TrademarkScreen, v.Engine.Palettes["sky"])
			v.TrademarkBackground.MoveTo(0, 0)
			v.Engine.LoadingProgress += loadStep
		}
		{
			v.DiabloLogoLeft = v.Engine.LoadSprite(ResourcePaths.Diablo2LogoFireLeft, v.Engine.Palettes["units"])
			v.DiabloLogoLeft.Blend = true
			v.DiabloLogoLeft.Animate = true
			v.DiabloLogoLeft.MoveTo(400, 120)
			v.Engine.LoadingProgress += loadStep
		}
		{
			v.DiabloLogoRight = v.Engine.LoadSprite(ResourcePaths.Diablo2LogoFireRight, v.Engine.Palettes["units"])
			v.DiabloLogoRight.Blend = true
			v.DiabloLogoRight.Animate = true
			v.DiabloLogoRight.MoveTo(400, 120)
			v.Engine.LoadingProgress += loadStep
		}
		{
			v.DiabloLogoLeftBack = v.Engine.LoadSprite(ResourcePaths.Diablo2LogoBlackLeft, v.Engine.Palettes["units"])
			v.DiabloLogoLeftBack.MoveTo(400, 120)
			v.Engine.LoadingProgress += loadStep
		}
		{
			v.DiabloLogoRightBack = v.Engine.LoadSprite(ResourcePaths.Diablo2LogoBlackRight, v.Engine.Palettes["units"])
			v.DiabloLogoRightBack.MoveTo(400, 120)
			v.Engine.LoadingProgress = 1.0
		}
	}()
}

func (v *MainMenu) Unload() {

}

func (v *MainMenu) Render(screen *ebiten.Image) {
	if v.ShowTrademarkScreen {
		v.TrademarkBackground.DrawSegments(screen, 4, 3, 0)
	} else {
		v.Background.DrawSegments(screen, 4, 3, 0)
	}
	v.DiabloLogoLeftBack.Draw(screen)
	v.DiabloLogoRightBack.Draw(screen)
	v.DiabloLogoLeft.Draw(screen)
	v.DiabloLogoRight.Draw(screen)

	if v.ShowTrademarkScreen {
		v.CopyrightLabel.Draw(screen)
		v.CopyrightLabel2.Draw(screen)
	} else {

	}
}

func (v *MainMenu) Update() {

}
