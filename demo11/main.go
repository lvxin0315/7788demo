package main

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
)

func main() {
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "Test",
		BaseDirectoryPath: "example",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// New window
	var w *astilectron.Window

	if w, err = a.NewWindow("index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	//var w1 *astilectron.Window
	//if w1, err = a.NewWindow("index1.html", &astilectron.WindowOptions{
	//	Center: astikit.BoolPtr(true),
	//	Height: astikit.IntPtr(700),
	//	Width:  astikit.IntPtr(700),
	//}); err != nil {
	//	l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	//}
	//
	//// Create windows
	//if err = w1.Create(); err != nil {
	//	l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	//}

	// Blocking pattern

	sendMessage(w)

	menus(a)

	// Listen to login events
	//w.OnLogin(func(i astilectron.Event) (username, password string, err error) {
	//	// Process the request and auth info
	//	if i.Request.Method == "GET" && i.AuthInfo.Scheme == "http://" {
	//		username = "username"
	//		password = "password"
	//	}
	//	return
	//})

	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		// Unmarshal
		var s string
		m.Unmarshal(&s)

		// Process message
		if s == "demo1" {
			//w1.Show()
			return "ok"
		} else {
			//w1.Hide()
			return "ok"
		}
		return nil
	})

	notifications(a)

	dock(a)

	a.Wait()
}

func notifications(a *astilectron.Astilectron) {
	// Create the notification
	var n = a.NewNotification(&astilectron.NotificationOptions{
		Body:             "My Body",
		HasReply:         astikit.BoolPtr(true), // Only MacOSX
		Icon:             "/path/to/icon",
		ReplyPlaceholder: "type your reply here", // Only MacOSX
		Title:            "My title",
	})

	// Add listeners
	n.On(astilectron.EventNameNotificationEventClicked, func(e astilectron.Event) (deleteListener bool) {
		log.Println("the notification has been clicked!")
		return
	})
	// Only for MacOSX
	n.On(astilectron.EventNameNotificationEventReplied, func(e astilectron.Event) (deleteListener bool) {
		log.Printf("the user has replied to the notification: %s\n", e.Reply)
		return
	})

	// Create notification
	n.Create()

	// Show notification
	n.Show()
}

func menus(a *astilectron.Astilectron) {
	// Init a new app menu
	// You can do the same thing with a window
	var m = a.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astikit.StrPtr("Separator"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astikit.StrPtr("Normal 1")},
				{
					Label: astikit.StrPtr("Normal 2"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						log.Println("Normal 2 item has been clicked")
						return
					},
				},
				{Type: astilectron.MenuItemTypeSeparator},
				{Label: astikit.StrPtr("Normal 3")},
			},
		},
		{
			Label: astikit.StrPtr("Checkbox"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Checked: astikit.BoolPtr(true), Label: astikit.StrPtr("Checkbox 1"), Type: astilectron.MenuItemTypeCheckbox},
				{Label: astikit.StrPtr("Checkbox 2"), Type: astilectron.MenuItemTypeCheckbox},
				{Label: astikit.StrPtr("Checkbox 3"), Type: astilectron.MenuItemTypeCheckbox},
			},
		},
		{
			Label: astikit.StrPtr("Radio"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Checked: astikit.BoolPtr(true), Label: astikit.StrPtr("Radio 1"), Type: astilectron.MenuItemTypeRadio},
				{Label: astikit.StrPtr("Radio 2"), Type: astilectron.MenuItemTypeRadio},
				{Label: astikit.StrPtr("Radio 3"), Type: astilectron.MenuItemTypeRadio},
			},
		},
		{
			Label: astikit.StrPtr("Roles"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astikit.StrPtr("Minimize"), Role: astilectron.MenuItemRoleMinimize},
				{Label: astikit.StrPtr("Close"), Role: astilectron.MenuItemRoleClose},
			},
		},
	})

	// Retrieve a menu item
	// This will retrieve the "Checkbox 1" item
	mi, _ := m.Item(1, 0)

	// Add listener manually
	// An OnClick listener has already been added in the options directly for another menu item
	mi.On(astilectron.EventNameMenuItemEventClicked, func(e astilectron.Event) bool {
		log.Printf("Menu item has been clicked. 'Checked' status is now %t\n", *e.MenuItemOptions.Checked)
		return false
	})

	// Create the menu
	m.Create()

	// Manipulate a menu item
	mi.SetChecked(true)

	// Init a new menu item
	var ni = m.NewItem(&astilectron.MenuItemOptions{
		Label: astikit.StrPtr("Inserted"),
		SubMenu: []*astilectron.MenuItemOptions{
			{Label: astikit.StrPtr("Inserted 1")},
			{Label: astikit.StrPtr("Inserted 2")},
		},
	})

	// Insert the menu item at position "1"
	m.Insert(1, ni)

	// Fetch a sub menu
	s, _ := m.SubMenu(0)

	// Init a new menu item
	ni = s.NewItem(&astilectron.MenuItemOptions{
		Label: astikit.StrPtr("Appended"),
		SubMenu: []*astilectron.MenuItemOptions{
			{Label: astikit.StrPtr("Appended 1")},
			{Label: astikit.StrPtr("Appended 2")},
		},
	})

	// Append menu item dynamically
	s.Append(ni)

	// Pop up sub menu as a context menu
	s.Popup(&astilectron.MenuPopupOptions{PositionOptions: astilectron.PositionOptions{X: astikit.IntPtr(50), Y: astikit.IntPtr(50)}})

	// Close popup
	//s.ClosePopup()

	// Destroy the menu
	//m.Destroy()
}

func sendMessage(w *astilectron.Window) {
	w.SendMessage("hello", func(m *astilectron.EventMessage) {
		// Unmarshal
		var s string
		m.Unmarshal(&s)

		// Process message
		log.Printf("received %s\n", s)
	})
}

func dock(a *astilectron.Astilectron) {
	// Get the dock
	var d = a.Dock()

	// Hide and show the dock
	d.Hide()
	d.Show()

	// Make the Dock bounce
	id, _ := d.Bounce(astilectron.DockBounceTypeCritical)

	// Cancel the bounce
	d.CancelBounce(id)

	// Update badge and icon
	d.SetBadge("test")
	//d.SetIcon("/path/to/icon")

	// New dock menu
	var m = d.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astikit.StrPtr("Root 1"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astikit.StrPtr("Item 1")},
				{Label: astikit.StrPtr("Item 2")},
				{Type: astilectron.MenuItemTypeSeparator},
				{Label: astikit.StrPtr("Item 3")},
			},
		},
		{
			Label: astikit.StrPtr("Root 2"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astikit.StrPtr("Item 1")},
				{Label: astikit.StrPtr("Item 2")},
			},
		},
	})

	// Create the menu
	m.Create()
}
