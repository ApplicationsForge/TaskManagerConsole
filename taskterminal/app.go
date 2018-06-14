package taskterminal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type App struct {
	TaskStore Store
	Printer   Printer
	TaskTerminal  *TaskTerminal
}

func NewApp() *App {
	app := &App{
		TaskTerminal:  &TaskTerminal{},
		Printer:   NewScreenPrinter(),
		TaskStore: NewFileStore(),
	}
	return app
}

func (a *App) InitializeRepo() {
	a.TaskStore.Initialize()
}

func (a *App) AddTask(input string) {
	a.Load()
	parser := &Parser{}
	task := parser.ParseNewTask(input)
	if task == nil {
		fmt.Println("I need more information. Try something like 'task a chat with @bob due tom'")
		return
	}

	id := a.TaskTerminal.NextId()
	a.TaskTerminal.Add(task)
	a.Save()
	fmt.Printf("Task %d added.\n", id)
}

// AddDoneTask Adds a task and immediately completed it.
func (a *App) AddDoneTask(input string) {
	a.Load()

	r, _ := regexp.Compile(`^(done)(\s*|)`)
	input = r.ReplaceAllString(input, "")
	parser := &Parser{}
	task := parser.ParseNewTask(input)
	if task == nil {
		fmt.Println("I need more information. Try something like 'task done chating with @bob'")
		return
	}

	id := a.TaskTerminal.NextId()
	a.TaskTerminal.Add(task)
	a.TaskTerminal.ChangeTaskStatus("Done", id)
	a.Save()
	fmt.Printf("Completed Task %d added.\n", id)
}

func (a *App) DeleteTask(input string) {
	a.Load()
	ids := a.getIds(input)
	if len(ids) == 0 {
		return
	}
	a.TaskTerminal.Delete(ids...)
	a.Save()
	fmt.Printf("%s deleted.\n", pluralize(len(ids), "Task", "Tasks"))
}

func (a *App) ChangeTaskStatus(input string) {
	a.Load()
	ids := a.getIds(input)
	if len(ids) == 0 {
		return
	}

	statusIndex := a.getStatus(input)
	a.TaskTerminal.ChangeTaskStatus(statusIndex, ids...)
	a.Save()
	fmt.Println("Task status changed.")
}

func (a *App) ArchiveTask(input string) {
	a.Load()
	ids := a.getIds(input)
	if len(ids) == 0 {
		return
	}
	a.TaskTerminal.Archive(ids...)
	a.Save()
	fmt.Println("Task archived.")
}

func (a *App) UnarchiveTask(input string) {
	a.Load()
	ids := a.getIds(input)
	if len(ids) == 0 {
		return
	}
	a.TaskTerminal.Unarchive(ids...)
	a.Save()
	fmt.Println("Task unarchived.")
}

func (a *App) EditTask(input string) {
	a.Load()
	id := a.getId(input)
	if id == -1 {
		return
	}
	task := a.TaskTerminal.FindById(id)
	if task == nil {
		fmt.Println("No such id.")
		return
	}
	parser := &Parser{}

	if parser.ParseEditTask(task, input) {
		a.Save()
		fmt.Println("Task updated.")
	}
}

func (a *App) ExpandTask(input string) {
	a.Load()
	id := a.getId(input)
	parser := &Parser{}
	if id == -1 {
		return
	}

	commonTag := parser.ExpandTag(input)
	tasks := strings.LastIndex(input, ":")
	if commonTag == "" || len(input) <= tasks+1 || tasks == -1 {
		fmt.Println("I'm expecting a format like \"taskterminal ex <tag>: <task1>, <task2>, ...\"")
		return
	}

	newTasks := strings.Split(input[tasks+1:], ",")

	for _, task := range newTasks {
		args := []string{"add ", commonTag, " ", task}
		a.AddTask(strings.Join(args, ""))
	}

	a.TaskTerminal.Delete(id)
	a.Save()
	fmt.Println("Task expanded.")
}

func (a *App) HandleNotes(input string) {
	a.Load()
	id := a.getId(input)
	if id == -1 {
		return
	}
	task := a.TaskTerminal.FindById(id)
	if task == nil {
		fmt.Println("No such id.")
		return
	}
	parser := &Parser{}

	if parser.ParseAddNote(task, input) {
		fmt.Println("Note added.")
	} else if parser.ParseDeleteNote(task, input) {
		fmt.Println("Note deleted.")
	} else if parser.ParseEditNote(task, input) {
		fmt.Println("Note edited.")
	} else if parser.ParseShowNote(task, input) {
		groups := map[string][]*Task{}
		groups[""] = append(groups[""], task)
		a.Printer.Print(&GroupedTasks{Groups: groups}, true)
		return
	}
	a.Save()
}

func (a *App) ArchiveCompleted() {
	a.Load()
	for _, task := range a.TaskTerminal.Tasks() {
		if (task.Status != "Task") {
			task.Archive()
		}
	}
	a.Save()
	fmt.Println("All completed tasks have been archived.")
}

func (a *App) ListTasks(input string) {
	a.Load()
	filtered := NewFilter(a.TaskTerminal.Tasks()).Filter(input)
	grouped := a.getGroups(input, filtered)

	re, _ := regexp.Compile(`^ln`)
	a.Printer.Print(grouped, re.MatchString(input))
}

func (a *App) PrioritizeTask(input string) {
	a.Load()
	ids := a.getIds(input)
	if len(ids) == 0 {
		return
	}
	a.TaskTerminal.Prioritize(ids...)
	a.Save()
	fmt.Println("Task prioritized.")
}

func (a *App) UnprioritizeTask(input string) {
	a.Load()
	ids := a.getIds(input)
	if len(ids) == 0 {
		return
	}
	a.TaskTerminal.Unprioritize(ids...)
	a.Save()
	fmt.Println("Task un-prioritized.")
}

func (a *App) getId(input string) int {
	re, _ := regexp.Compile("\\d+")
	if re.MatchString(input) {
		id, _ := strconv.Atoi(re.FindString(input))
		return id
	}

	fmt.Println("Invalid id.")
	return -1
}

func (a *App) getIds(input string) (ids []int) {

	idGroups := strings.Split(input, ",")
	for _, idGroup := range idGroups {
		if rangedIds, err := a.parseRangedIds(idGroup); len(rangedIds) > 0 || err != nil {
			if err != nil {
				fmt.Printf("Invalid id group: %s.\n", input)
				continue
			}
			ids = append(ids, rangedIds...)
		} else if id := a.getId(idGroup); id != -1 {
			ids = append(ids, id)
		} else {
			fmt.Printf("Invalid id: %s.\n", idGroup)
		}
	}
	return ids
}

func (a *App) getStatus(input string) string {
	status := "Undefined"
	parsed := strings.Split(input, " ")
	if len(parsed) > 2 {
		status = parsed[2]
	}
	return status
}

func (a *App) parseRangedIds(input string) (ids []int, err error) {
	rangeNumberRE, _ := regexp.Compile("(\\d+)-(\\d+)")
	if matches := rangeNumberRE.FindStringSubmatch(input); len(matches) > 0 {
		lowerID, _ := strconv.Atoi(matches[1])
		upperID, _ := strconv.Atoi(matches[2])
		if lowerID >= upperID {
			return ids, fmt.Errorf("Invalid id group: %s.\n", input)
		}
		for id := lowerID; id <= upperID; id++ {
			ids = append(ids, id)
		}
	}
	return ids, err
}

func (a *App) getGroups(input string, tasks []*Task) *GroupedTasks {
	grouper := &Grouper{}
	userRegex, _ := regexp.Compile(`by u.*$`)
	tagRegex, _ := regexp.Compile(`by t.*$`)

	var grouped *GroupedTasks

	if userRegex.MatchString(input) {
		grouped = grouper.GroupByUser(tasks)
	} else if tagRegex.MatchString(input) {
		grouped = grouper.GroupByTag(tasks)
	} else {
		grouped = grouper.GroupByNothing(tasks)
	}
	return grouped
}

func (a *App) GarbageCollect() {
	a.Load()
	a.TaskTerminal.GarbageCollect()
	a.Save()
	fmt.Println("Garbage collection complete.")
}

func (a *App) Load() error {
	tasks, err := a.TaskStore.Load()
	if err != nil {
		return err
	}
	a.TaskTerminal.Load(tasks)
	return nil
}

func (a *App) Save() {
	a.TaskStore.Save(a.TaskTerminal.Data)
}
