Overview
========

From the directory tree one can see that the code is organized into
separate namespaces. In the example implementation this is achieved by
splitting the code into different folders, since this is a one-to-one
mapping to packages (namespaces) in the Go programming language.

The ``api`` folder contains the API, the ``host`` web servers or GUI
apps, the ``service`` contains the boundary layer with the request and
responses models, the ``core`` layer contains the core program
architecture hidden from view.

As mentioned previously, the purpose of the program should be visible by
looking at it. By exploring the ``service`` directory (containing
``gophers.go`` *et al.*) we can immediately see the services this
program provides.

