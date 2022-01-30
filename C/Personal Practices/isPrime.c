#include <stdio.h>
#include <stdlib.h>
#include <conio.h>

int main(){
    // defined two integers to keep the limit values
    int first_value;
    int second_value;

    // getting the values from user input
    printf("Ilk Deger: ");
    scanf("%d", &first_value);
    printf("Ikinci Deger: ");
    scanf("%d", &second_value);

    if(first_value == second_value){
        printf("Bir deger araligi girmediniz...");
        getch();
        return 0;
    }

    // print statement to print title
    printf("%d ile %d Arasindaki Asal Sayilar: \n", first_value, second_value);

    // a kind of Boolean to be changed in primes
    int isPrime;

    // to avoid bugs about the special prime number: 2
    if(first_value == 1 && second_value != 2){
        printf("2, ");
    }

    // loop for wandering between the limit values to find prime numbers
    for(int i = first_value + 1; i < second_value; i++) {
        // setting our controller to default value for each number
        isPrime = 0;
        // checking 'Is this integer prime?'
        // starting from the smallest prime number 2, to the number that is being controling
        for( int j = 2; j < i; j++){
            if(i%j == 0){
                isPrime = 0;
                break;
            } else {
               isPrime = 1; 
            }
        }

        //printing the prime numbers
        if(isPrime==1){
            printf("%d, ", i);
        }
    }

    getch();
    return 0;
}
