## Command-Line Interface (CLI) Programming

The command-line is a versatile, efficient, and portable interface that has
seen continuous, ubiquitous use since the mainframe computing era. To this day,
it remains the primary interface for managing server and other production
systems, and is regularly used by computer professionals to complete both
simple and complex tasks with ease.

The capabilities of command-line systems have evolved over time, but the
underlying principles have remained much the same: a command line program
accepts program *arguments*, inherits an *environment*, and operates on input
and output text streams.

## The Command Shell

A program that interprets command descriptions and accordingly arranges for
programs to run is known as a *shell*. It is so named because it encapsulates
the lower level operating system *kernel*; the software that directly manages
system hardware, process invocation and scheduling, and other critical tasks,
but which, unlike a shell, does not any direct user interface.

Shells vary in the syntax they accept and the features they provide, but
typically provide at least: command pipelining, stream redirection, conditional
control flow, pattern matching, job control, and, for interactive shells,
command history.

The [Bourne Shell][1], released by Bell Labs in 1977, was the first Unix shell
with excellent programming features; it set the standard in shell syntax and
capabilities, and has left a strong legacy of similar shells that are referred
to as "Bourne shell compatible" (or "sh compatible" for short, after the
shell's command name). The following sections are described in terms of Bourne
shell compatible semantics and syntax.

  [1]: http://en.wikipedia.org/wiki/Bourne_shell

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
