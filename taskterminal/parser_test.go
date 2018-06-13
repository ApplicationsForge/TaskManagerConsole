package taskterminal

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseSubject(t *testing.T) {
	parser := &Parser{}
	task := parser.ParseNewTask("do this thing")
	if task.Subject != "do this thing" {
		t.Error("Expected task.Subject to equal 'do this thing'")
	}
}

func TestParseSubjectWithDue(t *testing.T) {
	parser := &Parser{}
	task := parser.ParseNewTask("do this thing due tomorrow")
	if task.Subject != "do this thing" {
		t.Error("Expected task.Subject to equal 'do this thing', got ", task.Subject)
	}
}

func TestParseExpandTags(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	correctFormat := parser.ExpandTag("ex 113 +meeting: figures, slides, coffee, suger")
	assert.Equal("+meeting", correctFormat)
	wrongFormat1 := parser.ExpandTag("ex 114 +meeting figures, slides, coffee, suger")
	assert.Equal("", wrongFormat1)
	wrongFormat2 := parser.ExpandTag("ex 115 meeting: figures, slides, coffee, suger")
	assert.Equal("", wrongFormat2)
	wrongFormat3 := parser.ExpandTag("ex 116 meeting figures, slides, coffee, suger")
	assert.Equal("", wrongFormat3)
	wrongFormat4 := parser.ExpandTag("ex 117 +重要な會議: 図, コーヒー, 砂糖")
	assert.Equal("+重要な會議", wrongFormat4)
}

func TestParseTags(t *testing.T) {
	parser := &Parser{}
	task := parser.ParseNewTask("do this thing +tag1 +tag2 +專案3 +tag-name due tomorrow")
	if len(task.Tags) != 4 {
		t.Error("Expected Tags length to be 3")
	}
	if task.Tags[0] != "tag1" {
		t.Error("task.Tags[0] should equal 'tag1' but got", task.Tags[0])
	}
	if task.Tags[1] != "tag2" {
		t.Error("task.Tags[1] should equal 'tag2' but got", task.Tags[1])
	}
	if task.Tags[2] != "專案3" {
		t.Error("task.Tags[2] should equal '專案3' but got", task.Tags[2])
	}
	if task.Tags[3] != "tag-name" {
		t.Error("task.Tags[3] should equal 'tag-name' but got", task.Tags[3])
	}
}

func TestParseContexts(t *testing.T) {
	parser := &Parser{}
	task := parser.ParseNewTask("do this thing with @bob and @mary due tomorrow")
	if len(task.Contexts) != 2 {
		t.Error("Expected Tags length to be 2")
	}
	if task.Contexts[0] != "bob" {
		t.Error("task.Contexts[0] should equal 'mary' but got", task.Contexts[0])
	}
	if task.Contexts[1] != "mary" {
		t.Error("task.Contexts[1] should equal 'mary' but got", task.Contexts[1])
	}
}

func TestParseAddNote(t *testing.T) {
	parser := &Parser{}
	task := parser.ParseNewTask("add write the test functions")

	b1 := parser.ParseAddNote(task, "an 1 TestPasrseAddNote")
	b2 := parser.ParseAddNote(task, "an 1 TestPasrseDeleteNote")
	b3 := parser.ParseAddNote(task, "an 1 TestPasrseEditNote")

	if !b1 || !b2 || !b3 {
		t.Error("Fail adding notes, expected 3 notes but", len(task.Notes))
	}
}

func TestParseDeleteNote(t *testing.T) {
	parser := &Parser{}
	task := parser.ParseNewTask("add buy notebook")

	task.Notes = append(task.Notes, "ASUStek")
	task.Notes = append(task.Notes, "Apple")
	task.Notes = append(task.Notes, "Dell")
	task.Notes = append(task.Notes, "Acer")

	b1 := parser.ParseDeleteNote(task, "dn 1 1")
	b2 := parser.ParseDeleteNote(task, "dn 1 1")

	if !b1 || !b2 {
		t.Error("Fail deleting notes, expected 2 notes left but", len(task.Notes))
	}

	if task.Notes[0] != "ASUStek" || task.Notes[1] != "Acer" {
		t.Error("Fail deleting notes,", task.Notes[0], "and", task.Notes[1], "are left")
	}
}

func TestParseEditNote(t *testing.T) {
	parser := &Parser{}
	task := parser.ParseNewTask("add record the weather")

	task.Notes = append(task.Notes, "Aug 29 Wed")
	task.Notes = append(task.Notes, "Cloudy")
	task.Notes = append(task.Notes, "40°C")
	task.Notes = append(task.Notes, "Tokyo")

	parser.ParseEditNote(task, "en 1 0 Aug 29 Tue")
	if task.Notes[0] != "Aug 29 Tue" {
		t.Error("Fail editing notes, note 0 should be \"Aug 29 Tue\" but got", task.Notes[0])
	}

	parser.ParseEditNote(task, "en 1 1 Sunny")
	if task.Notes[1] != "Sunny" {
		t.Error("Fail editing notes, note 1 should be \"Sunny\" but got", task.Notes[1])
	}

	parser.ParseEditNote(task, "en 1 2 22°C")
	if task.Notes[2] != "22°C" {
		t.Error("Fail editing notes, note 2 should be \"22°C\" but got", task.Notes[2])
	}

	parser.ParseEditNote(task, "en 1 3 Seoul")
	if task.Notes[3] != "Seoul" {
		t.Error("Fail editing notes, note 3 should be \"Seoul\" but got", task.Notes[3])
	}
}

func TestHandleNotes(t *testing.T) {
	parser := &Parser{}
	task := parser.ParseNewTask("add search engine survey")

	if !parser.ParseAddNote(task, "an 1 www.google.com") {
		t.Error("Expected Notes to be added")
	}
	if task.Notes[0] != "www.google.com" {
		t.Error("Expected note 1 to be 'www.google.com' but got", task.Notes[0])
	}

	if !parser.ParseEditNote(task, "en 1 0 www.duckduckgo.com") {
		t.Error("Expected Notes to be editted")
	}
	if task.Notes[0] != "www.duckduckgo.com" {
		t.Error("Expected note 1 to be 'www.duckduckgo.com' but got", task.Notes[0])
	}

	if !parser.ParseDeleteNote(task, "dn 1 0") {
		t.Error("Expected Notes to be deleted")
	}
	if len(task.Notes) != 0 {
		t.Error("Expected no note")
	}
}

func TestDueToday(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	expectedDate := bod(time.Now()).Format("2006-01-02")

	task := parser.ParseNewTask("do this thing with @bob and @mary due today")
	assert.Equal(expectedDate, task.Due)

	task = parser.ParseNewTask("do this thing with @bob and @mary due tod")
	assert.Equal(expectedDate, task.Due)
}

func TestDueTomorrow(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	expectedDate := bod(time.Now()).AddDate(0, 0, 1).Format("2006-01-02")

	task := parser.ParseNewTask("do this thing with @bob and @mary due tomorrow")
	assert.Equal(expectedDate, task.Due)

	task = parser.ParseNewTask("do this thing with @bob and @mary due tom")
	assert.Equal(expectedDate, task.Due)
}

func TestDueSpecific(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	task := parser.ParseNewTask("do this thing with @bob and @mary due jun 1")
	year := strconv.Itoa(time.Now().Year())
	assert.Equal(fmt.Sprintf("%s-06-01", year), task.Due)
}

func TestDueSpecificEuropeanDate(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	task := parser.ParseNewTask("do this thing with @bob and @mary due 1 jun")
	year := strconv.Itoa(time.Now().Year())
	assert.Equal(fmt.Sprintf("%s-06-01", year), task.Due)
}

func TestMondayOnSunday(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-24")
	assert.Equal("2016-04-25", parser.monday(now))
}

func TestMondayOnMonday(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-25")
	assert.Equal("2016-04-25", parser.monday(now))
}

func TestMondayOnTuesday(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-26")
	assert.Equal("2016-05-02", parser.monday(now))
}

func TestTuesdayOnMonday(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-25")
	assert.Equal("2016-04-26", parser.tuesday(now))
}

func TestTuesdayOnWednesday(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	now, _ := time.Parse("2006-01-02", "2016-04-27")
	assert.Equal("2016-05-03", parser.tuesday(now))
}

func TestDueOnSpecificDate(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	year := strconv.Itoa(time.Now().Year())
	assert.Equal(fmt.Sprintf("%s-05-02", year), parser.Due("due may 2", time.Now()))
	assert.Equal(fmt.Sprintf("%s-06-01", year), parser.Due("due jun 1", time.Now()))
}

func TestDueOnSpecificDateEuropeFormat(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	year := strconv.Itoa(time.Now().Year())
	assert.Equal(fmt.Sprintf("%s-05-02", year), parser.Due("due 2 may", time.Now()))
	assert.Equal(fmt.Sprintf("%s-06-01", year), parser.Due("due 1 jun", time.Now()))
}

func TestDueOnSpecificDateEuropean(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	year := strconv.Itoa(time.Now().Year())
	assert.Equal(fmt.Sprintf("%s-05-02", year), parser.Due("due 2 may", time.Now()))
}

func TestDueIntelligentlyChoosesCorrectYear(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	marchTime, _ := time.Parse("2006-01-02", "2016-03-25")
	januaryTime, _ := time.Parse("2006-01-02", "2016-01-05")
	septemberTime, _ := time.Parse("2006-01-02", "2016-09-25")
	decemberTime, _ := time.Parse("2006-01-02", "2016-12-25")

	assert.Equal("2016-01-10", parser.parseArbitraryDate("jan 10", januaryTime))
	assert.Equal("2016-01-10", parser.parseArbitraryDate("jan 10", marchTime))
	assert.Equal("2017-01-10", parser.parseArbitraryDate("jan 10", septemberTime))
	assert.Equal("2017-01-10", parser.parseArbitraryDate("jan 10", decemberTime))
}

func TestParseEditTaskJustDate(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	task := NewTask()
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	parser.ParseEditTask(task, "e 24 due tom")

	assert.Equal(task.Due, tomorrow)
}

func TestParseEditTaskJustDateDoesNotEditExistingSubject(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	task := NewTask()
	task.Subject = "pick up the trash"
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	parser.ParseEditTask(task, "e 24 due tom")

	assert.Equal(task.Due, tomorrow)
	assert.Equal(task.Subject, "pick up the trash")
}

func TestParseEditTaskJustSubject(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	task := &Task{Subject: "pick up the trash", Due: "2016-11-25"}

	parser.ParseEditTask(task, "e 24 changed the task")

	assert.Equal(task.Due, "2016-11-25")
	assert.Equal(task.Subject, "changed the task")
}

func TestParseEditTaskSubjectUpdatesTagsAndContexts(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	task := &Task{
		Subject:  "pick up the +trash with @dad",
		Due:      "2016-11-25",
		Tags: []string{"trash"},
		Contexts: []string{"dad"},
	}

	parser.ParseEditTask(task, "e 24 get the +garbage with @mom")

	assert.Equal(task.Due, "2016-11-25")
	assert.Equal(task.Subject, "get the +garbage with @mom")
	assert.Equal(task.Tags, []string{"garbage"})
	assert.Equal(task.Contexts, []string{"mom"})
}

func TestParseEditTaskWithSubjectAndDue(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	task := &Task{
		Subject:  "pick up the +trash with @dad",
		Due:      "2016-11-25",
		Tags: []string{"trash"},
		Contexts: []string{"dad"},
	}
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	parser.ParseEditTask(task, "e 24 get the +garbage with @mom due tom")

	assert.Equal(task.Due, tomorrow)
	assert.Equal(task.Subject, "get the +garbage with @mom")
}
