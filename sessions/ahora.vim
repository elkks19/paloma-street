let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +1 services/CategoriasService.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/negocios.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/productos.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/errors.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/middleware.go
badd +27 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/run.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/reviews.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/usuarios.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/favoritos.go
badd +8 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/auth.go
badd +18 .dockerignore
badd +23 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/Dockerfile
badd +26 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/fly.toml
badd +14 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/.github/workflows/fly-deploy.yml
badd +2 .gitignore
badd +4 README.md
badd +110 services/postgres/AuthService.go
argglobal
%argdel
set stal=2
tabnew +setlocal\ bufhidden=wipe
tabrewind
edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/auth.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 105) / 210)
exe 'vert 2resize ' . ((&columns * 179 + 105) / 210)
argglobal
enew
file neo-tree\ filesystem\ \[2]
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/auth.go
setlocal foldmethod=manual
setlocal foldexpr=0
setlocal foldmarker={{{,}}}
setlocal foldignore=#
setlocal foldlevel=0
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal foldenable
wincmd w
argglobal
setlocal foldmethod=manual
setlocal foldexpr=0
setlocal foldmarker={{{,}}}
setlocal foldignore=#
setlocal foldlevel=0
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal foldenable
silent! normal! zE
let &fdl = &fdl
let s:l = 6 - ((5 * winheight(0) + 28) / 56)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 6
normal! 014|
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 105) / 210)
exe 'vert 2resize ' . ((&columns * 179 + 105) / 210)
tabnext
edit services/postgres/AuthService.go
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
argglobal
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/auth.go
setlocal foldmethod=manual
setlocal foldexpr=0
setlocal foldmarker={{{,}}}
setlocal foldignore=#
setlocal foldlevel=0
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal foldenable
silent! normal! zE
let &fdl = &fdl
let s:l = 110 - ((40 * winheight(0) + 27) / 55)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 110
normal! 05|
tabnext 2
set stal=1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
