package main

import (
	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	type Task struct {
		Task string
	}

	var myTaskData []Task

	data_from_file, _ := ioutil.ReadFile("taskList.txt")

	json.Unmarshal(data_from_file, &myTaskData)

	a := app.New()
	w := a.NewWindow("Task List")
	// Setting the size of the window
	w.Resize(fyne.NewSize(500, 500))

	l_task := widget.NewLabel("My tasks...")

	// User entry field
	e_task := widget.NewEntry()
	e_task.SetPlaceHolder("Enter your task")

	// Submit function, button adds tasks to list
	submit := widget.NewButton("submit", func() {
		myData := &Task{
			Task: e_task.Text,
		}

		myTaskData = append(myTaskData, *myData)
		final_data, _ := json.MarshalIndent(myTaskData, "", " ")
		ioutil.WriteFile("taskList.txt", final_data, 0644)

		e_task.Text = ""
		e_task.Refresh()
	})

	// Delete function
	delete := widget.NewButton("delete", func() {

		var TempData []Task
		for _, e := range myTaskData {
			if l_task.Text != e.Task {
				TempData = append(TempData, e)
			}
		}
		myTaskData = TempData
		result, _ := json.MarshalIndent(myTaskData, "", " ")
		ioutil.WriteFile("taskList.txt", result, 0644)

		e_task.Text = ""
		e_task.Refresh()
	})

	list := widget.NewList(
		func() int { return len(myTaskData) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(myTaskData[lii].Task)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		l_task.Text = myTaskData[id].Task
		l_task.Refresh()
	}

	w.SetContent(

		container.NewHSplit(
			list,
			container.NewVBox(l_task, e_task, submit, delete),
		),
	)

	w.ShowAndRun()
}
