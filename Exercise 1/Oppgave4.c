#include <pthread.h>
#include <stdio.h>

int i = 0;

void* thread1(){
    int j;
    for(j=0; j<1000000;j++){
        i++;
    }
    return NULL;
}


void* thread2(){
    int j;
    for(j=0; j<1000000;j++){
        i--;
    }
    return NULL;
}


int main(){
    pthread_t firstThread;
    pthread_t secondThread;

    pthread_create(&firstThread, NULL, thread1, NULL);
    pthread_create(&secondThread, NULL, thread2, NULL);

    pthread_join(firstThread, NULL);
    pthread_join(secondThread, NULL);

    printf("The value is: %i \n", i);

    return 0;
}