# vlog
Package vlog add leveled log on std log(golang.org/pkg/log/)

It implements most std log functions(except logger), variables and add provides V-style logging controlled by the -v flag or SetLogLevel()  
If flag.Parse be called before any logging, -v flag(default 0) use automaticlly.

## Basic examples:
```go
if !vlog.IsLevelParsed() {
	vlog.SetLogLevel(3)
}
vlog.GetLogLevel()

vlog.Println("Prepare to repel boarders")

vlog.Fatalf("Initialization failed: %s", err)
```

See the documentation for the V function for an explanation of these examples:
```go
if vlog.V(2) {
	vlog.Print("Starting transaction...")
}

vlog.V(2).Println("Processed", nItems, "elements")
```

## License
vlog is available under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)
