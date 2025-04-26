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
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/cmd/app/main.go
badd +55 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/errors.go
badd +21 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/run.go
badd +49 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/server.go
badd +4 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/internal/config/env/env.go
badd +5 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/internal/db/db.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/internal/db/dsn.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/internal/db/tx.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/internal/utils/printStruct.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/pkg/validate/validate.go
badd +2 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/pkg/validate/rules.go
badd +2 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/.env
badd +2 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/AuthService.go
badd +32 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/ProductosService.go
badd +27 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/ReviewsService.go
badd +40 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/UsuariosService.go
badd +15 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/CategoriasService.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/FavoritosService.go
badd +54 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/NegociosService.go
badd +27 makefile
badd +3 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/modd.conf
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/.gitignore
badd +38 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/SessionService.go
badd +36 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/goth/AuthService.go
badd +88 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/SessionService.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/build/migrations.mk
badd +134 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/AuthService.go
badd +25 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/CategoriasService.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/UsuariosService.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/NegociosService.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/ProductosService.go
badd +9 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/FavoritosService.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/ReviewsService.go
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/000_create_users_table.sql
badd +18 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/001_create_roles_table.sql
badd +12 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/002_create_usuarios_table.sql
badd +4 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/003_create_sessions_table.sql
badd +15 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/004_create_negocios_table.sql
badd +6 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/006_create_productos_table.sql
badd +4 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/007_create_categorias_table.sql
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/008_create_productos_categorias_table.sql
badd +11 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/009_create_favoritos_table.sql
badd +16 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/010_create_reviews_table.sql
badd +1 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/005_create_negocio_usuarios_table.sql
badd +11 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/errors.go
badd +16 /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/middleware.go
badd +108 ~/go/pkg/mod/github.com/labstack/echo-jwt/v4@v4.3.1/jwt.go
badd +10 ~/go/pkg/mod/github.com/golang-jwt/jwt/v5@v5.2.2/map_claims.go
badd +9 ~/go/pkg/mod/github.com/golang-jwt/jwt/v5@v5.2.2/claims.go
badd +24 ~/go/pkg/mod/github.com/golang-jwt/jwt/v5@v5.2.2/registered_claims.go
badd +32 ~/go/pkg/mod/github.com/golang-jwt/jwt/v5@v5.2.2/types.go
badd +31 ~/go/pkg/mod/github.com/golang-jwt/jwt/v5@v5.2.2/token.go
argglobal
%argdel
set stal=2
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabrewind
edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/AuthService.go
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
exe '1resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 1resize ' . ((&columns * 30 + 105) / 210)
exe '2resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 2resize ' . ((&columns * 158 + 105) / 210)
argglobal
enew
file neo-tree\ filesystem\ \[4]
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/AuthService.go
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
let s:l = 114 - ((45 * winheight(0) + 23) / 47)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 114
normal! 05|
wincmd w
exe '1resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 1resize ' . ((&columns * 30 + 105) / 210)
exe '2resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 2resize ' . ((&columns * 158 + 105) / 210)
tabnext
edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/SessionService.go
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
exe '1resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 1resize ' . ((&columns * 94 + 105) / 210)
exe '2resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 2resize ' . ((&columns * 94 + 105) / 210)
argglobal
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/003_create_sessions_table.sql
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
let s:l = 54 - ((39 * winheight(0) + 23) / 47)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 54
normal! 0
wincmd w
argglobal
if bufexists(fnamemodify("/share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/003_create_sessions_table.sql", ":p")) | buffer /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/003_create_sessions_table.sql | else | edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/003_create_sessions_table.sql | endif
if &buftype ==# 'terminal'
  silent file /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/db/migrations/003_create_sessions_table.sql
endif
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/SessionService.go
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
let s:l = 4 - ((3 * winheight(0) + 23) / 47)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 4
normal! 0
wincmd w
exe '1resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 1resize ' . ((&columns * 94 + 105) / 210)
exe '2resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 2resize ' . ((&columns * 94 + 105) / 210)
tabnext
edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/run.go
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
exe '1resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 1resize ' . ((&columns * 1 + 105) / 210)
exe '2resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 2resize ' . ((&columns * 187 + 105) / 210)
argglobal
enew
file neo-tree\ filesystem\ \[7]
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/errors.go
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
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/AuthService.go
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
let s:l = 23 - ((17 * winheight(0) + 23) / 47)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 23
normal! 028|
wincmd w
exe '1resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 1resize ' . ((&columns * 1 + 105) / 210)
exe '2resize ' . ((&lines * 47 + 28) / 57)
exe 'vert 2resize ' . ((&columns * 187 + 105) / 210)
tabnext
edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/AuthService.go
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
let s:l = 102 - ((29 * winheight(0) + 27) / 55)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 102
normal! 015|
tabnext
edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/http/server.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
2wincmd h
wincmd w
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
wincmd =
argglobal
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/FavoritosService.go
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
let s:l = 49 - ((38 * winheight(0) + 27) / 55)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 49
normal! 058|
wincmd w
argglobal
if bufexists(fnamemodify("/share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/FavoritosService.go", ":p")) | buffer /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/FavoritosService.go | else | edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/FavoritosService.go | endif
if &buftype ==# 'terminal'
  silent file /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/FavoritosService.go
endif
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/FavoritosService.go
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
let s:l = 1 - ((0 * winheight(0) + 27) / 55)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
wincmd w
argglobal
if bufexists(fnamemodify("/share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/FavoritosService.go", ":p")) | buffer /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/FavoritosService.go | else | edit /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/FavoritosService.go | endif
if &buftype ==# 'terminal'
  silent file /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/postgres/FavoritosService.go
endif
balt /share/Unifranz/Semestre\ 7/Programacion\ de\ dispositivos\ moviles/PROYECTO_FINAL/services/FavoritosService.go
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
let s:l = 9 - ((8 * winheight(0) + 27) / 55)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 9
normal! 05|
wincmd w
3wincmd w
wincmd =
tabnext 5
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
