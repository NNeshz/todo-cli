# TODO APP IN CLI

## Description

This is a simple todo app in CLI. It is written in Golang and uses a JSON file to store the tasks.

## Demo

https://github.com/user-attachments/assets/f0611ce1-010f-48aa-bbaa-23dc80443e04

## Installation

### Setup

To set up the project, follow these steps:

1. Clone the repository to your local machine:

   ```
   git clone https://github.com/your-username/golang-todo.git
   ```

2. Navigate to the project directory:

   ```
   cd golang-todo
   ```

3. Install the necessary dependencies:

   ```
   go mod download
   ```

4. Build the executable:

   ```
   go build -o todo
   ```

5. Run the application:
   ```
   ./todo -h
   ```

Now you have successfully set up the todo app in your CLI. Enjoy managing your tasks!

### Usage

To use the todo app, you can run the following commands:

- To add a task:
  ```
  ./todo -add "Task description"
  ```
- To list all tasks:
  ```
  ./todo -list
  ```
- To complete a task:
  ```
  ./todo -mark <task-id int>
  ```
- To delete a task:
  ```
  ./todo -delete <task-id int>
  ```
- To clear all tasks:
  ```
  ./todo -deleteall
  ```
