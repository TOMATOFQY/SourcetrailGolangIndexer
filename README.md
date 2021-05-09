# SourcetrailGolangIndexer

Golang Indexer for Sourcetrail.

![](./img/println.png)

## Description

The SourcetrailGolangIndexer is a Sourcetrail language extension supply Golang support to Sourcetrail.

Currently, this project is still in a very early state, but you can already build correct Call Graph on your Golang project.

## Setup

> It could only be excuted on linux now (exactly, Debian), the whole bindings will be published very soon then you can recompile the bindings on your platform .

Just check `run.sh` or run it directly. This project absolutely depends on native tools supported by golang.org without third-party dependency.

After executing `runs.sh`, you should have this:

![](./img/example.png)

For usage of Sourcetrail, please check its official website.

## TOOD

- Add support for class hierachy.
