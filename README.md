## tag-checkout-branch-git
easier git tag checkout

## usage
add following alias to `.gitconfig`
```
[alias]
  tag-checkout-by-branch = !git tag | peco | tag-checkout-branch-git
```
## peco (filter tool)
- peco/peco: Simplistic interactive filtering tool https://github.com/peco/peco

## other filter tools
- junegunn/fzf: A command-line fuzzy finder https://github.com/junegunn/fzf
- lotabout/skim: Fuzzy Finder in rust! https://github.com/lotabout/skim
