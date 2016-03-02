Entity—Boundary—Interactor
==========================

.. rubric:: A modern application architecture

This repository contains a description and an example implementation examples of the
**Entity—Boundary—Interactor** (**EBI**) application architecture, derived from ideas initially
conceived by Uncle Bob in his series of talks titled `Architecture: The Lost Years
<https://www.youtube.com/watch?v=HhNIttd87xs>`__ and `his book
<http://www.amazon.com/Software-Development-Principles-Patterns-Practices/dp/0135974445/ref=asap_bc?ie=UTF8>`__.

EBI aims to **separate business logic from presentation logic**, creating applications that can be
easily extended to a wide variety of platforms and use cases.

EBI suited for a wide range of application styles. It is especially suitable for web application
APIS, but the idea of EBI is to produce an implementation agnostic architecture, it is not tied to a
specific platform, application, language or framework. It is a way to design programs, not a
library.

The name **Entity–Boundary—Interactor** originates from `a master's
thesis <https://jyx.jyu.fi/dspace/bitstream/handle/123456789/41024/URN:NBN:fi:jyu-201303071297.pdf?sequence=1>`__
where this architecture is studied in depth. Names that are common or
synonymous are **EBC** where *C* stands for **Controller**. Another similar idea is the
**Hexagonal** architecture.

Examples of how to implement the architecture are given in this document
and are written in *Elixir*, a dynamically typed language with a simple
and powerful syntax.

.. warning::

   This is still very much a work in progress. Contributions are
   welcome on `Github <https://github.com/ane/ebi>`_.

.. toctree::
   :caption: Entity–Boundary–Interactor

   introduction
   design
   architecture
   structuring
   lifecycles

.. toctree::
   :caption: Layers

   overview
   presentation
   boundary
   core

.. toctree::
   :caption: In Practice

   external

.. toctree::
   :caption: Examples

   A REST API <example>

.. toctree::
   :caption: Reference

   about
   license
   faq
   

