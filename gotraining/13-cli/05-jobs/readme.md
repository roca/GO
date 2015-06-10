## Job Control

Command, pipelines, and lists are all jobs, and jobs, by default, run in the
*foreground*; there may only be foreground job running in a shell at any given
time, and it will receive input and signals from the controlling terminal.

Jobs can be alternatively be started in or moved to the *background*.
Background jobs run normally, except that they do not receive terminal input,
and in that state they can only be controlled explicitly by PID or job number.
A job may be suffixed with a single ampersand to run it in the background:

    sleep 5 && echo "wakeup!" &

A foreground job can be suspended by typing `^Z` (Control-Z). Upon doing so,
control will be returned to the shell, and the job can be resumed in the
foreground or background using the `fg` or `bg` commands, respectively.  The
`fg` command may also be used to move a running background job to the
foreground.

The `jobs` command displays suspended and background jobs; each of these will
have a *job number*, beginning with 1. The first background job can therefore
be killed with `kill %1`, the second can be moved to the foreground with
`fg %2`, and so on. The percent syntax is known as a *jobspec*.

Finally, the `wait` command can be used wait for one or more background jobs,
which may be specified using jobspecs. If called without arguments, wait will
block until all active background pipelines have completed.

When a shell (or shell script) exits, all background processes are disowned,
and will continue to run.

Correct use of job control can allow tasks to run concurrently, and make much
more efficient use of computing resources. For example, to build several go
binaries, a good technique would be:

    go build cmd1 &
    go build cmd2 &
    wait

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
