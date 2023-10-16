# GOpt: Optional Values for Go

Originally, go could only represent a missing value by one of two methods:

* A pointer (missing <=> nil) - This requires using the heap and adding pointer chases.
* The Zero value - This requires removing the zero value as a valid value, or being able to distinguish it from "not present".

Neither of these are acceptable answers. With Go 1.18 comes generics, and the ability to represent optional values in a reasonable way.

This library is that reasonable way.
