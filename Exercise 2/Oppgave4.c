#include <pthread.h>
#include <stdio.h>
#include <sys/types.h>

int i = 0;
pthread_mutex_t lock;



void* thread1(){
    int j;
    for(j=0; j<1000000;j++){
        pthread_mutex_lock(&lock);
        i++;
    	pthread_mutex_unlock(&lock);
    }	
    return NULL;
}


void* thread2(){
    int j;
    for(j=0; j<999999;j++){
        pthread_mutex_lock(&lock);
        i--;
    	pthread_mutex_unlock(&lock);
    }
    return NULL;
}


int main(){
	pthread_mutex_init(&lock, NULL);	
	
    pthread_t firstThread;
    pthread_t secondThread;

    pthread_create(&firstThread, NULL, thread1, NULL);
    pthread_create(&secondThread, NULL, thread2, NULL);

    pthread_join(firstThread, NULL);
    pthread_join(secondThread, NULL);

    printf("The value is: %i \n", i);

    return 0;
}
