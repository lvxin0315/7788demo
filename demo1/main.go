package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
	"io/ioutil"
)

var (
	red   = &color.RGBA{R: 255, G: 0, B: 0, A: 255}
	green = &color.RGBA{R: 0, G: 255, B: 0, A: 255}
	blue  = &color.RGBA{R: 0, G: 0, B: 255, A: 255}
)

//实现fyne.resource,这是一个小图标
type demo1Icon struct{}

func (t *demo1Icon) Name() string {
	return "test21_icon"
}

func (t *demo1Icon) Content() []byte {
	iconB, _ := ioutil.ReadFile("tmp/20200610155355.png")
	return iconB
}

//字体
type demo1Font struct {
}

func (t *demo1Font) Name() string {
	return "PuHui"
}

func (t *demo1Font) Content() []byte {
	b, _ := ioutil.ReadFile("tmp/Alibaba-PuHuiTi-Medium.ttf")
	return b
}

//加粗字体

type demo1BoldFont struct {
}

func (t *demo1BoldFont) Name() string {
	return "PuHui"
}

func (t *demo1BoldFont) Content() []byte {
	b, _ := ioutil.ReadFile("tmp/Alibaba-PuHuiTi-Bold.ttf")
	return b
}

//自定义主题
type demo1Theme struct{}

func (demo1Theme) BackgroundColor() color.Color {
	return red
}

func (demo1Theme) ButtonColor() color.Color {
	return color.Black
}

func (demo1Theme) DisabledButtonColor() color.Color {
	return color.White
}

func (demo1Theme) DisabledIconColor() color.Color {
	return color.Black
}

func (demo1Theme) DisabledTextColor() color.Color {
	return color.Black
}

func (demo1Theme) FocusColor() color.Color {
	return green
}

func (demo1Theme) HoverColor() color.Color {
	return green
}

func (demo1Theme) HyperlinkColor() color.Color {
	return green
}

func (demo1Theme) IconColor() color.Color {
	return color.White
}

func (demo1Theme) IconInlineSize() int {
	return 24
}

func (demo1Theme) Padding() int {
	return 10
}

func (demo1Theme) PlaceHolderColor() color.Color {
	return blue
}

func (demo1Theme) PrimaryColor() color.Color {
	return green
}

func (demo1Theme) ScrollBarColor() color.Color {
	return blue
}

func (demo1Theme) ScrollBarSize() int {
	return 10
}

func (demo1Theme) ScrollBarSmallSize() int {
	return 2
}

func (demo1Theme) ShadowColor() color.Color {
	return blue
}

func (demo1Theme) TextBoldFont() fyne.Resource {
	return &demo1BoldFont{}
}

func (demo1Theme) TextBoldItalicFont() fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func (demo1Theme) TextColor() color.Color {
	return color.White
}

func (demo1Theme) TextFont() fyne.Resource {
	return &demo1Font{}
}

func (demo1Theme) TextItalicFont() fyne.Resource {
	return theme.DefaultTextBoldItalicFont()
}

func (demo1Theme) TextMonospaceFont() fyne.Resource {
	return theme.DefaultTextFont()
}

func (demo1Theme) TextSize() int {
	return 24
}

func main() {
	//初始化一个app
	a := app.New()
	//设置自己的主题，虽然有点丑，但是能解决中文的问题
	a.Settings().SetTheme(&demo1Theme{})
	//初始化一个窗口
	window := a.NewWindow("Hello")
	//设置一下属性
	window.SetTitle("我是title")
	window.Resize(fyne.Size{
		Width:  300,
		Height: 600,
	})
	//一个label组件
	hello := widget.NewLabel("Hello World!")
	btn := widget.NewButtonWithIcon("按钮", &demo1Icon{}, func() {
		//点击之后，把hello的内容修改了
		hello.SetText("我是点击button后的label内容")
	})

	//NewHBox使用指定的子对象列表创建一个新的水平对齐的box小部件
	//hBox := widget.NewHBox()
	//hBox.Children = append(hBox.Children, hello)
	//hBox.Children = append(hBox.Children, btn)

	//NewVBox使用指定的子对象列表创建一个新的垂直对齐框小部件
	vBox := widget.NewVBox()
	vBox.Children = append(vBox.Children, hello)
	vBox.Children = append(vBox.Children, btn)
	//把内容放到窗口
	window.SetContent(vBox)
	window.ShowAndRun()
}
