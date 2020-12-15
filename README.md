# Bugospots
Bugospots is a Go implementation of the reference [igrigorik/bugspots](https://github.com/igrigorik/bugspots) - a bug prediction tool.  
The bug prediction algorithm is stated on chapter V. in the paper - [Does Bug Prediction Support Human Developers? Findings from a Google Case Study](https://research.google/pubs/pub41145/).

## Note
Bugospots is buggy and unstable.  
Some behavior of Bugospots is not equal to the reference's.

## Building and Dependency
Bugospots uses [go-git/go-git](https://github.com/go-git/go-git).

```
$ go get github.com/go-git/go-git
$ go build -o bugospots bugospots.go
```

## Usage
```
$ bugospots -path <A path to the target Git repository>
```

|option|default|description|
|----|----|----|
|-regexp|(?i)(^\| )(fi(x\|xed\|xes)\|clos(e\|es\|ed))|A regurar expression specifying bug fix commit message|
|-o|./bugospots.csv|Full result of csv file|

You will get top 10 hotspots score and its relative file path in the target repository, and full result in csv file.  
Higher score represents higher possibility of including bugs.
Following shows sample output.

```
2020/12/31 19:43:33 oldest bug fix: 2015-01-01 00:00:01 +0900 JST
2020/12/31 19:43:33 latest bug fix: 2020-12-29 13:24:35 +0900 JST
2020/12/31 19:43:33 current: 2020-12-29 14:00:00.1542176 +0900 JST
2020/12/31 19:43:33 bug fixes: 987
1.0000000123456789,edit/and/crash.c
0.9123456789123456,want/to/throw/away.mk
0.8012345678901234,poopy.py
0.7654321098765432,time/spoiler.js
0.6060606060606060,stinker.cpp
0.5353535353535353,massive_logic.h
0.4321431431243124,SingletonLover.java
0.3939393939393939,cannot/read.yml
0.21021021021021021,god.go
0.10000001234567890,shutup.sh
```

## License
[MIT](LICENSE)
