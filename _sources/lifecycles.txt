Event Lifecycles
================

.. todo:: write summary
          
Web: the life-cycle of a HTTP request
-------------------------------------

.. figure:: ./images/life-cycle.png
   :align: right

   The request life-cycle of a request object in the HTTP context.

A *request DTO* enters the application via the request boundary. This is
usually the API layer sitting on top of some interactor. In the pictured
example, we have a ``GetGopher`` interactor whose task is to retrieve
information about a store of gophers, accepting ``GopherRequest``\ s and
returning ``GopherResponse``\ s. The *user interaction* is the request
DTO and in this example is in plain JSON.

The interactor ``GetGopher`` then can be seen as a mapping of
``GetGopherRequest``\ s to ``GopherResponse``\ s. Because the requests
and responses are **plain dumb objects**, this implementation is not
dependent of any technology. It is the duty of the API layer to
translate the request from, e.g., JSON, to the request DTO, but the
interactor doesn't know anything about the protocol or its environment.

GUI: the life-cycle of an event
-------------------------------

In the GUI model, the architecture looks a bit simpler, since the host
layer does all the important work for us. We don't need to worry about
serialization so much anymore.

The API layer of a GUI application translate user interactions to
request objects. The Host layer sends an ``AddAuthorButtonClick`` event
to the API. The API also listens to the ``AuthorNameFieldEdited`` event
using which it keeps track of the contents of the form field. Once the
``AddAuthorButtonClick`` event is received, the API creates a
``CreateAuthor`` request DTO with the appropriate details and sends it
to the interactor for processing. Once the interactor returns, the API layer pushes a response to the
Host via some mechanism, e.g., by formating and sending a
``UserAdded`` event to the Host layer, which in turn handles the
updating of the user list component using its own logic.

As you can see, a GUI application implemented with EBI is *instantly*
a lot more complicated than a simple web server. This isn't a
coincidence, it's natural: GUI apps **are** really complicated
underneath.

The extra cost of such abstractions is that as the interactor and
entity layer remain unmodified, you can easily swap the API and Host
layer for another implementation. So you're never tied to any certain
interaction or delivery mechanism.
