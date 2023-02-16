package main

import (
	"testing"

	"github.com/jason0x43/go-reminders"
)

func TestAddReminder(t *testing.T) {
	r, err := reminders.NewReminders()
	if err != nil {
		t.Error("Error creating reminders object:", err)
	}

	newRem := reminders.Reminder{
		Title: "New Reminder",
	}

	err = addReminder(r, newRem)
	if err != nil {
		t.Error("Error adding reminder:", err)
	}

	reminders, err := r.List()
	if err != nil {
		t.Error("Error listing reminders:", err)
	}

	if len(reminders) < 1 {
		t.Error("Expected at least one reminder, got none")
	}

	// Clean up by deleting the added reminder
	err = deleteReminder(r, reminders[0].ID)
	if err != nil {
		t.Error("Error deleting reminder:", err)
	}
}

func TestCompleteReminder(t *testing.T) {
	r, err := reminders.NewReminders()
	if err != nil {
		t.Error("Error creating reminders object:", err)
	}

	newRem := reminders.Reminder{
		Title: "New Reminder",
	}

	err = addReminder(r, newRem)
	if err != nil {
		t.Error("Error adding reminder:", err)
	}

	reminders, err := r.List()
	if err != nil {
		t.Error("Error listing reminders:", err)
	}

	if len(reminders) < 1 {
		t.Error("Expected at least one reminder, got none")
	}

	err = completeReminder(r, reminders[0].ID)
	if err != nil {
		t.Error("Error completing reminder:", err)
	}

	// Clean up by deleting the added reminder
	err = deleteReminder(r, reminders[0].ID)
	if err != nil {
		t.Error("Error deleting reminder:", err)
	}
}

func TestDeleteReminder(t *testing.T) {
	r, err := reminders.NewReminders()
	if err != nil {
		t.Error("Error creating reminders object:", err)
	}

	newRem := reminders.Reminder{
		Title: "New Reminder",
	}

	err = addReminder(r, newRem)
	if err != nil {
		t.Error("Error adding reminder:", err)
	}

	reminders, err := r.List()
	if err != nil {
		t.Error("Error listing reminders:", err)
	}

	if len(reminders) < 1 {
		t.Error("Expected at least one reminder, got none")
	}

	err = deleteReminder(r, reminders[0].ID)
	if err != nil {
		t.Error("Error deleting reminder:", err)
	}

	reminders, err = r.List()
	if err != nil {
		t.Error("Error listing reminders:", err)
	}

	if len(reminders) > 0 {
		t.Error("Expected no reminders, got at least one")
	}
}
