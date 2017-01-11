//Elev_StateMachine.h

#include "elev.h"

typedef enum tag_state{
	STATE_INITIALIZE = 0,
	STATE_DORMENT    = 1,
	STATE_STOP       = 2,
	STATE_UP         = 3,
	STATE_DOWN       = 4,
	STATE_EMERGENCY  = 5
}state_t;


typedef enum tag_elev_Direction_type{
	Down = -1,
	Up   =  1
} elev_Direction_type_t;



void sm_initiate();


void sm_new_order(elev_button_type_t orderType, int floor);


void sm_emergency_stop(int emergencyButtonState, int floorSensor);


void sm_floor_sensor(int floor);


void sm_time_done();
