# tl
=======
### An Essentalist Task List

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
Tasks are stored in .tasks.json per directory.

### New list
```tl init```

### Show list
```tl```

### Add task
```tl add "New Task :)"```

### Complete task
```tl done 0```
If the task is already done, this will revert it.

### Edit task
```tl edit 0 "No longer new D:"```

### Store task
```tl store 0```
Store only works on tasks marked done.
If the task is already stored, this will un-store it.

### Delete task
```tl del 0```

### Delete all stored tasks
```tl del store```

## Contributions
Any and all contributions or bug fixes are appreciated.
