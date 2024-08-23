# TodoMan - A Todo App

Another cli based todo app. This time, I'm building it with Go.

## Note

Im using a python script to build and distribute the app. Sadly, I won't be able to provide it.
If you want to build it yourself, you can use the following command:

```bash
    go clean && go build -o bin/todoman
```

Also, I've changed the package urls to go.smsk.dev, this is my personal domain and I'll be using it for my personal projects.

## Run Locally

Clone the project

```bash
  git clone https://github.com/devsimsek/todoman.git my-project
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod tidy
```

Start the application

```bash
  go run .
```

## Tech Stack

Using Go for the backend.
Sqlite3 for the database.
Tablewriter for the list view.
Gorm for the database operations.
And a simple framework for handling cli interface. (Maybe I'll publish it as a separate package in the future.)

## Features

- Add a todo
- List all todos
- Mark a todo as done
- Delete a todo

## Roadmap

- Publish minimal viable product (current state)
- Add workspaces
- Add tags
- Add priorities
- Add due dates
- Add reminders
- Add recurring todos
- Add subtasks

## Feedback

If you have any feedback, please reach out to me using the issues tab.

## License

[MIT](https://devsimsek.mit-license.org/)
or
license.md file in the root directory.

## Authors

- [@devsimsek](https://github.com/devsimsek)
