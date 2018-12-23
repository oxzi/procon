# procon
Terminal UI based Pros and Cons List.

[![asciicast](https://asciinema.org/a/jJ3T3oiOdp9VSZxPCeGnxZAhH.svg)](https://asciinema.org/a/jJ3T3oiOdp9VSZxPCeGnxZAhH)


## Usage

Launch `procon` with an existing *Pros and Cons List* as parameter or a new
filename, where your list will be created.

```bash
procon moving.pc
```

The usage is kind of Vi-like.

- Navigation in the list is possible both with `hjkl` and the arrow keys.
- A selected entry can be deleted by pressing `x` or edited by pressing `Enter`.
- A new entry can be added by pressing `a`.
- The list can be saved by pressing `w`.
- The application can be closed by pressing `q`. If there are unsaved changes,
  the user will be asked if they should be saved.
