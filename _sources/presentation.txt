The Presentation Layer
======================

The introduction of the API layer at this point may seem a bit
heavy-handed. Why not map the interactor methods directly to routes? I
mean, you could spin up a web server that handles Sinatra-like requests
and then points them to the right interactor and returns a serialized
version of whatever the interactor spewed out.

Indeed, if you have *one* interactor or a couple that don't ever do
business together, this seems like the right approach. Once you get too
many, it gets useful to wrap them underneath a single unit.

Suppose you have an API endpoint of a book catalogue, and you want to
implement functionality that that modifies a particular author and at
the same time transfers these modifications to the publications. You
receive the new author name as input, and then you must update the
author itself and their book catalogue in one go.

Sure, sounds easy, just create a ``ModifyCatalogue`` functionality into
the ``Author`` interactor. The interactor, in this case, would modify
the author's name, then loop over its ``Book``\ s and modify them
individually, finally sending the updates to a database. This system
works as long as the ``Book`` entities are under the sole ownership of
an ``Author``--that is, there is no way of adding, creating, deleting,
or modifying a book from outside.

As soon as you introduce a ``Book`` interactor into the mix, things
start to get hairy. The ``Author`` service, retaining its book
modification logic, now overlaps with the ``Book`` service. The imminent
solution to this is to lift this logic from the ``Author`` interactor to
the ``Book`` interactor, making the layout look like this.

.. figure:: ./images/book-author-problem.png
   :alt: a problem

   Could the blue arrow be removed, and contained inside the arrows
   from the API layer pointing towards the Service layer?

The blue dashed arrow can be lifted into the API layer with little extra
work. It's a good idea to push such arrows as far "up" as possible,
because this helps keep one thing in check: not violating the **single
responsibility principle**, which roughly means that your interactor
should do one thing and **one thing only**. So the Author interactor
should only care about author logic, and the Book interactor should care
only about book issues.

In the above example this process would not be violated if there was no
Book service, such that book-related logic was underneath the Author
interactor. But, as soon as you start sharing responsibilities, and they
start to overlap, you will run into problems.

Hence, the API layer is there to provide additional logic that ties two
interactors together. You could think of it as a *meta-interactor*,
something that operates on interactors only, but contains no low-level
business logic.

What is more, the API layer usually has some knowledge of the
application domain: while interactors deal with dumb objects (DTOs), the
API may be dealing with HTTP request objects. Thus, the API is closer to
the actual implementation.

Consequently, the **Core** layer is the non-duplicated, non-overlapping
part of the application: you may have multiple APIs for the same set of
interactors, and multiple *hosts* for each API, but at the fundamental
level, there's only one canonical implementation of the core.

To conclude, the key differences between an API and an interactor are
the following:

-  An API is domain-specific and knows about the target implementation.
   The API knows it is talking to a web server. It just doesn't know
   *which kind* of web server it is talking to, acting as a bridge
   between interactors and the delivery mechanism.
-  The API layer may tie a multitude of interactors together, without
   making them dependent on each other, enforcing loose coupling.
-  APIs can be seen as "meta-interactors", operating on interactors the
   same way interactors operate on entities.
