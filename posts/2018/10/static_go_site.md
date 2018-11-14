Building a Static Site Generator with Go
2018/10/29
golang,meta
@@@@@@@@@@
# Building a Static Site Generator with Go
This post relates to the creation of this website.

## Stack


## Dependencies


## Anatomy of a Post
Every post, such as this one, is simply markdown... well, almost. It's really a slightly
basterdized markdown in order to allow for some extra information to prepended to the beginning of
the file. The header section uses the following format:

``` markdown
title
date (YYYY/MM/DD)
tags (CSV e.g. golang,meta)
@@@@@@@@@@

# Post Content...
```

The reason for this is primarily because I am lazy and did not want to deal with the overhead of
having to keep track of multiple files for each post. Also, having a separate YAML, JSON, etc. file
for such a small amount of information is unnecessary at this point.