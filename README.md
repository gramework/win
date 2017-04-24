# gramework.win

Sources for the win! ;)

# Setup and run for development

### 1. Clone

```
git clone --recursive -j8 git@github.com:gramework/win.git
```

### 2. Build

```
go build
```

### 3. Run

```
sudo setcap cap_net_bind_service=+ep && ./win || sudo ./win || ./win
```

### 1, 2, 3 for laziest

```
git clone --recursive -j8 git@github.com:gramework/win.git
go build
sudo setcap cap_net_bind_service=+ep && ./win || sudo ./win || ./win
```

# Project structure

Documented in docs/README.md

# Contributing

Documented in .github/CONTRIBUTING.md
