// Elev_EventManager.c

#include "Elev_StateMachine.h"
#include "elev.h"
#include "Elev_Timer.h"


/*-----------------Private-----------------*/

static void check_button_signals(){
	int i;
		for(i = 0; i < 4; i ++){
			if(i < 3 && elev_get_button_signal(BUTTON_CALL_UP,i) != 0){
				sm_new_order(BUTTON_CALL_UP,i);
			}
			if (i > 0 && elev_get_button_signal(BUTTON_CALL_DOWN,i) != 0){
				sm_new_order(BUTTON_CALL_DOWN, i);
			}
			if(elev_get_button_signal(BUTTON_COMMAND,i) != 0){
				sm_new_order(BUTTON_COMMAND, i);
			}
		}
}




int main(){
	sm_initiate();
	while( !elev_get_obstruction_signal() ){
        	check_button_signals();
		sm_emergency_stop(elev_get_stop_signal(), elev_get_floor_sensor_signal());
		sm_floor_sensor(elev_get_floor_sensor_signal());
		if(time_get_is_done()){
			sm_time_done();
		}
 	}
	elev_init();
	elev_set_motor_direction(DIRN_STOP);
	return 0;
}
