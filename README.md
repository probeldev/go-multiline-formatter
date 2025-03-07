# Go Multiline Formatter

Помошник для редактирования кода на языке go для Vim/NeoVim/Emacs


## Возможности

### Преобразование однострочной функции в многострочную

Используется флаг: 

    -multiline

Пример работы:

    func myfunc(a int, b int, c int) error {

Преобразуется в:

    func myfunc(
        a int,
        b int,
        c int,
    ) error {
    

### Добавление переменной с названием структуры и функции

Исользуется флаг:
    
    -logfunc
    
Пример работы:

    func (s *structname) myfunc(a int, b int, c int) error {
    
Преобразуется в:

    func (s *structname) myfunc(a int, b int, c int) error {
        fn :="structname:myfunc"
    

## Установка

     go install github.com/probeldev/go-multiline-formatter@latest

## Использование в Vim/NeoVim

1. Выделить строку

2. использовать комманду !go-multiline-formatter 

что бы использовать hotkey(пробел + f и пробел + l) необходимо в ~/.config/nvim/init.vim добавить:

    vnoremap <space>f :'<,'>!go-multiline-formatter -multiline<CR>
    vnoremap <space>l :'<,'>!go-multiline-formatter -logfunc<CR>
