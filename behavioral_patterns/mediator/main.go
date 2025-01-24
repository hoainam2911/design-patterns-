package main

import "fmt"

type Mediator interface {
	Notify(sender Component, event string)
}

type Component interface {
	SetMediator(mediator Mediator)
}

type BaseComponent struct {
	mediator Mediator
}

func (c *BaseComponent) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

type Button struct {
	BaseComponent
}

func (b *Button) Click() {
	fmt.Println("Button clicked.")
	if b.mediator != nil {
		b.mediator.Notify(b, "click")
	}
}

type Checkbox struct {
	BaseComponent
	Checked bool
}

func (cb *Checkbox) Check() {
	cb.Checked = !cb.Checked
	fmt.Println("Checkbox state changed to:", cb.Checked)
	if cb.mediator != nil {
		cb.mediator.Notify(cb, "check")
	}
}

type Textbox struct {
	BaseComponent
	Text string
}

func (tb *Textbox) Input(text string) {
	tb.Text = text
	fmt.Println("Textbox input:", tb.Text)
	if tb.mediator != nil {
		tb.mediator.Notify(tb, "input")
	}
}

type Dialog struct {
	loginCheckbox  *Checkbox
	usernameTextbox *Textbox
	submitButton   *Button
}

func (d *Dialog) Notify(sender Component, event string) {
	switch sender.(type) {
	case *Checkbox:
		if event == "check" {
			if d.loginCheckbox.Checked {
				fmt.Println("Checkbox is checked. Showing login fields.")
			} else {
				fmt.Println("Checkbox is unchecked. Showing registration fields.")
			}
		}
	case *Textbox:
		if event == "input" {
			fmt.Println("Text entered:", d.usernameTextbox.Text)
		}
	case *Button:
		if event == "click" {
			fmt.Println("Submit button clicked.")
			if d.loginCheckbox.Checked {
				fmt.Println("Perform login with username:", d.usernameTextbox.Text)
			} else {
				fmt.Println("Perform registration with username:", d.usernameTextbox.Text)
			}
		}
	}
}

func main() {
	loginCheckbox := &Checkbox{}
	usernameTextbox := &Textbox{}
	submitButton := &Button{}
	dialog := &Dialog{
		loginCheckbox:  loginCheckbox,
		usernameTextbox: usernameTextbox,
		submitButton:   submitButton,
	}

	loginCheckbox.SetMediator(dialog)
	usernameTextbox.SetMediator(dialog)
	submitButton.SetMediator(dialog)

	loginCheckbox.Check()
	usernameTextbox.Input("JohnDoe")
	submitButton.Click()
}
