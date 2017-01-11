//Elev_Timer.c

#include <time.h>
#include <sys/time.h>
#include <stdbool.h>

int targetTime = -1;
struct timeval tv;


void time_start_timer(int time){
	gettimeofday(&tv, NULL);
	targetTime = tv.tv_sec + time;
}


bool time_get_is_done(){
	gettimeofday(&tv,NULL);
	int curentTime = tv.tv_sec;
	if (targetTime <= curentTime && targetTime != -1){
		return true;
	}
	else{
		return false;
	}
}


void time_reset_timer(){
	targetTime = -1;
}


