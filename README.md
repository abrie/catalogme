# (<sub><sup>experimental</sup></sub>) Catalog Site Builder

A tool to build/generate a digital catalog according to a flat JSON schema.

This project requires data files not currently checked into the repository.<br> So, at present, it's not useful to anyone other than myself.
## Requirements
NPM, Golang 1.13, HAProxy ~2, Python ~3.

## Clone and install
1) Clone this project
2) `make initialize-dependencies`

## Dev environment
1) Launch `tmux`
2) `make up` to start the servers and hot compilers/reloader.
3) Open `http://localhost:8888`

## Customize
Edit `./config.json` to change ports.

<hr>
#### dev note: experiments using grid layout for item lists
https://codepen.io/anon/pen/xNNXBR
