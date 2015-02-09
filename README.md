# Entity--Boundary--Interactor: A modern application architecture

This repository is the example of the **Entity-Boundary-Interactor**
(**EBI**) application architecture as presented by Uncle Bob in his
series of talks titled
[Architecture: The Lost Years](https://www.youtube.com/watch?v=HhNIttd87xs).

The EBI architecture is a modern application architecture suited for a
wide range of application styles. It is especially suitable for web
applications and APIs, but the idea of EBI is to produce an
implementation agnostic architecture, it is not tied to a specific
platform, application, language or framework. It is a way to design
programs, not a library.

Ultimately, the goal is the *separation of concerns* between application layers, this architecture and many like it aren't dependent on presentation models or platforms.

This repository contains a range of examples of how to implement the
architecture in different languages.

# Glossary

![An illustration](http://i.imgur.com/4K9BM2D.png)

The architecture is best described as a *functional data-driven*
architecture, where requests are processed into results. The
architecture consists of three different components.

* **Entities** are the core of the architecture. Entities are abstract
  data structures independent of any representation, they represent
  the models and data of the program. They could be `Book`s in a
  library or `Employee` in an employee registry.

* **Boundaries** are abstractions for transferring data from the outside world to actual business logic side. Succintly put, boundaries are *delivery mechanisms*. bA boundary can implement a graphical user interface of an application or its HTTP API. Boundaries aren't usually concrete implementations but usually defined as abstractions that *interactors* implement.

* **Interactors** manipulate entities. Their job is to accept requests through the boundaries and manipulate application state. Interactors is the business logic layer of the application: interactors *act* on requests and decide what to do with them. Interactors know of request and response models called **DTOs**, data transfer objects. Any interactor function is a mapping from a request DTO to a response DTO.

Additionally, there is one particular specialized boundary called an *entity gateway*. 

* **Entity gateways** are *internal* unexposed components that interactors depend upon for their program logic. An internal database is an entity gateway, a web service that interactors need is an entity gateway. The distinction between a boundary, of which a gateway is a specialized abstraction, is that boundaries are usually visible to the outside world, or they are often *the* link to the outside world, but gateways are totally invisible to the end user.
