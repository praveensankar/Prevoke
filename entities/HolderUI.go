package entities

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/praveensankar/Revocation-Service/models"
	"image/color"
)

func (h *Holder) setupUIForHolder(app fyne.App){
	//s.ui = UI.UserInterface{Window: window}

	window := app.NewWindow("holder")
	window.Resize(fyne.NewSize(300,300))

	homeContainer := h.HomeTab()
	shareVPContainer := h.ShareVPContainer()
	tabs := container.NewAppTabs(
		container.NewTabItem("Home", homeContainer),
		container.NewTabItem("Share VPs", shareVPContainer),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	window.SetContent(tabs)
	window.Show()
}

func (h *Holder) HomeTab() *fyne.Container{
	name := h.getName()
	red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	welcomeText := canvas.NewText("Holder Portal of "+name, red)

	vcsList := binding.NewString()
	vcsList.Set("No Verifiable Credentials")
	vcs := widget.NewLabelWithData(vcsList)

	requestVCsButton := widget.NewButton("Request VCs", func() {
		h.RequestVCFromIssuer()
	})

	viewVCsButton := widget.NewButton("View VCs", func() {
		var res string
		var vcs []models.VerifiableCredential

		vcs = h.verfiableCredentials

		for _, vc := range vcs{
			res = res + vc.GetId()+"\n"
		}
		vcsList.Set(res)
	})

	homeContainer := container.New(layout.NewVBoxLayout(),welcomeText,  requestVCsButton, viewVCsButton, vcs)
	return homeContainer
}


func (h *Holder) ShareVPContainer() *fyne.Container{
	shareVPButton := widget.NewButton("share all VPs", func() {
		h.ShareallVPs()
	})

	shareVPContainer := container.New(layout.NewVBoxLayout(),shareVPButton)
	return shareVPContainer
}

