# about

`promfmt` is a very naive PromQL formatter. One can pipe a promql query in it, and
it gets formatted.


I really hope someone will create a proper formatter some day, but for now this
one will have to do in my vimrc :).

## Example

```
echo 'sum by (bla, bladiebla, woot) (rate(somemetric{labelx="asdfasdf",labely="jajanounou"}[5m])) * 1024' | ./promfmt

sum by (bla, bladiebla, woot) (
    rate(
      somemetric{labelx="asdfasdf",labely="jajanounou"}[5m]
    )
  )
*
  1024
```

### or in `.vimrc`
```vimrc
augroup promql
  autocmd!
  autocmd FileType promql nmap <F3> :%!cat - \| promfmt<CR>
augroup END 
```

## Building

```
go build
```

## Copyright

The parser code is shamelessly copied from the [prometheus source code](https://github.com/prometheus/prometheus) which
is licensed under the apache license. I've only copied it because I needed
to access an unexported variable.



