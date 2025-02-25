# InterfaceTest

`InterfaceTest` is a simple project used to test a way to use an Interface in Go to handle different kinds of data sources for a single type of data.

Here, there are users, both in `JSON` and `SQL database` format, with a single `UserCRUD` interface meant to give the **CRUD** functions definitions to any data-source related to the application's data.

`DbModel` and `JsonModel` both implement the `UserCRUD` interface to retrieve the **Users**, and they are present in the `application` struct for ease of use across multiple functions. It would be easy to add a new data-source to this project, or even to modify or remove one.

## Getting Started

#### Prerequisites
You need `Go v1.24.0` or superior to run this project.

#### Instructions
- Open a terminal in the project folder (`InterfaceTest`)
- Type and execute:
	```bash
	go run .
	```