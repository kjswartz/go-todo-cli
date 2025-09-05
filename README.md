# Go Todo App

Simple command-line todo application written in Go using Cobra + SQLite.

## Current Features
* Add todo items with priority (p1 highest -> p3 lowest; default p3)
* List items (incomplete by default) with filters for completed / all
* Mark items complete
* Update description of an item
* Delete items
* Persistent storage via SQLite at `$HOME/go/data/todo.db`

> Note: Code comments mention updating priority on an existing item, but that capability is **not yet implemented** (no flag/logic present). Roadmap below.

## Install
Clone then build (produces `todo` binary):

```bash
git clone https://github.com/kjswartz/go-todo-cli.git
cd go-todo-cli
go build -o todo .
```

Optionally install to your `$GOBIN` (or `$GOPATH/bin`):

```bash
go install ./...
```

Run `./todo --help` or `todo --help` after install.

## Database Location
The SQLite database is created (if missing) at:

```
$HOME/go/data/todo.db
```

Ensure the parent directory exists if you run into path issues (the app will create the file but expects the directory tree to exist).

## Commands & Flags

### Root
```
todo --help
```
Global flags currently include an unused sample `--toggle/-t` (from Cobra scaffold).

### Add
Add a new todo (description required, priority optional):
```
todo add -d "Read a book" -p 1
```
Flags:
* `-d, --description` (string, required)
* `-p, --priority` (int: 1,2,3; default 3)

### List
List items ordered by priority ascending (p1 first). Default shows only incomplete items.
```
todo list          # incomplete only
todo list -c       # only completed
todo list -a       # all items
```
Flags:
* `-c, --completed` show only completed
* `-a, --all` show all (overrides default filter)

Output format:
```
P1 | (3) Read a book
P2 | (4) Wash car [c]
```
`[c]` indicates completed.

### Update
Update a single field on an item by ID. Currently supports marking complete OR changing description.
```
todo update 3 -c                 # mark complete
todo update 3 -d "Read two books" # change description
```
Flags:
* `-c, --complete` mark item complete
* `-d, --description` new description text

### Delete
Delete an item by ID:
```
todo delete 3
```

## Example Session
```bash
$ todo add -d "Read a book" -p 1
Added todo: Read a book with priority 1

$ todo add -d "Wash car"
Added todo: Wash car with priority 3

$ todo list
P1 | (1) Read a book
P3 | (2) Wash car

$ todo update 2 -c
Todo item 2 marked as complete

$ todo list -a
P1 | (1) Read a book
P3 | (2) Wash car [c]

$ todo delete 1
Todo item 1 has been deleted
```

## Data Model (SQLite `todos` table)
| column      | type     | notes                       |
|-------------|----------|-----------------------------|
| id          | INTEGER  | primary key autoincrement   |
| description | TEXT     | task text                   |
| priority    | INTEGER  | 1..3 (lower = higher prio)  |
| completed   | BOOLEAN  | 0 = incomplete, 1 = done    |

## Roadmap / Ideas
* Implement priority update on existing items (flag like `-p` on update)
* Add tests for all commands & edge cases
* Validation for empty arguments / missing ID
* Support changing completion back to incomplete
* Colored output (e.g., green for completed)
* Configurable database location
* Bulk operations (complete/delete multiple IDs)
* Export / import (JSON / CSV)

## Development Notes
Generated with Cobra; learning project for experimenting with Go CLI patterns + SQLite. PRs / suggestions welcome.

## License
See `LICENSE`.
