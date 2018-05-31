Example for golang/go#25671

Program 25671 demonstrates a bug in shiny/windriver (#25671)

The program creates a main window (blue).
For any click into the window and a green client window appears.
Pressing the ESC key asks shiny to close the current window
by sending a lifecycle event.
The event loop exits and calls the deferred Release method on the window.
If this is done on the main window, it works as expected.
However if it is done on a client window, it does not disappear.
It is left in an unusable state (try to make it bigger), as the event loop is
not running anymore.

The fix is mentioned in the issue.
