Structuring Applications
========================

.. figure:: ./images/hierarchy.png
   :alt: Organization

   The code-level organization of modules. Each vertical section is an
   separate module of the program.

Furthermore, it is good practice to separate the EBI architecture itself
into five different layers. These layers correspond to namespaces or
packages in your language of choice.

-  The **Host** layer implements a physical manifestation of the API,
   e.g., a web server
-  The **API** layer is the interface to the program itself, which
   accepts input and translates it into DTOs, passing them to
-  The **Service** layer that contains **boundaries** and **response**
   and **request** models
-  The **Core** layer that contains a concrete implementation of the
   service layer
-  **Interactors** which implement boundaries and form the core business
   logic of the application
-  **Entities** which represent the data models of the program

Thus, when a program is constructed, the API is built top-down using
dependency injection. The **Host** layer is the one doing the DI of the
concrete interactors.

And that's it. The interactors do not know what protocol its requests
come from or are sent to, and the API doesn't know what sort of an
interactor implements the service boundary.
