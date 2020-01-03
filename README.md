# gomailctl
Simple email smtp CLI written in go.


## Usage
```
mime=$(cat << EOF
From: me@myself.com
To: you0@yourself.com
Subject: Hello you from me

hello world
EOF
)


go run main.go -s smtpserver:25 --from me@myself.com --recipient you@yourself.com -username user --password supersecret -c "$mime"
```

## Todo
- have run without ssl
- create `go build`
