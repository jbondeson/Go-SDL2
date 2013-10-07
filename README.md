# Introduction

This a fork of [Go-SDL2](http://github.com/scottferg/Go-SDL2) which
supports SDL version 2.0 in a slightly different way from the version
of which this is a fork. That version, in turn, forked
[Go-SDL](http://github.com/Banthar/Go-SDL2) to add SDL version 2.0
support. 

The SDL2 support is still very much work in progress! It is, in fact,
completely broken right now.

Differences from scottfergs library are:

* No mutex handling.

* No custom event loop adaptation to use a goroutine and channel:
  Catch events in the same way as using the C API directly.

* Remove the audio/ submodule since the file headers have dubious
  licensing terms (and it probably doesn't work for SDL2).

* Import paths are "github.com/krig/Go-SDL2/..."

Differences from Banthar's original library are:

* SDL functions (except for SDL-mixer) can be safely called from concurrently
  running goroutines
* All SDL events are delivered via a Go channel
* Support for low-level SDL sound functions

* Can be installed in parallel to Banthar's Go-SDL
* The import paths are "github.com/scottferg/Go-SDL2/..."


# Installation

Make sure you have SDL2, SDL2-image, SDL2-mixer and SDL2-ttf (all in -dev version).

Installing libraries and examples:

    go get -v github.com/krig/Go-SDL2/sdl


# Credits

Music to test SDL2-mixer is by Kevin MacLeod.
