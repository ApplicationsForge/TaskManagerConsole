package taskterminal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type FileStore struct {
	FileLocation string
	Loaded       bool
}

func NewFileStore() *FileStore {
	return &FileStore{FileLocation: "", Loaded: false}
}

func (f *FileStore) Initialize() {
	if f.FileLocation == "" {
		f.FileLocation = "tasks.json"
	}

	_, err := ioutil.ReadFile(f.FileLocation)
	if err == nil {
		fmt.Println("It looks like a tasks.json file already exists!  Doing nothing.")
		os.Exit(0)
	}
	if err := ioutil.WriteFile(f.FileLocation, []byte("[]"), 0644); err != nil {
		fmt.Println("Error writing json file", err)
		os.Exit(1)
	}
	fmt.Println("Task repo initialized.")
}

func (f *FileStore) Load() ([]*Task, error) {
	if f.FileLocation == "" {
		f.FileLocation = getLocation()
	}

	data, err := ioutil.ReadFile(f.FileLocation)
	if err != nil {
		fmt.Println("No task file found!")
		fmt.Println("Initialize a new task repo by running './TaskTerminal init'")
		os.Exit(0)
		return nil, err
	}

	var tasks []*Task
	jerr := json.Unmarshal(data, &tasks)
	if jerr != nil {
		fmt.Println("Error reading json data", jerr)
		os.Exit(1)
		return nil, jerr
	}
	f.Loaded = true

	return tasks, nil
}

func (f *FileStore) Save(tasks []*Task) {
	data, _ := json.Marshal(tasks)
	if err := ioutil.WriteFile(f.FileLocation, []byte(data), 0644); err != nil {
		fmt.Println("Error writing json file", err)
	}
}

func getLocation() string {
	localrepo := "tasks.json"
	usr, _ := user.Current()
	homerepo := fmt.Sprintf("%s/tasks.json", usr.HomeDir)
	_, ferr := os.Stat(localrepo)

	if ferr == nil {
		return localrepo
	} else {
		return homerepo
	}
}
