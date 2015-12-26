
Talking to the External World
=============================

One part I haven't yet addressed in this overview is how to talk to
external dependencies, like a database. The answer is remarkably simple:
create them behind a boundary and build them like an interactor. This
enforces loose coupling, and the interactors *still* talk to each other
using interfaces.

Similarly, if you're building a GUI application and want to use events,
the interactor can push events to an event broker boundary, or the API
layer can handle the responses from the interactor, and call other
interactors through their boundary interfaces. This brings us to the API
layer.
