# LeetCode CLI

LeetCode CLI is a Go-based command-line tool designed to fetch LeetCode problem details listed in a `README.md` file and save them locally. It utilizes the LeetCode API to retrieve problem information and stores it in a structured format.

## Features

- Parses a `README.md` file to extract problem slugs.
- Fetches problem details from the LeetCode API.
- Saves problem details in a local directory structure.
- Skips problems that are marked as "paid only".

## Installation

1. Install the tool using `go install`:
```bash
go install github.com/Baticaly/leetcode-cli@latest
```

2. Ensure your `GOPATH/bin` is in your `PATH` environment variable. You can add the following line to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.):
```bash
export PATH=$PATH:$GOPATH/bin
```
Then, reload your shell profile:
```bash
source ~/.bashrc
```

## Usage

1. Ensure your `README.md` file is formatted correctly with problem slugs. Example:
```
## LeetCode Archive

- [x] 1.Two Sum
- [ ] 2.Add Two Numbers
- [ ] 3.Longest Substring Without Repeating Characters
- [x] 4.Median of Two Sorted Arrays
```

2. Run the CLI tool:
```sh
leetcode-cli <README.md>
```

3. The tool will fetch the specified problems and save their details in the `problems` directory.

## Example

Given the following `README.md`:
```
## LeetCode Archive

- [x] 1.Two Sum
- [ ] 2.Add Two Numbers
- [ ] 3.Longest Substring Without Repeating Characters
- [x] 4.Median of Two Sorted Arrays
```
The tool will fetch details for "Add Two Numbers" and "Longest Substring Without Repeating Characters" and save them in the `problems` directory.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

The MIT License (MIT) - See the [LICENSE](LICENSE) for more details