Golang delta compression algorithm
===
Fossil achieves efficient storage and low-bandwidth synchronization through the
use of delta-compression. Instead of storing or transmitting the complete
content of an artifact, fossil stores or transmits only the changes relative to
a related artifact.

* [Format](http://www.fossil-scm.org/index.html/doc/tip/www/delta_format.wiki)
* [Algorithm](http://www.fossil-scm.org/index.html/doc/tip/www/delta_encoder_algorithm.wiki)
* [Original implementation](http://www.fossil-scm.org/index.html/artifact/f3002e96cc35f37b)

Other implementations:
- [C#](https://github.com/endel/FossilDelta)
- [Haxe](https://github.com/endel/fossil-delta-hx)
- [Python](https://github.com/ggicci/python-fossil-delta)
- [JavaScript](https://github.com/dchest/fossil-delta-js) ([Online demo](https://dchest.github.io/fossil-delta-js/))

### Installation:
#### Install:
```go
$ go get -u github.com/mdvan/fdelta
```
#### Import:
```go
import "github.com/mdvan/fdelta"
```
### Example
```go
package main

import (
	"fmt"

	"github.com/mdvan/fdelta"
)

func main() {
	origin := []byte("abcdefghijklmnopqrstuvwxyz1234567890")

	target := []byte("abcdefghijklmnopqrstuvwxyz0987654321")

	//Create delta
	delta := fdelta.Create(origin, target)
	
	//Create target by patching origin with delta
	target2, err := fdelta.Apply(origin, delta)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Origin : `%s`\n", origin)
	fmt.Printf("Target : `%s`\n", target)
	fmt.Printf("Delta  : `%s`\n", delta)
	fmt.Printf("Target2: `%s`\n", target2)
}
```