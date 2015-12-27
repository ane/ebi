Design Goals
============

.. todo:: write summary

Extensibility
-------------

Testability
-----------

Stability
---------

Summary
-------

The above architecture is suited for any language and **any use case**.
One only needs an ability to define abstractions, were they type
classes, interfaces, OCaml modules, Rust traits, or Clojure protocols.
Static typing is not required here: you just need *one* way of creating
clear and verifiable functionality definitions. In dynamically typed
languages like Clojure and Elixir you can use protocols (with runtime
assertions), or even just plain old documentation. The boundary layer
needs only to be *specified*, it's not a strict language requirement.

The arrows in this architecture tend to point inwards. Only the middle
layer (the service layer) is seen by both the Core and the API layer is
because it describes the language of the system, but none of its
functionality.

Keeping the arrows unidirectional will make the system more robust and
scalable. If you decide to port your GUI app to a web service the
interactors will stay the same.

Moreover, unit testing is easy: you can mock *anything*, and what is
more, the unit tests will be fast and simple. Entities will only test
their internal business logic, interactors will not fumble with web
services, the presentation layer will only deal with handling requests
and responses and calling the right interactor, the host layer will
contain system-specific tests (e.g. HTTP tests), but **all** of these
components can be tested separately in a horizontal fashion.
