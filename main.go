package main

import (
   "fmt"
   "runtime"
   "os"
   "time"
   "sort"
   u "github.com/ardeshir/version"
)

var (
 debug bool = false
 version string = "0.0.1"
)

func printEnv() {
    
    fmt.Print("Using ", runtime.Compiler, " ")
    fmt.Println("on a", runtime.GOARCH, "machine")
    fmt.Println("Go version", runtime.Version())
    fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())
    fmt.Println(" ")
}

func printStats(mem runtime.MemStats) {
    runtime.ReadMemStats(&mem)

    fmt.Println("mem.Alloc: ", mem.Alloc)
    fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
    fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
    fmt.Println("mem.NumGC: ", mem.NumGC)
    fmt.Println("------")
}

func checkGC() {
    var mem runtime.MemStats
    printStats(mem)

    for i := 0; i< 10; i++ {
        s := make([]byte, 100000000)

        if s == nil {
           fmt.Println("Operation failed")
        }
     }

     printStats(mem)

      for i := 0; i< 10; i++ {

         s := make([]byte, 100000000)
         if s == nil {
            fmt.Println("Operation failed")
         }
         time.Sleep(2 * time.Second)
    }

     printStats(mem)

}

// This is used for map -> slice example

type NameAge struct {
     Name string
     Age int
}


func main() {

fmt.Println("Program to make unique a slice of integers with duplicate values:")

intSlice := []int{1,4,5,4,3,3,44,3,2,2,3,4,88,8,7,0,9,9,5,2,3,5}

fmt.Println("Slice: ", intSlice)
unqSlice := uniqueInt(intSlice)

fmt.Println("Making it unique(slice):", unqSlice)

unqSorted := sort.IntSlice(unqSlice)
sort.Sort(unqSorted)

fmt.Println("Sorting it: ", unqSorted)


str := []string{"Ardeshir", "Ardeshir","Sepahsalar", "Amir","Amir", "Khademzadeh","Casey","Casey","Depasquale", "Kayhan","Anoush", "Anoush"}
fmt.Println("Before uniqueStr: ", str)

str = uniqueStr(str)
fmt.Println("After uniqueStr: ", str)

for i, v := range str {
    if v == "Kayhan" {
        fmt.Println("index:", i , " value:", v)
    }
}

// sort them 
sorted := sort.StringSlice(str)
sorted.Sort()
fmt.Println(sorted)

var nameAgeSlice []NameAge

ourAges := map[string]int{
    "Ardesir": 47,
    "Casey":   35,
    "Kayan":    8,
    "Anoush":   6 }

// append to the slice a map for each value
for key, value := range ourAges {
    nameAgeSlice = append(nameAgeSlice, NameAge {key, value} )
}

// range over the slice and print the Struct 
fmt.Println(nameAgeSlice)

 for _, val := range nameAgeSlice {
     
    //  fmt.Println("Value: ", val)
     fmt.Println("Name: ", val.Name, " Age: ", val.Age)

 }



///=++++++++++++++++++++++++++++++++++++++++++=///
///=+++++++++++++++++++++++++++++++++++++++++++=///
///=++++++++++  footer 
 if debugTrue() {

    printEnv()

    checkGC()

     u.V(version)
  }

} // end of main
 
// Find an element from an array or slice 
 
 
 
 
// Define uniqueInt()

func uniqueInt(slc []int) []int {
    // use a map of int->bool to capture if int exists
    keys := make(map[int]bool)
    uniqElems := []int{}
    
    for _, valInt := range slc {
         // fmt.Println("Index: ", _, " ValueInt: ", valInt)
         if _, valueMap := keys[valInt]; !valueMap {
             // fmt.Println("ixMap", _, " ValueMap:", valueMap)
             keys[valInt] = true
             uniqElems = append(uniqElems, valInt)
         }
    } // end of for 
    
    return uniqElems
} // end of unique()


// Define unique()

func uniqueStr(slc []string) []string {
    // use a map of int->bool to capture if int exists
    keys := make(map[string]bool)
    uniqElems := []string{}
    
    for _, valInt := range slc {
         // fmt.Println("Index: ", _, " ValueInt: ", valInt)
         if _, valueMap := keys[valInt]; !valueMap {
             // fmt.Println("ixMap", _, " ValueMap:", valueMap)
             keys[valInt] = true
             uniqElems = append(uniqElems, valInt)
         }
    } // end of for 
    
    return uniqElems
} // end of uniqueStr()



// Function to check env variable DEFAULT_DEBUG bool is set
func debugTrue() bool {

     if os.Getenv("DEFAULT_DEBUG") != "" {
         return true
     }
return false 
}
