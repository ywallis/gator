# Gator

This is an go-based RSS aggrgator written as a boot.dev guided project

## Requirements:

- Golang 1.23+
- Postgres 16+ 

The project so far assumes that postgres runs on the localhost.

## Installation:

You can install my version of gator by running:

```
go install github.com/ywallis/gator
```

Then manually initialize a .gatorconfig.json file in your home directory with the following content:

```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}
```

In case your postgres configuration differs from this, adapt accordingly.

## Usage

- First, a user must be registered with: 

`gator register $username`

- A feed can then be added with:

`gator addfeed $feedname $feedurl`

- Posts are collected using the long lasting process:

`gator agg`

- Posts can be displayed with:

`gator browse $limit`
