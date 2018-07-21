package main

import (
   "fmt"
   "runtime"
   "os"
   "time"
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





func main() {

fmt.Println("Program to make unique a slice of integers with duplicate values:")

intSlice := []int{1,4,5,4,3,3,44,3,2,2,3,4,88,8,7,0,9,9,5,2,3,5}

fmt.Println("Slice: ", intSlice)
unqSlice := unique(intSlice)

fmt.Println("Making it unique(slice): ")
fmt.Println(unqSlice)

 //++++++++++  footer 
 if debugTrue() {

    printEnv()

    checkGC()

     u.V(version)
  }

} // end of main
 
// Define unique()

func unique(slc []int) []int {
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



// Function to check env variable DEFAULT_DEBUG bool is set
func debugTrue() bool {

     if os.Getenv("DEFAULT_DEBUG") != "" {
         return true
     }
return false 
}
