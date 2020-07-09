# Shakespeare Autocomplete
Coding Interview for Zesty.io

## Installation/Setup
- Build project with `go build server.go`.
- Run server with `./server.exe` on Windows or `./server` on a Unix environment.
- Server runs on port 9000.
- Service can be accessed with `curl --location --request GET 'http://localhost:9000/autocomplete?term=<term>' with <term> replaced with whatever term you want to search for.
## Other information
- Pre-specified results can be found in `results.txt`.
### Known limitations:
- Words separated by punctuation and no spaces are considered a single word
- With a larger word bank, rerunning populateDictionary could become time consuming, so there could be a separate service that just generates the word list in usage-order, and it wouldn't have to repopulate that data.
- Wrong endpoints aren't properly handled.