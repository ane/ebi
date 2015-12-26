Example Implementation
======================

.. rubric:: A REST API in Elixir

This document features a small example implementation of this
architecture in *Elixir*. Elixir is a dynamically typed language
leveraging the Erlang virtual machine.

I chose Elixir because of its simple but powerful syntax. I originally
wanted to implement this in Ruby but I wanted clear examples of
*interfaces* and Ruby doesn't really have them. Thankfully, Elixir has
*protocols*, which let me write the boundary descriptions using a
high-level abstraction.

.. note:: Interfaces aren't absolutely necessary.

   You don't really *need* interfaces to implement boundaries, the
   language-level abstractions make it easier to understand in its own
   terms. Since no boundary object is an actual, concrete
   implementation, it quickly becomes obvious that the boundary
   objects, and thus the service layer, act as a *data model* inside
   the system.

   For Ruby and Python you could easily write a dummy abstract class
   with ``NoMethodImplementation`` exceptions being thrown left and
   right, in case of an unsatisfied boundary.

The Elixir implementation makes the use of the `Spirit
<https://github.com/citrusbyte/spirit>`_ microframework for
Elixir. Equivalent frameworks in other applications:

- **Ruby**: Sinatra, Cuba
- **JavaScript**: Express
- **Go**: net/http
- **C#**: ServiceStack
- **Java**: SparkJava

...and so on.

.. code:: bash

   .
   ├── api
   │   └── web_api.ex
   ├── entities
   │   ├── author.ex
   │   └── publication.ex
   ├── host
   │   └── server.ex
   ├── interactors
   │   ├── author_service.ex
   │   └── publication_service.ex
   └── service
       ├── protocols.ex
       ├── requests.ex
       └── responses.ex
