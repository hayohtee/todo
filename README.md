# todo
A command-line tool for creating and managing a list of "to-do" item, 
keep track of items left in the todo list and saving the list in a file
using the JSON format.

## Features
- Add a new "to-do" item.
- Save the list of "to-do" item to the current working directory in JSON format.
- Marking a "to-do" item as completed.
- Listing "to-do" items.
- Deleting a "to-do" item.

## Building from source
1. Clone the repository
   ```bash
   git clone git@github.com:hayohtee/todo.git
   ```
2. Change into the project directory
   ```bash
   cd todo
   ```
3. Compile
   ```bash
   go build -o todo ./cmd/todo
   ```

## Usage
1. Add new task
   You can add new task either by using command-line arguments or from STDIN\
   Here is an example of using command-line arguments:\
   ```bash
   ./todo -add The name of the task you wanted to add
   ```
   Here is an example of using STDIN
   ```bash
   echo "The name of the task you wanted to add" | ./todo -add
   ```
2. List all tasks
   ```bash
   ./todo -list
   ```
3. Mark a task as completed
   To mark a task as completed, you supply the position of the task in the list\
   *Note* positon starts from 1\
   Here is an example of completing task 1
   ```bash
   ./todo -complete 1
   ```
4. Show all available options
   ```bash
   ./todo -h
   ```
5. Delete a task
   To delete a task, you supply the position of the task in the list to -del flag\
   Here is an example of deleting task 1
   ```bash
   ./todo -del 1
   ```
