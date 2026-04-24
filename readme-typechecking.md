# Go Format Verbs Cheat Sheet

This document lists all the important format verbs (`%` specifiers) used with `fmt.Printf`, `fmt.Sprintf`, etc. in Go.  

---

## 1. General
| Verb | Meaning | Example |
|------|---------|---------|
| `%v` | Default format | `fmt.Printf("%v", 123)` → `123` |
| `%+v` | Struct with field names | `fmt.Printf("%+v", p)` → `{Name:Farhan Age:25}` |
| `%#v` | Go-syntax representation | `fmt.Printf("%#v", []int{1,2})` → `[]int{1, 2}` |
| `%T` | Type of the value | `fmt.Printf("%T", 123)` → `int` |
| `%%` | A literal percent sign | `fmt.Printf("100%%")` → `100%` |

---

## 2. Boolean
| Verb | Meaning |
|------|---------|
| `%t` | `true` or `false` |

---

## 3. Integer
| Verb | Meaning | Example |
|------|---------|---------|
| `%d` | Base 10 | `255` |
| `%b` | Binary | `11111111` |
| `%o` | Octal | `377` |
| `%x` | Hex (lowercase) | `ff` |
| `%X` | Hex (uppercase) | `FF` |
| `%c` | Character (Unicode code point) | `'A'` for `65` |
| `%U` | Unicode format | `U+0041` |

---

## 4. Floating-point and Complex
| Verb | Meaning | Example |
|------|---------|---------|
| `%f` | Decimal (no exponent) | `123.456000` |
| `%F` | Same as `%f` | |
| `%e` | Scientific (lowercase) | `1.234560e+02` |
| `%E` | Scientific (uppercase) | `1.234560E+02` |
| `%g` | Compact (uses `%e` or `%f`) | `123.456` |
| `%G` | Compact (uses `%E` or `%F`) | |

---

## 5. String and Byte Slice
| Verb | Meaning | Example |
|------|---------|---------|
| `%s` | String | `"hello"` |
| `%q` | Quoted string | `"\"hello\""` |
| `%x` | Hex dump (lowercase) | `"68656c6c6f"` |
| `%X` | Hex dump (uppercase) | `"68656C6C6F"` |

---

## 6. Pointer
| Verb | Meaning | Example |
|------|---------|---------|
| `%p` | Pointer address (hex) | `0xc000010230` |

---

## 7. Width & Precision Modifiers
- `%6d` → pad integer to width **6**  
- `%.2f` → **2 decimal places** for floats  
- `%6.2f` → width **6**, 2 decimals  
- `%-6d` → left-aligned within width **6**

### Example
```go
fmt.Printf("|%6d|%-6d|%06d|\n", 123, 123, 123)
