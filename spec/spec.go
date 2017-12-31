package spec

//go:generate sh -c "npm list spectacle-docs || npm install spectacle-docs"
//go:generate ./node_modules/.bin/spectacle api.yaml
//go:generate go-bindata -nometadata -pkg spec -o assets.go -prefix public ./public/...
