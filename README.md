# Small tcp server written in go

## Architecture details:

- The 'manager' package encapsulates the server initialization logic (listener configuration), the client connection logic and a high-level request router to the correct handler
- The 'handlers' package contains project requirements specific logic
- The 'config' package contains a 'reader' component responsible with the json configuration deserialization
- The 'utils' package contains a tokenizer for the client requests and a parser that extract numeric data from a given token 