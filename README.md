# gowgo

**Go with Go** is a simple and easy tutorial-like website with a minimal static webpage structure. The
purpose of gowgo is to help people learn or serve as a reference to certain features in [Go](https://golang.org).

# development

I created a small templating language for this for the fun of it when creating these tutorials. Originally, they were
all completely hand-formatted which made it really hard to create and maintain. I could've used something like
*Velocity* or *Handlebars*, but that would not have been as fun.

Now there is syntax highlighting, css linking, a single template file, and a little state.

There is also a little vanilla JS for copying code examples to clipboard and executing examples against the 
[playground](https://play.golang.org). I would like to keep this pretty minimal though, because I want this to be
as bug-free as possible.
