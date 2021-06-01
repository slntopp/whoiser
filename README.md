# WHOISER

Quick `whois` lookups.

## Usage

### Get WHOIS data as JSON

```shell
sh# whoiser -domain example.com
{
    "domain": {
        "expiration_date": "2021-08-13T04:00:00Z"
        ***
    }
}
```

### Make queries

Pass queries under `-query` key as the comma-separated list of Pascal-Case written keys path, like:

```shell
sh# whoiser -domain example.com -query Domain/ExpirationDate
ExpirationDate: 2021-08-13T04:00:00Z
```

### Output format

1. Results only

    ```shell
    sh# whoiser -domain example.com -query Domain/ExpirationDate -print-keys=0
    2021-08-13T04:00:00Z
    ```

2. Keys only

    ```shell
    sh# whoiser -domain example.com -query Domain/ExpirationDate
    ExpirationDate: 2021-08-13T04:00:00Z
    ```

    or

    ```shell
    sh# whoiser -domain example.com -query Domain/ExpirationDate -print-keys
    ExpirationDate: 2021-08-13T04:00:00Z
    ```

3. Full path

    ```shell
    sh# whoiser -domain example.com -query Domain/ExpirationDate -queries
    ExpirationDate: 2021-08-13T04:00:00Z
    ```
