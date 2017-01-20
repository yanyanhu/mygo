package main


func main() {
    var n int;
    n = 0;
    println(n)

    switch {
	case n == 0 || n == 1:
	    println("n is equal to 0 or 1")
	default:
	    println("n is not equal to 0 and 1")
    }

/*    for i := 0; i < 10; i++ {
	println(i)
    } */

    //JLoop:
    for i := 0; i < 10; i++ {
	println(i)
	for j := 0;; {
	    j++
	    if j > 3 {
		//break JLoop
		goto Done
	    }
	}
    }
    Done:
    println("done")
}
