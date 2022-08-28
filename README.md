# go-interfaces-with-constructors

This repository is a demo of how to keep your packages decoupled from dependencies that use method chaining,
like logrus.Logger for example.

Until Go 1.18 there was no way to define an interface
for external types that use method chaining without coupling your package to them.
For example, let's say TODO

```

```

Since Go 1.18 we can define interfaces that include constructors of themselves in a generic way:

```

```

The approach leads to convoluted code.

## References

- [proposal: Go 2: spec: add self type for use within interfaces](https://github.com/golang/go/issues/28254).
- [Self referencing interfaces in golang 1.18](https://medium.com/@mier85/self-referencing-interfaces-in-golang-1-18-bcd6b5701992).
- [How to use generics for creating self-referring interfaces](https://appliedgo.com/blog/generic-interface-functions).
