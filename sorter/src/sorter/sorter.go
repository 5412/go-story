package main

import "flag"
import "fmt"
import "bufio"
import "io"
import "os"
import "strconv"
import "time"

import "algorithms/bubblesort"
import "algorithms/qsort"

// infile 是一个指针变量， *infile 表示此指针变量指向的变量的值， &infile 表示此指针变量的内存地址， 结论指针变量是一种特殊的变量
var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "sort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
    file, err := os.Open(infile)
    if err != nil {
        fmt.Println("Failed to open the input file ", infile)
        return 
        //return GoLang不允许在if里面写return,理解错误并非是不允许写return，而是写在if条件里的return编译器不能识别为是此函数的最终return，如果不写最终return，程序会报错
    }
    
    defer file.Close() //确保出现任何异常文件正常关闭

    br := bufio.NewReader(file)

    values = make([]int, 0)

    for {
        line, isPrefix, err1 := br.ReadLine()
        if err1 != nil {
            if err1 != io.EOF {
                err = err1
            }
            break
        }
        if isPrefix {
            fmt.Println("A too long line, seems unexpected.")
            return
        }

        str := string(line)

        value, err1 := strconv.Atoi(str)

        if err1 != nil {
            err = err1
            return
        }

        values = append(values, value)
    }
    return
}

func writeValues(values []int, outfile string) error {
    file, err := os.Create(outfile)
    if err != nil {
        fmt.Println("Failed to create the outfile ", outfile)
        return err
    }

    defer file.Close()

    for _, value := range values {
        str := strconv.Itoa(value)
        file.WriteString(str + "\n")
    }
    return nil
}

func main() {
    flag.Parse()
    if infile != nil {
        fmt.Println("infile = ", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
    }

    values, err := readValues(*infile)

    if err == nil {
        fmt.Println("Read values:", values)
        t1 := time.Now()
        switch *algorithm {
            case "qsort":
                qsort.QuickSort(values)
            case "bubblesort":
                bubblesort.BubbleSort(values)
            default:
                fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
        }
        t2 := time.Now()

        fmt.Println("The sorting process costs", t2.Sub(t1), "to complete")

        writeValues(values, *outfile)
    } else {
        fmt.Println(err)
    }
}

