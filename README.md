# tl

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
tl looks for tasks in `.tasks.json` in the cwd.

### New list
```tl init```

### Add task
```tl add New Task```

### Show list
```tl```

### Complete task
```tl done 0```

If the task is already done, this will revert it.

### Edit task
```tl edit 0 No longer new D:```


### Store task
```tl store 0```

Store only works on tasks marked done.

If the task is already stored, this will un-store it.

### Delete task
```tl del 0```

This is permanent.

### Show specific lists
```tl list all```
```tl list store```
```tl list done```

## Using Projects
Each task file can store multiple projects.

The default project is called main.

### Add project
```tl p add new-proj```

### List projects
```tl p```

### Switch current project
```tl p new-proj```

### Edit project 
```tl p edit new-proj bettername```

### Delete project
```tl p del bettername```

Deleting a project is permanent.


## Contributions
Any and all contributions or bug fixes are appreciated.
