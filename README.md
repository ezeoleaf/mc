# mc
MC

## Validation of chunks
The program takes input via a file that consists of a single line of chunks. A chunk contains zero or more other chunks. Every chunk opens and closes with matching characters in the following way.

- A chunk opening with (, must close with ).
- A chunk opening with [, must close with ].
- A chunk opening with {, must close with }.
- A chunk opening with <, must close with >.

## How to run it
You can run it using Docker directly.

```bash
docker run -it $(docker build -q .)
```

Or using the make target.
```bash
make run
```
