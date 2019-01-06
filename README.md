# envexample
A tool to make .env.example

## Installation

```
$ go get github.com/ikedaosushi/envexample
```

## Usage

Let's say you have a .env file following.

```
# This is a sample enviroment variables
FOO=BAR
HOGE=FUGAFUGA

# Second section
# SAMPLE=ENV
HELLO=WORLD
TEST=TEST
```

You can use

```
$ envexample
```

then, you get a .env.example following

```
# This is a sample enviroment variables
FOO=
HOGE=

# Second section
# SAMPLE=ENV
HELLO=
TEST=
```

You also can use

```
$ envexample .env .env.sample
```

or

```
$ envexample .myenv .myenv.sample
```