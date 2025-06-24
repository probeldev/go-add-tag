# go-add-tag

Утилита для автоматического добавления JSON-тегов к структурам Go

![Before](screenshots/before.png) → ![After](screenshots/after.png)

## Установка

### Go

    go install github.com/probeldev/go-add-tag@latest


### Nix

    nix profile install github:probeldev/go-add-tag


## Использование

Через Vim

Выделите структуру в визуальном режиме

Выполните команду:

    :'<,'>!go-add-tag
