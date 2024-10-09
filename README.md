# Currency Monitor

A small CLI application written in Golang to display exchange rates against the russian ruble.

## Usage

**Clone the repository:**

```sh
git clone https://github.com/jdeeline/currency-monitor.git
cd currency-monitor
```

**Build:**

```sh
go build ./cmd/monitor
```

**Run on Linux or Mac:**

```sh
./monitor
```

**Run on Windows:**

```sh
./monitor.exe
```

Use the -i flag to change the update interval in minutes (default: 60):

```sh
./monitor -i 240
```
