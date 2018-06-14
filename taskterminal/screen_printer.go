package taskterminal

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/NonerKao/color-aware-tabwriter"
	"github.com/fatih/color"
)

type ScreenPrinter struct {
	Writer *tabwriter.Writer
}

func NewScreenPrinter() *ScreenPrinter {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	formatter := &ScreenPrinter{Writer: w}
	return formatter
}

func (f *ScreenPrinter) Print(groupedTasks *GroupedTasks, printNotes bool) {
	cyan := color.New(color.FgCyan).SprintFunc()

	var keys []string
	for key := range groupedTasks.Groups {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Fprintf(f.Writer, "\n %s\n", cyan(key))
		for _, task := range groupedTasks.Groups[key] {
			f.printTask(task)
			if printNotes {
				for nid, note := range task.Notes {
					fmt.Fprintf(f.Writer, "   %s\t%s\t\n",
						cyan(strconv.Itoa(nid)), note)
				}
			}
		}
	}
	f.Writer.Flush()
}

func (f *ScreenPrinter) printTask(task *Task) {
	yellow := color.New(color.FgYellow)
	if task.IsPriority {
		yellow.Add(color.Bold, color.Italic)
	}
	fmt.Fprintf(f.Writer, " %s\t%s\t%s\t%s\t\n",
		yellow.SprintFunc()(strconv.Itoa(task.Id)),
		f.formatCompleted(task.Status),
		f.formatDue(task.Due, task.IsPriority),
		f.formatSubject(task.Subject, task.IsPriority))
}

func (f *ScreenPrinter) formatDue(due string, isPriority bool) string {
	blue := color.New(color.FgBlue)
	red := color.New(color.FgRed)

	if isPriority {
		blue.Add(color.Bold, color.Italic)
		red.Add(color.Bold, color.Italic)
	}

	if due == "" {
		return blue.SprintFunc()(" ")
	}
	dueTime, err := time.Parse("2006-01-02", due)

	if err != nil {
		fmt.Println(err)
		fmt.Println("This may due to the corruption of tasks.json file.")
		os.Exit(-1)
	}

	if isToday(dueTime) {
		return blue.SprintFunc()(dueTime.Format("Mon Jan 2"))
	} else if isTomorrow(dueTime) {
		return blue.SprintFunc()(dueTime.Format("Mon Jan 2"))
	} else if isPastDue(dueTime) {
		return red.SprintFunc()(dueTime.Format("Mon Jan 2"))
	} else {
		return blue.SprintFunc()(dueTime.Format("Mon Jan 2"))
	}
}

func (f *ScreenPrinter) formatSubject(subject string, isPriority bool) string {

	red := color.New(color.FgRed)
	magenta := color.New(color.FgMagenta)
	white := color.New(color.FgWhite)

	if isPriority {
		red.Add(color.Bold, color.Italic)
		magenta.Add(color.Bold, color.Italic)
		white.Add(color.Bold, color.Italic)
	}

	splitted := strings.Split(subject, " ")
	tagRegex, _ := regexp.Compile(`\+[\p{L}\d_]+`)
	contextRegex, _ := regexp.Compile(`\@[\p{L}\d_]+`)

	coloredWords := []string{}

	for _, word := range splitted {
		if tagRegex.MatchString(word) {
			coloredWords = append(coloredWords, magenta.SprintFunc()(word))
		} else if contextRegex.MatchString(word) {
			coloredWords = append(coloredWords, red.SprintFunc()(word))
		} else {
			coloredWords = append(coloredWords, white.SprintFunc()(word))
		}
	}
	return strings.Join(coloredWords, " ")

}

func (f *ScreenPrinter) formatCompleted(status string) string {
	return "[" + status + "]";
}
