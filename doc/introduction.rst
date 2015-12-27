Introduction
============

.. epigraph::

    The architecture of something *screams* the intent.

    -- Robert C. Martin

Often when looking at the code of web applications, you are greeted
with a mass of folders, library installation, and tooling
configurations. The code is structured haphazardly into nondescriptive
folders like ``app``, and the dependency layout of the application is
a jungle.

As a result, the purpose and architecture of the program become
opaque. 

.. epigraph::

    The architecture of an application is driven by its use cases.

    -- Ivar Jacobsen

The idea is to design programs so that their architectures immediately
present their use case. It's a way to design programs so that its
internal dependency graph is organized cleanly and its elements are
joined together with as loose coupling as possible. 

Ultimately, the goal is the *separation of concerns* between
application layers, this architecture and many like it aren't
dependent on presentation models or platforms. All the arrows, or
dependencies, point inwards in the abstraction chain, each successive
layer less abstract than the one before it.

Keeping the arrows pointing inwards makes the code easy to maintain,
extend, test and refactor. EBI imposes some architectural requirements
on the programmers, that is, you must navigate around its rules, but
this is kept at a minimum.

How does this archtecture differ from MVC?
------------------------------------------

The difference between EBI and MVC is that an EBI architecture is
that the business logic of the application is designed to be
platform-agnostic of its delivery mechanism.

To paraphrase, this means that the business logic parts, interactors and
entities, do not know in which medium they're being accessed from. It
could be from a web server, a unit test, or a GUI application.

Contrast this to MVC, where there is *always* a dependency on the
delivery mechanism. No matter how hard you tried, you cannot tear
Rails controllers out of the web world.

What makes decoupling the business logic from the delivery mechanism a
good thing? This is outlined in the next section, `ref`:design:.

