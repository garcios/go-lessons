# Pointers
- Variables that store value addresses instead of values.

## Why pointers?
1. Avoid unnecessary value copies
 - By default, Go creates a copy when passing values to functions.
 - For very large & complex values, this may take up too much memory space unnecessarily.
 - With pointers, only one valu is stored in memory (and the address is passed around).

2. Directly mutate values
- Pass a pointer (address) instead of value to a function.
- The function can then directly edit the underlying value - no return value is required.
- Can lead to less code (but also to less understandable code or unexpected behaviors).


