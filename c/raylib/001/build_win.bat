
@echo off

set originalDir=%CD%

if not defined DevEnvDir (
	call "C:\Program Files\Microsoft Visual Studio\18\Community\VC\Auxiliary\Build\vcvarsall.bat" x64
)

mkdir build
cd build

cl ..\src\main_win.c /link ..\lib\raylib.lib gdi32.lib msvcrt.lib winmm.lib user32.lib shell32.lib /NODEFAULTLIB:libcmt /NODEFAULTLIB:msvcrtd

cd /d "%originalDir%"
