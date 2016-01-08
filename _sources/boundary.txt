The Boundary Layer
==================

.. todo:: flesh out

Designing good DTOs
-------------------

DTOs have no business logic. Think of them as language constructs around
simple requests not dependent of any protocol.

In our Go program, the naming convention is to have a service "Foobar"
(in caps, can be a pluralized noun), and have it in
``service/foobar.go``, and its request and response models are *all* in
``service/requests/foobar.go`` and ``service/responses/foobar.go``.

Though these interfaces are named similarly, in Go, we refer to these
types as ``requests.FindGopher``, hence it is never ambiguous as to what
the structures are. The ``requests`` (or responses) packages contain
only structures like these, hence there will never be any confusion
between the two.

In other languages, you would usually have a suffix of some sorts or use
a namespace explicitly to avoid repetition.

Wrapping up
-----------

The service layer is the common language of the application
architecture. When the API and core speak to each other, they do so via
an abstract boundary. They use DTOs (data transfer objects), simple
structures of data, for communication. We now move on to the core layer
of the architecture.
