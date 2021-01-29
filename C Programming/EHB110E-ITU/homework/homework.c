// This file main.c has set to just compile and run the Question 5. Other questions are in comment.
// Also, Question 4 can be compiled and runned by deleting its comment signs and commenting Question 5.


/*

Q1: Determine the values of variables product and x after the following calculation is performed. Assume that product and x each have the value 5 when the statement begins executing.
product *= x++;
x = 5;
product = 5;
product *= x++;
printf("%d", product);

Output >>> 25

*/

/*

Q2. Write a single C statement to accomplish each of the following:
a) Assign the sum of x and y to z and increment the value of x by 1 after the calculation.
    z = x++ + y;
b) Multiply the variable product by 2 using the *= operator.
    x *= 2;
c) Multiply the variable product by 2 using the = and * operators.
    x = x*2;
d) Test if the value of the variable count is greater than 10. If it is, print �Count is greater than 10.�
    if(value>10){ printf("Count is greater than 10.")}
e) Decrement the variable x by 1, then subtract it from the variable total.
    total -= --x;
f) Add the variable x to the variable total, then decrement x by 1.
    total += x++;
g) Calculate the remainder after q is divided by divisor and assign the result to q. Write this statement two different ways.
    q = q%divisor;
    q %= divisor;
h) Print the value 123.4567 with 2 digits of precision. What value is printed?
    printf("%.2f", 123.4567);
    Output >>> 123.46
i) Print the floating-point value 3.14159 with three digits to the right of the decimal point. What value is printed?
    printf("%.3f", 3.14159)

*/

/*

Q3 . Identify and correct the errors in each of the following.
a)
if ( age >= 65 ){
    printf( "Age is greater than or equal to 65\n" );
} else {
    printf( "Age is less than 65\n" );
}

b)
int x = 1;
int total = 5; //should be assigned to a value to make addition
while ( x <= 10 ) {
    total+= x;
    ++x;
}

c)
while ( x <= 100 ) {
    total += x;
    ++x;
}

d)
while ( y < 10 ) { // loop never ends on previous situation
    printf( "%d\n", y );
    ++y;
}

*/

/*

//Q4. (Credit Limit Calculator) Develop a C program that will determine if a department store customer has exceeded the credit limit on a charge account.

#include <stdio.h>
#include <stdlib.h>

int main() {

    while(1) {
        int account_ID;
        int balance_of_begin;
        int total_charged;
        int total_credits;
        int credit_limit;
        int final_balance;
        int total_exceeded;

        printf("*********** Welcome to Credit Limit Calculator ***********\n");

        printf("Please enter your account number (-1 to exit!): ");
        scanf("%d", &account_ID);

        if(account_ID == -1) {
            printf("********************* See you later! *********************\n");
            break;
        }

        printf("\nPlease enter your balance at the beginning of the month: ");
        scanf("%d", &balance_of_begin);

        printf("\nPlease enter your total of all items charged this month: ");
        scanf("%d", &total_charged);

        printf("\nPlease enter total of all credits applied to your account this month:  ");
        scanf("%d", &total_credits);

        printf("\nPlease enter your allowed credit limit: ");
        scanf("%d", &credit_limit);


        final_balance = balance_of_begin + total_charged - total_credits;
        total_exceeded = credit_limit - final_balance;


        if(credit_limit < final_balance) {
            printf("**********************************************\n");
            printf(" You have exceeded your credit limit :( ");
            printf("Account number: %d\nYour final balance: %d $\nRequired fund: %d $\n", account_ID, final_balance, -total_exceeded);
            printf("**********************************************\n");
        } else{
            printf("**********************************************\n");
            printf(" You are in your limits, the Economist :D ");
            printf("Account number: %d $\nYour final balance: %d $\n You can still spend %d $\n", account_ID, final_balance, total_exceeded);
            printf("**********************************************\n");

        }

    }

    return 0;
}

*/









//MY MAX VALUE PROGRAM STARTS FROM HERE. ABOVE IS FOR PREVIOUS QUESTIONS

//Q5. Write a program that takes 10 integers from the user and prints out the largest two integers entered by the user.
#include <stdio.h>
#include <stdlib.h>

int main() {
    int buffer_value1 = -1;
    int buffer_value2 = -1;
    int input;
    int x = 0;

    while(x<10) {
        printf("Please enter your %dth integer: ", x+1);
        scanf("%d:", &input);

        if(x==0) {
            buffer_value1 = input;
        } else if(x==1) {
            buffer_value2 = input;
            if(buffer_value1 < buffer_value2) {
                int w = buffer_value1;
                buffer_value1 = buffer_value2;
                buffer_value2 = w;
            }
        }

        if(input > buffer_value1 && input > buffer_value2  ) {
            buffer_value2 = buffer_value1;
            buffer_value1 = input;

        } else if(input <= buffer_value1 && input > buffer_value1) {
            buffer_value2 = input;
        }

        x++;
    }

    printf("The biggest value is: %d\n", buffer_value1);
    printf("Second biggest value is: %d\n", buffer_value2);


    return 0;
}
