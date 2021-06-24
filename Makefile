README.md: HEADER.md skills.json education.json references.json experience.json resume.go
	go run resume.go

pdf: README.md
	pandoc --template=resume.latex --variable mainfont="FreeSerif" --variable sansfont="FreeSans" --variable monofont="Inconsolata Nerd Font Mono" -M date="`date "+%B %e, %Y"`" --variable fontsize=12pt --variable version=2.0 README.md --pdf-engine=xelatex --toc -o resume.pdf

clean:
	rm -rf README.md
