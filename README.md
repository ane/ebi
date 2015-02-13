## Entity&ndash;Boundary&ndash;Interactor

### A modern application architecture

This repository contains implementation examples of the **Entity-Boundary-Interactor**
(**EBI**) application architecture as presented by Uncle Bob in his
series of talks titled
[Architecture: The Lost Years](https://www.youtube.com/watch?v=HhNIttd87xs) and [his book](http://www.amazon.com/Software-Development-Principles-Patterns-Practices/dp/0135974445/ref=asap_bc?ie=UTF8).

The EBI architecture is a modern application architecture suited for a
wide range of application styles. It is especially suitable for web application APIS, 
but the idea of EBI is to produce an implementation agnostic architecture, it is not tied to a specific
platform, application, language or framework. It is a way to design
programs, not a library.

The name **Entity&ndash;Boundary&ndash;Interactor** originates from [a master's thesis](https://jyx.jyu.fi/dspace/bitstream/handle/123456789/41024/URN:NBN:fi:jyu-201303071297.pdf?sequence=1) where this architecture is studied in depth. Names that are common or synonymous are **EBC** where *C* stands for **Controller**.

Examples of how to implement the architecture are given in this document and are written in Go.

## Goals & Motivation

> "The architecture of something screams the intent." &mdash;Robert C. Martin

As Martin points out, a lot of the times when looking at web applications you see library and tooling extrusions, but the *purpose* of the program is opaque. 

> "The architecture of an application is driven by its use cases." &mdash;Ivar Jacobsen

The idea is to design programs so that their architectures immediately present their use case. EBI is a way to do that. It's a way to design programs so that its modules are organized cleanly and its architecture uses loose coupling to remain extensible.

Ultimately, the goal is the *separation of concerns* between application layers, this architecture and many like it aren't dependent on presentation models or platforms.

# Glossary

![An illustration](docs/images/overview.png)

The architecture can be approached from two different perspectives. The first is the depedency graph, as you can see above. The second is the hierarchy graph, which presents a concrete separation in a program.

The architecture is best described as a *functional data-driven*
architecture, where requests are processed into results. The
architecture consists of three different components.

* **Entities** are the core of the architecture. Entities represent 
business objects that have application independent business rules.
They could be `Book`s in a library or `Employee` in an employee registry.
All the application agnostic business rules should be located in the entities.

* **Boundaries** are the link to the outside world. A boundary can implement functionality for processing data for a graphical user interface or a web API. Boundaries are functional in nature: they accept data *requests* and produce *responses* as result. These abstractions are concretely implemented by interactors.

* **Interactors** manipulate entities. Their job is to accept requests through the boundaries and manipulate application state. Interactors is the business logic layer of the application: interactors *act* on requests and decide what to do with them. Interactors know of request and response models called **DTOs**, data transfer objects. Interactors are **concrete** implementations of boundaries.

![API](docs/images/boundary.png)
*What the object diagram of the program looks like.*

# Request and Response Lifecycle for Interactors

![Request lifecycle](docs/images/lifecycle.png)

A *request DTO* enters the application via the request boundary. This is usually the API layer sitting on top of some interactor. In the pictured example, we have a `GetGopher` interactor whose task is to retrieve information about a store of gophers, accepting `GopherRequest`s and returning `GopherResponse`s. The *user interaction* is the request DTO and in this example is in plain JSON.

The interactor `GetGopher` then can be seen as a mapping of `GetGopherRequest`s to `GopherResponse`s. Because the requests and responses are **plain dumb objects**, this implementation is not dependent of any technology. It is the duty of the API layer to translate the request from, e.g., JSON, to the request DTO, but the interactor doesn't know anything about the protocol or its environment.

What does a program using this architecture look like?

# Module Hierarchy

![Organization](docs/images/hierarchy.png)

Furthermore, it is good practice to separate the EBI architecture itself into five different layers. These layers correspond to namespaces or packages in your language of choice.

* The **Host** layer implements a physical manifestation of the API, e.g., a web server
* The **API** layer is the interface to the program itself, which accepts input and translates it into DTOs, passing them to 
* The **Service** layer that contains **boundaries** and **response** and **request** models
* The **Core** layer that contains a concrete implementation of the service layer
  * **Interactors** which implement boundaries and form the core business logic of the application
  * **Entities** which represent the data models of the program

Thus, when a program is constructed, the API is given 

* A set of boundaries it needs to talk to
* A set of interactors that implement these functionalities

And that's it. The interactors do not know what protocol its requests come from or are sent to, and the API doesn't know what sort of an interactor implements the service boundary.

## Example SOA implementation in Go

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
│       └── gophers.go
├── host
│   └── webserver.go
├── main.go
└── service
    ├── requests
    │   └── gopher.go
    ├── responses
    │   └── gopher.go
    └── service.go
    └── gophers.go
```

## Implementation 

The `api` folder contains the API, the `host` web servers or GUI apps, the `service` contains the boundary layer with the request and responses models, the `core` layer contains the core program architecture hidden from view. 

As mentioned previously, the purpose of the program should be visible by looking at it. By exploring the `service` directory (containing `gophers.go` *et al.*) we can immediately see the services this program provides.

### Service layer

The common language spoken by the boundaries and interactors are requests and responses. Both interfaces are defined in `service.go`.

```Go
package service

// Request is a request to any service.
type Request interface{}

// Response is a response to a request.
type Response interface{}
```

These are empty interfaces. As a result, in Go, any type implements this interface, so this is just naming sugar for now, as logic can be added into these interfaces later when this architecture spec develops further.

The service layer also contains definitions for different sorts of boundaries. The simplest of which is a `Finder` boundary that finds resources based on a request, in `service/boundaries.go`.

```Go
package service

// Creator is a boundary that creates resources.
type Creator interface {
	Create(Request) (Response, error)
}

// Finder is a boundary that finds resources.
type Finder interface {
	Find(Request) (Response, error)
}
```

We can now implement the Gophers service (which finds and stores gophers) in `service/gophers.go`.

```Go
package service

type Gophers interface {
	Finder
	Creator
}
```

The response and request models live in `responses/gophesr.go` and `requests/gophers.go`.

```Go
package requests

type FindGopher struct {
	ID int
}

type CreateGopher struct {
	Name string
	Age  int
}
```

```Go
package responses

type FindGopher struct {
	ID   int
	Name string
	Age  int
}

type CreateGopher struct{}
```

The naming convention is to have a service "Foobar"  (in caps, can be a pluralized noun), and have it in `service/foobar.go`, and its request and response models are *all* in `service/requests/foobar.go` and `service/responses/foobar.go`.

Though these interfaces are named similarly, in Go, we refer to these types as `requests.GetGopher`, hence it is never ambiguous as to what the structures are. The `requests` (or responses) packages contain only structures like these, hence there will never be any confusion between the two.

This concludes the definition for the Service layer. This the the common "runtime" of the application, if you will, these are the elements both the API and internal logic see. It can also be exported as a library, hence it is easy to extend.

### Core layer

The core layer contains actual business logic. First we start off with the entity, the rich business objects of the application. In `core/entities/entity.go`,

```Go
package entities

import "github.com/ane/ebi/service"

// Validator is an interface that validates incoming requests to see if these are valid transformations.
type Validator interface {
	Validate(service.Request) error
}

// Translator is an interface that translates itself into a simpler structure.
type Translator interface {
	As(service.Response) error
}

// Entity is an interface for an object that contains business rules. It can validate incoming
// transformations to itself and it can also convert itself into simpler data structures.
type Entity interface {
	Validator
	Translator
}
```
