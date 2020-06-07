# tl - Essentialist Tasklist

tl is a lightweight command line task list without the bells and whistles of heavier lists like Taskwarrior.

## Installation

### Compile from Source
```bash
$ go get github.com/kelr/tl
$ go install github.com/kelr/tl
```
## Usage
tl looks for a task file in your current directory. 

### Create a new list
`$ tl init`

### Print the list
`$ tl`

### Add a new task
`$ tl add "hi im a new task"`

### Complete a task
`$ tl done 0`

### Delete a task
`$ tl del 0`

### Edit a task
`$ tl edit 0`
