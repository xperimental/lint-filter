lint-filter
===========

This is a small command-line tool for filtering XML report files generated by the [Android lint](http://developer.android.com/tools/help/lint.html) tool. Currently it can only filter by location, as this was the only filter I currently needed.

Usage
-----

````
usage: lint-filter [options] input-file output-file

Options:
  -config="filterconfig.json": Path to filter configuration file.
  -pretty=false: If present, XML output will be indented.
````

Configuration
-------------

The filter configuration is done using a simple JSON file. A sample configuration for filtering out all issues in the `target/` directory is included.
