## Entity&ndash;Boundary&ndash;Interactor: A modern application architecture

This repository contains implementation examples of the **Entity-Boundary-Interactor**
(**EBI**) application architecture as presented by Uncle Bob in his
series of talks titled
[Architecture: The Lost Years](https://www.youtube.com/watch?v=HhNIttd87xs) and [his book](http://www.amazon.com/Software-Development-Principles-Patterns-Practices/dp/0135974445/ref=asap_bc?ie=UTF8).

The EBI architecture is a modern application architecture suited for a
wide range of application styles. It is especially suitable for web application APIS, 
but the idea of EBI is to produce an implementation agnostic architecture, it is not tied to a specific
platform, application, language or framework. It is a way to design
programs, not a library.

Examples of how to implement the architecture are given in this document and are written in Go.

## Goals & Motivation

> "The architecture of something screams the intent." &mdash;Robert C. Martin

As Martin points out, a lot of the times when looking at web applications you see library and tooling extrusions, but the *purpose* of the program is opaque. 

> "The architecture of an application is driven by its use cases." &mdash;Ivar Jacobsen

The idea is to design programs so that their architectures immediately present their use case. EBI is a way to do that. It's a way to design programs so that its modules are organized cleanly and its architecture uses loose coupling to remain extensible.

Ultimately, the goal is the *separation of concerns* between application layers, this architecture and many like it aren't dependent on presentation models or platforms.

# Glossary

![An illustration](https://dl.dropboxusercontent.com/u/11213781/ebi/overview.png)

The architecture can be approached from two different perspectives. The first is the depedency graph, as you can see above. The second is the hierarchy graph, which presents a concrete separation in a program.

The architecture is best described as a *functional data-driven*
architecture, where requests are processed into results. The
architecture consists of three different components.

* **Entities** are the core of the architecture. Entities are abstract
  data structures independent of any representation, they represent
  the models and data of the program. They could be `Book`s in a
  library or `Employee` in an employee registry.

* **Boundaries** are the link to the outside world. A boundary can implement functionality for processing data for a graphical user interface or a web API. Boundaries are functional in nature: they accept data *requests* and produce *responses* as result. These abstractions are concretely implemented by interactors.

* **Interactors** manipulate entities. Their job is to accept requests through the boundaries and manipulate application state. Interactors is the business logic layer of the application: interactors *act* on requests and decide what to do with them. Interactors know of request and response models called **DTOs**, data transfer objects. Interactors are **concrete** implementations of boundaries.

![API](https://dl.dropboxusercontent.com/u/11213781/ebi/api.png)
*What the object diagram of the program looks like.*

# Request and Response Lifecycle for Interactors

![Request lifecycle](https://dl.dropboxusercontent.com/u/11213781/ebi/lifecycle.png)

A *request DTO* enters the application via the request boundary. This is usually the API layer sitting on top of some interactor. In the pictured example, we have a `GetGopher` interactor whose task is to retrieve information about a store of gophers, accepting `GopherRequest`s and returning `GopherResponse`s. The *user interaction* is the request DTO and in this example is in plain JSON.

The interactor `GetGopher` then can be seen as a mapping of `GetGopherRequest`s to `GopherResponse`s. Because the requests and responses are **plain dumb objects**, this implementation is not dependent of any technology. It is the duty of the API layer to translate the request from, e.g., JSON, to the request DTO, but the interactor doesn't know anything about the protocol or its environment.

What does a program using this architecture look like?

# Module Hierarchy

![Organization](https://dl.dropboxusercontent.com/u/11213781/ebi/hierarchy.png)

Furthermore, it is good practice to separate the EBI architecture itself into five different layers. These layers correspond to namespaces or packages in your language of choice.

* The **Host** layer implements a physical manifestation of the API, e.g., a web server
* The **API** layer is the interface to the program itself, which accepts input and translates it into DTOs, passing them to 
* The **Boundary** layer which is an **abstract** interface between interactors and the API, the boundary layer is concretely is implemented by 
* The **Service** layer which receives information from and delivers results to the boundary, containing the main program logic which manipulates 
* The **Entity** layer which contains objects that represent program models and data

Thus, when a program is constructed, the API is given 

* A set of boundaries it needs to talk to
* A set of interactors that implement these functionalities

## File structure (example)

Using the above list, the application can be structured as follows. 

```
.
├── api
│   └── gopher.go
├── core
│   ├── entities
│   │   ├── entity.go
│   │   └── gopher.go
│   └── interactors
│       └── gopher.go
├── host
│   └── webserver.go
├── main.go
└── service
    ├── boundaries
    │   └── gopher.go
    ├── requests
    │   └── gopher.go
    ├── responses
    │   └── gopher.go
    └── service.go

```

## Implementation 

The `api` folder contains the API, the `host` web servers or GUI apps, the `service` contains the boundary layer with the request and responses models, the `core` layer contains the core program architecture hidden from view. 

As mentioned previously, the purpose of the program should be visible by looking at it. We can now explore the boundaries in `boundaries/gopher.go` to see what the program is supposed to do.

### Service layer

The common language spoken by the boundaries and interactors are requests and responses. Both interfaces are defined in `service.go`.

asdfasdf

Do note that due to the Go [package naming convention](http://blog.golang.org/package-names) the files are often the same. This deliberate, as the request and response models are bound to represent a single target entity, in this case, a `Gopher`, which is defined in `service/gopher.go`. Entities are objects that contain their internal validation logic, thus they implement the `Validate`

```Go
package entities

type Gopher struct {
	Name	string
	Age	int
}
```

Entities are manipulated by interactors, which implement boundaries. An example boundary could be as follows.

```Go
package boundaries

type GopherFinder  {
	Find(requests.GetGopher) (responses.GetGopher, error)
}
```

The boundary also contains request and response models.

```Go
package requests

type GetGopher struct {
	Id	int
}
```
```Go
package responses

type GetGopher struct {
	Id	int
	Name	string
}
```

An interactor that implements this boundary simply need to implement the `Find` method.

```Go
package services

type GopherService struct {
	Gophers map[int]Gopher
}

// Registering this method to the GopherService struct, "(gs GopherService)",
// makes GopherService implement the GopherFinder interface. Yay for structural typing!
func (gs GopherService) Find(req requests.GetGopher) (responses.GetGopher, error) {
	if gopher, exists := gs.Gophers[req.Id]; exists {
		return responses.GetGopher{Id: req.Id, Name: gopher.Name}, nil
	} else {
		return responses.GetGopher{}, errors.New("Gopher not found")
	}
}
```

When an API is constructed, it is given the interactors as parameters.

```Go
package api

type GopherAPI struct {
	GopherFinder boundaries.GopherFinder
}
```

Finally, the host is a simple web server that builds the API.

