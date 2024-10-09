# Task Tracker

Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on.

## Tech stack
Built mainly with Go. No external package.

## How to run
- Cline repository
```bash
git clone https://github.com/arikchakma/backend-projects.git
cd backend-projects/task-tracker
```

- Run the following command to build and run the project:
```bash
# To build the app
go build -o task-tracker

# To see a list of avaiable command
./task-tracker --help

# To add a new task
./task-tracker --add "Buy groceries"

# To update a task
./task-tracker --add "1~Buy groceries and cook dinner"

# To delete a task
./task-tracker delete 1

# To mark a task as in progress/done/todo
./task-tracker mark-in-progress 1
./task-tracker mark-done 1
./task-tracker mark-todo 1

# To list all tasks
./task-tracker list
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```

## Reference
https://roadmap.sh/projects/task-tracker