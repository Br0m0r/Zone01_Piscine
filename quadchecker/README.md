Quad Pattern Detection - Manual
Overview

This manual provides an explanation of the installation, usage, and structure of the quad pattern project, a collection of Go functions that generate and identify rectangular patterns (quads). It offers step-by-step instructions for building, running, and checking which quad function(s) correspond to given input patterns.
Introduction

The quad pattern project contains several Go functions to create rectangular patterns, each with distinct edge and corner characters. It also includes a checker function to identify which quad function matches a given pattern.
Installation Instructions

To install and build the project, follow these steps:

    Obtain the Project Files:

      Ensure you have all the necessary files: main.go, quadchecker.go, build.sh, and the quad function files (quadA.go, quadB.go, quadC.go, quadD.go, quadE.go).

    Compile the Project:
        Use the provided build.sh script to compile all quad functions and the checker into executables:

    Set Execution Permission: Use the chmod command to set execution permission on your build.sh script.
        chmod +x build.sh

    
    ./build.sh

    The script compiles the Go files, sets permissions, and organizes the source files.

Explanation of Key Components

Here's a summary of the primary components in this project:

    Main Script (main.go):
        Serves as the interface for interacting with the quad checker. It reads from standard input (stdin) and executes the quadchecker to determine which quad function(s) match the given pattern.

    Quad Functions (quadA.go, quadB.go, quadC.go, quadD.go, quadE.go):
        These functions generate rectangular patterns, each with unique edge and corner characters. The functions take width and height as parameters and output the corresponding pattern.

    Quad Checker (quadchecker.go):
        Contains the logic to compare an input pattern with the output from the quad functions. It executes each quad function with specified dimensions to identify which matches the given input.

    Build Script (build.sh):
        Compiles the quad functions and the checker into executables, sets permissions, and cleans up after compilation.

Usage Instructions

Identifying a Quad Pattern

To identify which quad function corresponds to a pattern, use the main.go script to run the quadchecker. For example, to generate a pattern with quadA and then check it with the quad checker, use:



./quadA 3 3 | go run .

The output will indicate which quad function matches the pattern along with its dimensions:



[quadA] [3] [3]

Handling Multiple Matches

Sometimes, a pattern may match more than one quad function. For example, running quadC 1 1 and piping the output into the checker may produce multiple matches:



./quadC 1 1 | go run .

The checker will return all possible matches:



[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]

Non-Quad Patterns

If a pattern does not correspond to any known quad function, the checker will indicate this. You can test with arbitrary patterns to see if they match any quad. For example:



echo -n "o--o"$'\n'"|"$'\n'"o" | go run .

The checker will return:



Not a quad function

    Additional Notes

    Ensure that the quadchecker and the quad function executables are in the same directory as main.go.
    Use the examples above to guide your understanding of quad patterns and how to use the project.