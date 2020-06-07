# tl

### An essentalist task list

tl is a lightweight CLI task list without the bells and whistles of heavier lists like Taskwarrior.

tl aims to provide the simplicity of [t][tlink] with some of the features of [ultralist][ultralink]. 

[tlink]: https://github.com/sjl/t
[ultralink]: https://github.com/ultralist/ultralist

## Installation

### Compile from Source
```bash
go get github.com/kelr/tl
go install github.com/kelr/tl
```

### Precompiled Binaries
Binaries can be found at [releases][rel].

[rel]: https://github.com/kelr/tl/releases

## Usage
Tasks are stored in a .tasks.json in each directory.

### Create a new list
```tl init```

### Show the list
```tl```

### Add a new task
```tl add "hi im a new task"```

### Complete a task
```tl done 0```

### Delete a task
```tl del 0```

### Edit a task
```tl edit 0 "hi im no longer new"```
