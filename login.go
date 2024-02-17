package main

import (
	"errors"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	errDialogOnEmptyAccountOrPassword   = errors.New("empty account or password")
	errDialogOnInvalidAccountOrPassword = errors.New("invalid account or password")

	colorGreen = color.NRGBA{R: 7, G: 253, B: 73, A: 255}
)

const (
	appTitle = "FyLogin"

	rememberMeKey = "remember_me"
)

type LoginWindow struct {
	app    fyne.App
	window fyne.Window

	mainContainer *fyne.Container

	welcomeText *canvas.Text

	accountEntry  *widget.Entry
	passwordEntry *widget.Entry

	supportContainer   *fyne.Container
	rememberMeCheck    *widget.Check
	forgotPasswordLink *widget.Hyperlink

	loginButton *widget.Button

	registerContainer *fyne.Container
	registerLabel     *widget.Label
	registerLink      *widget.Hyperlink
}

func NewLoginWindow(app fyne.App) *LoginWindow {
	loginWindow := app.NewWindow(appTitle)
	loginWindow.CenterOnScreen()
	loginWindow.SetFixedSize(true)

	return &LoginWindow{
		app:    app,
		window: loginWindow,
	}
}

func (lw *LoginWindow) BuildAndShow() {
	lw.setMainContainer()
	lw.window.SetContent(lw.mainContainer)
	lw.window.Show()
}

func (lw *LoginWindow) setMainContainer() {
	var objs []fyne.CanvasObject

	lw.setWelcomeText()
	lw.setAccountEntry()
	lw.setPasswordEntry()
	lw.setSupportContainer()
	lw.setLoginButton()
	lw.setRegisterContainer()

	objs = append(
		objs,
		lw.welcomeText,
		lw.accountEntry,
		lw.passwordEntry,
		lw.supportContainer,
		lw.loginButton,
		lw.registerContainer,
	)

	lw.mainContainer = container.NewGridWithRows(len(objs), objs...)

	// Is possible to use VBox or HBox
	// 	lw.mainContainer = container.NewVBox(
	//      lw.welcomeText,
	// 		lw.accountEntry,
	// 		lw.passwordEntry,
	// 		lw.supportContainer,
	// 		lw.loginButton,
	// 		lw.registerContainer
	// )
}

func (lw *LoginWindow) setWelcomeText() {
	lw.welcomeText = canvas.NewText("Welcome To FyLogin", colorGreen)
	lw.welcomeText.Alignment = fyne.TextAlignCenter
	lw.welcomeText.TextSize = 18
	lw.welcomeText.TextStyle.Bold = true
}

func (lw *LoginWindow) setAccountEntry() {
	lw.accountEntry = widget.NewEntry()
	lw.accountEntry.SetPlaceHolder("Account")
}

func (lw *LoginWindow) setPasswordEntry() {
	lw.passwordEntry = widget.NewPasswordEntry()
	lw.passwordEntry.SetPlaceHolder("Password")
}

func (lw *LoginWindow) setSupportContainer() {
	var objs []fyne.CanvasObject

	lw.setRememberMeCheck()
	lw.setForgotPasswordLink()

	objs = append(objs, lw.rememberMeCheck, lw.forgotPasswordLink)

	lw.supportContainer = container.NewGridWithColumns(len(objs), objs...)
}

func (lw *LoginWindow) setRememberMeCheck() {
	lw.rememberMeCheck = widget.NewCheck("Remember me", func(b bool) {
		lw.app.Preferences().SetBool(rememberMeKey, b)
	})
	lw.rememberMeCheck.Checked = lw.app.Preferences().BoolWithFallback(rememberMeKey, false)
}

func (lw *LoginWindow) setForgotPasswordLink() {
	lw.forgotPasswordLink = widget.NewHyperlink("Forgot Password?", nil)
	lw.forgotPasswordLink.Alignment = fyne.TextAlignCenter
	lw.forgotPasswordLink.TextStyle.Bold = true

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("e-mail")

	lw.forgotPasswordLink.OnTapped = func() {
		dialog.NewCustomConfirm("Forgot Password", "Ok", "Cancel", emailEntry, func(b bool) {}, lw.window).Show()
	}
}

func (lw *LoginWindow) setLoginButton() {
	lw.loginButton = widget.NewButtonWithIcon("Login", theme.LoginIcon(), func() { login(lw) })
	lw.loginButton.Alignment = widget.ButtonAlignCenter
	lw.loginButton.Importance = widget.SuccessImportance
}

func (lw *LoginWindow) setRegisterContainer() {
	var objs []fyne.CanvasObject

	lw.setRegisterLabel()
	lw.setRegisterLink()

	objs = append(objs, lw.registerLabel, lw.registerLink)

	lw.registerContainer = container.NewGridWithColumns(len(objs), objs...)
}

func (lw *LoginWindow) setRegisterLabel() {
	lw.registerLabel = widget.NewLabel("Don't have an account?")
	lw.registerLabel.Alignment = fyne.TextAlignCenter
}

func (lw *LoginWindow) setRegisterLink() {
	lw.registerLink = widget.NewHyperlink("Register", nil)
	lw.registerLink.Alignment = fyne.TextAlignCenter
	lw.registerLink.TextStyle.Bold = true
}

// THIS IS ONLY EXAMPLEEEEE
type Credentials struct {
	Account  string
	Password string
}

func login(lw *LoginWindow) {
	account := lw.accountEntry.Text
	password := lw.passwordEntry.Text

	var credentials = Credentials{
		Account:  "account",
		Password: "password",
	}

	if account == "" || password == "" {
		dialog.ShowError(errDialogOnEmptyAccountOrPassword, lw.window)
		return
	}

	if account == credentials.Account && password == credentials.Password {
		showLoggedWindow(lw)
	} else {
		dialog.ShowError(errDialogOnInvalidAccountOrPassword, lw.window)
	}
}

func showLoggedWindow(lw *LoginWindow) {
	w := lw.app.NewWindow(appTitle)
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 300))

	text := canvas.NewText("Congratulations, you are logged in!!!", colorGreen)
	text.TextSize = 14
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle.Bold = true

	w.SetContent(text)
	w.Show()

	lw.window.Close()
}
