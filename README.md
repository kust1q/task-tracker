# The task manager

## Description

First simple roadmap project, CLI

## Usage 

### Just compile file to use 
```bash
go build -o task-tracker main.go 
```

### Command list 
- add - adding a new task
- delete - deleting tasks
- update - updating task
- mark-in-progress - marking a task as in progress
- mark-done - marking a task as done
- list [done / todo / in-progress] - listing all tasks / by status

### Example

```bash
./task-tracker add "Buy groceries"
```
```bash
./task-tracker update 1 "Buy groceries and cook dinner"
```
```bash
./task-tracker delete 1
```
```bash
./task-tracker mark-in-progress 1
```
```bash
./task-tracker mark-done 1
```
```bash
./task-tracker list
```
```bash
./task-tracker list done
```
```bash
./task-tracker list todo
```
```bash
./task-tracker list in-progress
```

