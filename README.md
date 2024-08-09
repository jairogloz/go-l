# ⚽ GO-L

## Introduction

Introducing "GO-L", a cutting-edge RESTful API backend, meticulously crafted with Go. This project is not just about managing soccer tournaments and leagues, it's about redefining the way we interact with sports data.  

Built on the robust foundation of the Gin framework and employing the Hexagonal Architecture, GO-L ensures clean, maintainable, and testable code. It's a beacon of best industry practices, designed to handle the complexities of real-world applications.  

But what truly sets GO-L apart is its collaborative spirit. It's a project by Go enthusiasts, for Go enthusiasts. It's a learning platform, a place to explore, experiment, and elevate your Go skills. With GO-L, you're not just coding, you're contributing to a community, and shaping the future of Go programming.  Join us in this exciting journey and let's redefine the game together with GO-L!

## Creating mocks

To create a mock for a specific interface using mockgen, execute the following command:

```bash
mockgen -destination=mocks/<name of the mock file>.go -package=mocks --build_flags=--mod=mod <path to the package where the interface is> NameOfTheInterfaceToMock
```

Example:

```bash
mockgen -destination=mocks/mock_team_repository.go -package=mocks --build_flags=--mod=mod github.com/jairogloz/go-l/pkg/ports TeamRepository
```

## Running tests

To run tests for the whole project execute the following command:

```bash
go test ./...
```

If you want to skip integration tests (because those require docker) you can run:

```bash
go test -short ./...
```

## Contributing

Thank you for considering contributing to [GO-L]! We appreciate your efforts and contributions to improving our project. Please follow these guidelines to contribute effectively at [Contributing Guide](CONTRIBUTING.md).