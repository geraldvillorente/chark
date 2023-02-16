package main

import (
	"fmt"
	"os"

	"github.com/deanishe/awgo"
	"github.com/jason0x43/go-reminders"
)

func main() {
	wf := aw.New()

	r, err := reminders.NewReminders()
	if err != nil {
		wf.FatalError(err)
	}

	// List all reminders
	listReminders(r)

	// Add a new reminder
	newRem := reminders.Reminder{
		Title: "New Reminder",
	}
	err = addReminder(r, newRem)
	if err != nil {
		wf.FatalError(err)
	}

	// Mark a reminder as complete
	reminders, err := r.List()
	if err != nil {
		wf.FatalError(err)
	}
	if len(reminders) > 0 {
		err = completeReminder(r, reminders[0].ID)
		if err != nil {
			wf.FatalError(err)
		}
	}

	// Delete a reminder
	if len(reminders) > 1 {
		err = deleteReminder(r, reminders[1].ID)
		if err != nil {
			wf.FatalError(err)
		}
	}

	os.Exit(0)
}

// List all reminders
func listReminders(r *reminders.Reminders) {
	reminders, err := r.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, rem := range reminders {
		fmt.Println(rem.Title)
	}
}

// Add a new reminder
func addReminder(r *reminders.Reminders, newRem reminders.Reminder) error {
	err := r.Add(newRem)
	if err != nil {
		return err
	}
	return nil
}

// Mark a reminder as complete
func completeReminder(r *reminders.Reminders, id int) error {
	err := r.Complete(id)
	if err != nil {
		return err
	}
	return nil
}

// Delete a reminder
func deleteReminder(r *reminders.Reminders, id int) error {
	err := r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
