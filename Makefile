.PHONY: pdf clean lint

README.md: resume.toml resume.go
	go run resume.go

pdf: README.md
	nix-shell -p pandoc 'texlive.combine { inherit (texlive) scheme-medium collection-fontsextra collection-latexextra; }' --run \
		'pandoc --template=resume.latex --variable mainfont="FreeSerif" --variable sansfont="FreeSans" --variable monofont="FreeMono" -M date="`date "+%B %e, %Y"`" --variable fontsize=10pt --variable version=2.0 README.md --pdf-engine=xelatex -o resume.pdf'

clean:
	rm -f README.md resume.pdf

lint:
	go tool golangci-lint run --timeout 5m
