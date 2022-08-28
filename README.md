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