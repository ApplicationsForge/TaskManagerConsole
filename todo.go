package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/ApplicationsForge/TaskTerminal/taskterminal"
)

func main() {
	if len(os.Args) <= 1 {
		usage()
		os.Exit(0)
	}
	input := strings.Join(os.Args[1:], " ")
	routeInput(os.Args[1], input)
}

func usage() {
	blue := color.New(color.FgBlue)
	cyan := color.New(color.FgCyan)
	yellow := color.New(color.FgYellow)

	blueBold := blue.Add(color.Bold)

	fmt.Printf("TaskTerminal is a simple, command line based, GTD-style task manager.\n")

	blueBold.Println("\nInitializing Repo")
	yellow.Println("\tTaskTerminal init")
	fmt.Println("\tCreates json file with tasks.")

	blueBold.Println("\nAdding tasks")
	yellow.Println("\tTaskTerminal a Task")
	yellow.Println("\tTaskTerminal add Task")
	fmt.Println("\tAdds a task\n")
	yellow.Println("\tTaskTerminal done Task")
	fmt.Println("\tAdds completed tasks\n")
	fmt.Println("\tYou can also optionally specify a due date.")
	fmt.Println("\tSpecify a due date by putting 'due <date>' at the end, where <date> is in (tod|today|tom|tomorrow|mon|tue|wed|thu|fri|sat|sun)")
	fmt.Println("\n  Examples for adding a task:")
	yellow.Println("\tTaskTerminal a Meeting '#Title1' with @bob about +importantTag due today")
	yellow.Println("\tTaskTerminal a '#Task1' +work +verify did @john fix the build\\?")

	blueBold.Println("\nListing tasks")
	yellow.Println("\tTaskTerminal l")
	yellow.Println("\tTaskTerminal list")
	fmt.Println("\tWhen listing tasks, you can filter and group the output.\n")

	fmt.Println("\tTaskTerminal l due (tod|today|tom|tomorrow|overdue|this week|next week|last week|mon|tue|wed|thu|fri|sat|sun|none)")
	fmt.Println("tTaskTerminal l overdue\n")

	cyan.Println("\tFiltering by date:")
	yellow.Println("\tTaskTerminal l due tod")
	fmt.Println("\tLists all tasks due today\n")
	yellow.Println("\tTaskTerminal l due tom")
	fmt.Println("\tLists all tasks due tomorrow\n")
	yellow.Println("\tTaskTerminal l due mon")
	fmt.Println("\tLists all tasks due monday\n")
	yellow.Println("\tTaskTerminal l overdue")
	fmt.Println("\tLists all tasks where the due date is in the past\n")
	yellow.Println("\tTaskTerminal agenda")
	fmt.Println("\tLists all tasks where the due date is today or in the past\n")

	fmt.Println("\tTaskTerminal l completed (tod|today|this week)")

	cyan.Println("\nFiltering by date:")
	yellow.Println("\tTaskTerminal l completed (tod|today)")
	fmt.Println("\tLists all tasks that were completed today\n")
	yellow.Println("\tTaskTerminal l completed this week")
	fmt.Println("\tLists all tasks that were completed this week\n")

	cyan.Println("\tGrouping:")
	fmt.Println("\tYou can group tasks by user or tag.\n")
	yellow.Println("\tTaskTerminal l by u")
	fmt.Println("\tLists all tasks grouped by user\n")
	yellow.Println("\tTaskTerminal l by t")
	fmt.Println("\tLists all tasks grouped by tag\n")

	cyan.Println("\tGrouping and filtering")
	fmt.Println("\tOf course, you can combine grouping and filtering to get a nice formatted list.\n")
	yellow.Println("\tTaskTerminal l due today by u")
	fmt.Println("\tLists all tasks due today grouped by user\n")
	yellow.Println("\tTaskTerminal l +tag due this week by u")
	fmt.Println("\tLists all tasks due today for +tag, grouped by user\n")
	yellow.Println("\tTaskTerminal l @frank due tom by t")
	fmt.Println("\tLists all tasks due tomorrow concerining @frank for +tag, grouped by tag\n")

	blueBold.Println("\nCompleting and uncompleting")
	yellow.Println("\tTaskTerminal c 33")
	fmt.Println("\tCompletes the task with id 33\n")
	yellow.Println("\tTaskTerminal uc 33")
	fmt.Println("\tUncompletes the task with id 33\n")

	blueBold.Println("\nPrioritizing")
	fmt.Println("Tasks have a priority flag, which will make them bold when listed.\n")
	yellow.Println("\tTaskTerminal p 33")
	yellow.Println("\tTaskTerminal prioritize 33")
	fmt.Println("\tPrioritizes the task with id 33\n")
	yellow.Println("\tTaskTerminal up 33")
	yellow.Println("\tTaskTerminal unprioritize 33")
	fmt.Println("\tUn-prioritizes the task with id 33\n")
	yellow.Println("\tTaskTerminal l p")
	fmt.Println("\tList all priority tasks\n")

	blueBold.Println("\nArchiving")
	fmt.Println("You can archive tasks once they are done, or if you might come back to them.")
	fmt.Println("By default, the list will only show unarchived tasks.\n")
	yellow.Println("\tTaskTerminal ar 33")
	yellow.Println("\tTaskTerminal archive 33")
	fmt.Println("\tArchives the task with id 33\n")
	yellow.Println("\tTaskTerminal uar 33")
	yellow.Println("\tTaskTerminal unarchive 33")
	fmt.Println("\tUnarchives the task with id 33\n")
	yellow.Println("\tTaskTerminal ac")
	fmt.Println("\tArchives all completed tasks\n")
	yellow.Println("\tTaskTerminal as Testing")
	yellow.Println("\tTaskTerminal ar_status Testing")
	yellow.Println("\tTaskTerminal archive_by_status Testing")
	fmt.Println("\tArchives all tasks with Testing status\n")
	yellow.Println("\tTaskTerminal l archived")
	fmt.Println("\tLists all archived tasks\n")

	blueBold.Println("\nEditing")
	yellow.Println("\tTaskTerminal e 33 due mon")
	fmt.Println("\tEdits the task with id 33 and sets the due date to this coming Monday\n")
	yellow.Println("\tTaskTerminal edit 33 due none")
	fmt.Println("\tEdits the task with 33 and removes the due date\n")

	blueBold.Println("\nChanging status")
	yellow.Println("\tTaskTerminal 33 cs Testing")
	yellow.Println("\tTaskTerminal 33 ch_status Testing")
	yellow.Println("\tTaskTerminal 33 change_status Testing")
	fmt.Println("\tСhanges the status of the task with id 33 for Testing")

	blueBold.Println("\nExpanding existing tasks")
	yellow.Println("\tTaskTerminal ex 39 +final: read physics due mon, do literature report due fri")
	fmt.Println("\tRemoves the task with id 39 and adds following two tasks\n")

	blueBold.Println("\nManipulating notes")
	yellow.Println("\tTaskTerminal ln")
	fmt.Println("\tLists all tasks with their notes")
	yellow.Println("\tTaskTerminal n 12")
	fmt.Println("\tLists notes of the task with id 12\n")
	yellow.Println("\tTaskTerminal dn 12 3")
	fmt.Println("\tDeletes the 3rd note of the task with id 12\n")

	blueBold.Println("\nDeleting")
	yellow.Println("\tTaskTerminal d 33")
	yellow.Println("\tTaskTerminal delete 33")
	fmt.Println("\tDeletes the task with id 33\n")

	blueBold.Println("\nGarbage Collection")
	yellow.Println("\tTaskTerminal gc")
	fmt.Println("\tDeletes all archived tasks\n")

	fmt.Println("TaskTerminal is developed by ApplicationsForge (https://github.com/ApplicationsForge).")
	fmt.Println("TaskTerminal is based on todolist (https://github.com/gammons/todolist) by Grant Ammons (https://twitter.com/gammons).")
}

func routeInput(command string, input string) {
	app := taskterminal.NewApp()
	switch command {
	case "l", "ln", "list", "agenda":
		app.ListTasks(input)
	case "a", "add":
		app.AddTask(input)
	case "done":
		app.AddDoneTask(input)
	case "d", "delete":
		app.DeleteTask(input)
	case "cs", "ch_status", "change_status":
		app.ChangeTaskStatus(input)
	case "ar", "archive":
		app.ArchiveTask(input)
	case "as", "ar_status", "archive_by_status":
		app.ArchiveByStatus(input)
	case "uar", "unarchive":
		app.UnarchiveTask(input)
	case "ac":
		app.ArchiveCompleted()
	case "e", "edit":
		app.EditTask(input)
	case "ex", "expand":
		app.ExpandTask(input)
	case "an", "n", "dn", "en":
		app.HandleNotes(input)
	case "gc":
		app.GarbageCollect()
	case "p", "prioritize":
		app.PrioritizeTask(input)
	case "up", "unprioritize":
		app.UnprioritizeTask(input)
	case "init":
		app.InitializeRepo()
	}
}
