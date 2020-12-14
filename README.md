# Bugospots
Bugospots is a Go implementation of the reference [igrigorik/bugspots](https://github.com/igrigorik/bugspots) - a bug prediction tool.
The bug prediction algorithm is stated on chapter V. in the paper - [Does Bug Prediction Support Human Developers? Findings from a Google Case Study](https://research.google/pubs/pub41145/).

## Note
Bugospots is buggy and unstable.  
Some behavior of Bugospots is not equal to the reference's.

## Building and Dependency
Bugospots uses [go-git/go-git](https://github.com/go-git/go-git).

```
$ go build -o bugospots bugospots.go
```

## Usage
```
$ bugospots -path <A path to the target Git repository>
```

|option|default|description|
|----|----|----|
|-regexp|`(?i)(^| )(fi(x|xed|xes)|clos(e|es|ed))`|A regurar expression of bug fix|
|-o|./bugospots.csv|An output of csv file|
	

## License
[MIT](LICENSE)
