# SourcetrailGolangIndexer

Golang Indexer for Sourcetrail.

![](./img/println.png)

## Description

The SourcetrailGolangIndexer is a language extension for Sourcetrail that provides support for Golang. Although this project is still in its early stages, it allows you to generate an accurate call graph for your Golang project.
## Setup

Currently, the SourcetrailGolangIndexer can only be executed on Linux, specifically Debian. However, the complete bindings will be released soon, which will allow you to recompile them on your platform.

To run the SourcetrailGolangIndexer, you can either check the `run.sh` script or execute it directly. It's worth noting that this project relies solely on native tools supported by golang.org and doesn't have any third-party dependencies.

Once you have executed the `run.sh` script, you should have the following:

![](./img/example.png)

For information on how to use Sourcetrail, please refer to its official website.

## TOOD

- Add support for class hierachy.

## WARNING

As of March 10th, 2022, I have completed a runnable version of this extension, which was finalized in April 2021. However, I discovered that the master branch of the project has since been archived, meaning I cannot make a pull request. Unfortunately, this also means that there will be no further progress on the project.

If you're interested in this extension and find that it lacks a CMakeFiles.txt for Golang, you can simply copy the CMakeFiles.txt from the official repository and replace the language name with Golang (e.g. Python => Golang). Since this part hasn't changed much, it shouldn't be a big issue.

Additionally, there are other heuristic projects like goplantuml that perform similar functions, but with standard plantuml output. You may find these projects to be more suitable for your needs.
