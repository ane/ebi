The Core Layer
==============

The core layer contains actual business logic. First we start off with
the entity, the rich business objects of the application. In
``core/entities/entity.go``,

Entities
--------

.. code:: go

    type Gopher struct {
        Name string
        Age int
    }

Entities are completely invisible to the outside layers. Nothing
but the interactors know about them. Entities contain business logic,
e.g., a ``Gopher`` entity can modify itself or contain functions related
to it, but the distinction between entities and interactors is the
following:

-  entities modify *themselves* vs.
-  interactors modify *entities*

An entity can contain other entities: a ``Gopher``, could technically
possess a ``Tail`` and two ``Eye``\ s, and it can modify them at will.
This hierarchy is strictly unidirectional: a ``Gopher`` doesn't know
about other gophers, more importantly, *it doesn't know about the
interactor*.

Interactors
-----------

Interactors contain rich business logic. They can manipulate entities
and they implement boundaries. Here, we have the ``Gophers`` boundary
from above to implement, so we implement a smallish interactor that
implements it.

.. code:: go

    type Gophers struct {
        burrow map[int]entities.Gopher
    }

    func NewGophers() *Gophers {
        return &Gophers{
            burrow: make(map[int]entities.Gopher),
        }
    }

It implements the three methods as defined by the ``Gophers`` boundary

.. code:: go

    // Find finds a gopher from storage.
    func (g Gophers) Find(req requests.FindGopher) (responses.FindGopher, error) {
        gopher, exists := g.burrow[req.ID]
        if !exists {
            return responses.FindGopher{}, errors.New("Not found.")
        }

        return gopher.ToFindGopher()
    }

    func (g Gophers) FindAll(req requests.FindGopher) ([]responses.FindGopher, error) {
        var resps []responses.FindGopher
        for _, gopher := range g.burrow {
            fg, err := gopher.ToFindGopher()
            if err != nil {
                return []responses.FindGopher{}, err
            }
            resps = append(resps, fg)
        }
        return resps, nil
    }

    // Create creates a gopher.
    func (g Gophers) Create(req requests.CreateGopher) (responses.CreateGopher, error) {
        var gopher entities.Gopher
        if err := gopher.Validate(req); err != nil {
            return responses.CreateGopher{}, err
        }

        gopher.ID = g.getFreeKey()
        gopher.Name = req.Name
        gopher.Age = req.Age
        g.burrow[gopher.ID] = gopher

        return responses.CreateGopher{ID: gopher.ID}, nil
    }

As one can see, the interactor is completely unaware of any protocol
dependencies. The relation to web applications is obvious: we are, after
all, talking about requests and responses, and the DTOs translate very
easily to JSON objects. But they can be used without JSON, in fact, the
whole point is that even a GUI application will pass the same objects
around.

The interactors (and by extension, entities) are completely oblivious to
their environment: they don't care whether they are running inside a GUI
application, a system-level daemon, or a web server.

Beware of Behemoths
-------------------

Interactors are business logic units. How much business logic is too
much business logic? The best rule of thumb is the **single
responsibility principle**: an interactor should only do one thing, and
one thing only. I'm also going to address this
`below <#the-api-layer>`__, but the most important thing to understand
about interactors is that they should operate only one *one* aspect of
the business logic.

What this means may not be immediately clear. If you are building a REST
API, you will generally have some separation of concerns already going
on at the external API level, in the form of URIs. To use a book
catalogue as an ad hoc example, you could have a URI for book authors at
``/authors`` and ``/books``, these clearly indicate---to the API user,
anyway---what lies beneath.

At the code level, this distinction must be maintained. An author may
contain a collection of books they have, but whose responsibility is
modifying them? Obviously, since we have two URIs here, one for books,
one for authors, we must decide which one handles the logic of modifying
book entities. In this case, any internal *modification logic* of the
book entities must reside underneath a **single** interactor. There can
be two cases here:

-  **One interactor does everything**. The ``/books`` URI is just an
   alias underneath the Author interactor, or vice versa.

   -  **Pros**: no overlap in logic, no conflicts, since everything is
      contained under one unit (a single interactor).
   -  **Cons**: must be split eventually, since otherwise it will grow
      to monstrous proportions.

-  **Two interactors, ``AuthorInteractor`` and ``BookInteractor``**. The
   ``AuthorInteractor`` calls methods of the ``IBookService`` (which
   ``BookInteractor`` implements) to modify the ``Book`` entities
   contained (or *owned*) by an ``Author`` entity.

   -  **Pros**: no chance of overlap since the responsibilities are
      split.
   -  **Cons**: risk of introducing circular dependencies between
      boundaries (see `below <#the-api-layer>`__).

If you're building a really simple service, you don't *have* to split
interactor duties, but it's a good idea. Be careful of choosing
short-term practicality in favor of long-term abstractions, it may bite
you in the rear one day!

As a summary, in the presented example, the ``AuthorInteractor`` should
only modify things related to ``Author``\ s, and preferably only *read*
data about ``Book``\ s, leaving modification and updates to the
``BookInteractor``. There are two ways on how to implement the necessary
communication, that is, how the ``AuthorInteractor`` calls the
``BookInteractor``, and this will be resolved later, but now we have a
small interlude about something equally vital: the external world.
