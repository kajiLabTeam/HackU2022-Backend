# HackU2022-Backend

# Introduction
このサービスは、Xclothesというサービスで、すれ違った人の服を評価できるAndroidアプリである。
そのXclothesのバックエンドの実装方法について記す。

# Environment
 | Tool |  Version |
| ---- | ---- |
|  Go  |  1.18.5  |
|  Docker  |  20.10.17 |

# Usage

# go

brew install go

go get github.com/gin-gonic/gin

go get -u github.com/jinzhu/gorm

go get github.com/go-sql-driver/mysql

go get -u github.com/teris-io/shortid

# docker

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

brew install --cask docker

open /Applications/Docker.app

docker-compose build

docker-compose up -d

# execution

go run main.go


