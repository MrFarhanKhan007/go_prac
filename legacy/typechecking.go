package main

import (
	"fmt"
    "reflect"
)

type Person struct {
	Name string
	Age  int
}

func main3() {
	// ---------- 1. General ----------
	p := Person{Name: "Farhan", Age: 25}
	fmt.Printf("%%v: %v\n", 123)
	fmt.Printf("%%+v: %+v\n", p)
	fmt.Printf("%%#v: %#v\n", []int{1, 2})
	fmt.Printf("%%T: %T\n", 123)
	fmt.Printf("%%%%: 100%%\n")

	// ---------- 2. Boolean ----------
	fmt.Printf("%%t: %t\n", true)

	// ---------- 3. Integer ----------
	num := 255
	fmt.Printf("%%d: %d\n", num)
	fmt.Printf("%%b: %b\n", num)
	fmt.Printf("%%o: %o\n", num)
	fmt.Printf("%%x: %x\n", num)
	fmt.Printf("%%X: %X\n", num)
	fmt.Printf("%%c: %c\n", 65) // Unicode code point → 'A'
	fmt.Printf("%%U: %U\n", 65)

	// ---------- 4. Floating-point ----------
	f := 123.456
	fmt.Printf("%%f: %f\n", f)
	fmt.Printf("%%F: %F\n", f)
	fmt.Printf("%%e: %e\n", f)
	fmt.Printf("%%E: %E\n", f)
	fmt.Printf("%%g: %g\n", f)
	fmt.Printf("%%G: %G\n", f)

	// ---------- 5. String & Byte Slice ----------
	str := "hello"
	fmt.Printf("%%s: %s\n", str)
	fmt.Printf("%%q: %q\n", str)
	fmt.Printf("%%x: %x\n", str)
	fmt.Printf("%%X: %X\n", str)

	// ---------- 6. Pointer ----------
	fmt.Printf("%%p: %p\n", &num)

	// ---------- 7. Width & Precision ----------
	fmt.Printf("|%%6d| → |%6d|\n", 123)
	fmt.Printf("|%%-6d| → |%-6d|\n", 123)
	fmt.Printf("|%%06d| → |%06d|\n", 123)
	fmt.Printf("|%%6.2f| → |%6.2f|\n", f)

    // ---------- 8. Reflection ----------
    s:= "hello"
    s1:= 'A'
    s2:= 123
    s3:= 123.456
    s4:= true
    s5:= []int{1, 2, 3}
    s6:= Person{Name: "Farhan", Age: 25}

	// string
	fmt.Println("reflect:", reflect.TypeOf(s))
	fmt.Println("reflect:", reflect.TypeOf(s1))
	fmt.Println("reflect:", reflect.TypeOf(s2))
	fmt.Println("reflect:", reflect.TypeOf(s3))
	fmt.Println("reflect:", reflect.TypeOf(s4))
	fmt.Println("reflect:", reflect.TypeOf(s5))
	fmt.Println("reflect:", reflect.TypeOf(s6))
}
