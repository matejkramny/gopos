# gopos
ESC/POS library for Golang

## How to use

Use REDIS publish

## JSON:

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