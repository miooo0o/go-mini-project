#### About
A mini project to study the Go language.  
The goal is to reflect my personal needs as much as possible.

#### Plan
- [ ] **Set the CSV path via `config.json`**  
- [ ] **Create a TODO list based on the configured CSV file**  
- [ ] **Structure the project into three main parts:**  
  - `cmd/` → Command-line interface (CLI) logic  
  - `tasks/` → Task management logic (CRUD operations)  
  - `config/` → Configuration handling (load/save settings)  
- [ ] **Add `alarm` command** (Trigger alerts for scheduled tasks)  
- [ ] **Add `remind` command** (Remind about upcoming tasks)


```
~/.tasks-cli/         // or <root>/.tasks-cli
│── tasks             // executable
│── config.json       // config file
│── tasks.csv         // data base

install.go	    		  -> create/generate "tasks-cli && files"
```