# gopos
Distributed ESC/POS library for Golang with templating

## How to use

Use REDIS publish

## Features..

- ESC/POS command library
- Templating (best feature imho)
- You can print by pushing to redis, its very useful if there are many printers.

## JSON pushed to redis:

```
{
	"print": "... template ..."
}
```

```
Redis> publish print "...blob of json..."
```

## Example Template

```
[[justify 1]]Kitchen receipt[[lf]][[justify 2]]
justified just right

[[justify 0]]justified left..
[[cut]]
```

Prints:

```
         Kitchen receipt

               justified just right
justified left..
        {paper cut here}
```

The commands in the templates are in `gopos.go`.
