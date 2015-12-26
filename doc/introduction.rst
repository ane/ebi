Introduction
============

.. epigraph::

    The architecture of something *screams* the intent.

    -- Robert C. Martin

As Martin points out, a lot of the times when looking at web
applications you see library and tooling artifacts, but the *purpose* of
the program is opaque. You open up a repository of a web application
and what you see is a load of configuration files, complicated
directory structures and lots of extraneous cruft that one day become
invisible. Could this opacity be avoided?

.. epigraph::

    The architecture of an application is driven by its use cases.

    -- Ivar Jacobsen

The idea is to design programs so that their architectures immediately
present their use case. EBI is a way to do that. It's a way to design
programs so that its modules are organized cleanly and its architecture
uses loose coupling to remain extensible.

Ultimately, the goal is the *separation of concerns* between application
layers, this architecture and many like it aren't dependent on
presentation models or platforms. All the arrows, or dependencies,
point inwards in the abstraction chain, each successive layer is
less abstract than the one before it.

Keeping the arrows pointing inwards makes the code easy to maintain,
extend, test and refactor. EBI imposes some architectural requirements
on the programmers, that is, you must navigate around its rules, but
this is kept at a minimum.



