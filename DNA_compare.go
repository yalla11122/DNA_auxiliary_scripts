package main

import (
    "fmt"
)

// Function to compare two DNA sequences
func compareDNA(seq1, seq2 string) int {
    if len(seq1) != len(seq2) {
        return -1 // Return -1 if sequences are of different lengths
    }
    mismatches := 0
    for i := range seq1 {
        if seq1[i] != seq2[i] {
            mismatches++
        }
    }
    return mismatches
}

func main() {
    seq1 := "AGCT"
    seq2 := "AGGT"
    mismatches := compareDNA(seq1, seq2)
    if mismatches == -1 {
        fmt.Println("Sequences are of different lengths")
    } else {
        fmt.Printf("Number of mismatches: %d\n", mismatches)
    }
}
