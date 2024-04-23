$env:CGO_CFLAGS="-I$env:TIBRV_HOME/include"
$env:CGO_LDFLAGS="-L$env:TIBRV_HOME/lib"
$env:TESTDIR="C:/tmp"
$path = "C:\tmp"
If(!(test-path -PathType container $path))
{
    New-Item -ItemType Directory -Path $path
}
