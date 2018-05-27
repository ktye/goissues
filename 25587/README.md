Example for golang/go#25587

Program 25576 demonstrates a bug in shiny/windriver (#25576)

The program creates a main window (blue).
For any click into the window and a green client window appears.
Trying to close the client window does not make it disappear,
but leaves it around unfunctinally:
It does not redraw when resizing (try to make it bigger).
Closing the main window always works as expected.

The fix is mentioned in the issue
