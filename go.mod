module GoModule

go 1.22.5

// replace github.com/Pravin-Jalodiya/Puppy v2.0.0-20240718085847-044b42daa27e => C:\GoCodes\Puppy

require (
	example2 v0.0.0
	github.com/Pravin-Jalodiya/Puppy v1.1.0
	github.com/Pravin-Jalodiya/Version-Testing/v2 v2.0.0-pre.1
)

replace example2 v0.0.0 => C:\GoCodes\Example2
