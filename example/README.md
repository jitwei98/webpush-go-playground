# example

## Quickstart

1. Generate VAPID keys
    ```
    go run keygen/main.go
    ```
1. Replace VAPID keys
   -  in `index.html`
   -  in `main.go`
1. Run a simple server to server index.html
   - Python 2
        ```
        python -m SimpleHTTPServer 8000
        ```
   - Python 3
        ```
        python -m http.server 8000
        ```
1. Go to `http://localhost:8000` and copy the logged subscription from the console. Hardcode `s` in `main.go` (this should be sent to the server)
   - Before
        ```go
            s := &webpush.Subscription{}
        ```
   - After
        ```go
            s := &webpush.Subscription{
                    Endpoint: "<ENDPOINT FROM CONSOLE>",
                    Keys: webpush.Keys{
                        Auth:   "<AUTH FROM CONSOLE>",
                        P256dh: "<P256dh FROM CONSOLE>",
                    },
            }
        ```


### 

## Access index.html

Replace the public VAPID key in index.html.

Use a tool such as SimpleHTTPServer to run a web server:

```
python -m SimpleHTTPServer 8000

or

python -m http.server 8000
```

Go to `http://localhost:8000` and copy the logged subsciption from the console.

## Test send a notification

Replace the public/private VAPID keys. Use the subscription you had from the first section.

```bash
go run main.go
```
