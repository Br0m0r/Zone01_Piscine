    _____ _    _ _____   ____  _  ___    _ 
   / ____| |  | |  __ \ / __ \| |/ / |  | |
  | (___ | |  | | |  | | |  | | ' /| |  | |
   \___ \| |  | | |  | | |  | |  < | |  | |
   ____) | |__| | |__| | |__| | . \| |__| |
  |_____/ \____/|_____/ \____/|_|\_\\____/ 
                                          
 
MAIN: This is the entry point of the program. It initializes the puzzle grid based on input arguments, and triggers the solving process using backtracking.

BACKTRACKING: This recursive function attempts to solve the Sudoku puzzle by placing numbers in empty cells (denoted by 0). It chooses possible candidates for each cell and recursively tries to solve the puzzle with each candidate. If a candidate leads to a solution, it returns true. If none do, it backtracks by resetting the cell and trying the next candidate.

GETCANDIDATES: Given a cell's position in the grid, this function returns a list of valid numbers (1-9) that can be placed in that cell without violating Sudoku rules.

CONTAINS: Checks if a given number is already in a specific row of the grid.

ISCOLUMNDUPLICATE: Checks if placing a given number in a specific column would violate the Sudoku rules.

ISSUBGRIDDUPLICATE: Checks if placing a given number in a specific 3x3 subgrid would violate the Sudoku rules.

PRINTGRID: Prints the grid to the console. This function formats the grid for easy reading, displaying each row on a new line.

GRIDCOMPLETE: Checks if the grid is fully solved by ensuring there are no cells containing 0.



+----------------+
|    main()      |
+----------------+
         |
         v
+----------------+      +-----------------------+
|  backtracking()|----->| gridComplete()        |
+----------------+      +-----------------------+
         |                       |
         v                       v
+----------------+      +-----------------------+
| getCandidates()|      |    printGrid()        |
+----------------+      +-----------------------+
         |
         v
+----------------+      +-----------------------+
|  contains()    |      | isColumnDuplicate()   |
+----------------+      +-----------------------+
         |                       |
         v                       v
+-----------------------+
| isSubgridDuplicate()  |
+-----------------------+



This flowchart starts at main(), where the grid is set up and backtracking() is called. The backtracking() function uses gridComplete() to check if the grid is fully solved. If not, it calls getCandidates() to get valid numbers for each cell, and for each candidate, it recurses. If the recursion returns true, the solution is printed using printGrid(). Helper functions contains(), isColumnDuplicate(), and isSubgridDuplicate() assist in determining valid candidates. If no solution is found, it backtracks by resetting the cell and continues with the next candidate.


