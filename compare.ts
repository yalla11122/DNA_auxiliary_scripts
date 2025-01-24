function compareDNA(seq1: string, seq2: string): number {
  if (seq1.length !== seq2.length) {
    return -1; // Return -1 if sequences are of different lengths
  }
  let mismatches = 0;
  for (let i = 0; i < seq1.length; i++) {
    if (seq1[i] !== seq2[i]) {
      mismatches++;
    }
  }
  return mismatches;
}

// Example usage
const seq1 = "AGCT";
const seq2 = "AGGT";
const mismatches = compareDNA(seq1, seq2);
if (mismatches === -1) {
  console.log("Sequences are of different lengths");
} else {
  console.log(`Number of mismatches: ${mismatches}`);
}
