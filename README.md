# Entity&ndash;Boundary&ndash;Interactor: A modern application architecture

This repository contains implementation examples of the **Entity-Boundary-Interactor**
(**EBI**) application architecture as presented by Uncle Bob in his
series of talks titled
[Architecture: The Lost Years](https://www.youtube.com/watch?v=HhNIttd87xs) and [his book](http://www.amazon.com/Software-Development-Principles-Patterns-Practices/dp/0135974445/ref=asap_bc?ie=UTF8).

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

![An illustration](https://dl.dropboxusercontent.com/u/11213781/ebi/overview.png)

The architecture is best described as a *functional data-driven*
architecture, where requests are processed into results. The
architecture consists of three different components.

* **Entities** are the core of the architecture. Entities are abstract
  data structures independent of any representation, they represent
  the models and data of the program. They could be `Book`s in a
  library or `Employee` in an employee registry.

* **Boundaries** are abstractions for transferring data from the outside world to actual business logic side. Succintly put, boundaries are *delivery mechanisms*. A boundary can implement a graphical user interface of an application or its HTTP API. Boundaries are abstractions that *interactors* implement. A boundary is a mapping from a *request* to a *response*.

* **Interactors** manipulate entities. Their job is to accept requests through the boundaries and manipulate application state. Interactors is the business logic layer of the application: interactors *act* on requests and decide what to do with them. Interactors know of request and response models called **DTOs**, data transfer objects. Interactors are **concrete** implementations of boundaries.

Additionally, there is one particular specialized boundary called an *entity gateway*. 

* **Entity gateways** are *internal* unexposed components that interactors depend upon for their program logic. An internal database is an entity gateway, a web service that interactors need is an entity gateway. The distinction between a boundary, of which a gateway is a specialized abstraction, is that boundaries are usually visible to the outside world, or they are often *the* link to the outside world, but gateways are totally invisible to the end user.

# Request and Response Lifecycle for Interactors

![Request lifecycle](https://dl.dropboxusercontent.com/u/11213781/ebi/lifecycle.png)

A *request DTO* enters the application via the request boundary. This is usually the API layer sitting on top of some interactor. In the pictured example, we have a `GetGopher` interactor whose task is to retrieve information about a store of gophers, accepting `GopherRequest`s and returning `GopherResponse`s. The *user interaction* is the request DTO and in this example is in plain JSON.

The request DTO canand be modelled as an object like so.
```Go
type GetGopherRequest struct {
	Id		int
}

type GopherResponse struct {
	Id		int
	Name	string
}
```

The interactor `GetGopher` then can be seen as a mapping of `GetGopherRequest`s to `GopherResponse`s. Because the requests and responses are **plain dumb objects**, this implementation is not dependent of any technology. It is the duty of the API layer to translate the request from, e.g., JSON, to the request DTO, but the interactor doesn't know anything about the protocol or its environment.

# Implementational hierarchy

![Organization](https://dl.dropboxusercontent.com/u/11213781/ebi/hierarchy.png)

Furthermore, it is good practice to separate the EBI architecture itself into its own *service* layer, and provide a wrapping around that in the API layer. The API layer functionalities would commonly be registered to different HTTP verbs or user interactions in a GUI application, but still speak in the clean data language of DTOs. In summary, the API layer:

* contains a set of **boundaries** that one or more **interactors** implement
* translates data from the outside world into request and response DTOs, passing them to boundaries,
* is not dependent of any protocol, only data formats (JSON, XML),
* is unit testable with simple DTOs.

The **service** layer contains the **boundaries** that the API layer calls. Their actual implementation is implemented by interactors. The manipulation can occur in a variety of ways, but any database should be abstracted via some delivery mechanism in a gateway, viz., databases and other web services interactors rely on.

![API](https://dl.dropboxusercontent.com/u/11213781/ebi/api.png)
