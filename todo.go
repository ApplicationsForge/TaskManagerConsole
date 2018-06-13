package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/ApplicationsForge/TaskTerminal/src"
	"github.com/skratchdot/open-golang/open"
)

const (
	VERSION = "0.8"
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

	fmt.Printf("TaskTerminal v%s, a simple, command line based, GTD-style task manager\n", VERSION)

	blueBold.Println("\nAdding tasks")
	fmt.Println("  the 'a' command adds tasks.")
	fmt.Println("  You can also optionally specify a due date.")
	fmt.Println("  Specify a due date by putting 'due <date>' at the end, where <date> is in (tod|today|tom|tomorrow|mon|tue|wed|thu|fri|sat|sun)")
	fmt.Println("\n  Examples for adding a task:")
	yellow.Println("\t./TaskTerminal a Meeting with @bob about +importantTag due today")
	yellow.Println("\t./TaskTerminal a +work +verify did @john fix the build\\?")

	blueBold.Println("\nListing tasks")
	fmt.Println("  When listing tasks, you can filter and group the output.\n")

	fmt.Println("  ./TaskTerminal l due (tod|today|tom|tomorrow|overdue|this week|next week|last week|mon|tue|wed|thu|fri|sat|sun|none)")
	fmt.Println("  ./TaskTerminal l overdue\n")

	cyan.Println("  Filtering by date:\n")
	yellow.Println("\t./TaskTerminal l due tod")
	fmt.Println("\tlists all tasks due today\n")
	yellow.Println("\t./TaskTerminal l due tom")
	fmt.Println("\tlists all tasks due tomorrow\n")
	yellow.Println("\t./TaskTerminal l due mon")
	fmt.Println("\tlists all tasks due monday\n")
	yellow.Println("\t./TaskTerminal l overdue")
	fmt.Println("\tlists all tasks where the due date is in the past\n")
	yellow.Println("\t./TaskTerminal agenda")
	fmt.Println("\tlists all tasks where the due date is today or in the past\n")

	fmt.Println("  ./TaskTerminal l completed (tod|today|this week)")
	cyan.Println("  Filtering by date:\n")

	yellow.Println("\t./TaskTerminal l completed (tod|today)")
	fmt.Println("\tlists all tasks that were completed today\n")
	yellow.Println("\t./TaskTerminal l completed this week")
	fmt.Println("\tlists all tasks that were completed this week\n")

	cyan.Println("  Grouping:")
	fmt.Println("  You can group tasks by context or tag.")
	yellow.Println("\t./TaskTerminal l by c")
	fmt.Println("\tlists all tasks grouped by context\n")
	yellow.Println("\t./TaskTerminal l by t")
	fmt.Println("\tlists all tasks grouped by tag\n")

	cyan.Println("  Grouping and filtering:")
	fmt.Println("  Of course, you can combine grouping and filtering to get a nice formatted list.\n")
	yellow.Println("\t./TaskTerminal l due today by c")
	fmt.Println("\tlists all tasks due today grouped by context\n")
	yellow.Println("\t./TaskTerminal l +tag due this week by c")
	fmt.Println("\tlists all tasks due today for +tag, grouped by context\n")
	yellow.Println("\t./TaskTerminal l @frank due tom by t")
	fmt.Println("\tlists all tasks due tomorrow concerining @frank for +tag, grouped by tag\n")

	blueBold.Println("\nCompleting and uncompleting ")
	fmt.Println("Complete and Uncomplete a task by its Id:\n")
	yellow.Println("\t./TaskTerminal c 33")
	fmt.Println("\tCompletes a task with id 33\n")
	yellow.Println("\t./TaskTerminal uc 33")
	fmt.Println("\tUncompletes a task with id 33\n")

	blueBold.Println("\nPrioritizing")
	fmt.Println("tasks have a priority flag, which will make them bold when listed.\n")
	yellow.Println("\t./TaskTerminal p 33")
	fmt.Println("\tPrioritizes a task with id 33\n")
	yellow.Println("\t./TaskTerminal up 33")
	fmt.Println("\tUn-prioritizes a task with id 33\n")
	yellow.Println("\t./TaskTerminal l p")
	fmt.Println("\tlist all priority tasks\n")

	blueBold.Println("\nArchiving")
	fmt.Println("You can archive tasks once they are done, or if you might come back to them.")
	fmt.Println("By default, task will only show unarchived tasks.\n")
	yellow.Println("\t./TaskTerminal ar 33")
	fmt.Println("\tArchives a task with id 33\n")
	yellow.Println("\t./TaskTerminal ac")
	fmt.Println("\tArchives all completed tasks\n")
	yellow.Println("\t./TaskTerminal l archived")
	fmt.Println("\tlist all archived tasks\n")

	blueBold.Println("\nEditing due dates")
	yellow.Println("\t./TaskTerminal e 33 due mon")
	fmt.Println("\tEdits the task with 33 and sets the due date to this coming Monday\n")
	yellow.Println("\t./TaskTerminal e 33 due none")
	fmt.Println("\tEdits the task with 33 and removes the due date\n")

	blueBold.Println("\nExpanding existing tasks")
	yellow.Println("\t./TaskTerminal ex 39 +final: read physics due mon, do literature report due fri")
	fmt.Println("\tRemoves the task with id 39, and adds following two tasks\n")

	blueBold.Println("\nDeleting")
	yellow.Println("\t./TaskTerminal d 33")
	fmt.Println("\tDeletes a task with id 33\n")

	blueBold.Println("\nManipulating notes")
	yellow.Println("\t./TaskTerminal ln")
	fmt.Println("\tlists all tasks with their notes")
	yellow.Println("\t./TaskTerminal an 12 check http://this.web.site")
	fmt.Println("\tAdds notes \"check http://this.web.site\" to the task with id 12\n")
	yellow.Println("\t./TaskTerminal n 12")
	fmt.Println("\tLists notes of the task with id 12\n")
	yellow.Println("\t./TaskTerminal dn 12 3")
	fmt.Println("\tDeletes the 3rd note of the task with id 12\n")
	yellow.Println("\t./TaskTerminal en 12 3 check http://that.web.site")
	fmt.Println("\tEditing the 3rd note of the task with id 12 to \"http://that.web.site\" \n")

	blueBold.Println("\nGarbage Collection")
	yellow.Println("\t./TaskTerminal gc")
	fmt.Println("\tDeletes all archived tasks.\n")

	fmt.Println("\tTaskTerminal is developed by ApplicationsForge (https://github.com/ApplicationsForge).")
	fmt.Println("\tTaskTerminal is based on todolist (https://github.com/gammons/todolist) by Grant Ammons (https://twitter.com/gammons).")
	//fmt.Println("For full documentation, please visit http://todolist.site")
}

func routeInput(command string, input string) {
	app := todolist.NewApp()
	switch command {
	case "l", "ln", "list", "agenda":
		app.ListTodos(input)
	case "a", "add":
		app.AddTodo(input)
	case "done":
		app.AddDoneTodo(input)
	case "d", "delete":
		app.DeleteTodo(input)
	case "cs", "ch_status", "change_status":
		app.ChangeTodoStatus(input)
	case "ar", "archive":
		app.ArchiveTodo(input)
	case "uar", "unarchive":
		app.UnarchiveTodo(input)
	case "ac":
		app.ArchiveCompleted()
	case "e", "edit":
		app.EditTodo(input)
	case "ex", "expand":
		app.ExpandTodo(input)
	case "an", "n", "dn", "en":
		app.HandleNotes(input)
	case "gc":
		app.GarbageCollect()
	case "p", "prioritize":
		app.PrioritizeTodo(input)
	case "up", "unprioritize":
		app.UnprioritizeTodo(input)
	case "init":
		app.InitializeRepo()
	case "web":
		if err := app.Load(); err != nil {
			os.Exit(1)
		} else {
			web := todolist.NewWebapp()
			fmt.Println("Now serving todolist web.\nHead to http://localhost:7890 to see your todo list!")
			open.Start("http://localhost:7890")
			web.Run()
		}
	}
}
