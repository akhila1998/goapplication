package main

import (
    "bufio"
    "compress/gzip"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    // Open file on disk.
    name := "file.txt"
    f, _ := os.Open("C:\\programs\\" + name)

    // Create a Reader and use ReadAll to get all the bytes from the file.
    reader := bufio.NewReader(f)
    content, _ := ioutil.ReadAll(reader)

    // Replace txt extension with gz extension.
    name = strings.Replace(name, ".txt", ".gz", -1)

    // Open file for writing.
    f, _ = os.Create("C:\\programs\\" + name)

    // Write compressed data.
    w := gzip.NewWriter(f)
    w.Write(content)
    w.Close()

    // Done.
    fmt.Println("DONE")
}