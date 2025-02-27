package main

import (
	"fmt"

	"github.com/d3417/tochka-source/modules/marketplace"
	_ "github.com/d3417/tochka-source/modules/util"
)

func manageRole(username, action, role string) {
	user, _ := marketplace.FindUserByUsername(username)
	if user == nil {
		fmt.Println("No such user")
		return
	}

	if action == "grant" && role == "admin" {
		user.IsAdmin = !user.IsAdmin
	} else {
		fmt.Println("Wrong action")
		return
	}
	user.Save()
}

func indexItems() {
	println("[Index] Indexing items...")
	for _, item := range marketplace.GetItemsForIndexing() {
		println("[Index] ", item.Name)
		err := item.Index()
		if err != nil {
			println("Error: ", err)
		}
	}
}

func syncModels() {
	marketplace.SyncModels()
}

func syncDatabaseViews() {
	marketplace.SyncDatabaseViews()
}

func staffStats() {
	interval := "2019-05-06 00:16"
	sTickets, err := marketplace.StaffSupportTicketsResolutionStats(interval)
	if err != nil {
		return
	}
	sDisputes, err := marketplace.StaffSupportDisputesResolutionStats(interval)
	if err != nil {
		return
	}
	sItems, err := marketplace.StaffItemApprovalStats(interval)
	if err != nil {
		return
	}

	var (
		text = fmt.Sprintf(
			`
Support Agent | Ticket Status | Number Of Tickets
--- | --- | ---
`)
	)
	for _, si := range sTickets {
		text += fmt.Sprintf("%s | TICKET %s | %d\n", si.ResolverUsername, si.CurrentStatus, si.TicketCount)
	}
	for _, si := range sDisputes {
		text += fmt.Sprintf("%s | DISPUTE %s | %d\n", si.ResolverUsername, si.CurrentStatus, si.TicketCount)
	}
	for _, si := range sItems {
		text += fmt.Sprintf("%s | ITEM %s | %d\n", si.ResolverUsername, si.CurrentStatus, si.TicketCount)
	}

	println(text)
}

func importMetroStations() {
	marketplace.ImportCityMetroStations(524901, "./dumps/moscow-metro.json")
	marketplace.ImportCityMetroStations(498817, "./dumps/spb-metro.json")
}

func updateDeposits() {
	marketplace.CommandUpdateDeposits()
}
