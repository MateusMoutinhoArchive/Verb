# API Samples List

## Description
A reference list of all Go examples provided in this repository. Each one is a self-contained `package main` that hands the real process argv to the library and drives it from code. Running one is [RunApiSample.md](/docs/Tutorials/RunApiSample.md).

## Examples

| Example | Description |
| --- | --- |
| [Presence](/examples/libraryExamples/Presence) | Checks a boolean flag with `IsPresent`, consuming it from the argv. |
| [Options](/examples/libraryExamples/Options) | Reads every occurrence of a repeatable `--flag value` option. |
| [KeyValues](/examples/libraryExamples/KeyValues) | Reads every occurrence of a repeatable `key=value` argument. |
| [StringArg](/examples/libraryExamples/StringArg) | Reads a positional argument by its absolute index in the argv. |
| [NextArg](/examples/libraryExamples/NextArg) | Drains the leftover positional arguments with the Unused Mechanic. |
