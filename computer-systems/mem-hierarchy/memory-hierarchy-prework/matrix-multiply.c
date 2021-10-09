/*
Naive code for multiplying two matrices together.

There must be a better way!
*/

#include <stdio.h>
#include <stdlib.h>

/*
  A naive implementation of matrix multiplication.

  DO NOT MODIFY THIS FUNCTION, the tests assume it works correctly, which it
  currently does
*/
void matrix_multiply(double **C, double **A, double **B, int a_rows, int a_cols,
                     int b_cols) {
  for (int i = 0; i < a_rows; i++) {
    for (int j = 0; j < b_cols; j++) {
      C[i][j] = 0;
      for (int k = 0; k < a_cols; k++)
        C[i][j] += A[i][k] * B[k][j];
    }
  }
}

void fast_matrix_multiply(double **c, double **a, double **b, int a_rows,
                          int a_cols, int b_cols) {
  // Before transposing:
  // c[i][j] is the dot product of ith row of a, jth col of b
  
  // transpose B
  for (int row = 0; row < a_cols; row++) {
    double* brow = b[row];
    for (int col = row+1; col < b_cols; col++) {
      double temp = brow[col];
      brow[col] = b[col][row];
      b[col][row] = temp;
    }
  }

  // After tranposing:
  // c[i][j] is the dot product of ith row of a, jth row of b
  double dot_product;
  for (int a_row = 0; a_row < a_rows; a_row++) {
    double* crow = c[a_row];
    for (int b_row = 0; b_row < b_cols; b_row++) {
      double* brow = b[b_row];
      double* arow = a[a_row];
      dot_product = 0;
      for (int col = 0; col < a_cols; col++) {
        dot_product += arow[col] * brow[col];
      }
      crow[b_row] = dot_product;
    }
  }
}


