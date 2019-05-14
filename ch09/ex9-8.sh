# Listing 9.8  Cross-compile an application with gox

gox \
-os="linux darwin windows " \
-arch="amd64 386" \
-output="dist/{{.OS}}-{.Arch}}/{{.Dir}}" .
