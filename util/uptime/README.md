## Uptime calculator for range of blocks

### How to use


1. Rename config.toml.example to config.toml
    ```sh
        mv config.toml.example config.toml
    ```
2. Edit mongodb credentials

3. Run the script with startblock, endblock flags

```sh
go run main.go --start 0 --end 1000
```