The Service Layer
=================

The common language spoken by the boundaries and interactors are
requests and responses. Both interfaces are defined in ``service.go``.


We can now implement the Gophers service (which finds and stores
gophers) in ``service/gophers.go``.

.. code:: go

    package service

    import (
        "github.com/ane/ebi/service/requests"
        "github.com/ane/ebi/service/responses"
    )

    // Gophers is a boundary that can do things with gophers.
    type Gophers interface {
        Create(requests.CreateGopher) (responses.CreateGopher, error)
        Find(requests.FindGopher) (responses.FindGopher, error)
        FindAll(requests.FindGopher) ([]responses.FindGopher, error)
    }

Boundary complexity
-------------------

The above code presented a rather simple boundary, composed of just two
methods. This is obviously suitable for a simple web application, but
this is not the design goal of boundaries. The purpose of boundaries is
to *decouple* the application interface and its implementation from each
other.

When writing boundaries, there aren't any limits to their complexities.
They can contain just one method or a dozen method.

.. tip::
   
   In Go, it is idiomatic to aim for interface composition. The ``Gophers``
   boundary above is composed of two distinct interfaces. This allows for
   extensibility.

Though similar to multiple inheritance, Go interfaces allow for
decomposition. In Java you could define a class
``FinderCreator implements Finder, Creator`` but you **cannot decompose
them**. This means that in Go, it is entirely valid to define a function
``func Foo(c Creator)`` yet pass a ``FinderCreatorRemoverUpdater`` to it
as a parameter. In Java or its family you can't decompose multiple
inheriting classes or interfaces into their constituent interfaces.

The take-away points of boundary design are these:

1. Make loose coupling easy by defining abstract interfaces that aren't
   too monolithic.
2. Decompose if you can if your interfaces are too big, think about
   splitting them into modular parts.
3. Make boundaries synchronous. Calling them asynchronously in the API
   layer is easy. Make them mappings from requests to responses.

Request models
--------------

Once the boundaries are complete, then we can move to the request and
response models. These are implemented with simple structures that
contain no validation logic. They are simply information vectors.

In our example, the response and request models live in
``responses/gophers.go`` and ``requests/gophers.go``.

.. code:: go

    package requests

    type FindGopher struct {
        ID int
    }

    type CreateGopher struct {
        Name string
        Age  int
    }

.. code:: go

    package responses

    type FindGopher struct {
        ID   int
        Name string
        Age  int
    }

    type CreateGopher struct{
        ID  int
    }

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
