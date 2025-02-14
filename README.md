# go-multiline-formatter

Форматер для neovim что бы преобразовывать однострочные конструкции в языке GO в многострочные

Например:

    func myfunc(a int, b int, c int) error {

Преобразуется в 

    func myfunc(
        a int,
        b int,
        c int,
    ) error {


# Установка

     go install github.com/probeldev/go-multiline-formatter@latest

# Использование в NeoVim

1. Выделить строку

2. использовать комманду !go-multiline-formatter 

что бы использовать hotkey(пробел + f) необходимо в ~/.config/nvim/init.vim добавить:

    vnoremap <space>f :'<,'>!go-multiline-formatter \| gofmt<CR>



