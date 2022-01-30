// Ulaş Erdoğan - 040190230

#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#define ROLLING 10000 //keeps the number of times the dice will be rolled
#define GRAPH1SCALING 70 // scaling coefficient to scale graph 1
#define GRAPH2SCALING 200 // scaling coefficient to scale graph 2

// function prototypes
int rollDice();
int findMax(int array[], int size);

int main() {

    // arrray that keep the frequency data of rolling dice operations results
    int results[12] = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0};
    
    srand(time(NULL)); // takes computer time as seed unsigned integer for dice rolling operation

    // a program that rolls dice 10,000 times and draw its frequency table

    //rolling dice 'ROLLING' times
    for(int i = 0; i < ROLLING; i++) {

        // rolling dice and adding the result the counters
        int diceFace; // variable that keeps one time rolling result temporarily
        diceFace = rollDice(); // rolling dice

        switch(diceFace) {
        case 1: 
            results[0]++;
            break;
        case 2:
            results[1]++;
            break;
        case 3:
            results[2]++;
            break;
        case 4:
            results[3]++;
            break;
        case 5:
            results[4]++;
            break;
        case 6:
            results[5]++;
            break;
        }
    }

    //scaling operations, the results will be scaled by "GRAPH1SCALING" and the graph will get better look
    for(int i=0; i < sizeof(results) / sizeof(results[0]); i++ ){
        results[i] /= GRAPH1SCALING;
    }

    // defined findMax function to observe the data and find the maximum (scaled) frequency (maximum rows) to draw the table
    int max = findMax(results, sizeof(results) / sizeof(results[0]));
    
    /*
    final - drawing frequency table part:
    the program continues row by row and column by column, determines the space and * pieces for every row-column parts (cells)
    */

    // loop to draw each row, there are 'max' number of rows, i values match to row numbers
    for(int i = 1; i <= max; i++) {
        // loop to fill 6 columns pieces, k values match to column numbers
        for(int k = 1; k <= 6; k++) {
            //understands that if the frequency is equal to the max value, it puts * there.
            // if not equal puts space and increment frequency value by one to find which row the *'s of this columns start
            switch(k) {
            case 1:
                if(results[0] != max) {
                    printf("  ");
                    results[0]++;
                } else {
                    printf("* ");
                }
                break;
            case 2:
                if(results[1] != max) {
                    printf("  ");
                    results[1]++;
                } else {
                    printf("* ");
                }
                break;
            case 3:
                if(results[2] != max) {
                    printf("  ");
                    results[2]++;
                } else {
                    printf("* ");
                }
                break;
            case 4:
                if(results[3] != max) {
                    printf("  ");
                    results[3]++;
                } else {
                    printf("* ");
                }
                break;
            case 5:
                if(results[4] != max) {
                    printf("  ");
                    results[4]++;
                } else {
                    printf("* ");
                }
                break;
            case 6:
                if(results[5] != max) {
                    printf("  ");
                    results[5]++;
                } else {
                    printf("* ");
                }
                break;
            }

        }
        // passing from one row to other row by a new line
        printf("\n");
    }

    // printing the dice faces
    printf("%d%2d%2d%2d%2d%2d\n\n", 1, 2, 3, 4, 5, 6);


    // SECOND GRAPH

    //rolling dices 'ROLLING's times
    for(int i = 0; i < ROLLING; i++) {

        // rolling dices and adding the result the counters
        int sumDice; // variable that keeps the sum of dice faces
        sumDice = rollDice() + rollDice(); // rolling dice abd keep the face on 

        switch(sumDice) {
        case 2:
            results[1]++;
            break;
        case 3:
            results[2]++;
            break;
        case 4:
            results[3]++;
            break;
        case 5:
            results[4]++;
            break;
        case 6:
            results[5]++;
            break;
        case 7:
            results[6]++;
            break;
        case 8:
            results[7]++;
            break;
        case 9:
            results[8]++;
            break;
        case 10:
            results[9]++;
            break;
        case 11:
            results[10]++;
            break;
        case 12:
            results[11]++;
            break;
        }
    }

    //scaling operations, the results will be scaled by "GRAPH2SCALING" and the graph will get better look
    for(int i=0; i < sizeof(results) / sizeof(results[0]); i++ ){
        results[i] /= GRAPH2SCALING;
    }

    // defined findMax function to observe the data and find the maximum (scaled) frequency (maximum rows) to draw the table
    max = findMax(results, sizeof(results) / sizeof(results[0]));
 
    /*
    final - drawing frequency table part:
    the program continues row by row and column by column, determines the space and * pieces for every row-column parts (cells)
    */
   
    // loop to draw each row, there are 'max' number of rows, i values match to row numbers
    for(int i = 1; i <= max; i++) {
        // loop to fill 6 columns pieces, k values match to column numbers
        for(int k = 2; k <= 12; k++) {
            //understands that if the frequency is equal to the max value, it puts * there.
            // if not equal puts space and increment frequency value by one to find which row the *'s of this columns start
            switch(k) {
            case 2:
                if(results[1] != max) {
                    printf("  ");
                    results[1]++;
                } else {
                    printf("* ");
                }
                break;
            case 3:
                if(results[2] != max) {
                    printf("  ");
                    results[2]++;
                } else {
                    printf("* ");
                }
                break;
            case 4:
                if(results[3] != max) {
                    printf("  ");
                    results[3]++;
                } else {
                    printf("* ");
                }
                break;
            case 5:
                if(results[4] != max) {
                    printf("  ");
                    results[4]++;
                } else {
                    printf("* ");
                }
                break;
            case 6:
                if(results[5] != max) {
                    printf("  ");
                    results[5]++;
                } else {
                    printf("* ");
                }
                break;
            case 7:
                if(results[6] != max) {
                    printf("  ");
                    results[6]++;
                } else {
                    printf("* ");
                }
                break;
            case 8:
                if(results[7] != max) {
                    printf("  ");
                    results[7]++;
                } else {
                    printf("* ");
                }
                break;
            case 9:
                if(results[8] != max) {
                    printf("  ");
                    results[8]++;
                } else {
                    printf("* ");
                }
                break;
            case 10:
                if(results[9] != max) {
                    printf("   ");
                    results[9]++;
                } else {
                    printf("*  ");
                }
                break;
            case 11:
                if(results[10] != max) {
                    printf("   ");
                    results[10]++;
                } else {
                    printf("*  ");
                }
                break;
            case 12:
                if(results[11] != max) {
                    printf("   ");
                    results[11]++;
                } else {
                    printf("*  ");
                }
                break;
            }

        }
        // passing from one row to other row by a new line
        printf("\n");
    }

    // printing the dice faces
    printf("%d%2d%2d%2d%2d%2d%2d%2d%3d%3d%3d\n", 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12);

    return 0;
}


// function to get a result for dice rolling operation
int rollDice() {
    int diceFace; // variable to keep rolling result
    diceFace = 1 + rand()%6; // random integer operation between 1-6 range

    return diceFace; // returns an integer which is the result of the operation
}

// function to find max frequencied rolling result in specified sized array
int findMax(int array[], int size) {
    int maxValue = 0; // defined a temporary value to find the maximum. The lowest frequency can be 0, so it has been chosen 0.

    for(int i = 0; i < size; i++) {
        if(array[i] > maxValue) {
            maxValue = array[i];
        }
    }

    return maxValue; // at the final, return the maximum frequency of the results
}
